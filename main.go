package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type user struct {
	Id       int    `json:"id" form:"id"`
	Nama     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"Password"`
}

var users []user

//--------------controllers------------------------------------------

func GetUsercontroller(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"Messages": "succes get all users",
		"users":    users,
	}) //hardcodedata
}

func SearchUserController(c echo.Context) error {
	Index, _ := strconv.Atoi(c.Param("id"))
	for _, value := range users {
		if value.Id == Index {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "user found!!",
				"user":     value,
			})

		}

	}
	return c.JSON(http.StatusNotFound, "Tidak Ada")
}

func CreateUsersController(c echo.Context) error {
	user := user{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})

}

func remove(users []int, i int) []int {
	copy(users[i:], users[i+1:])
	return users[:len(users)-1]
}

func deleteUserController(c echo.Context) error {
	Index, _ := strconv.Atoi(c.Param("id"))
	for _, value := range users {
		if value.Id == Index {
			NewUser = remove 
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "user deleted!!",
				"user":
			})

		}

	}
	return c.JSON(http.StatusNotFound, "Tidak Ada")
}

func main() {
	c := echo.New()
	c.GET("/user", GetUsercontroller)
	c.POST("/user", CreateUsersController)
	c.GET("/user/:id", SearchUserController)
	c.DELETE("/user/:id", deleteUserController)

	c.Logger.Fatal(c.Start(":8000"))
}
