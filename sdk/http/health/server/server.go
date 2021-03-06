// Code generated by goa v3.0.6, DO NOT EDIT.
//
// Health HTTP server
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package server

import (
	"context"
	"net/http"
	"regexp"

	health "github.com/rightscale/policy_sdk/sdk/health"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the Health service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Check  http.Handler
	Report http.Handler
	CORS   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the Health service endpoints.
func New(
	e *health.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Check", "GET", "/health-check"},
			{"Report", "GET", "/health-report"},
			{"CORS", "OPTIONS", "/health-check"},
			{"CORS", "OPTIONS", "/health-report"},
		},
		Check:  NewCheckHandler(e.Check, mux, dec, enc, eh),
		Report: NewReportHandler(e.Report, mux, dec, enc, eh),
		CORS:   NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Health" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Check = m(s.Check)
	s.Report = m(s.Report)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the Health endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCheckHandler(mux, h.Check)
	MountReportHandler(mux, h.Report)
	MountCORSHandler(mux, h.CORS)
}

// MountCheckHandler configures the mux to serve the "Health" service "check"
// endpoint.
func MountCheckHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleHealthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/health-check", f)
}

// NewCheckHandler creates a HTTP handler which loads the HTTP request and
// calls the "Health" service "check" endpoint.
func NewCheckHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeCheckResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "check")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Health")
		var err error

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountReportHandler configures the mux to serve the "Health" service "report"
// endpoint.
func MountReportHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleHealthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/health-report", f)
}

// NewReportHandler creates a HTTP handler which loads the HTTP request and
// calls the "Health" service "report" endpoint.
func NewReportHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeReportResponse(enc)
		encodeError    = goahttp.ErrorEncoder(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "report")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Health")
		var err error

		res, err := endpoint(ctx, nil)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service Health.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleHealthOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/health-check", f)
	mux.Handle("OPTIONS", "/health-report", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleHealthOrigin applies the CORS response headers corresponding to the
// origin for the service Health.
func handleHealthOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile("[.]rightscale[.]com$")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "ETag, Location")
			w.Header().Set("Access-Control-Max-Age", "900")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, HEAD, POST, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Api-Version, Authorization, Client-Name, Content-Type, Csrf-Token, Prefer, If-Modified-Since, If-None-Match, If-Unmodified-Since, If-Match, X-Api-Version, X-Csrf-Token, X-Requested-With")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
