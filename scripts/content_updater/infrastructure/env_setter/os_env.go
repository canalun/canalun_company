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

func GetHatenaEnvFromOSEnv() (HatenaEnv, error) {
	var hatenaEnv HatenaEnv
	if err := envconfig.Process("HATENA", &hatenaEnv); err != nil {
		return HatenaEnv{}, err
	}
	return hatenaEnv, nil
}

type ZennEnv struct {
	User_id string
}

func GetZennEnvFromOSEnv() (ZennEnv, error) {
	var zennEnv ZennEnv
	if err := envconfig.Process("ZENN", &zennEnv); err != nil {
		return ZennEnv{}, err
	}
	return zennEnv, nil
}
