package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"strconv"
)

func main() {
	fmt.Printf("Я родился\n")

	router := gin.Default()
	router.GET("/blockchain/:height", getBlockchain)

	router.Run("localhost:8080")
}

func getBlockchain(c *gin.Context) {
	var height int
	if _, err := fmt.Sscanf(c.Param("height"), "%d", &height); err != nil {
		panic(err)
	}
	//fmt.Printf("height: %d\n", height)

	whatsonchainResp, err := http.Get("https://api.whatsonchain.com/v1/bsv/test/block/height/" + strconv.Itoa(height))
	if err != nil {
		panic(err)
	}
	defer whatsonchainResp.Body.Close()

	body, err := io.ReadAll(whatsonchainResp.Body)

	if err := set_db(height, string(body)); err != nil {
		fmt.Println("Error during db update:", err)
	}

	c.String(http.StatusOK, string(body))
}

func set_db(height int, blockchain string) error {
	db, err := sql.Open("mysql", "root:root@/blockchain_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO blockchains VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(height, blockchain)
	if err != nil {
		return err
	}

	return nil
}
