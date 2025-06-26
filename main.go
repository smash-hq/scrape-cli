/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/smash-hq/scrape-cli/cmd"
	"os"
	"path"
	"runtime"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		DisableColors: false,
		ForceColors:   true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := path.Base(f.File)
			fc := path.Base(f.Function)
			return fmt.Sprintf("%s()", fc), fmt.Sprintf(" - %s:%d", filename, f.Line)
		},
		TimestampFormat: time.DateTime,
	})
	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)

}

func main() {

	cmd.Execute()
}
