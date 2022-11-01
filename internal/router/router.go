package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"techTask/internal/controller"
	"techTask/internal/models"
)

func Router(c *controller.Controller, r *gin.Engine) {
	//r.GET("/", c.Data.Welcome)
	r.GET("/download", func(ctx *gin.Context) {
		url := "https://drive.google.com/u/0/uc?id=1IwZ3uUCHGpSL2OoQu4mtbw7Ew3ZamcGB&export=download"
		filename := "dataParse.csv"
		fmt.Println(123)
		err := c.Data.DownloadFile(url, filename)
		if err != nil {
			log.Println(err)
		} else {
			d, _ := json.Marshal("Success")
			_, err = ctx.Writer.Write(d)
			if err != nil {
				log.Println(err)
			}
		}
	})

	r.GET("/search", func(ctx *gin.Context) {
		//param1 := ctx.Query("transaction")
		resp := models.Queries{
			Transactions: strings.Split(ctx.Query("transaction"), ","),
			Terminal:     strings.Split(ctx.Query("terminal"), ","),
			Status:       strings.Split(ctx.Query("status"), ","),
			Payment:      strings.Split(ctx.Query("payment"), ","),
			FromDate:     []string{ctx.Query("from")},
			ToDate:       []string{ctx.Query("to")},
			Narrative:    []string{ctx.Query("narrative")},
		}
		fmt.Println(resp)
		data, err := c.Data.Search(&resp)
		if err != nil {
			log.Println(err)
		}
		if data == nil {
			d, _ := json.Marshal("invalid data")
			_, err = ctx.Writer.Write(d)
			if err != nil {
				log.Println(err)
			}
		} else {
			d, _ := json.Marshal(data)
			_, err = ctx.Writer.Write(d)
			if err != nil {
				log.Println(err)
			}
		}

	})
}
