package common

import (
	cfg "github.com/Unknwon/goconfig"

)

var Cfg *cfg.ConfigFile

func SetConf()  {

	var err error
	Cfg, err = cfg.LoadConfigFile("config.ini")
	if err != nil {
		Cfg, err = cfg.LoadConfigFile("../config.ini")
	}
}