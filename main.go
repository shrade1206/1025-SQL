package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gocv.io/x/gocv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Main_DB *gorm.DB

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/camdb?charset=utf8mb4&parseTime=True&loc=Local",
		// DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true, //不自動建立FK，處理速度會比較快
	})
	if err != nil {
		log.Fatal("Failed To Connect : ", err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(20) //
	sqlDB.SetConnMaxLifetime(20)
	defer sqlDB.Close()
	// db.LogMode(true)
	Main_DB = db
	// TestCreate()
	// Find()

	r := gin.New()
	// r.GET("/ws", wsEndpoint)
	r.GET("/ws", selectName)
	// r.GET("/get", gin.WrapH(http.FileServer(http.Dir("./public/select.html"))))

	r.NoRoute(gin.WrapH(http.FileServer(http.Dir("./public"))), func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		fmt.Println(path)
		fmt.Println(method)
		//檢查path的開頭使是否為"/"
		if strings.HasPrefix(path, "/") {
			fmt.Println("ok")
		}
	})
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("8080 err : ", err.Error())
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var newImg []byte
var data string

func selectName(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket 連線異常 : %s", err.Error())
		return
	}
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		if p != nil {
			var Aa Pic
			_, s, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
			}

			Main_DB.Where("Name = ? ", "%s", s).Find(Aa)
			log.Println(Aa)
			// if err := ws.WriteMessage(messageType, []byte()); err != nil {
			// 	log.Println(err)
			// 	return
			// }
		}
	}
}

func wsEndpoint(c *gin.Context) {
	// 透過http請求程序調用upgrader.Upgrade，來獲取*Conn (代表WebSocket連接)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket 連線異常 : %s", err.Error())
		return
	}
	log.Println("使用者已連線")

	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
		}
		if string(p) == "select" {
			var Aa Pic
			// _, s, err := ws.ReadMessage()
			// if err != nil {
			// 	log.Println(err)
			// }
			Main_DB.Model(&Pic{}).First(&Aa)

			// Main_DB.Where("Name = ? ", "%"+"%s"+"%", s).Find(Aa)
			log.Println(Aa)
			// if err := ws.WriteMessage(messageType, []byte()); err != nil {
			// 	log.Println(err)
			// 	return
			// }

		}
		if string(p) == "run" {

			func() {
				webcam, err := gocv.VideoCaptureDevice(0)
				if err != nil {
					log.Println(err)
				}

				time.Sleep(time.Second)
				img := gocv.NewMat()
				defer img.Close()

				webcam.Read(&img)
				defer webcam.Close()

				buf, err := gocv.IMEncode(".jpg", img)
				if err != nil {
					log.Fatal(err)
				}
				defer buf.Close() //nolint

				newImg = buf.GetBytes()
				// d, _ := os.ReadFile(a)

				data = base64.StdEncoding.EncodeToString(newImg)
				if err := ws.WriteMessage(messageType, []byte(data)); err != nil {
					log.Println(err)
					return
				}
			}()
		}
		if string(p) == "save" {
			a := []byte(data)

			Main_DB.Model(&pic).Create(map[string]interface{}{
				"Name": "GGGG", "Picture": a,
			})
		}
		log.Println("使用者訊息: " + string(p))
	}
}

var pic []Pic

// func save() {

// }
