package main

import (
	"GO-Basic/initializers"
	"GO-Basic/migrate"
	"GO-Basic/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDb()
	migrate.Migrate()
}

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(); err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080

}
