package openapi

const (
	REF_PREFIX = "#/components/schemas/"
)

var (
	METHODS_WITH_BODY         = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH"}
	STATUS_CODES_WITH_NO_BODY = []int{100, 101, 102, 103, 204, 304}
)
