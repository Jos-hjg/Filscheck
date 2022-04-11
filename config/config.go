package config

import (
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
	"os"
)

type Config struct {
	Debug  bool         `yaml:"debug,omitempty" default:"false" `
	Logger LoggerConfig `yaml:"logger,omitempty"`
	Router RouterInfo `yaml:"router"`
	Filscout   FilscoutData   `yaml:"filscout,omitempty"`
}

type LoggerConfig struct {
	Level string `yaml:"level,omitempty" default:"debug"`
	// json or text
	Format string `yaml:"format,omitempty" default:"json"`
	// file
	Output string `yaml:"output,omitempty" default:""`
}

type RouterInfo struct {
	Port string `yaml:"port" default:"8082"`
}

type FilscoutData struct {
	Verbose       bool   `yaml:"verbose" default:"false"`
	//TODO: any other config?
	Scheme string `yaml:"scheme"`
	Host string `yaml:"host"`
	URI string `yaml:"uri"`
	Miners []string `yaml:"miners"`
}

var C *Config = &Config{}

func initLogger(debug bool, level, output string) {
	var conf zap.Config
	if debug {
		conf = zap.NewDevelopmentConfig()
	} else {
		conf = zap.NewProductionConfig()
	}

	var zapLevel = zap.NewAtomicLevel()
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zap.L().Panic("set logger level fail",
			zap.Strings("only", []string{"debug", "info", "warn", "error", "dpanic", "panic", "fatal"}),
			zap.Error(err),
		)
	}

	conf.Level = zapLevel
	conf.Encoding = "json"

	if output != "" {
		conf.OutputPaths = []string{output}
		conf.ErrorOutputPaths = []string{output}
	}

	logger, _ := conf.Build()

	zap.RedirectStdLog(logger)
	zap.ReplaceGlobals(logger)
}

func Init(cfgFile string) {
	_ = os.Setenv("CONFIGOR_ENV_PREFIX", "-")
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	if cfgFile != "" {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C, cfgFile); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	} else {
		if err := configor.New(&configor.Config{AutoReload: true}).Load(C); err != nil {
			zap.L().Panic("init config fail", zap.Error(err))
		}
	}

	initLogger(C.Debug, C.Logger.Level, C.Logger.Output)
	zap.L().Debug("loaded config")
}
