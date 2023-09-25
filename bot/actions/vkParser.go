package actions

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func VkParser(url string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	vkToken, exists := os.LookupEnv("VK_TOKEN")
	if !exists {
		fmt.Println("vkToken doesent exist")
	}
	//fmt.Println(vkToken)

	params := req.URL.Query()
	params.Add("access_token", vkToken)
	//params.Add("owner_id", "-174672969")
	params.Add("domain", "blankclub_spb")
	params.Add("count", "1")
	params.Add("v", "5.131")
	req.URL.RawQuery = params.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	//decodedData, err := base64.StdEncoding.DecodeString(string(resBody))
	//if err != nil {
	//	log.Fatal(err, " error with decoding response body")
	//}

	//fmt.Println(resBody, string(resBody))
	//fmt.Println(len(resBody), string(resBody))

	data, err := json.Marshal(string(resBody))
	if err != nil {
		log.Fatal(err, " error with Marshaling data")
	}

	fmt.Println("SJON????? \n", string(data), "SJON????? \n", string(resBody))
}
