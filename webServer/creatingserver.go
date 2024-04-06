package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
 
type Albums struct {
	Id    string   `json:"id"`
	Name  string   `json:"name"`
	Price string   `json:"price"`
}

var allAlbums=[]Albums{
	// {Id:1,Name:"abc",Price:20},
	// {Id:2,Name:"xyz",Price:30},
}


func getAlbums(c *gin.Context){
	// *gin Context contains about the request data
	// c for sending data
	fmt.Println("c is ",c)
	c.JSON(http.StatusOK,allAlbums)
	      // status      dataThat we sending
}

func addAlbum(c *gin.Context){
	var newAlbum Albums
	if err:=c.BindJSON(&newAlbum);err!=nil{
		 fmt.Println("error is",err)
          return
	}
      fmt.Println("new album is ",newAlbum)
	allAlbums=append(allAlbums, newAlbum)
	c.JSON(http.StatusCreated,allAlbums)
}

func getAlbumById(c *gin.Context){
    id:=c.Param("id")
	fmt.Println("id is ",id)
	for _,value:=range allAlbums {
      if(value.Id==id){
		c.IndentedJSON(http.StatusOK,value)
		return 
	  }
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{
		"message":"id not find",
	})
}

func deleteAlbumById(c *gin.Context){
    id:=c.Param("id")
	var index int=-1
	for ind,value:=range allAlbums{
		if value.Id==id{
			index=ind
			break
		}
	}
	fmt.Println(index)
	if index==-1{
		c.IndentedJSON(http.StatusNotFound,gin.H{
			"message":"id not find",
		})
		return 
	}
    allAlbums=append(allAlbums[:index],allAlbums[index+1:]...)
	c.IndentedJSON(http.StatusOK,allAlbums)
}

func main()  {
	fmt.Println("creating server")
	router:=gin.Default()
	router.GET("/albums",getAlbums)
	router.POST("/addalbum",addAlbum)
	router.GET("/getbyid/:id",getAlbumById)
	router.DELETE("/deletebyid/:id",deleteAlbumById)
	// server set up in local host 4000 
	// output-> Listening and serving HTTP on localhost:4000
	router.Run("localhost:4000")
}