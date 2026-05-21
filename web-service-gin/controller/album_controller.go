package controller

import (
	"github.com/stefano2007/aprendendo-go/dto"
	"github.com/stefano2007/aprendendo-go/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	service service.AlbumService
}

func NewAlbumHandler(service service.AlbumService) *AlbumHandler {
	return &AlbumHandler{
		service: service,
	}
}

func (h *AlbumHandler) GetAlbums(context *gin.Context) {
	albums := h.service.ListAlbums()
	context.IndentedJSON(http.StatusOK, albums)
}

func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := h.service.GetAlbumByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "album not found",
		})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) CreateAlbum(context *gin.Context) {
	var albumRequest dto.AlbumCreateRequest
	if err := context.BindJSON(&albumRequest); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album data"})
		return
	}
	if _, err := h.service.GetAlbumByID(albumRequest.ID); err == nil {
		context.IndentedJSON(http.StatusConflict, gin.H{"message": "já existe um album com esse id"})
		return
	}

	h.service.CreateAlbum(albumRequest)
	context.IndentedJSON(http.StatusCreated, albumRequest)
}

func (h *AlbumHandler) UpdateAlbum(context *gin.Context) {
	id := context.Param("id")
	var updateAlbum dto.AlbumUpdateRequest
	if err := context.BindJSON(&updateAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album data"})
		return
	}

	if _, error := h.service.GetAlbumByID(id); error != nil {
		context.IndentedJSON(http.StatusConflict, gin.H{"message": "não existe um album com esse id"})
		return
	}
	h.service.UpdateAlbum(updateAlbum, id)
	context.IndentedJSON(http.StatusOK, updateAlbum)
}

func (h *AlbumHandler) DeleteAlbumByID(context *gin.Context) {
	id := context.Param("id")
	if _, error := h.service.GetAlbumByID(id); error == nil {
		h.service.DeleteAlbumByID(id)
		context.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
		return
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
