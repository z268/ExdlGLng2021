package main

import (
	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/app"
)

// @title Book order service API
// @version 0.1
// @BasePath /api/v1/
// @query.collection.format multi
func main() {
	app.RunOrderService("configs/order.yml")
}
