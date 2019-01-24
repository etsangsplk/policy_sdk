// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the Incident service.
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"fmt"
)

// ShowIncidentPath returns the URL path to the Incident service show HTTP endpoint.
func ShowIncidentPath(projectID uint, incidentID string) string {
	return fmt.Sprintf("/api/governance/projects/%v/incidents/%v", projectID, incidentID)
}

// IndexIncidentPath returns the URL path to the Incident service index HTTP endpoint.
func IndexIncidentPath(projectID uint) string {
	return fmt.Sprintf("/api/governance/projects/%v/incidents", projectID)
}

// ResolveIncidentPath returns the URL path to the Incident service resolve HTTP endpoint.
func ResolveIncidentPath(projectID uint, incidentID string) string {
	return fmt.Sprintf("/api/governance/projects/%v/incidents/%v/resolve", projectID, incidentID)
}

// IndexEscalationsIncidentPath returns the URL path to the Incident service index_escalations HTTP endpoint.
func IndexEscalationsIncidentPath(projectID uint, incidentID string) string {
	return fmt.Sprintf("/api/governance/projects/%v/incidents/%v/escalations", projectID, incidentID)
}

// IndexResolutionsIncidentPath returns the URL path to the Incident service index_resolutions HTTP endpoint.
func IndexResolutionsIncidentPath(projectID uint, incidentID string) string {
	return fmt.Sprintf("/api/governance/projects/%v/incidents/%v/resolutions", projectID, incidentID)
}
