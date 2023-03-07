package api

import (
	"log"
	"net/http"
	"time"

	"github.com/baribari2/pulp-calculator/common/types"
	api "github.com/baribari2/pulp-calculator/common/types/api"
	"github.com/baribari2/pulp-calculator/simulator"
	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
)

func SimulateTree(cfg *types.Config, line *charts.Line) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := api.NewSimulateThreadRequest(0, 0, 0)

		err := c.BindJSON(req)
		if err != nil {
			log.Printf("failed to bind json: %v", err.Error())

			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			return
		}

		_, _, _, err = simulator.SimulateThread(cfg, line, 0, time.Duration(req.Tick), time.Duration(req.EndTime), req.Frequency)
		if err != nil {
			log.Printf("failed to simulate thread: %v", err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
			return
		}

		// c.JSON(http.StatusOK, api.NewSimulateThreadResponse(tree.Root, tree.Timestamps, tree.LastScore, tree.InactiveCount, tree.Nodes))
	}
}

func GetTree() {
	// return func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// }
}
