package icinga

import (
	"bytes"
	"fmt"
	"os"
)

type (
	// Results contains multiple results for service checks
	Results interface {
		All() []Result
		Add(Result)
		CalculateStatus() Status
		GenerateMessage() string
		Exit()
	}

	resultsImpl struct {
		results             map[string]Result
		statusPolicy        StatusPolicy
		statusMessagePolicy StatusMessagePolicy
	}

	// ResultsOptions options to generate a new instance of Results
	ResultsOptions struct {
		StatusPolicy        StatusPolicy
		StatusMessagePolicy StatusMessagePolicy
	}
)

// NewResults creates a new instance of Results
func NewResults() Results {
	return &resultsImpl{make(map[string]Result), NewDefaultStatusPolicy(), NewDefaultStatusMessagePolicy()}
}

// NewResultsWithOptions creates a new instance of Results with options
func NewResultsWithOptions(options ResultsOptions) Results {
	var statusPolicy StatusPolicy
	if options.StatusPolicy != nil {
		statusPolicy = options.StatusPolicy
	} else {
		statusPolicy = NewDefaultStatusPolicy()
	}

	var statusMessagePolicy StatusMessagePolicy
	if options.StatusMessagePolicy != nil {
		statusMessagePolicy = options.StatusMessagePolicy
	} else {
		statusMessagePolicy = NewDefaultStatusMessagePolicy()
	}
	return &resultsImpl{make(map[string]Result), statusPolicy, statusMessagePolicy}
}

// Add adds a element to the set
func (r *resultsImpl) Add(result Result) {
	r.results[result.Name()] = result
}

// All returns all values from the set
func (r *resultsImpl) All() []Result {
	values := []Result{}
	for _, value := range r.results {
		values = append(values, value)
	}
	return values
}

// CalculateStatus calculates the service state for multiple checks
func (r *resultsImpl) CalculateStatus() Status {
	return r.statusPolicy.Calculate(r)
}

// GenerateMessage generates the overall message for multiple checks
func (r *resultsImpl) GenerateMessage() string {
	return r.statusMessagePolicy.Generate(r)
}

var resultOrder = []Status{
	ServiceStatusUnknown,
	ServiceStatusCritical,
	ServiceStatusWarning,
	ServiceStatusOk,
}

// Exit prints the check result and exits the program
func (r *resultsImpl) Exit() {
	fmt.Println(r)
	os.Exit(r.CalculateStatus().Ordinal())
}

func (r *resultsImpl) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s\n", r.GenerateMessage()))

	// group all checks by status
	statusMap := make(map[Status][]Result)
	for _, result := range r.All() {
		statusMap[result.Status()] = append(statusMap[result.Status()], result)
	}

	for _, status := range resultOrder {
		if results, found := statusMap[status]; found {
			for _, result := range results {
				buffer.WriteString(fmt.Sprintf("%s: %s: %s\n", status, result.Name(), result.Message()))
			}
		}
	}

	return buffer.String()
}
