package main

import (
	"crud_go/routes"
	"crud_go/services"
	"github.com/gin-gonic/gin"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	services.InitDatabase()

	r := gin.Default()
	routes.InitRoutes(r)

	// Jalankan server di port 8080
	log.Println("Server is running on http://localhost:8080")

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
