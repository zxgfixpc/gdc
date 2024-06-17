package ginsugar

import (
	"net/http"

	"_gdc_/lib/errors"
	"_gdc_/lib/logger"
	"_gdc_/lib/trace"

	"github.com/gin-gonic/gin"
)

type JsonResult struct {
	Errno   int64       `json:"errno"`
	ErrMsg  string      `json:"err_msg"`
	TraceID string      `json:"trace_id"`
	Data    interface{} `json:"data"`
}

func buildJsonResult(c *gin.Context, httpCode int, result interface{}, err error) {
	traceID := trace.GetTraceID(c)
	ret := &JsonResult{
		TraceID: traceID,
		Data:    result,
	}
	if err != nil {
		logger.Error(c.Request.Context(), err.Error())

		ret.Errno, ret.ErrMsg = errors.CodeMsg(err)
		if ret.Errno == 0 {
			ret.Errno = int64(httpCode)
		}
	}

	c.JSON(httpCode, ret)
}

func Success(c *gin.Context, result interface{}) {
	buildJsonResult(c, http.StatusOK, result, nil)
}

func InputError(c *gin.Context, err error) {
	buildJsonResult(c, http.StatusBadRequest, nil, err)
}

func Fail(c *gin.Context, result interface{}, err error) {
	buildJsonResult(c, http.StatusInternalServerError, result, err)
}
