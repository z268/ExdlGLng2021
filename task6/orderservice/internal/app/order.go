package app

import (
	"log"
	"net/http"
	"fmt"

	"google.golang.org/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/config"
	"github.com/z268/ExdlGLng2021/task6/orderservice/internal/database"
	delivery "github.com/z268/ExdlGLng2021/task6/orderservice/internal/delivery/http"
	pb "github.com/z268/ExdlGLng2021/task6/orderservice/internal/delivery/grpc"
	repo_mysql "github.com/z268/ExdlGLng2021/task6/orderservice/internal/repository/mysql"
)

func RunOrderService(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// DB connection
	db, err := database.Init(cfg)
	if err != nil {
		log.Fatalf("No connection to database: %v", err)
	}
	defer database.Close(db)

	// Apply SQL migrations
	driver, _ := mysql_migrate.WithInstance(db.DB, &mysql_migrate.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://./db/migrations", "mysql", driver)
	if err != nil {
		log.Fatal("Migration init error: ", err)
	}
	_ = m.Steps(1)

	// GRPC
	grpcAddr := fmt.Sprintf("%v:%v", cfg.CatalogGrpc.Host, cfg.CatalogGrpc.Port)
	fmt.Printf("Connecting to %v...\n", grpcAddr)
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpc_client := pb.NewBookServiceClient(conn)

	// Services, Repos & API Handlers
	repo := repo_mysql.NewOrderRepository(db)
	router := delivery.InitOrderRouter(repo, grpc_client)

	listenAddr := fmt.Sprintf("%v:%v", cfg.Http.Host, cfg.Http.Port)
	fmt.Println("Running HTTP on address", listenAddr)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
