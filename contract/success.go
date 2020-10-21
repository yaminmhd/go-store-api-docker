package contract

import (
	"github.com/yaminmhd/go-hardware-store/constant"
	"net/http"
)

type successDetails struct{
	status int
}

var successObjects = map[string]successDetails{
	constant.SuccessOK: {
		status: http.StatusOK,
	},
	constant.SuccessNoContent: {
		status: http.StatusNoContent,
	},
	constant.SuccessCreated: {
		status: http.StatusCreated,
	},
}
