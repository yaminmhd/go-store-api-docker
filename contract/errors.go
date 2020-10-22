package contract

import (
	"net/http"

	"github.com/yaminmhd/go-hardware-store/constant"
)

type errorDetails struct {
	message string
	Status  int
}

var ErrorObjects = map[string]errorDetails{
	constant.ErrorInvalidRequest: {
		message: "invalid request",
		Status:  http.StatusBadRequest,
	},
	constant.ErrorInternalServerError: {
		message: "something went wrong",
		Status:  http.StatusInternalServerError,
	},
	constant.ErrorNotFound: {
		message: "not found",
		Status:  http.StatusNotFound,
	},
}
