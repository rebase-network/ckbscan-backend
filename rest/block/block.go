package rest

import (
	"ckbscan-backend/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetBlockReq struct {
	Hash string `json:"hash"`
}

type GetBlockResp struct {
	Number string `json:"number"`
}

func GetBlock(c *gin.Context) {
	var req GetBlockReq
	if err := c.Bind(&req); err != nil {
		log.Error("c.Bind")
		c.JSON(http.StatusBadRequest, nil)
	}

	blockresp := &GetBlockResp{
		Number: "Genesis",
	}
	c.JSON(http.StatusOK, blockresp)
}
