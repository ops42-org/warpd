package main

import (
	"flag"

	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/ops42-org/warpd/pkg/deployer"
	"github.com/ops42-org/warpd/pkg/helm"
)

func main() {
	flagLogLevel := zap.LevelFlag("log-level", zapcore.InfoLevel, "Set log level. The following levels are available: 'debug', 'info', 'warn', 'error', 'panic' and 'fatal'")
	flagDevLog := flag.Bool("dev", false, "Enable development logger mode")
	flag.Parse()

	// Logger
	var loggerConfig zap.Config
	if *flagDevLog {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
	}
	loggerConfig.Level = zap.NewAtomicLevelAt(*flagLogLevel)
	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err.Error())
	}
	defer logger.Sync()
	log := zapr.NewLogger(logger)

	log.Info("Running warpd-deployer...")

	deployer := deployer.NewDeployer("./deploy", &log)
	deployer.InstallTools()
	deployer.DeployComponents()

	helm := helm.NewHelm(&log)
	helm.Version()

	/*
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Error(err, "Failed to read directory contents")
		} else {
			for _, f := range files {
				if f.IsDir() {
					log.Info("Dir: " + f.Name())
					helm.DepUp(f.Name() + "/chart")
				}
			}
		}
	*/
}
