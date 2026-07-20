package controllers

import (
	"net/http"
	"github.com/car-rental-template/api/internal/services"
	"github.com/gin-gonic/gin"
)

type HealthController struct { service *services.HealthService }
func NewHealthController(service *services.HealthService) *HealthController { return &HealthController{service: service} }
func (c *HealthController) GetStatus(ctx *gin.Context) { status := c.service.Status(ctx.Request.Context()); code := http.StatusOK; if status.Status != "ok" { code = http.StatusServiceUnavailable }; ctx.JSON(code, status) }
