package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	SchoolName string `json:"schoolName"`
}
type DataRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func giveCalltoService2(dataRequest []byte) (DataResponse, error) {
	client := &http.Client{}
	url := "http://localhost:8080/getData"
	method := "POST"
	body := dataRequest

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	var output DataResponse
	if resp.StatusCode == 200 {
		body, _ := io.ReadAll(resp.Body)
		_ = json.Unmarshal(body, &output)

	}
	return output, nil
}
func GetData(c *gin.Context) {
	var dataResponse DataResponse
	jsonbyte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error:", err)
	}

	dataResponse, err = giveCalltoService2(jsonbyte)
	if err != nil {
		fmt.Println("Error occured in dataResponse:", dataResponse)

	}

	c.JSON(http.StatusOK, dataResponse)
}
