package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ISongwuts/go-crud-elasticsearch/internal/models"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type (
	IBookRepository interface {
		Create(client *elasticsearch.TypedClient, user *models.Book) (string, error)
		GetBooks(client *elasticsearch.TypedClient) ([]models.Book, error)
		GetByID(client *elasticsearch.TypedClient, id string) (models.Book, error)
		Update(client *elasticsearch.TypedClient, id string, modify map[string]string) (string, error)
		Delete(client *elasticsearch.TypedClient, id string) (string, error)
	}

	BookRepository struct {}
)

func NewBookRepository() IBookRepository {
	return &BookRepository{}
}

func (br *BookRepository) Create(client *elasticsearch.TypedClient, users *models.Book) (string, error) {
	return "", nil
}

func (br *BookRepository) GetBooks(client *elasticsearch.TypedClient) ([]models.Book, error) {
	books := make([]models.Book, 0)
	fmt.Println("in bookusecase")
	// Prepare the search request
	req := search.Request{
		Query: &types.Query{
			MatchAll: &types.MatchAllQuery{},
		},
	}

	// Perform the search request
	res, err := client.Search().Index("books").Request(&req).Do(context.Background())
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("error performing search request: %w", err)
	}

	// Parse the hits and append them to the books slice
	for _, hit := range res.Hits.Hits {
		var book models.Book
		if err := json.Unmarshal(hit.Source_, &book); err != nil {
			return nil, fmt.Errorf("error unmarshalling book: %w", err)
		}
		books = append(books, book)
	}
	fmt.Println(books)
	return books, nil
}

func (br *BookRepository) GetByID(client *elasticsearch.TypedClient, id string) (models.Book, error) {
	books := models.Book{}
	return books, nil
}

func (br *BookRepository) Update(client *elasticsearch.TypedClient, id string, modify map[string]string) (string, error){
	return "", nil
}

func (br *BookRepository) Delete(client *elasticsearch.TypedClient, id string) (string, error) {
	return "", nil
}