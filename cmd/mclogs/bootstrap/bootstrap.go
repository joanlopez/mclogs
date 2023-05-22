package bootstrap

import "github.com/joanlopez/mclogs"

func Run() error {
	cfg := configFromFlags()
	return mclogs.Run(cfg)
}
