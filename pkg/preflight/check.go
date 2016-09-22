package preflight

// A Check is a preflight check that verifies some condition on the node.
type Check interface {
	// Check returns nil if the check was successful. Otherwise, returns an error message
	// describing the problem, and potential remediation steps.
	Check() error
	// Name returns the name of the check
	Name() string
}

// A ClosableCheck is a preflight check that needs to be closed.
type ClosableCheck interface {
	Check
	Close() error
}

// CheckRequest contains the list of checks that should be run
type CheckRequest struct {
	BinaryDependencies  []string
	PackageDependencies []string
	TCPPorts            []int
}
