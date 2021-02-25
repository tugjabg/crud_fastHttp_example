package service_user

import (
	"fmt"
	"service/service_user"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"log"
)

func Index(ctx *fasthttp.RequestCtx) {
    fmt.Fprint(ctx, "Welcome!\n")
}

func main() {
	r := router.New()
	r.GET("/", Index)
	r.GET("/user/{id}", service_user.GetUser)
	r.GET("/users", service_user.GetUsers)
	r.POST("/user", service_user.CreateStudent)
	r.DELETE("/user/{id}", service_user.DeleteStudent)
	r.PUT("/user/{id}", service_user.UpdateStudent)
	log.Fatal(fasthttp.ListenAndServe(":1342", r.Handler))
}
