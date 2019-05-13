//+build linux darwin

package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

// Disklog logs to disk w/ this instance.  Logfile
// location is setup via the "logfile" flag
var Disklog = log.New()

// SetupLogfile will attempt to create a logfile if one does not exist
func SetupLogfile() {
	logfile := "/tmp/logfile"
	Disklog.SetFormatter(&log.TextFormatter{})
	Disklog.SetLevel(log.DebugLevel)
	Disklog.Out = os.Stdout
	Disklog.Info("Looking for logfile: ", logfile)
	if _, err := os.Stat(logfile); err == nil {
		Disklog.Info("Using existing file: ", logfile)
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		multi := io.MultiWriter(os.Stdout, file)

		if err != nil {
			log.Info("Failed to log to file, using default stdout")
		} else {
			Disklog.Out = multi
			Disklog.Info("Started program and now writing to ", logfile)
		}
	} else if os.IsNotExist(err) {
		log.Debug(logfile, " NOT found: ", err)
		file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		multi := io.MultiWriter(os.Stdout, file)

		if err != nil {
			log.Info("Failed to log to file, using default stderr")
		} else {
			Disklog.Out = multi
			Disklog.Infof("Started program and now writing to %s.", logfile)
		}
	} else {
		log.Warn("unable to set the logfile to", logfile)
	}
}
