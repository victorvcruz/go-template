package main

import (
	"github.com/fasthttp/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go-template/cmd/api/handlers"
	"go-template/cmd/docs"
	"go-template/internal/config"
	"go-template/internal/platform"
	"go-template/internal/user"
	"log"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalln("failed to load configuration")
	}

	db, err := platform.PostgresConnect(cfg)
	if err != nil {
		log.Fatalln("failed to load configuration")
	}

	err = platform.RunMigrations(db)
	if err != nil {
		log.Fatalln("failed to run migrations")
	}

	userRepository := user.NewRepository(db)

	userService := user.NewService(userRepository)

	userHandler := handlers.NewUser(userService)

	docs.SwaggerInfo.Title = "Fasthttp Swagger"
	docs.SwaggerInfo.Description = "Fasthttp Swagger"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = ""

	// routes
	r := router.New()
	r.GET("/{filepath:*}", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
	r.GET("/user/{id}", userHandler.GetUser)
	r.POST("/user", userHandler.InsertUser)

	log.Println("server started on port 8080")
	if err := fasthttp.ListenAndServe(":8080", r.Handler); err != nil {
		log.Fatalln("failed to start server on port 8080")
	}
}
