package main

import (
	"fmt"
	"io"
	"strings"
)

const (
	allSign        = '*'
	rangeSeparator = '-'
	stepsSeparator = '/'
	commaSeparator = ','
)

type cron struct {
	fields  []field
	command command
}

func newCron(cronTask []string) *cron {
	return &cron{
		fields: []field{
			{
				name:   "minute",
				base:   cronTask[0],
				ranges: [2]int{0, 59},
			},
			{
				name:   "hour",
				base:   cronTask[1],
				ranges: [2]int{0, 23},
			},
			{
				name:   "day of month",
				base:   cronTask[2],
				ranges: [2]int{1, 31},
			},
			{
				name:   "month",
				base:   cronTask[3],
				ranges: [2]int{1, 12},
			},
			{
				name:   "day of week",
				base:   cronTask[4],
				ranges: [2]int{1, 7},
			},
		},
		command: command{
			name: "command",
			base: strings.Join(cronTask[5:], " "),
		},
	}
}

func (c *cron) parseFields() error {
	for i := 0; i < len(c.fields); i++ {
		err := c.fields[i].parse()
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *cron) prettyPrint(w io.Writer) error {
	longestName := len(c.command.name)
	for _, v := range c.fields {
		if len(v.name) > longestName {
			longestName = len(v.name)
		}
	}

	for _, v := range c.fields {
		indent := strings.Repeat(" ", longestName-len(v.name))
		_, err := fmt.Fprintf(w, "%s %s %s\n", v.name, indent, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(v.values)), " "), "[]"))
		if err != nil {
			return err
		}
	}
	indent := strings.Repeat(" ", longestName-len(c.command.name))
	_, err := fmt.Fprintf(w, "%s %s %s", c.command.name, indent, c.command.base)
	if err != nil {
		return err
	}

	return nil
}
