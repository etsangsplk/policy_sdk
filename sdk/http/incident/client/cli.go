// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Incident HTTP client CLI support package
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	incident "github.com/rightscale/right_pt/sdk/incident"
	goa "goa.design/goa"
)

// BuildShowPayload builds the payload for the Incident show endpoint from CLI
// flags.
func BuildShowPayload(incidentShowProjectID string, incidentShowIncidentID string, incidentShowView string, incidentShowAPIVersion string, incidentShowEtag string, incidentShowToken string) (*incident.ShowPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(incidentShowProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var incidentID string
	{
		incidentID = incidentShowIncidentID
	}
	var view *string
	{
		if incidentShowView != "" {
			view = &incidentShowView
		}
	}
	var apiVersion string
	{
		apiVersion = incidentShowAPIVersion
	}
	var etag *string
	{
		if incidentShowEtag != "" {
			etag = &incidentShowEtag
		}
	}
	var token *string
	{
		if incidentShowToken != "" {
			token = &incidentShowToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &incident.ShowPayload{
		ProjectID:  projectID,
		IncidentID: incidentID,
		View:       view,
		APIVersion: apiVersion,
		Etag:       etag,
		Token:      token,
	}
	return payload, nil
}

// BuildIndexPayload builds the payload for the Incident index endpoint from
// CLI flags.
func BuildIndexPayload(incidentIndexProjectID string, incidentIndexView string, incidentIndexState string, incidentIndexAppliedPolicyID string, incidentIndexAPIVersion string, incidentIndexEtag string, incidentIndexToken string) (*incident.IndexPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(incidentIndexProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var view *string
	{
		if incidentIndexView != "" {
			view = &incidentIndexView
		}
	}
	var state []string
	{
		if incidentIndexState != "" {
			err = json.Unmarshal([]byte(incidentIndexState), &state)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for state, example of valid JSON:\n%s", "'[\n      \"triggered\"\n   ]'")
			}
			for _, e := range state {
				if !(e == "triggered" || e == "resolved" || e == "terminated") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("state[*]", e, []interface{}{"triggered", "resolved", "terminated"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var appliedPolicyID *string
	{
		if incidentIndexAppliedPolicyID != "" {
			appliedPolicyID = &incidentIndexAppliedPolicyID
		}
	}
	var apiVersion string
	{
		apiVersion = incidentIndexAPIVersion
	}
	var etag *string
	{
		if incidentIndexEtag != "" {
			etag = &incidentIndexEtag
		}
	}
	var token *string
	{
		if incidentIndexToken != "" {
			token = &incidentIndexToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &incident.IndexPayload{
		ProjectID:       projectID,
		View:            view,
		State:           state,
		AppliedPolicyID: appliedPolicyID,
		APIVersion:      apiVersion,
		Etag:            etag,
		Token:           token,
	}
	return payload, nil
}

// BuildResolvePayload builds the payload for the Incident resolve endpoint
// from CLI flags.
func BuildResolvePayload(incidentResolveProjectID string, incidentResolveIncidentID string, incidentResolveAPIVersion string, incidentResolveEtag string, incidentResolveToken string) (*incident.ResolvePayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(incidentResolveProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var incidentID string
	{
		incidentID = incidentResolveIncidentID
	}
	var apiVersion string
	{
		apiVersion = incidentResolveAPIVersion
	}
	var etag *string
	{
		if incidentResolveEtag != "" {
			etag = &incidentResolveEtag
		}
	}
	var token *string
	{
		if incidentResolveToken != "" {
			token = &incidentResolveToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &incident.ResolvePayload{
		ProjectID:  projectID,
		IncidentID: incidentID,
		APIVersion: apiVersion,
		Etag:       etag,
		Token:      token,
	}
	return payload, nil
}

// BuildIndexEscalationsPayload builds the payload for the Incident
// index_escalations endpoint from CLI flags.
func BuildIndexEscalationsPayload(incidentIndexEscalationsProjectID string, incidentIndexEscalationsIncidentID string, incidentIndexEscalationsAPIVersion string, incidentIndexEscalationsToken string) (*incident.IndexEscalationsPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(incidentIndexEscalationsProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var incidentID string
	{
		incidentID = incidentIndexEscalationsIncidentID
	}
	var apiVersion string
	{
		apiVersion = incidentIndexEscalationsAPIVersion
	}
	var token *string
	{
		if incidentIndexEscalationsToken != "" {
			token = &incidentIndexEscalationsToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &incident.IndexEscalationsPayload{
		ProjectID:  projectID,
		IncidentID: incidentID,
		APIVersion: apiVersion,
		Token:      token,
	}
	return payload, nil
}

// BuildIndexResolutionsPayload builds the payload for the Incident
// index_resolutions endpoint from CLI flags.
func BuildIndexResolutionsPayload(incidentIndexResolutionsProjectID string, incidentIndexResolutionsIncidentID string, incidentIndexResolutionsAPIVersion string, incidentIndexResolutionsToken string) (*incident.IndexResolutionsPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(incidentIndexResolutionsProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var incidentID string
	{
		incidentID = incidentIndexResolutionsIncidentID
	}
	var apiVersion string
	{
		apiVersion = incidentIndexResolutionsAPIVersion
	}
	var token *string
	{
		if incidentIndexResolutionsToken != "" {
			token = &incidentIndexResolutionsToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &incident.IndexResolutionsPayload{
		ProjectID:  projectID,
		IncidentID: incidentID,
		APIVersion: apiVersion,
		Token:      token,
	}
	return payload, nil
}
