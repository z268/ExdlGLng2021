package main

import (
	"flag"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/app"
)

// @title Book catalog service API
// @version 0.1
// @BasePath /api/v1/
// @query.collection.format multi
func main() {
	serverMode := flag.String("mode", "http", "Server mode: http or grpc")
	flag.Parse()

	app.RunCatalogService("configs/catalog.yml", *serverMode)
}
