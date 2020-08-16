package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiUrl = "http://www.sojson.com/open/api/weather/json.shtml?city="

func Print(date string, r Response) {
	fmt.Println("城市是:", r.CityName)
	if date == "今天" {
		fmt.Println("湿度:", r.Data.ShiDu)
		fmt.Println("空气质量:", r.Data.Quality)
		fmt.Println("温馨提示:", r.Data.Ganmao)
	} else if date == "昨天" {
		fmt.Println("日期:", r.Data.Yesterday.Date)
		fmt.Println("温度:", r.Data.Yesterday.Low, r.Data.Yesterday.High)
		fmt.Println("风量:", r.Data.Yesterday.Fx, r.Data.Yesterday.Fl)
		fmt.Println("天气:", r.Data.Yesterday.Type)
		fmt.Println("温馨提示:", r.Data.Yesterday.Notice)
	} else if date == "预测" {
		fmt.Println("====================================")
		for _, item := range r.Data.Forecast {
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

func Request(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}
