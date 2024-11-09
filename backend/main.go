package main

import (
	"fmt"
	"williams/internal/service"
)

func main() {
	svc := service.NewService(8080)

	fmt.Println("Gathering data")
	svc.InitDatasets()

	fmt.Println("Serving")
	svc.ListenAndServe()
}
