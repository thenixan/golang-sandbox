package userdatabase

import (
	"errors"
	"sync"
	"time"
)

var Users = map[string]string {
	"that": "hjNwDTmw",
	"lazaret": "5jPhVcaE",
	"hospitable": "S42e4Thb",
	"amber": "h3CrFtv9",
	"wealth": "FrUgr8bE",
	"globe": "XNTQABRZ",
	"pied": "XNTQABRZ",
	"petulant": "43fnuqzk",
	"whoop": "2GBehYW9",
	"congratulations": "hWdS6EYb",
	"glands": "Pfwv6XrH",
	"peach": "ZUDnKZNW",
	"vela": "ZUDnKZNW",
	"muppet": "mXM6qy4v",
	"pascal": "Y5VnBcHg",
	"dial": "ESXbgvqY",
	"scalene": "gXfykv28",
	"mechanic": "GvPZzp83",
	"conducting": "Tz8ztghH",
	"thorough": "6gFrHAWs",
	"share": "6gFrHAWs",
	"dolls": "QfMCHnxm",
	"amazed": "YxprWcry",
	"parker": "374Q66Ld",
	"ocean": "ZHacuMcB",
	"lively": "EHcqFmdm",
	"icecream": "y4WFcM6n",
	"morello": "6jhG66gn",
	"length": "N7vukEXp",
	"dozens": "JzedFe5q",
	"compiler": "QaBjyAFV",
}

var mutex = &sync.Mutex{}

func CheckAuth(login string, password string) (err error)  {
	mutex.Lock()
	time.Sleep(1 * time.Second)
	mutex.Unlock()
	if Users[login] == password {
		return nil
	} else {
		return errors.New("wrong auth")
	}
}

func ChangePassword(login string, password string, newPassword string) (err error)  {
	authErr := CheckAuth(login, password)
	if authErr == nil {
		mutex.Lock()
		time.Sleep(1 * time.Second)
		mutex.Unlock()
		Users[login] = newPassword
		return nil
	} else {
		return authErr
	}
}