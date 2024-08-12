package dto

type Product struct {
	ID              int64   `json:"id"`
	Title           string  `json:"title"`
	Author          string  `json:"author"`
	Publicatio_year int64   `json:"publication_year"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	ISBN            string  `json:"isbn"`
	Stoct           int64   `json:"stoct"`
	Price           float64 `json:"price"`
}

type NewProduct struct {
	Title           string  `json:"title"`
	Author          string  `json:"author"`
	Publicatio_year int64   `json:"publication_year"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	ISBN            string  `json:"isbn"`
	Stoct           int64   `json:"stoct"`
	Price           float64 `json:"price"`
}