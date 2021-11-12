package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gasser707/go-gql-server/helpers"
	"github.com/gin-gonic/gin"
)

type CookieAccess struct {
    Writer     http.ResponseWriter
    encodedCookie   string
}
// method to write cookie
func (ca *CookieAccess) SetToken(at string, rt string) {

    value := map[string]string{
		"access_token": at,
        "refresh_token": rt,
	}

	if encoded, err := helpers.SecureCookieManager.Encode("cookie-name", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: encoded,
			Path:  "/*",
			// Secure: true,
			HttpOnly: true,
            Expires: time.Now().Add(time.Hour*24*7),
		}
        http.SetCookie(ca.Writer,cookie)
        ca.encodedCookie = encoded
	} else{
        println(err.Error())
    }
    
}


func setValInCtx(ctx *gin.Context, val interface{}) {
    newCtx := context.WithValue(ctx.Request.Context(), "cookie-name", val)
    ctx.Request = ctx.Request.WithContext(newCtx)
}

func GetCookieAccess(ctx context.Context) (*CookieAccess, error) {

    ca, ok :=  ctx.Value("cookie-name").(*CookieAccess)
    if(!ok){
        return nil, errors.New("cookie not found")
    }
    return ca, nil
}

func Middleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ec := ""
        ca,err := ctx.Request.Cookie("cookie-name")
        if(err==nil){
            ec = ca.Value
        }
        cookieA := CookieAccess{
            Writer: ctx.Writer,
            encodedCookie: ec,
        }

        // &cookieA is a pointer so any changes in future is changing cookieA is context
        setValInCtx(ctx, &cookieA)

       // calling the actual resolver
        ctx.Next()
       // here will execute after resolver and all other middlewares was called
       // so &cookieA is safe from garbage collector
    }
}