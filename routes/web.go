package routes

import (
	"main/app/controllers"
	"net/http"
)

func WebRoutes() {

	// Manejador para servir archivos estáticos
	fileServer := http.FileServer(http.Dir("public"))
	// Ruta para acceder a los archivos estáticos
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))

	// Configurar rutas
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/show", controllers.Show)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
}
