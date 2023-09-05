package auth

import "log"
import "time"

import "github.com/gin-gonic/gin"
import jwt "github.com/appleboy/gin-jwt/v2"

import "crud"
import "models"

type LoginPasswordForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

const realm = "test zone"

var secret = []byte("secret key")

const identityKey = "name"
const emailKey = "email"
const isAdminKey = "is_admin"

var Middleware *jwt.GinJWTMiddleware

func init() {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       realm,
		Key:         secret,
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Name,
					emailKey:    v.Email,
					isAdminKey:  v.IsAdmin,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			is_admin := claims[isAdminKey].(bool)
			return &models.User{
				Name:    claims[identityKey].(string),
				Email:   claims[emailKey].(string),
				IsAdmin: &is_admin,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginForm LoginPasswordForm
			if err := c.ShouldBind(&loginForm); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			var user = crud.GetUserByName(loginForm.Username)
			if user == nil {
				return nil, jwt.ErrFailedAuthentication
			}

			if *user.Password != loginForm.Password {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization, query: token",
		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := middleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("MiddlewareInit() Error:" + errInit.Error())
	}

	Middleware = middleware
}
