package cmd

import (
	"fmt"
	"it-test/config"

	"os"

	vipercfg "it-test/pkg/config"

	"github.com/go-pg/pg/v10"

	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var (
	logLevel string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "it-test",
	Short: "it-test for the win",
	Long:  `it-test cli gives possibility to start servers, migrate schemas based on the given parameters.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Dependencies struct {
	DB *pg.DB
}

func init() {
	serverCmd.PersistentFlags().StringVarP(
		&logLevel,
		"log-level",
		"l",
		"info",
		"Logging level. Default: info",
	)
}

func initLogLevel(cfg *config.AppConfig) {
	logrus.
		WithField("logLevel", logLevel).
		WithField("logLevelEnv", cfg.LogLevel).
		Info("Init logging")

	if cfg.LogLevel != "" {
		logLevel = cfg.LogLevel
	}
	parseLevel, err := logrus.ParseLevel(logLevel)
	check(err)

	logrus.SetLevel(parseLevel)
}

func configLoader() config.Loader {
	return &vipercfg.ViperConfig{}
}

func initGocore(cfg *config.AppConfig) {

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
