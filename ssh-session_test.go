package icinga

import "testing"

func TestNewSshSessionError(t *testing.T) {

	tests := []struct {
		host string
		port int
		ssh_user string
		msg string
	}{
		{"", 22, "ssh-user", "missing host"},
		{"host", 0, "ssh-user", "missing port"},
		{"host", 22, "", "missing ssh user"},
		{"host", 22, "ssh-user", "missing ssh directory"},
	}
	for _, test := range tests {
		_, err := NewSshSession(test.host, test.port, test.ssh_user)
		if err == nil {
			t.Fatalf("Failed to test missing mandatory parameter: %v", test.msg)
		}
	}
}
