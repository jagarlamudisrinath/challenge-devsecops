package api

import (
	"challenge/dummy"
	"fmt"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	var users []dummy.User
	c.BindJSON(&users)

	users, n, err := dummy.UserList()
	if err != nil {
		log.Errorf("Impossible to fetch all users")
		RespondJSON(c, 404, users)
	} else {
		log.Infof("List users returned %v values\n", n)
		RespondJSON(c, 200, users)
	}

}

func AddUser(c *gin.Context) {

	var user dummy.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		fmt.Println("Impoosible to bind json")
	} else {
		fmt.Println(user)

		err := dummy.UserAdd(&user)
		if err != nil {
			log.Errorf("Failed to add user: %v\n", user)
			RespondJSON(c, 404, user)
		} else {
			RespondJSON(c, 200, user)
		}
	}

}

func DelUser(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := dummy.UserDel(id); err != nil {
		log.Errorf("Could not delete user: %v", id)
		RespondJSON(c, 400, id)
	} else {
		RespondJSON(c, 200, id)
	}
}

func GetUser(c *gin.Context) {
	var user dummy.User
	id := c.Params.ByName("id")
	if err := dummy.UserGet(&user, id); err != nil {
		RespondJSON(c, 404, id)
	} else {
		RespondJSON(c, 200, user)
	}
}
