// Code generated by goa v3.0.6, DO NOT EDIT.
//
// Health service
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package health

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Service is the Health service interface.
type Service interface {
	// Health check
	Check(context.Context) (err error)
	// Health report
	Report(context.Context) (res map[string]string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "Health"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"check", "report"}

// MakeServiceUnavailable builds a goa.ServiceError from an error.
func MakeServiceUnavailable(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "service_unavailable",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
