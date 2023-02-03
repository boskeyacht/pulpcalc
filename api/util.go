package api

import (
	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
)

func SetupRouter(cfg *types.Config, line *charts.Line) *gin.Engine {
	r := gin.Default()

	// r.GET("/tree", GetTree())
	r.GET("/simulate", SimulateTree(cfg, line))

	return r
}
