package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"techTask/internal/models"

	//"net/http"
	"techTask/internal/controller"
)

func Router(c *controller.Controller, r *gin.Engine) {
	//r.GET("/", c.Data.Welcome)
	r.GET("/download", func(ctx *gin.Context) {
		url := "https://drive.google.com/u/0/uc?id=1IwZ3uUCHGpSL2OoQu4mtbw7Ew3ZamcGB&export=download"
		filename := "dataParse.csv"

		err := c.Data.DownloadFile(url, filename)
		if err != nil {
			fmt.Println(err)
			fmt.Println(err)
			fmt.Println(err)
			fmt.Println(err)
			fmt.Println(err)
			fmt.Println(err)
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
		d, _ := json.Marshal(data)
		_, err = ctx.Writer.Write(d)
		if err != nil {
			log.Println(err)
		}
	})

	//r.GET("/transaction/{id}", c.Data.GetByTransactionID)
	//r.GET("/terminal/{id}", c.Data.GetByTerminalID)
	//r.GET("/status={s}", c.Data.GetByStatus)
	//r.GET("/payment-type={t}", c.Data.GetByPaymentType)
	//r.GET("/from={start}&to={end}", c.Data.GetByTime)
	//r.GET("/narrative={text}", c.Data.GetByNarrative)
}
