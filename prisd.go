package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type configuration struct {
	Port      int      `yaml:"port"`
	Prefix    string   `yaml:"prefix"`
	PrefixAlt []string `yaml:"prefix-alit"`
	Help      string   `yaml:"help-command"`
	Secret    string   `yaml:"secret"`
	LogLevel  string   `yaml:"loglevel"`
	LogFile   string   `yaml:"logfile"`
	prefixLen int
}

var config configuration
var version, build string

func init() {

	confFile := flag.String("conf", "", "Conf files, you know, conf files")
	dev := flag.Bool("dev", false, "invoke dev mode (prettier log, etc.)")
	showversion := flag.Bool("version", false, "show version and exit")

	var loglevel string
	flag.StringVar(&loglevel, "loglevel", "", "log level override")

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

	zerolog.TimeFieldFormat = ""

	if *dev {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if *confFile == "" {
		log.Fatal().Msg("Config file not specified")
	}

	confRaw, err := ioutil.ReadFile(*confFile)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error reading config file")
	}
	log.Debug().Msgf("Config file read: %s", *confFile)

	err = yaml.Unmarshal(confRaw, &config)

	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Error parsing config file")
	}
	log.Debug().Msg("Config file loaded")

	if loglevel == "" {
		if config.LogLevel == "" {
			loglevel = "info"
		}
	}
	log.Debug().Msgf("Switching log level to: %s", loglevel)

	switch strings.ToLower(loglevel) {
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Warn().Msgf("Unknown loglevel %s, default to INFO", loglevel)
	}

}

func main() {
	fmt.Println("Started")
}
