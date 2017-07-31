package main

import (
	"github.com/kardianos/service"
	"neton/utils"
	"log"
	"os"
	"strconv"
	"fmt"
)

type program struct{
	exit    chan struct{}
	*utils.Config
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	//检查网络
	c := -1
	code, _ := utils.GetStatusCode("https://www.aliyun.com")
	if code == -1 {
		code, _ := utils.GetStatusCode("https://www.baidu.com")
		if code != -1 {
			c = code
		}
	}else {
		c = code
	}

	if c == 200 {
		fmt.Println("Network is healthy")
		utils.Info("Network is healthy")
		return
	}
	fmt.Println("Network is unreasonable")
	utils.Info("Network is unreasonable: " + strconv.Itoa(c))
}

func (p *program) Stop(s service.Service) error {
	return nil
}

//Main
func main() {
	configPath, err := utils.GetConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	config, err := utils.GetConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	svcConfig := &service.Config{
		Name:        config.Name,
		DisplayName: config.DisplayName,
		Description: config.Description,
	}

	prg := &program{
		exit: make(chan struct{}),

		Config: config,
	}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "install" {
			err := s.Install()
			if err != nil{
				log.Fatal(err)
			}
			log.Fatal("Service install success")
		}
		if os.Args[1] == "uninstall" {
			err := s.Uninstall()
			if err != nil{
				log.Fatal(err)
			}
			log.Fatal("Service uninstall success")
		}
	}

	err = s.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

