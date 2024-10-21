package main

import (
	"fmt"
	"database/sql"
	"log"
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
	"github.com/joho/godotenv"

	_ "github.com/microsoft/go-mssqldb"
)

var books =[]Book { 
	{Isbn: "9780441172719", Title: "Dune", Author: "Frank Herbert", ReleaseDate: "1987", Availability: false}, 
}

func getBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books);
}

func main() {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	e := echo.New()
	
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	server := os.Getenv("SERVER_NAME")
	database := os.Getenv("DATABASE_NAME")
	user := os.Getenv("USER_ID")

	log.Printf("Beep boop, connecting to database")
	log.Printf("Server:%s - database:%s", server, database)


	connStr := fmt.Sprintf("server=%s;database=%s;user id=%s;", server, database, user)

	db, err := sql.Open("mssql", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
    fmt.Println("Connected!")

	defer db.Close()

	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.GET("/Books", getBooks)

	e.Logger.Fatal(e.Start(":1323"))
	
}

type Book struct {
	Title  string `json:"title" query:"title"`
	Isbn string `json:"isbn" query:"isbn"`
	Author string `json:"author" query:"author"`
	ReleaseDate string `json:"releaseDate" query:"releaseDate"`
	Availability bool `json:"availability" query:"availability"`
}
