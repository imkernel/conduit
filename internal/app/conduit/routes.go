package conduit

import (
	"github.com/gorilla/mux"
	"net/http"
)

func createRoutes(r *mux.Router) {
	createApplicationRoutes(r)
	createWorkflowRoutes(r)
}

func createApplicationRoutes(r *mux.Router) {
	r.Path("/application").Methods(http.MethodPost).HandlerFunc(createApplicationHandler)
	r.Path("/application/{application_id}").Methods(http.MethodGet).HandlerFunc(getApplicationHandler)
	r.Path("/application/{application_id}").Methods(http.MethodPut).HandlerFunc(updateApplicationHandler)
	r.Path("/application/{application_id}").Methods(http.MethodDelete).HandlerFunc(deleteApplicationHandler)
}

func createWorkflowRoutes(r *mux.Router) {
	r.Path("/application/{application_id}/workflow").Methods(http.MethodPost).HandlerFunc(createWorkflowHandler)
	r.Path("/application/{application_id}/workflow/{workflow_id}").Methods(http.MethodGet).HandlerFunc(getWorkflowHandler)
	r.Path("/application/{application_id}/workflow/{workflow_id}").Methods(http.MethodPut).HandlerFunc(updateWorkflowHandler)
	r.Path("/application/{application_id}/workflow/{workflow_id}").Methods(http.MethodDelete).HandlerFunc(deleteWorkflowHandler)
}
