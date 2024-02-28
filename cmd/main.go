package main

import (
	"bookstoreGo/internals/routes"
	"bookstoreGo/pkg"
	"log"
)

//dependency Injection (DI)

func main() {

	//inisialisasi db
	db, err := pkg.InitMysql()
	if err != nil {
		log.Fatal(err)
		//return
	}
	//inisialisasi router
	router := routes.InitRouter(db)

	//inisialisasi server
	server := pkg.InitServer(router)

	//jalankan servernya
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
