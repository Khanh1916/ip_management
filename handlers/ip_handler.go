package handlers
// Xu ly cac yeu cau API

import (
	"database/sql"			// Thu vien lam viec voi SQL database
	"net/http"			// Xu ly HTTP requests, phuc vu cho viec dinh nghia route cua Gin framework
	//"encoding/json"
	"log"				// Ghi lai log khi co loi
	// Cac package noi bo du an
	"ip_management/models"	// Dinh nghia du lieu	
	"ip_management/utils"	// Xac thuc du lieu
	"ip_management/config"	// Ket noi du lieu

	"github.com/gin-gonic/gin"	// Xay dung REST API
)

// Them dia chi IP
func AddIP(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var ip models.IP
		if err := c.BindJSON(&ip); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		} // Nhan JSON tu yeu cau HTTP, kiem tra dinh dang du lieu

		if !utils.IsValidIP(ip.Address) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IP address"})
			return
		} // Xac thuc dinh dang dia chi IP	

		// Kiem tra su ton tai cua dia chi IP
		var exists int64
		err := config.DB.QueryRow("SELECT COUNT(*) FROM ips WHERE ip_address = ?", ip.Address).Scan(&exists)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query error", "details": err.Error()})
			return
		} // Ghi lai loi khi truy van du lieu phuc vu check su ton tai cua IP 
		if exists > 0 {
			c.JSON(http.StatusConflict, gin.H{"error": "IP already exists"})
			return
		} // Neu IP ton tai dung dich vu

		// Chen IP vao database
		_, err = config.DB.Exec("INSERT INTO ips (ip_address, allocated) VALUES (?, ?)", ip.Address, false) // Trang thai ALLCOCATED mac dinh = false
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert IP", "details": err.Error()})
			return
		} Neu co loi khi them dia chi vao dtb, he thong se thong bao loi

		c.JSON(http.StatusCreated, gin.H{"message": "IP added successfully"})	// Tra ve tin nhan thanh cong
	}
}

// Lay danh sach dia chi IP
func GetAllIPs(c *gin.Context) {
    rows, err := config.DB.Query("SELECT ip_address, allocated FROM ips") // Truy van all IPs tu database
    if err != nil {
        log.Println("Query failed:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch IPs"})
        return
    } // Neu loi truy van, he thong se dung dich vu va thong bao
    defer rows.Close()

    var ips []map[string]interface{} // Bien luu tru du lieu gom ip va chi so allocated
    for rows.Next() {	// Doc du lieu tu dtb va tra ve IP list
        var ip string
        var allocated bool
        if err := rows.Scan(&ip, &allocated); err != nil {
            log.Println("Scan error:", err)
            continue
        }

        ips = append(ips, map[string]interface{}{	
            "ip_address": ip,
            "allocated":  allocated,
        }) // Them mot dia chi IP sau moi vong for
    }

    c.JSON(http.StatusOK, gin.H{"ips": ips})	// Tra ve danh sach IP
}

// Xoa 1 dia chi IP
func DeleteIP(c *gin.Context) {
    ip := c.Param("ip")

    tx, err := config.DB.Begin() // Bat dau transaction de xoa ip
    if err != nil {
    	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
    	return
    }

    result, err := tx.Exec("DELETE FROM ips WHERE ip_address = ?", ip) // Xoa dia chi IP tu database
    if err != nil {
        tx.Rollback() // Huy transaction neu co loi
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete IP"})
        return
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
    	tx.Rollback()
    	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get affected rows"})
    	return
    }
    if rowsAffected == 0 {
    	tx.Rollback()
    	c.JSON(http.StatusNotFound, gin.H{"error": "IP not found"})
    	return
    } // Neu khong tim thay IP, tra ve loi

    tx.Commit() // Commit transaction
    c.JSON(http.StatusOK, gin.H{"message": "IP deleted successfully"}) // Tra ve thong bao thanh cong
}

// Cap phat tu dong dia chi IP chua duoc cap phat
func AllocateIP(db *sql.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var ipAddress string // Luu tru available ip

        // Lay dia chi IP chua duoc cap phat
        err := config.DB.QueryRow("SELECT ip_address FROM ips WHERE allocated = false ORDER BY id ASC LIMIT 1").Scan(&ipAddress)
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "No available IPs"})
            return
        } else if err != nil {
            log.Println("Error querying available IP:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
            return
        }

        // Cap nhat trang thai cua IP thanh da duoc cap phat, tranh race condition
        result, err := db.Exec("UPDATE ips SET allocated = true WHERE ip_address = ? AND allocated = false LIMIT 1", ipAddress)
        if err != nil {
            log.Println("Error updating allocated status:", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to allocate IP"})
            return
        } // Loi trong viec cap nhat allocated

        rowsAffected, _ := result.RowsAffected()
        if rowsAffected == 0 {
            log.Println("No IP was updated. Check if the IP exists.")
            c.JSON(http.StatusInternalServerError, gin.H{"error": "No IP was updated"})
            return
        } // Neu loi trong viec update, he thong thong bao loi

        c.JSON(http.StatusOK, gin.H{
            "allocated_ip": ipAddress,
            "status":       "allocated",
        }) // Tra ve ip cho viec cap phat dong
    }
}

// Xac thuc IPv4 va IPv6
func ValidateIP(c *gin.Context) {
    var request struct {
        IPAddress string `json:"ip_address"`
    } // Nhan dia chi IP can kiem tra

    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    } // Loi dinh dang du lieu

    if utils.IsIPv4(request.IPAddress) {
        c.JSON(http.StatusOK, gin.H{"message": "Valid IPv4"})
    } else if utils.IsIPv6(request.IPAddress) {
        c.JSON(http.StatusOK, gin.H{"message": "Valid IPv6"})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IP address"})
    } // Goi ham xu ly xac thuc dinh dang cua IP va tra ve ket qua
}