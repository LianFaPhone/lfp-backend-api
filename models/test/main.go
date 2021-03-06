package main

import "github.com/LianFaPhone/lfp-backend-api/models"
import (
	"fmt"
	"github.com/LianFaPhone/lfp-backend-api/tools"
	"time"
)

func main() {
	fmt.Println("start")
	config := tools.Config{}
	rule := &tools.Rule{
		UserNotify: []string{"call in ('ni','wo','ta')"},
		TextId:     []string{"t1", "call in ('ni','wo','ta')"},
	}
	config.Rule = *rule
	models.GlobalRuleMgr.Init(&config)
	input := make(map[string]interface{})
	input["name"] = "ni"
	input["call"] = "ni"
	models.GlobalRuleMgr.Match(input)

	fmt.Println("end")
	time.Sleep(time.Second * 2)
}
