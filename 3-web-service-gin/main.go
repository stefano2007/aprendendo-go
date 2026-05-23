package main

import (
	"github.com/gin-gonic/gin"

	"github.com/stefano2007/aprendendo-go/web-service-gin/controller"
	"github.com/stefano2007/aprendendo-go/web-service-gin/repository"
	"github.com/stefano2007/aprendendo-go/web-service-gin/service"
)

/*
	type AlbumDTO struct {
		ID     string  `json:"id"     binding:"required,min=1"`
		Title  string  `json:"title"  binding:"required,min=1,max=100"`
		Artist string  `json:"artist" binding:"required,min=1,max=100"`
		Price  float64 `json:"price"  binding:"required,min=1,max=1000"`
	}

var albums map[string]AlbumDTO

	func init() {
		albums = map[string]AlbumDTO{
			"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		}

		fmt.Println("init() executada")
	}

	func getAlbums(context *gin.Context) {
		context.IndentedJSON(http.StatusOK, albums)
	}

	func getAlbumByID(context *gin.Context) {
		id := context.Param("id")
		if album, ok := obterAlbumID(id); ok {
			context.IndentedJSON(http.StatusOK, album)
			return
		}
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}

	func createAlbum(context *gin.Context) {
		var newAlbum AlbumDTO
		if err := context.BindJSON(&newAlbum); err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album data"})
			return
		}
		if _, existe := obterAlbumID(newAlbum.ID); existe {
			context.IndentedJSON(http.StatusConflict, gin.H{"message": "já existe um album com esse id"})
			return
		}
		albums[newAlbum.ID] = newAlbum
		context.IndentedJSON(http.StatusCreated, newAlbum)
	}

	func obterAlbumID(id string) (AlbumDTO, bool) {
		album, ok := albums[id]
		return album, ok
	}

	func updateAlbum(context *gin.Context) {
		id := context.Param("id")
		var updateAlbum AlbumDTO
		if err := context.BindJSON(&updateAlbum); err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid album data"})
			return
		}

		if id != updateAlbum.ID {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id da rota e id do corpo da requisição não coincidem"})
			return
		}

		if _, existe := obterAlbumID(updateAlbum.ID); !existe {
			context.IndentedJSON(http.StatusConflict, gin.H{"message": "não existe um album com esse id"})
			return
		}
		albums[updateAlbum.ID] = updateAlbum
		context.IndentedJSON(http.StatusOK, updateAlbum)
	}

	func deleteAlbumByID(context *gin.Context) {
		id := context.Param("id")
		if _, ok := obterAlbumID(id); ok {
			delete(albums, id)
			context.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
*/
func main() {
	router := gin.Default()

	repo := repository.NewInMemoryAlbumRepository()
	serv := service.NewAlbumService(repo)
	handler := controller.NewAlbumHandler(serv)

	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.CreateAlbum)
	router.PUT("/albums/:id", handler.UpdateAlbum)
	router.DELETE("/albums/:id", handler.DeleteAlbumByID)

	router.Run("localhost:8080")
}
