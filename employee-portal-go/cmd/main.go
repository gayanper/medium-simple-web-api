package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := initDb()

	if err != nil {
		panic("Error connecting to the database")
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()
	initRoute(db)
}

func initDb() (*sql.DB, error) {
	var HOST = os.Getenv("DB_HOST") + ":3306"
	var PW = os.Getenv("DB_PW")
	fmt.Println("Using database host ", HOST)
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               PW,
		Net:                  "tcp",
		Addr:                 HOST,
		DBName:               "employees",
		AllowNativePasswords: true,
	}
	return sql.Open("mysql", cfg.FormatDSN())
}

func initRoute(db *sql.DB) {
	gin.SetMode(gin.ReleaseMode)
 	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/employees", func(ctx *gin.Context) {
		afterDt, err := time.Parse(time.DateOnly, ctx.Query("after"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid after date"})
			return
		}
		result, err := findHiredAfter(db, afterDt.Format(time.DateOnly))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("server error : %v", err)})
			return
		}
		ctx.JSON(http.StatusOK, result)
	})
	//pprof.Register(router)
	router.Run("localhost:8080")
}

type Employee struct {
	EmpNo     string `json:"empNo"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	JoinedOn  string `json:"joinedOn"`
}

func findHiredAfter(db *sql.DB, after string) ([]Employee, error) {
	rows, err := db.Query("SELECT emp_no,first_name,last_name,hire_date FROM `employees` WHERE hire_date > ?",
		after)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Employee
	for rows.Next() {
		var emp Employee
		rows.Scan(&emp.EmpNo, &emp.FirstName, &emp.LastName, &emp.JoinedOn)
		result = append(result, emp)
	}

	return result, nil
}
