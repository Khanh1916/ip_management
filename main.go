package main

import (
	"log"				// Ghi log he thong
	"ip_management/config"		// Chua logic ket noi dtb
	"ip_management/handlers"	// Cac ham xu ly API

	"github.com/gin-gonic/gin"	// Xay dung REST API
)

func main() {
	err := config.InitDB()  // Khoi tao database
    	if err != nil {
        log.Fatal("Database initialization failed:", err)
    	}

	r := gin.Default()	// Tao router mac dinh cho phep Gin dinh nghia API endpoints

	// Dinh nghia cac API endpoints
	r.POST("/ip", handlers.AddIP(config.DB))       		// Them IP moi
	r.GET("/ips", handlers.GetAllIPs)   			// Lay danh sach IPs
	r.DELETE("/ip/:ip", handlers.DeleteIP) 			// Xoa IP
	r.GET("/ip/allocate", handlers.AllocateIP(config.DB)) 	// Cap phat IP tu đong
	r.POST("/validate-ip", handlers.ValidateIP)		// Xac thuc dia chi IP

	// Chay server tren cong 8080
	log.Println("Server running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	} // Thong bao loi khi va dung ctrinh khi server bij loi khoi dong
}