package main

import (
	"github.com/kardianos/service"
	"neton/utils"
	"log"
	"os"
	"strconv"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	//检查网络
	c := -1
	code, _ := utils.GetStatusCode("https://www.baidu.com")
	if code == -1 {
		code, _ := utils.GetStatusCode("https://www.aliyun.com")
		if code != -1 {
			c = code
		}
	}else {
		c = code
	}

	if c == 200 {
		utils.Info("网络健康")
		return
	}
	utils.Info("网络不通: " + strconv.Itoa(c))
}

func (p *program) Stop(s service.Service) error {
	return nil
}

//Main
func main() {
	svcConfig := &service.Config{
		Name:        "Neton", //服务显示名称
		DisplayName: "Neton", //服务名称
		Description: "执行网络适配工作.", //服务描述
	}
	prg := &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatalln(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			s.Install()
			log.Fatalln("服务安装成功")
			return
		}
		if os.Args[1] == "uninstall" {
			s.Uninstall()
			log.Fatalln("服务卸载成功")
			return
		}
	}

	err = s.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
