// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Approval HTTP client CLI support package
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	approval "github.com/rightscale/right_pt/sdk/approval"
	goa "goa.design/goa"
)

// BuildShowPayload builds the payload for the Approval show endpoint from CLI
// flags.
func BuildShowPayload(approvalShowProjectID string, approvalShowApprovalRequestID string, approvalShowView string, approvalShowAPIVersion string, approvalShowToken string) (*approval.ShowPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(approvalShowProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var approvalRequestID string
	{
		approvalRequestID = approvalShowApprovalRequestID
	}
	var view *string
	{
		if approvalShowView != "" {
			view = &approvalShowView
		}
	}
	var apiVersion string
	{
		apiVersion = approvalShowAPIVersion
	}
	var token *string
	{
		if approvalShowToken != "" {
			token = &approvalShowToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &approval.ShowPayload{
		ProjectID:         projectID,
		ApprovalRequestID: approvalRequestID,
		View:              view,
		APIVersion:        apiVersion,
		Token:             token,
	}
	return payload, nil
}

// BuildIndexPayload builds the payload for the Approval index endpoint from
// CLI flags.
func BuildIndexPayload(approvalIndexProjectID string, approvalIndexView string, approvalIndexID string, approvalIndexSubjectKind string, approvalIndexSubjectHref string, approvalIndexStatus string, approvalIndexAPIVersion string, approvalIndexEtag string, approvalIndexToken string) (*approval.IndexPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(approvalIndexProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var view *string
	{
		if approvalIndexView != "" {
			view = &approvalIndexView
		}
	}
	var id []string
	{
		if approvalIndexID != "" {
			err = json.Unmarshal([]byte(approvalIndexID), &id)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for id, example of valid JSON:\n%s", "'[\n      \"5b36ad26d4c1990070df726a\",\n      \"5b36ad26d4c2000070df726a\"\n   ]'")
			}
		}
	}
	var subjectKind *string
	{
		if approvalIndexSubjectKind != "" {
			subjectKind = &approvalIndexSubjectKind
		}
	}
	var subjectHref *string
	{
		if approvalIndexSubjectHref != "" {
			subjectHref = &approvalIndexSubjectHref
		}
	}
	var status []string
	{
		if approvalIndexStatus != "" {
			err = json.Unmarshal([]byte(approvalIndexStatus), &status)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for status, example of valid JSON:\n%s", "'[\n      \"pending\"\n   ]'")
			}
			for _, e := range status {
				if !(e == "pending" || e == "approved" || e == "denied") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("status[*]", e, []interface{}{"pending", "approved", "denied"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var apiVersion string
	{
		apiVersion = approvalIndexAPIVersion
	}
	var etag *string
	{
		if approvalIndexEtag != "" {
			etag = &approvalIndexEtag
		}
	}
	var token *string
	{
		if approvalIndexToken != "" {
			token = &approvalIndexToken
		}
	}
	if err != nil {
		return nil, err
	}
	payload := &approval.IndexPayload{
		ProjectID:   projectID,
		View:        view,
		ID:          id,
		SubjectKind: subjectKind,
		SubjectHref: subjectHref,
		Status:      status,
		APIVersion:  apiVersion,
		Etag:        etag,
		Token:       token,
	}
	return payload, nil
}

// BuildApprovePayload builds the payload for the Approval approve endpoint
// from CLI flags.
func BuildApprovePayload(approvalApproveBody string, approvalApproveProjectID string, approvalApproveApprovalRequestID string, approvalApproveAPIVersion string, approvalApproveToken string) (*approval.ApprovePayload, error) {
	var err error
	var body ApproveRequestBody
	{
		err = json.Unmarshal([]byte(approvalApproveBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"options\": [\n         {\n            \"name\": \"excluded_tags\",\n            \"value\": [\n               \"env:name=staging\",\n               \"custom:save=true\"\n            ]\n         }\n      ]\n   }'")
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(approvalApproveProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var approvalRequestID string
	{
		approvalRequestID = approvalApproveApprovalRequestID
	}
	var apiVersion string
	{
		apiVersion = approvalApproveAPIVersion
	}
	var token *string
	{
		if approvalApproveToken != "" {
			token = &approvalApproveToken
		}
	}
	if err != nil {
		return nil, err
	}
	v := &approval.ApprovePayload{}
	if body.Options != nil {
		v.Options = make([]*approval.ConfigurationOptionCreateType, len(body.Options))
		for i, val := range body.Options {
			v.Options[i] = &approval.ConfigurationOptionCreateType{
				Name:  val.Name,
				Value: val.Value,
			}
		}
	}
	v.ProjectID = projectID
	v.ApprovalRequestID = approvalRequestID
	v.APIVersion = apiVersion
	v.Token = token
	return v, nil
}

// BuildDenyPayload builds the payload for the Approval deny endpoint from CLI
// flags.
func BuildDenyPayload(approvalDenyBody string, approvalDenyProjectID string, approvalDenyApprovalRequestID string, approvalDenyAPIVersion string, approvalDenyToken string) (*approval.DenyPayload, error) {
	var err error
	var body DenyRequestBody
	{
		err = json.Unmarshal([]byte(approvalDenyBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"comment\": \"The volumes need to be backed up.\"\n   }'")
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(approvalDenyProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			err = fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var approvalRequestID string
	{
		approvalRequestID = approvalDenyApprovalRequestID
	}
	var apiVersion string
	{
		apiVersion = approvalDenyAPIVersion
	}
	var token *string
	{
		if approvalDenyToken != "" {
			token = &approvalDenyToken
		}
	}
	if err != nil {
		return nil, err
	}
	v := &approval.DenyPayload{
		Comment: body.Comment,
	}
	v.ProjectID = projectID
	v.ApprovalRequestID = approvalRequestID
	v.APIVersion = apiVersion
	v.Token = token
	return v, nil
}