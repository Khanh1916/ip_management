package config

import (
	"database/sql"		// Thu vien lam viec voi SQL dtb
	"fmt"			// Dung tao chuoi ket noi DSN cho database
	"log"			// Ghi log neu co van de
	"os"			// Doc cac bien moi truong tu .env

	"github.com/joho/godotenv"		// Ho tro load bien moi truong tu .env
	_ "github.com/go-sql-driver/mysql"	// Import driver MySQL cho database/sql dung
)

var DB *sql.DB	// Luu tru Database Connection

// Khoi tao database
func InitDB() error {
    err := godotenv.Load()	// Tai bien moi truong file .env
    if err != nil {
        log.Fatal("Error loading .env file") // Neu loi dung lai thong bao loi
	return err
    }

    // Lay thong tin tu bien moi truong
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    ) 	// Tao chuoi ket noi DSN

    db, err := sql.Open("mysql", dsn)	// Ket noi den database
    if err != nil {
        log.Fatal("Error connecting to database:", err)
        return err
    }

     
    err = db.Ping()	// Kiem tra ket noi   
    if err != nil {
        log.Fatal("Database connection failed:", err)
	return err
    }

    DB = db	// Luu ket noi database vao bien DB de package khac dung
    log.Println("Connected to database successfully")
    return nil
}