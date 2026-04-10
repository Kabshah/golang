package routes

import (
	"go-agent/chains"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

func generateVacation(r GenerateVacationIdeaRequest) GenerateVacationIdeaResponse {
	id := uuid.New()
	go chains.GenerateVacationIdeaChange(id, r.Budget, r.FavoriteSeason, r.Hobbies)
	return GenerateVacationIdeaResponse{Id: id, Completed: false}
}

// all funs related to db are in chains folder
// 1) yeh db sy fetch krni wali cheez h uuid aur yeh send back wali cheezein heinerr,getvac response
func getVacation(id uuid.UUID) (GetVacationIdeaResponse, error) {
	v, err := chains.GetVacationFromDb(id)
	if err != nil {
		return GetVacationIdeaResponse{}, err
	}
	return GetVacationIdeaResponse{Id: v.Id, Completed: v.Completed, Idea: v.Idea}, nil
}
func GetVacationRouter(router *gin.Engine) *gin.Engine {
	registrationRoutes := router.Group("/vacation")
	registrationRoutes.POST("/create", func(c *gin.Context) {
		// creating request obj
		var req GenerateVacationIdeaRequest
		err := c.BindJSON(&req)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
		} else {
			c.JSON(http.StatusOK, generateVacation(req))
		}

	})
	registrationRoutes.GET("/:id", func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Bad Request",
			})
		} else {
			resp, err := getVacation(id)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "ID not found",
				})
			} else {
				c.JSON(http.StatusOK, resp)
			}
		}
	})
	return router
}
