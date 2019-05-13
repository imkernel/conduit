package api

// CreateApplicationRequest is the request payload for creation of an application in conduit
type CreateApplicationRequest struct {
	ApplicationID   string `json:"application_id"`
	ApplicationName string `json:"application_name"`
	OrganizationID  string `json:"organization_id"`
	Security        struct {
		APIKey string `json:"api_key"`
	} `json:"security"`
}
