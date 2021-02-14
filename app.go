package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"reserve.com/m/util"
	"time"

	//"sourcegraph.com/sourcegraph/go-selenium"
)

type Info struct {
	StartTime string  //开始抢号时间
	Register  string  //注册按钮
	Login     string  //微信登陆按钮
	Hospital  string  //选择医院
	Category  string  //选择大类
	Item      string  //选择科室
	Before    string  //前一天
	Today     string  //今天
	Choose    string  //选择按钮
	Person    string  //就诊人
	Captcha   string  //验证码下方按钮
	Apply     string  //提交按钮
}

var hospital = Info{
	StartTime: "2020-06-04 07:00:00",
	Register:  "//*[@id=\"main\"]/div[1]/div/div[2]/span[2]",
	Login:     "//*[@id=\"app-container\"]/div[2]/div/div[2]/div[1]/div[1]/div[1]/div[2]/div/span",
	Hospital:  "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[3]/div[1]/div[2]/div[5]/div/div/div/div[1]",
	Category:  "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[4]/div[1]/div/div/div[1]/div/div[6]",
	Item:      "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[4]/div[2]/div[6]/div[2]/div[2]/span",
	Before:    "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[3]/div[2]/div[7]/div[1]",
	Today:     "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[3]/div[2]/div[8]/div[1]",
	Choose:    "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[4]/div[1]/div[2]/div/div[2]/div[2]",
	Person:    "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[2]/div[2]/div[1]/div/div/div/div/div[1]",
	Captcha:   "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[2]/div[5]/div[2]/form/div[2]/div/div/span/span/span",
	Apply:     "//*[@id=\"main\"]/div[2]/div/div[1]/div/div/div[2]/div/div[3]/div/div",
}

func main() {
	var err error
	var wd selenium.WebDriver

	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if wd, err = selenium.NewRemote(caps, "http://localhost:9515"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer wd.Quit()

	err = wd.Get("https://www.114yygh.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	//点击注册按钮
	fmt.Println("点击登陆注册按钮")
	click(wd, hospital.Register)

	//点击微信登陆按钮
	fmt.Println("点击微信登陆按钮")
	click(wd, hospital.Login)

	time.Sleep(10 * time.Second)

	//点击协和医院
	fmt.Println("点击协和医院")
	click(wd, hospital.Hospital)

	//点击妇产科
	fmt.Println("点击妇产科")
	click(wd, hospital.Category)

	time.Sleep(2 * time.Second)
	//点击妇科门诊
	fmt.Println("点击妇科门诊")
	click(wd, hospital.Item)

	for {
		if util.CheckTime(hospital.StartTime) {
			break
		}
		time.Sleep(100 * time.Millisecond)
		fmt.Println("waiting")
	}

	for {
		//点击前一天的日期 <span data-v-ed76ad7c="">06月10日</span>
		click(wd, hospital.Before)
		//点击后一天的日期
		click(wd, hospital.Today)

		btn, err := wd.FindElement(selenium.ByXPATH, hospital.Choose)
		if err == nil {
			if err := btn.Click(); err != nil {
				fmt.Println(err)
			}
			fmt.Println("点击挂号操作")
			break
		}
	}

	//选择挂号人
	fmt.Println("开始选择挂号人")
	click(wd, hospital.Person)

	////检测是否含有验证码
	//fmt.Println("检测是否含有验证码")
	//click(wd, hospital.Captcha)

	//点击挂号按钮
	fmt.Println("点击确认挂号按钮")
	click(wd, hospital.Apply)

	time.Sleep(1 + time.Minute)
}

func click(wd selenium.WebDriver, ele string) {
	for {
		time.Sleep(100 * time.Millisecond)
		btn, err := wd.FindElement(selenium.ByXPATH, ele)
		if err != nil {
			continue
		}
		if err := btn.Click(); err != nil {
			fmt.Println(err)
		}
		return
	}

}
