package cli

import (
	"fmt"
	"os"
	"text/tabwriter"
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

	const padding = 3
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.Debug)
	fmt.Fprintln(writer, "Project\t Task\t Hours\t Note")
	fmt.Fprintln(writer, "----\t ----\t ----\t ----")
	for _, entry := range logEntries {
		fmt.Fprintln(writer, fmt.Sprintf("%s\t %s\t %.1f\t %s", entry.Task.Name, entry.Project.Name, entry.Hours, entry.Note))
	}
	writer.Flush()
	fmt.Println()

	return nil
}
