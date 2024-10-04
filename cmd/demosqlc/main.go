package main

import "github.com/hoangnguyen/demo-sqlc/internal/initialize"
import ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
import swaggerFiles "github.com/swaggo/files" // swagger embed files
import  _ "github.com/hoangnguyen/demo-sqlc/docs"

// @title           Swagger Demo SQLC API
// @version         1.0.0
// @description     This is a sample server celler server.
// @termsOfService  github.com/hoangnguyen/demosqlc

// @contact.name   Tips Go
// @contact.url    github.com/hoangnguyen/demosqlc
// @contact.email  hoangnguyendang777@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema http
func main() {
	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}