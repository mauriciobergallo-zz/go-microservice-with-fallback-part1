package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/adding"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/deleting"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/deletingFallback"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/listing"
	"github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/updating"
)

func NewRestService(as adding.Service, ds deleting.Service, us updating.Service, ls listing.Service, dfs deletingFallback.Service) {
	r := gin.Default()
	rest := &restService{as, ds, us, ls, dfs}

	r.GET("/api/health", rest.getHealth)

	r.GET("/api/users/:id", rest.getUserById)
	r.POST("/api/users", rest.postUser)
	r.DELETE("/api/users/:id", rest.deleteUser)
	r.PUT("/api/users/:id", rest.putUser)
	r.DELETE("/api/users/fallback", rest.deleteFallback)

	r.Run()
}

type restService struct {
	addingService           adding.Service
	deletingService         deleting.Service
	updatingService         updating.Service
	listingService          listing.Service
	deletingFallbackService deletingFallback.Service
}

func (rs *restService) getHealth(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "I´m Alive",
	})
}

func (rs *restService) postUser(c *gin.Context) {
	var u adding.User
	err := c.BindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "BadRequest",
			"message": err.Error(),
		})

		return
	}

	cu, err := rs.addingService.AddUser(u)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "InternalServer",
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"data":   cu,
	})

	return
}

func (rs *restService) deleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "BadRequest",
			"message": err.Error(),
		})

		return
	}

	err = rs.deletingService.RemoveUser(id)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "InternalServer",
			"message": err.Error(),
		})

		return
	}

	c.JSON(202, gin.H{
		"status": "ok",
	})

	return
}

func (rs *restService) putUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "BadRequest",
			"message": err.Error(),
		})

		return
	}

	var u updating.User
	err = c.BindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "BadRequest",
			"message": err.Error(),
		})

		return
	}

	u.ID = id

	uu, err := rs.updatingService.UpdateUser(u)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "InternalServer",
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"data":   uu,
	})

	return
}

func (rs *restService) deleteFallback(c *gin.Context) {
	err := rs.deletingFallbackService.RemoveUsersFallback()
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "InternalServer",
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
	})

	return
}

func (rs *restService) getUserById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"status":  "BadRequest",
			"message": err.Error(),
		})

		return
	}

	u, err := rs.listingService.ObtainUserById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "InternalServer",
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"status": "ok",
		"data":   u,
	})

	return
}
