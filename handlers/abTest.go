package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	ID               int
	GroupID          int    //
	Mobile           string //
	InvitorMobile    string //邀请人手机号
	Password         string //
	GuesturePassword string //手势密码
	Salt             string //
	Hxname           string //环信用户名
	Status           int    //
	Investor         int    //
	GuideUid         uint   //引导用户注册的用户ID
	AeLabel          uint   //营销事件label
	CreateAt         uint
	UpdateAt         uint
}

type UsersInfo struct {
	ID             int
	UserID         uint //
	VipLevel       uint
	OpenID         string  //微信ID
	Channel        string  //渠道
	RealName       string  //
	NickName       string  //
	Avatar         string  //
	IDNumber       string  //
	Birthday       []uint8 //用户生日
	LastIp         string  //
	Signature      string  //用户个性签名
	Description    string  //
	MobileArea     string  //手机号所在地
	WechatID       string  //微信号
	IsWechatFriend uint    //是否微信好友:0=否，1=是
	CreateAt       uint
	UpdateAt       uint
}

type UserData struct {
	User      User
	UsersInfo UsersInfo
}

type UserActivity struct {
	ID        uint   //
	UserID    uint   //
	Channel   string //渠道
	Activity  string //
	Data      string //
	Ip        string //
	CreateAt  uint   //
	DayCreate string //
}

// GenerateRangeNum 生成一个区间范围的随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

func TestAb(c *gin.Context) {

	type ResponseData struct {
		errorCode int
		msg       string
		success   bool
		data      UserData
		signature string
	}

	userId := GenerateRangeNum(29789, 30418)
	//db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wealthbetter_com")
	db, err := sql.Open("mysql", "root:Tutululu@tcp(127.0.0.1:3306)/wealthbetter_dev")
	//db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.219:3306)/bonuolicai_com")
	defer db.Close()
	if err != nil {
		log.Panic(err)
		return
	}
	var user User
	var usersInfo UsersInfo
	var userData UserData
	var responseData ResponseData

	userData.UsersInfo = usersInfo
	userData.User = user

	responseData.errorCode = 0
	responseData.msg = ""
	responseData.data = userData
	responseData.success = true

	err = db.QueryRow("SELECT * FROM wb_users WHERE id = ?", userId).Scan(&user.ID, &user.GroupID, &user.Mobile, &user.InvitorMobile, &user.Password, &user.GuesturePassword, &user.Salt, &user.Hxname, &user.Status, &user.Investor, &user.GuideUid, &user.AeLabel, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"data": responseData})
	}

	err = db.QueryRow("SELECT * FROM wb_users_info WHERE user_id = ?", userId).Scan(&usersInfo.ID, &usersInfo.UserID, &usersInfo.VipLevel, &usersInfo.OpenID, &usersInfo.Channel, &usersInfo.RealName, &usersInfo.NickName, &usersInfo.Avatar, &usersInfo.IDNumber, &usersInfo.Birthday, &usersInfo.LastIp, &usersInfo.Signature, &usersInfo.Description, &usersInfo.MobileArea, &usersInfo.WechatID, &usersInfo.IsWechatFriend, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"data": responseData})
	}
	userData.UsersInfo = usersInfo
	userData.User = user
	responseData.data = userData
	log.Println(userData)
	c.JSON(http.StatusOK, gin.H{"data": userData})
}

func TestAb2(c *gin.Context) {
	var a UserActivity

	type ResponseData struct {
		errorCode int
		msg       string
		success   bool
		data      int
		signature string
	}

	//db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/wealthbetter_com")
	db, err := sql.Open("mysql", "root:Tutululu@tcp(127.0.0.1:3306)/wealthbetter_dev")
	//db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.219:3306)/bonuolicai_com")
	defer db.Close()
	if err != nil {
		log.Panic(err)
		return
	}

	a.UserID = 29778999
	a.Channel = "web"
	a.Activity = "register"
	a.Data = "{\"muname\":\"18500335694\",\"mupass\":\"Y2IwZDdjOTEwNmRlZDc2NzQyZDJjMDU4MjM3NjI2ODVkM2Y1ZDEwZg==\",\"mcode\":\"131762\",\"channel\":\"web\",\"aeLabel\":\"1000100040\",\"member_type\":\"1\",\"invitor_mobile\":\"\"}"
	a.Ip = "220.249.18.222"
	a.CreateAt = 1492483938
	a.DayCreate = "2017-04-18"

	stmt, err := db.Prepare("INSERT wb_users_activity (user_id,channel,activity,data,ip,create_at,day_create) values (?,?,?,?,?,?,?)")
	if err != nil {
		log.Panic(err)
		return
	}
	res, err := stmt.Exec(a.UserID, a.Channel, a.Activity, a.Data, a.Ip, a.CreateAt, a.DayCreate)
	if err != nil {
		log.Panic(err)
		return
	}
	id, err := res.LastInsertId()
	fmt.Println(id)
	c.JSON(http.StatusOK, gin.H{"data": id})
}
