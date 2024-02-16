package auth

import (
	"fmt"
	"log"
	"net/http"

	"todo_app/util"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"todo_app/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Register //
// @desc Create a new user from User struct and add it to the database
// @route POST /api/v1/auth/register
// @access Public
func Register(c *gin.Context, client *mongo.Client) {

	user := models.UserCred{}

	bindErr := c.ShouldBindJSON(&user)
	if bindErr != nil {
		log.Fatal(bindErr)
	}

	// checks to make sure that all fields are populated
	if user.Name == "" {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Please add a name",
		})
		return
	} else if user.Email == "" {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Please add an email",
		})
		return
	} else if user.Password == "" {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Please add a password",
		})
		return
	}

	user.Encrypt(user.Password)

	usersCollection := client.Database("TodosDB").Collection("users")
	fmt.Print(usersCollection)
	alreadyExists := models.User{}

	// search db to make sure provided email is unique
	findOneErr := usersCollection.FindOne(c.Request.Context(), bson.M{"email": user.Email}).Decode(&alreadyExists)
	if findOneErr == nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "That email already exists",
		})
		return
	}
	fmt.Print(user)
	insertOneResult, insertErr := usersCollection.InsertOne(c.Request.Context(), user)
	if insertErr != nil {
		c.JSON(400, util.ResError{
			Success: false,
			Error:   insertErr,
		})
	} else {

		token, getSignedErr := user.GetSignedJWT(insertOneResult.InsertedID.(primitive.ObjectID).Hex())
		if getSignedErr != nil {
			c.JSON(400, util.ResError{
				Success: false,
				Error:   getSignedErr,
			})
			return
		}

		// secure cookie unles in development env
		secure := true

		// strict for csrf safety
		c.SetSameSite(http.SameSiteStrictMode)

		c.SetCookie("token", token, 2000, "/", "", secure, true)

		c.JSON(200, util.ResUser{
			Success: true,
			Message: models.UserRes{
				Name:  user.Name,
				Email: user.Email,
			},
		})
	}
}
