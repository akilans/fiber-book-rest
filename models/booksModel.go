package models

// Book Type -> Books table
type Book struct {
	ID       uint64
	Title    string
	Author   string
	Price    float64
	ImageURL string
}
