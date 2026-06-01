package main

import (
	"github.com/gin-gonic/gin"
	"api/internal/database"
	"api/internal/models"
	"gorm.io/gorm"
	"net/http"
		"github.com/gin-contrib/cors"
		"time"
)

func main() {

	var r * gin.Engine = gin.Default()
db, err := database.Connect()
if err != nil {
    panic(err)
}

err = db.AutoMigrate(&models.SurveyResponse{})
if err != nil {
    panic(err)
}
r.Use(cors.New(cors.Config{
	AllowOrigins: []string{
		"http://localhost:3232",
		// ou seu domínio personalizado:
		// "https://pesquisa.seuprojeto.com",
	},

	AllowMethods: []string{
		"POST",
		"OPTIONS",
	},

	AllowHeaders: []string{
		"Origin",
		"Content-Type",
		"Accept",
	},

	MaxAge: 12 * time.Hour,
}))
r.Use(func(c *gin.Context) {
    println("Origin:", c.Request.Header.Get("Origin"))
    c.Next()
})
	r.POST("/survey",saveForm(db))	
	r.Run(":3000")


}

func saveForm(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {

        var response models.SurveyResponse

        if err := c.ShouldBindJSON(&response); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        if err := db.Create(&response).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "erro ao salvar resposta",
            })
            return
        }

        c.JSON(http.StatusCreated, gin.H{
            "message": "resposta salva com sucesso",
            "id":      response.ID,
        })
    }
}
