package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title GraphQL Go API
// @version 1.0
// @description API que muestra "Hi, I am Erick" usando GraphQL y Swagger
// @host localhost:8080
// @BasePath /

// @Summary Consulta de mensaje
// @Description Devuelve un mensaje simple "Hola soy Erick"
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Hola soy Erick"
// @Router /graphql [post]

func main() {
	// Definir esquema de GraphQL
	fields := graphql.Fields{
		"message": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "I am Erick. GraphQL with Go", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Error creando esquema GraphQL: %v", err)
	}

	// Configurar el manejador GraphQL
	graphQLHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Habilitar interfaz GraphiQL
	})

	// Configurar rutas con Gorilla Mux
	r := mux.NewRouter()

	// Ruta GraphQL
	r.Handle("/graphql", graphQLHandler)

	// Ruta Swagger - Apuntar correctamente a swagger.json
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler(httpSwagger.URL("http://localhost:8080/swagger/swagger.json")))

	// Iniciar el servidor
	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
