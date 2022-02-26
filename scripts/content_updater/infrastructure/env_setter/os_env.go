package env_setter

import (
	"github.com/kelseyhightower/envconfig"
)

type HatenaEnv struct {
	Id       string
	Blog_id  string
	User_id  string
	Password string
}

func GetHatenaEnvFromOSEnv() HatenaEnv {
	var hatenaEnv HatenaEnv
	envconfig.Process("HATENA", &hatenaEnv)
	return hatenaEnv
}

type ZennEnv struct {
	User_id string
}

func GetZennEnvFromOSEnv() ZennEnv {
	var zennEnv ZennEnv
	envconfig.Process("ZENN", &zennEnv)
	return zennEnv
}
