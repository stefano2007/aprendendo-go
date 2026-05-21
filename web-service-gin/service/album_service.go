package service

import (
	"github.com/stefano2007/aprendendo-go/domain"
	"github.com/stefano2007/aprendendo-go/dto"
)

type albumService struct {
	repository domain.AlbumRepository
}

type AlbumService interface {
	GetAlbumByID(id string) (dto.AlbumResponse, error)
	CreateAlbum(albumCreateRequest dto.AlbumCreateRequest) error
	ListAlbums() []dto.AlbumResponse
	UpdateAlbum(album dto.AlbumUpdateRequest, id string) error
	DeleteAlbumByID(id string) error
}

func NewAlbumService(r domain.AlbumRepository) *albumService {
	return &albumService{
		repository: r,
	}
}

func (s *albumService) GetAlbumByID(id string) (dto.AlbumResponse, error) {

	album, err := s.repository.FindByID(id)

	if err != nil {
		return dto.AlbumResponse{}, err
	}

	return domainToDTO(*album), nil
}

func (s *albumService) CreateAlbum(albumCreateRequest dto.AlbumCreateRequest) error {

	newAlbum := domain.Album{
		ID:     albumCreateRequest.ID,
		Title:  albumCreateRequest.Title,
		Artist: albumCreateRequest.Artist,
		Price:  albumCreateRequest.Price,
	}

	error := s.repository.Save(&newAlbum)

	if error != nil {
		return error
	}
	return nil
}

func (s *albumService) UpdateAlbum(albumUpdateRequest dto.AlbumUpdateRequest, id string) error {

	album, err := s.repository.FindByID(id)

	if err != nil {
		return err
	}

	album.Title = albumUpdateRequest.Title
	album.Artist = albumUpdateRequest.Artist
	album.Price = albumUpdateRequest.Price

	error := s.repository.Save(album)

	if error != nil {
		return error
	}
	return nil
}

func (s *albumService) ListAlbums() []dto.AlbumResponse {

	albums := make([]dto.AlbumResponse, 0)
	for _, album := range s.repository.FindAll() {
		albums = append(albums, domainToDTO(album))
	}

	return albums
}

func (s *albumService) DeleteAlbumByID(id string) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func domainToDTO(album domain.Album) dto.AlbumResponse {
	return dto.AlbumResponse{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
}
