package main

import (
	"github.com/phuong/go-ecomerce/internal/routers"
)

func main() {
	r := routers.NewRouter()

	r.Run()
}
