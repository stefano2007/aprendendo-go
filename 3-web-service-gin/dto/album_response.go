package dto

type AlbumResponse struct {
	ID     string  `json:"id"     binding:"required,min=1"`
	Title  string  `json:"title"  binding:"required,min=1,max=100"`
	Artist string  `json:"artist" binding:"required,min=1,max=100"`
	Price  float64 `json:"price"  binding:"required,min=1,max=1000"`
}
