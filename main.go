package main

import (
	"context"
	"fapi/database"
	"fapi/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

type Account struct {
	gorm.Model
	Code string
	Password uint
}

func main() {


	db:=database.InitDb()

migration := database.Migrations{
	DB: db,
	Models:[]interface{} {
		&Product{},
		&Account{},
		

	},
}

database.RunMigrations(migration)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routers.NewRouter(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
// kill (no param) default send syscall.SIGTERM

// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 1 seconds.
	 <-ctx.Done()

	
	log.Println("Server exiting")
}
