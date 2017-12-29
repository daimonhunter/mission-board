package handlers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

//'id as aid, title, content, update_at'

type Notice struct {
	id       int
	title    string
	content  string
	updateAt int
}

type ResponseData struct {
	errorCode int
	msg       string
	success   bool
	data      Notice
	signature string
}

func EmergencyNotice(c *gin.Context) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wealthbetter_com?charset=utf8")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var notice Notice
	var responseData ResponseData

	responseData.errorCode = 0
	responseData.msg = ""
	responseData.data = notice
	responseData.success = true
	noticeType := c.DefaultQuery("noticeType", "1")

	err = db.QueryRow("SELECT id as aid, title, content, update_at FROM wb_emergency_notice WHERE `type`=? AND `status` = 1 ORDER BY create_at DESC ", noticeType).Scan(&notice.id, &notice.title, &notice.content,&notice.updateAt)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"data": responseData})
	}
	responseData.data = notice
	log.Println(notice)
	log.Println(responseData)
	c.JSON(http.StatusOK, gin.H{"data": responseData})

}
