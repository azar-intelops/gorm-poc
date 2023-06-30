package main

// only need mysql OR sqlite
// both are included here for reference
import (
	"gromnew/pkg/rest/server/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)


func main() {
	r := gin.Default()
	userController, err := controllers.NewUserController()
	if err != nil {
		logrus.Debug(err)
	}
	r.GET("/people/", userController.ListUsers)
	r.GET("/people/:id", userController.FetchUser)
	r.POST("/people", userController.CreateUser)
	r.PUT("/people/:id", userController.UpdateUser)
	r.DELETE("/people/:id", userController.DeleteUser)

	r.Run(":8080")
}

// func DeletePerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	d := db.Where("id = ?", id).Delete(&person)
// 	fmt.Println(d)
// 	c.JSON(200, gin.H{"id #" + id: "deleted"})
// }

// func UpdatePerson(c *gin.Context) {

// 	var person Person
// 	id := c.Params.ByName("id")

// 	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	}
// 	c.BindJSON(&person)

// 	db.Save(&person)
// 	c.JSON(200, person)
// }

// func CreatePerson(c *gin.Context) {

// 	var person Person
// 	c.BindJSON(&person)

// 	db.Create(&person)
// 	c.JSON(200, person)
// }

// func GetPerson(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var person Person
// 	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, person)
// 	}
// }

// func GetPeople(c *gin.Context) {
// 	var people []Person
// 	if err := db.Find(&people).Error; err != nil {
// 		c.AbortWithStatus(404)
// 		fmt.Println(err)
// 	} else {
// 		c.JSON(200, people)
// 	}
// }
