package externalAPI

import "github.com/kelseyhightower/envconfig"

type HatenaEnv struct {
	Id        string
	Blog_id   string
	User_name string
	Password  string
}

func GetHatenaEnvFromGithub() HatenaEnv {
	var hatenaEnv HatenaEnv
	envconfig.Process("HATENA", &hatenaEnv)
	return hatenaEnv
}
