package handler

import (
	"net/http"

	"github.com/elitecodegroovy/gnetwork/apps/micro/conf-center/admin-api/config"
	"github.com/elitecodegroovy/gnetwork/apps/micro/conf-center/admin-api/format"
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func UpdateConfig(c *gin.Context) {
	var req = struct {
		AppName       string `json:"appName"        binding:"required"`
		ClusterName   string `json:"clusterName"    binding:"required"`
		NamespaceName string `json:"namespaceName"  binding:"required"`
		Format        string `json:"format"         binding:"required"`
		Value         string `json:"value"          binding:"required"`
	}{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	if err := format.CheckFormat(req.Format, req.Value); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	err := config.UpdateConfig(req.AppName, req.ClusterName, req.NamespaceName, req.Value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func ListSupportedFormat(c *gin.Context) {
	c.JSON(http.StatusOK, format.SupportedFormat())
}
