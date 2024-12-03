package routes

import (
	"net/http"
	"salao/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
