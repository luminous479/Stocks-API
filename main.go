package main

import (
	"fmt"
	"github.com/luminous479/Stocks-API/router"
)

func main() {
	r := router.Routers() 
	fmt.Println("Server is running on port 8080")
	if err := r.ListenAndServe(":8080"); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
