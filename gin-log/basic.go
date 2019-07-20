package gin_log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	router := gin.New()
	file, _ := os.Create("access.log")
	c := gin.LoggerConfig{
		Output:    file,
		SkipPaths: []string{"/test"},
		Formatter: func(params gin.LogFormatterParams) string {
			//return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			//return fmt.Sprintf("%s - [%s] \"%s %s %s %d %#vms \"%s\" %s\"\n",
			//	params.ClientIP,
			//	params.TimeStamp.Format(time.RFC1123),
			//	params.Method,
			//	params.Path,
			//	params.Request.Proto,
			//	params.StatusCode,
			//	params.Latency/1e6,
			//	params.Request.UserAgent(),
			//	params.ErrorMessage,
			//)
			return fmt.Sprintf("[GIN] %v |%3d| %#13v ms | %15s | %-7s %s\n%s",
				params.TimeStamp.Format("2006/01/02 - 15:04:05"),
				params.StatusCode,
				params.Latency/1e6,
				params.ClientIP,
				params.Method,
				params.Path,
				params.ErrorMessage,
			)
		},
	}
	router.Use(gin.LoggerWithConfig(c))
	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		time.Sleep(1200000000)
		c.String(200, "pong")
	})
	router.GET("/test", func(c *gin.Context) {
		c.String(200, "test")
	})
	fmt.Printf("%#v",1e6)
	_ = router.Run(":8080")
}
