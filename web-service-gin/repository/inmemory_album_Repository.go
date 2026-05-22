package repository

import (
	"fmt"

	"github.com/stefano2007/aprendendo-go/web-service-gin/domain"
)

type InMemoryAlbumRepository struct {
	data map[string]domain.Album
}

func NewInMemoryAlbumRepository() *InMemoryAlbumRepository {

	return &InMemoryAlbumRepository{
		data: cargaInicialDados(),
	}
}

func cargaInicialDados() map[string]domain.Album {
	return map[string]domain.Album{
		"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
}

func (r *InMemoryAlbumRepository) FindByID(id string) (*domain.Album, error) {
	album, ok := r.data[id]

	if !ok {
		return nil, fmt.Errorf("album not found")
	}

	return &album, nil
}

func (r *InMemoryAlbumRepository) FindAll() []domain.Album {
	albums := make([]domain.Album, 0, len(r.data))
	for _, album := range r.data {
		albums = append(albums, album)
	}
	return albums
}

func (r *InMemoryAlbumRepository) Save(album *domain.Album) error {
	r.data[album.ID] = *album
	return nil
}

func (r *InMemoryAlbumRepository) Delete(id string) error {
	if _, ok := r.data[id]; !ok {
		return fmt.Errorf("album not found")
	}
	delete(r.data, id)
	return nil
}
