package server

import (
	"backend/internal/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type logWriter struct {
	gin.ResponseWriter
	respBody *bytes.Buffer
}

func (l logWriter) Write(data []byte) (int, error) {
	l.respBody.Write(data)
	return l.ResponseWriter.Write(data)
}

func (s *Server) Logger(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(HttpError(http.StatusInternalServerError, err.Error()))
		return
	}
	logger.InfoFuncf("Request: ClientIP:%s-Path:%s-Header:%+v-Body:%s",
		ctx.ClientIP(),
		ctx.Request.URL.Path,
		ctx.Request.Header,
		string(reqBody))
	// Restore the request body
	ctx.Request.Body = io.NopCloser(bytes.NewReader(reqBody))
	writer := &logWriter{
		ResponseWriter: ctx.Writer,
		respBody:       bytes.NewBuffer(nil),
	}
	ctx.Writer = writer
	ctx.Next()
	logger.InfoFuncf("Response: ClientIP:%s-Path:%s-Header:%+v-Status:%d-Body:%s",
		ctx.ClientIP(),
		ctx.Request.URL.Path,
		ctx.Writer.Header(),
		ctx.Writer.Status(),
		writer.respBody.String())
}
