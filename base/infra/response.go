package infra

import (
	"encoding/json"
	"github.com/kataras/iris/v12/hero"
	"going"
	"strconv"
)

// JSONResponse.
type JSONResponse struct {
	Code             int
	Error            error
	Object           interface{}
	DisableLogOutput bool
}

// Dispatch
func (jrep JSONResponse) Dispatch(ctx going.Context) {
	contentType := "application/json"
	var content []byte

	var body struct {
		Code  int         `json:"code"`
		Error string      `json:"error"`
		Data  interface{} `json:"data,omitempty"`
	}
	body.Data = jrep.Object
	body.Code = jrep.Code

	if jrep.Error != nil {
		body.Error = jrep.Error.Error()
	}
	if jrep.Error != nil && body.Code == 0 {
		body.Code = 400
	}

	if content, jrep.Error = json.Marshal(body); jrep.Error != nil {
		content = []byte(jrep.Error.Error())
	}

	ctx.Values().Set("code", strconv.Itoa(body.Code))
	if !jrep.DisableLogOutput {
		ctx.Values().Set("response", string(content))
	}

	hero.DispatchCommon(ctx, 200, contentType, content, nil, nil, true)
}
