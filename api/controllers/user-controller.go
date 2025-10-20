package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
	"jester_api/repository"
)

var repo *repository.UserRepo

func SetDB (db *repository.UserRepo) {
	repo = db
}

func GetAllUsers(ctx *gin.Context){
	if repo == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database not initialized"})
		return
	}

	users, err := repo.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}