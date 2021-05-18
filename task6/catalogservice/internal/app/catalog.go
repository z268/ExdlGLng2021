package app

import (
	"fmt"
	"log"
	"net"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"github.com/golang-migrate/migrate/v4"
	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/config"
	"github.com/z268/ExdlGLng2021/task6/catalogservice/internal/database"
	pb "github.com/z268/ExdlGLng2021/task6/catalogservice/internal/delivery/grpc"
	delivery "github.com/z268/ExdlGLng2021/task6/catalogservice/internal/delivery/http"
	repo_mysql "github.com/z268/ExdlGLng2021/task6/catalogservice/internal/repository/mysql"
)


func RunCatalogService(configPath string, serverMode string) {
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
	_ = m.Steps(2)

	// Init repositories
	repo := repo_mysql.NewCatalogRepository(db)

	switch serverMode {
	case "grpc":
		lis, _ := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Grpc.Port))
		s := grpc.NewServer()
		pb.RegisterBookServiceServer(s, pb.NewBookServiceServer(repo))

		log.Println("Running gRPC server at", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	default:
		hander := delivery.NewCatalogHandler(repo, cfg.Api.PageSize)
		router := delivery.InitCatalogRouter(hander)
		http.Handle("/", router)

		listenAddr := fmt.Sprintf("%v:%v", cfg.Http.Host, cfg.Http.Port)
		fmt.Println("Running HTTP server at", listenAddr)
		log.Fatal(http.ListenAndServe(listenAddr, nil))
	}

}
