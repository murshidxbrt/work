package middileware

import(
	token "github.com/murshidxbrt/ecommerce-yt/tokens"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
	}
} 