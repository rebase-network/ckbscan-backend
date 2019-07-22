package address

import (
	"ckbscan-backend/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAddressReq struct {
	Hash string `json:"hash"`
}

type GetAddressResp struct {
	Number string `json:"number"`
}

func GetAddress(c *gin.Context) {
	var req GetAddressReq
	if err := c.Bind(&req); err != nil {
		log.Error("c.Bind")
		c.JSON(http.StatusBadRequest, nil)
	}

	addressresp := &GetAddressResp{
		Number: "My address",
	}
	c.JSON(http.StatusOK, addressresp)
}
