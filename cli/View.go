package cli

import (
	"fmt"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/stuart-xyz/harvest-cli/services"
)

func View(config services.Config, opts docopt.Opts) (err error) {
	isRequestForYesterday, err := opts.Bool("--yesterday")
	if err != nil {
		return err
	}

	var date string
	if isRequestForYesterday {
		date = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	} else {
		date = time.Now().Format("2006-01-02")
	}

	logEntries, err := services.ViewLog(config, date)
	if err != nil {
		return err
	}

	for _, entry := range logEntries {
		fmt.Println(entry)
	}

	return nil
}
