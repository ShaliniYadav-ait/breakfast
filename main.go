package main

import "net/http"
import "breakfast/middleware"

func main(){
	http.ListenAndServe(":9000", middleware.Handle())
}