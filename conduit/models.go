package conduit

// Application holds app related information
type Application struct {
	ID    string
	Name  string
	OrgID string
}

// Group is a logical grouping of multiple steps in a workflow
type Group struct {
	ID        string
	Name      string
	Message   string
	Completed bool
	Steps     []*Step
}

// Step holds the information of a step in a workflow
type Step struct {
	ID         string
	Name       string
	Completed  bool
	WorkflowID string
	GroupID    string
}

// Workflow holds workflow related information
type Workflow struct {
	ID            string
	Name          string
	ApplicationID string
	Groups        []*Group
}

// Transaction holds transaction related information
type Transaction struct {
	ID         string
	WorkflowID string
	Status     string
	Workflow   *Workflow
}

// Manager holds all the information of the application in memory
type Manager struct {
	Applications           map[string]*Application
	WorkflowApplicationMap map[string]string
	Workflows              map[string]*Workflow
	TransactionWorkflowMap map[string]string
	Transactions           map[string]*Transaction
}
