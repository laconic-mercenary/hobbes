package function

import (
	"net/http"

	handler "github.com/openfaas/templates-sdk/go-http"

	httpc "handler/function/internal/httpc"
)

// Handle a function invocation
func Handle(request handler.Request) (handler.Response, error) {
	var response handler.Response
	responseText, status, headers, err := httpc.HandleRequest(request.Body, request.Method, request.QueryString, request.Header)
	if err == nil {
		response = handler.Response{
			Body:       []byte(responseText),
			StatusCode: status,
			Header:     headers,
		}
	} else {
		response = handler.Response{
			Body:       []byte("error"),
			StatusCode: http.StatusInternalServerError,
			Header:     headers,
		}
	}
	return response, nil
}
