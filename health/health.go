package health

import (
	"github.com/gin-gonic/gin"
	"github.com/rhinoman/couchdb-go"
	"net/http"
	"os"
)

func CreateHealthService(router *gin.Engine, conn *couchdb.Connection) {
	health := router.Group("api/v1/health")
	{
		health.GET("", func(c *gin.Context) {
			if err := conn.Ping(); err == nil {
				host,_ := os.Hostname()
				c.JSON(http.StatusOK, host)
			} else {
				c.JSON(http.StatusInternalServerError, err.Error())
			}
		})
	}
}
