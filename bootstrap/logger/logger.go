package logger

import (
	"backend/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

var Default *logrus.Logger

// Init 根据配置文件初始化log
func Init(vconf *viper.Viper) error {
	folder := vconf.GetString("log.folder")
	if folder == "" {
		folder = util.LogFolder()
	}
	Default = logrus.New()
	logFile := path.Join(folder, "hade.log")
	fd, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	Default.Out = fd
	return nil
}
