package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var city string
	var day string
	flag.StringVar(&city, "c", "上海", "城市中文名")
	flag.StringVar(&day, "d", "今天", "可选 今天，昨天，预测")

	flag.Parse()

	body, err := Request(apiUrl + city)
	if err != nil {
		fmt.Printf("err was %v", err)
		return
	}

	var res Response
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Printf("err mesage %v", err)
		return
	}

	if res.Status != http.StatusOK {
		fmt.Printf("获取天气API出现错误, %s", res.Message)
		return
	}
	Print(day, res)
}

const apiUrl = "http://www.sojson.com/open/api/weather/json.shtml?city="

func Request(url string) (string, error) {
	client := http.DefaultClient
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err %v", err)
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	return string(body), nil
}

func Print(day string, res Response) {
	fmt.Println("城市：", res.CityName)

	if day == "今天" {
		fmt.Println("湿度:", res.Data.Shidu)
		fmt.Println("空气质量:", res.Data.Quality)
		fmt.Println("温馨提示:", res.Data.Ganmao)
	} else if day == "昨天" {
		fmt.Println("日期:", res.Data.Yesterday.Date)
		fmt.Println("温度:", res.Data.Yesterday.Low, res.Data.Yesterday.High)
		fmt.Println("风量:", res.Data.Yesterday.Fx, res.Data.Yesterday.Fl)
		fmt.Println("天气:", res.Data.Yesterday.Type)
		fmt.Println("温馨提示:", res.Data.Yesterday.Notice)
	} else if day == "预测" {
		fmt.Println("====================================")
		for _, item := range res.Data.Forecast {
			fmt.Println("日期:", item.Date)
			fmt.Println("温度:", item.Low, item.High)
			fmt.Println("风量:", item.Fx, item.Fl)
			fmt.Println("天气:", item.Type)
			fmt.Println("温馨提示:", item.Notice)
			fmt.Println("====================================")
		}
	} else {
		fmt.Println("大熊你是想刁难我胖虎吗?_?")
	}
}

type Response struct {
	Status   int    `json:"status"`
	CityName string `json:"city"`
	Data     Data   `json:"data"`
	Date     string `json:"date"`
	Message  string `json:"message"`
	Count    int    `json:"count"`
}

type Data struct {
	Shidu     string `json:"shidu"`
	Quality   string `json:"quality"`
	Ganmao    string `json:"ganmao"`
	Yesterday Day    `json:"yesterday"`
	Forecast  []Day  `json:"forecast"`
}

type Day struct {
	Date    string  `json:"date"`
	Sunrise string  `json:"sunrise"`
	High    string  `json:"high"`
	Low     string  `json:"low"`
	Sunset  string  `json:"sunset"`
	Aqi     float32 `json:"aqi"`
	Fx      string  `json:"fx"`
	Fl      string  `json:"fl"`
	Type    string  `json:"type"`
	Notice  string  `json:"notice"`
}
