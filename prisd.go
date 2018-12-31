package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
)

type configuration struct {
	Port       int              `yaml:"port"`
	Ip         string           `yaml:"ip,omitempty"`
	Prefix     string           `yaml:"prefix"`
	PrefixAlt  []string         `yaml:"prefix-alit"`
	Help       string           `yaml:"help-command"`
	Secret     string           `yaml:"secret"`
	LogLevel   string           `yaml:"loglevel"`
	LogFile    string           `yaml:"logfile"`
	Responders *responderConfig `yaml:"responders"`
	prefixLen  int
	helpRegex  *regexp.Regexp
}

var config configuration
var logger *zap.SugaredLogger

var version, build string

func init() {

	confFile := flag.String("conf", "", "Conf files, you know, conf files")
	showversion := flag.Bool("version", false, "show version and exit")
	dev := flag.Bool("dev", false, "dev mode (console friendly log, etc.)")
	loglevel := flat.String("loglevel", "", "log level override")

	flag.Parse()

	if *showversion {
		if version == "" {
			fmt.Println("Version: development")
		} else {
			fmt.Println("Version:", version)
		}

		if build == "" {
			fmt.Println("Build: development")
		} else {
			fmt.Println("Build:", build)
		}
		os.Exit(0)
	}

	var err error
	if dev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal("Unable to initialize zap logger: %v", err)
	}
	defer logger.Sync()

}
