package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Book struct {
	Title        string `json:"title" query:"title"`
	Isbn         string `json:"isbn" query:"isbn"`
	Author       string `json:"author" query:"author"`
	ReleaseDate  string `json:"releaseDate" query:"releaseDate"`
	Availability bool   `json:"availability" query:"availability"`
}

type GoogleBookResponse struct {
	Items []GoogleBookItem `json:"items"`
}

type GoogleBookItem struct {
	ID         string      `json:"id"`
	VolumeInfo VolumeInfo  `json:"volumeInfo"`
}

type VolumeInfo struct {
	Title               string              `json:"title"`
	Authors             []string            `json:"authors"`
	PublishedDate       string              `json:"publishedDate"`
	IndustryIdentifiers []IndustryIdentifier `json:"industryIdentifiers"`
}

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

func getBookByID(c echo.Context) error {
	bookID := c.QueryParam("id") 

	
	if bookID == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "No book ID provided",
		})
	}

	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes/%s", bookID)

	resp, err := http.Get(url)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch book from Google Books API",
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Google Books API returned non-200 status code",
		})
	}

	var bookItem GoogleBookItem
	if err := json.NewDecoder(resp.Body).Decode(&bookItem); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to parse response from Google Books API",
		})
	}
	authors := "Unknown Author"
	if len(bookItem.VolumeInfo.Authors) > 0 {
		authors = bookItem.VolumeInfo.Authors[0]
	}

	isbn := "Unknown ISBN"
	for _, identifier := range bookItem.VolumeInfo.IndustryIdentifiers {
		if identifier.Type == "ISBN_13" {
			isbn = identifier.Identifier
			break
		}
	}

	book := Book{
		Title:       bookItem.VolumeInfo.Title,
		Isbn:        isbn,
		Author:      authors,
		ReleaseDate: bookItem.VolumeInfo.PublishedDate,
		Availability: true,
	}

	return c.JSON(http.StatusOK, book)
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/Books", getBookByID) 

	e.Logger.Fatal(e.Start(":1323"))
}