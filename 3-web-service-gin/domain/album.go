package domain

type Album struct {
	ID     string
	Title  string
	Artist string
	Price  float64
}

type AlbumRepository interface {
	FindByID(id string) (*Album, error)
	FindAll() []Album
	Save(album *Album) error
	Delete(id string) error
}
