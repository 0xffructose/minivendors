package main

import (
	"log"
	"os"
	"time"

	"backend/database"
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	if err := godotenv.Load(); err != nil {
		log.Println("Uyarı: .env dosyası bulunamadı, sistem ortam değişkenleri kullanılacak.")
	}

	database.ConnectDatabase()

	r := gin.Default()

	// 3. CORS Konfigürasyonu
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		
		protected := api.Group("/")
		protected.Use(middleware.RequireAuth())
		{
			protected.GET("/dashboard-data", func(c *gin.Context) {
				c.JSON(200, gin.H{"data": "Veritabanı ve CORS yapılandırması tamamlanmış güvenli veri!"})
			})
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}