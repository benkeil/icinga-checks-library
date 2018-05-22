package icinga

type (
	// Results contains multiple results for service checks
	Results interface {
		All() []Result
		Add(Result)
		CalculateStatus() Status
	}

	resultsImpl struct {
		results      map[string]Result
		statusPolicy StatusPolicy
	}

	// ResultsOptions options to generate a new instance of Results
	ResultsOptions struct {
		StatusPolicy StatusPolicy
	}
)

// NewResults creates a new instance of Results
func NewResults() Results {
	return &resultsImpl{make(map[string]Result), NewDefaultStatusPolicy()}
}

// NewResultsWithOptions creates a new instance of Results with options
func NewResultsWithOptions(options ResultsOptions) Results {
	var statusPolicy StatusPolicy
	if options.StatusPolicy != nil {
		statusPolicy = options.StatusPolicy
	} else {
		statusPolicy = NewDefaultStatusPolicy()
	}
	return &resultsImpl{make(map[string]Result), statusPolicy}
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
