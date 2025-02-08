/**
 * @Author: chentong
 * @Date: 2025/02/08 22:01
 */

package handler

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ch3nnn/webstack-go/pkg/config"
	"github.com/ch3nnn/webstack-go/pkg/log"
)

func Test_Handler(t *testing.T) {
	err := os.Setenv("APP_CONF", "../../config/local.yml")
	if err != nil {
		fmt.Println("Setenv error", err)
	}
	envConf := flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logPath := filepath.Join("../../../", conf.GetString("log.log_file_name"))
	conf.Set("log.log_file_name", logPath)

	NewHandler(log.NewLog(conf))
}
