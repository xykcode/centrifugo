package callback

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
	"io/ioutil"
	"github.com/FZambia/viper-lite"
	"time"
)

type LeaveInfo struct {
	Channels []string	`json:"channels"`
	Uid 	string		`json:"uid"`
}

var LeaveCh chan LeaveInfo

func init()  {
	LeaveCh = make(chan LeaveInfo)


	//config.Getter(viper.GetViper())
	go callbackServer()
}

func callbackServer()  {
	var leaveurl string
	for true{
		c := viper.GetViper()
		leaveurl = c.GetString("callback_url_leave")
		if leaveurl != ""{
			break
		}
		time.Sleep(time.Second)
	}
	for true{
		select {
		case info := <- LeaveCh:
			content,_ := json.Marshal(info)

			resp, err := http.Post(leaveurl,
				"application/x-www-form-urlencoded",
				strings.NewReader(string(content)))
			if err != nil {
				fmt.Println(err)
			}

			defer resp.Body.Close()
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
		}
	}
}