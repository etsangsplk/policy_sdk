// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// Approval HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	approval "github.com/rightscale/right_pt/sdk/approval"
	approvalviews "github.com/rightscale/right_pt/sdk/approval/views"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// BuildShowRequest instantiates a HTTP request object with method and path set
// to call the "Approval" service "show" endpoint
func (c *Client) BuildShowRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		projectID         uint
		approvalRequestID string
	)
	{
		p, ok := v.(*approval.ShowPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Approval", "show", "*approval.ShowPayload", v)
		}
		projectID = p.ProjectID
		approvalRequestID = p.ApprovalRequestID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ShowApprovalPath(projectID, approvalRequestID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Approval", "show", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeShowRequest returns an encoder for requests sent to the Approval show
// server.
func EncodeShowRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*approval.ShowPayload)
		if !ok {
			return goahttp.ErrInvalidType("Approval", "show", "*approval.ShowPayload", v)
		}
		req.Header.Set("Api-Version", p.APIVersion)
		if p.Token != nil {
			if !strings.Contains(*p.Token, " ") {
				req.Header.Set("Authorization", "Bearer "+*p.Token)
			} else {
				req.Header.Set("Authorization", *p.Token)
			}
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeShowResponse returns a decoder for responses returned by the Approval
// show endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeShowResponse may return the following errors:
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeShowResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ShowResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "show", err)
			}
			p := NewShowApprovalRequestOK(&body)
			view := resp.Header.Get("goa-view")
			vres := &approvalviews.ApprovalRequest{p, view}
			if err = vres.Validate(); err != nil {
				return nil, goahttp.ErrValidationError("Approval", "show", err)
			}
			res := approval.NewApprovalRequest(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body ShowNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "show", err)
			}
			return nil, NewShowNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ShowUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "show", err)
			}
			return nil, NewShowUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body ShowForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "show", err)
			}
			return nil, NewShowForbidden(&body)
		case http.StatusBadRequest:
			var (
				body ShowBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "show", err)
			}
			return nil, NewShowBadRequest(&body)
		case http.StatusBadGateway:
			var (
				body ShowBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "show", err)
			}
			return nil, NewShowBadGateway(&body)
		case http.StatusInternalServerError:
			var (
				body ShowInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "show", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "show", err)
			}
			return nil, NewShowInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Approval", "show", resp.StatusCode, string(body))
		}
	}
}

// BuildIndexRequest instantiates a HTTP request object with method and path
// set to call the "Approval" service "index" endpoint
func (c *Client) BuildIndexRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		projectID uint
	)
	{
		p, ok := v.(*approval.IndexPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Approval", "index", "*approval.IndexPayload", v)
		}
		projectID = p.ProjectID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: IndexApprovalPath(projectID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Approval", "index", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeIndexRequest returns an encoder for requests sent to the Approval
// index server.
func EncodeIndexRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*approval.IndexPayload)
		if !ok {
			return goahttp.ErrInvalidType("Approval", "index", "*approval.IndexPayload", v)
		}
		req.Header.Set("Api-Version", p.APIVersion)
		if p.Etag != nil {
			req.Header.Set("If-None-Match", *p.Etag)
		}
		if p.Token != nil {
			if !strings.Contains(*p.Token, " ") {
				req.Header.Set("Authorization", "Bearer "+*p.Token)
			} else {
				req.Header.Set("Authorization", *p.Token)
			}
		}
		values := req.URL.Query()
		if p.View != nil {
			values.Add("view", *p.View)
		}
		for _, value := range p.ID {
			values.Add("id", value)
		}
		if p.SubjectKind != nil {
			values.Add("subject_kind", *p.SubjectKind)
		}
		if p.SubjectHref != nil {
			values.Add("subject_href", *p.SubjectHref)
		}
		for _, value := range p.Status {
			values.Add("status", value)
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeIndexResponse returns a decoder for responses returned by the Approval
// index endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeIndexResponse may return the following errors:
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeIndexResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNotModified:
			var (
				etag *string
				err  error
			)
			etagRaw := resp.Header.Get("ETag")
			if etagRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("ETag", "header"))
			}
			etag = &etagRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			p := NewIndexApprovalRequestListNotModified(etag)
			view := resp.Header.Get("goa-view")
			vres := &approvalviews.ApprovalRequestList{p, view}
			res := approval.NewApprovalRequestList(vres)
			tmp := "true"
			res.NotModified = &tmp
			return res, nil
		case http.StatusOK:
			var (
				body IndexOKResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "index", err)
			}
			var (
				etag *string
			)
			etagRaw := resp.Header.Get("ETag")
			if etagRaw == "" {
				err = goa.MergeErrors(err, goa.MissingFieldError("ETag", "header"))
			}
			etag = &etagRaw
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			p := NewIndexApprovalRequestListOK(&body, etag)
			view := resp.Header.Get("goa-view")
			vres := &approvalviews.ApprovalRequestList{p, view}
			if err = vres.Validate(); err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			res := approval.NewApprovalRequestList(vres)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body IndexUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "index", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			return nil, NewIndexUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body IndexForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "index", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			return nil, NewIndexForbidden(&body)
		case http.StatusBadRequest:
			var (
				body IndexBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "index", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			return nil, NewIndexBadRequest(&body)
		case http.StatusBadGateway:
			var (
				body IndexBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "index", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			return nil, NewIndexBadGateway(&body)
		case http.StatusInternalServerError:
			var (
				body IndexInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "index", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "index", err)
			}
			return nil, NewIndexInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Approval", "index", resp.StatusCode, string(body))
		}
	}
}

// BuildApproveRequest instantiates a HTTP request object with method and path
// set to call the "Approval" service "approve" endpoint
func (c *Client) BuildApproveRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		projectID         uint
		approvalRequestID string
	)
	{
		p, ok := v.(*approval.ApprovePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Approval", "approve", "*approval.ApprovePayload", v)
		}
		projectID = p.ProjectID
		approvalRequestID = p.ApprovalRequestID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ApproveApprovalPath(projectID, approvalRequestID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Approval", "approve", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeApproveRequest returns an encoder for requests sent to the Approval
// approve server.
func EncodeApproveRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*approval.ApprovePayload)
		if !ok {
			return goahttp.ErrInvalidType("Approval", "approve", "*approval.ApprovePayload", v)
		}
		req.Header.Set("Api-Version", p.APIVersion)
		if p.Token != nil {
			if !strings.Contains(*p.Token, " ") {
				req.Header.Set("Authorization", "Bearer "+*p.Token)
			} else {
				req.Header.Set("Authorization", *p.Token)
			}
		}
		body := NewApproveRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Approval", "approve", err)
		}
		return nil
	}
}

// DecodeApproveResponse returns a decoder for responses returned by the
// Approval approve endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeApproveResponse may return the following errors:
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeApproveResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusAccepted:
			return nil, nil
		case http.StatusNotFound:
			var (
				body ApproveNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "approve", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "approve", err)
			}
			return nil, NewApproveNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body ApproveUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "approve", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "approve", err)
			}
			return nil, NewApproveUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body ApproveForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "approve", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "approve", err)
			}
			return nil, NewApproveForbidden(&body)
		case http.StatusBadRequest:
			var (
				body ApproveBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "approve", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "approve", err)
			}
			return nil, NewApproveBadRequest(&body)
		case http.StatusBadGateway:
			var (
				body ApproveBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "approve", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "approve", err)
			}
			return nil, NewApproveBadGateway(&body)
		case http.StatusInternalServerError:
			var (
				body ApproveInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "approve", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "approve", err)
			}
			return nil, NewApproveInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Approval", "approve", resp.StatusCode, string(body))
		}
	}
}

// BuildDenyRequest instantiates a HTTP request object with method and path set
// to call the "Approval" service "deny" endpoint
func (c *Client) BuildDenyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		projectID         uint
		approvalRequestID string
	)
	{
		p, ok := v.(*approval.DenyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("Approval", "deny", "*approval.DenyPayload", v)
		}
		projectID = p.ProjectID
		approvalRequestID = p.ApprovalRequestID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DenyApprovalPath(projectID, approvalRequestID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("Approval", "deny", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDenyRequest returns an encoder for requests sent to the Approval deny
// server.
func EncodeDenyRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*approval.DenyPayload)
		if !ok {
			return goahttp.ErrInvalidType("Approval", "deny", "*approval.DenyPayload", v)
		}
		req.Header.Set("Api-Version", p.APIVersion)
		if p.Token != nil {
			if !strings.Contains(*p.Token, " ") {
				req.Header.Set("Authorization", "Bearer "+*p.Token)
			} else {
				req.Header.Set("Authorization", *p.Token)
			}
		}
		body := NewDenyRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("Approval", "deny", err)
		}
		return nil
	}
}

// DecodeDenyResponse returns a decoder for responses returned by the Approval
// deny endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDenyResponse may return the following errors:
//	- "not_found" (type *goa.ServiceError): http.StatusNotFound
//	- "unauthorized" (type *goa.ServiceError): http.StatusUnauthorized
//	- "forbidden" (type *goa.ServiceError): http.StatusForbidden
//	- "bad_request" (type *goa.ServiceError): http.StatusBadRequest
//	- "bad_gateway" (type *goa.ServiceError): http.StatusBadGateway
//	- "internal_error" (type *goa.ServiceError): http.StatusInternalServerError
//	- error: internal error
func DecodeDenyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusAccepted:
			return nil, nil
		case http.StatusNotFound:
			var (
				body DenyNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "deny", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "deny", err)
			}
			return nil, NewDenyNotFound(&body)
		case http.StatusUnauthorized:
			var (
				body DenyUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "deny", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "deny", err)
			}
			return nil, NewDenyUnauthorized(&body)
		case http.StatusForbidden:
			var (
				body DenyForbiddenResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "deny", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "deny", err)
			}
			return nil, NewDenyForbidden(&body)
		case http.StatusBadRequest:
			var (
				body DenyBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "deny", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "deny", err)
			}
			return nil, NewDenyBadRequest(&body)
		case http.StatusBadGateway:
			var (
				body DenyBadGatewayResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "deny", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "deny", err)
			}
			return nil, NewDenyBadGateway(&body)
		case http.StatusInternalServerError:
			var (
				body DenyInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("Approval", "deny", err)
			}
			err = body.Validate()
			if err != nil {
				return nil, goahttp.ErrValidationError("Approval", "deny", err)
			}
			return nil, NewDenyInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("Approval", "deny", resp.StatusCode, string(body))
		}
	}
}

// unmarshalApprovalSubject2ToApprovalSubject2 builds a value of type
// *approvalviews.ApprovalSubject2 from a value of type *ApprovalSubject2.
func unmarshalApprovalSubject2ToApprovalSubject2(v *ApprovalSubject2) *approvalviews.ApprovalSubject2 {
	res := &approvalviews.ApprovalSubject2{
		Kind: v.Kind,
		Href: v.Href,
	}

	return res
}

// unmarshalRegexpResponseBodyToRegexpView builds a value of type
// *approvalviews.RegexpView from a value of type *RegexpResponseBody.
func unmarshalRegexpResponseBodyToRegexpView(v *RegexpResponseBody) *approvalviews.RegexpView {
	if v == nil {
		return nil
	}
	res := &approvalviews.RegexpView{
		Pattern: v.Pattern,
		Options: v.Options,
	}

	return res
}

// unmarshalUserResponseBodyToUserView builds a value of type
// *approvalviews.UserView from a value of type *UserResponseBody.
func unmarshalUserResponseBodyToUserView(v *UserResponseBody) *approvalviews.UserView {
	if v == nil {
		return nil
	}
	res := &approvalviews.UserView{
		ID:    v.ID,
		Email: v.Email,
		Name:  v.Name,
	}

	return res
}
