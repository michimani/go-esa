package gesa

var (
	_ interface{}

	ExportNewRequest         = newRequest
	ExportResolveEsaAPIError = resolveEsaAPIError

	ExportWrapErr        = wrapErr
	ExportWrapWithAPIErr = wrapWithAPIErr
)
