package todo

import (
	"todo_app/util"

	"todo_app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteTodo(c *gin.Context, client *mongo.Client) {

	id, objErr := primitive.ObjectIDFromHex(c.Param("id"))
	if objErr != nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "That is not a valid id",
		})
		return
	}

	todosCollection := client.Database("TodosDB").Collection("todos")

	deletedDocument := models.Todo{}

	deleteErr := todosCollection.FindOneAndDelete(c.Request.Context(), bson.M{"_id": id}).Decode(&deletedDocument)
	if deleteErr != nil {
	
		if deleteErr == mongo.ErrNoDocuments {
			c.JSON(400, util.ResMessage{
				Success: false,
				Message: "There is no todo with that id",
			})
			return
		}
		c.JSON(500, util.ResError{
			Success: false,
			Error:   deleteErr,
		})
		return
	}

	c.JSON(200, util.ResMessage{
		Success: true,
		Message: "Todo deleted",
	})
}
