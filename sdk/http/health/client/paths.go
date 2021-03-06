// Code generated by goa v3.0.6, DO NOT EDIT.
//
// HTTP request path constructors for the Health service.
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

// CheckHealthPath returns the URL path to the Health service check HTTP endpoint.
func CheckHealthPath() string {
	return "/health-check"
}

// ReportHealthPath returns the URL path to the Health service report HTTP endpoint.
func ReportHealthPath() string {
	return "/health-report"
}
