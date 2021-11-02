package main

import (
	"bytes"
	"testing"
)

func Test_cron_prettyPrint(t *testing.T) {
	type fields struct {
		fields  []field
		command command
	}
	tests := []struct {
		name       string
		fields     fields
		wantErr    bool
		wantOutput string
	}{
		{
			name: "example print",
			fields: fields{
				fields: []field{
					{
						name:   "minute",
						values: []int{0, 15, 30, 45},
					},
					{
						name:   "hour",
						values: []int{0},
					},
					{
						name:   "day of month",
						values: []int{1, 15},
					},
					{
						name:   "month",
						values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
					},
					{
						name:   "day of week",
						values: []int{1, 2, 3, 4, 5},
					},
				},
				command: command{
					name: "command",
					base: "/usr/bin/find",
				},
			},
			wantErr: false,
			wantOutput: `minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer

			a := &cron{
				fields:  tt.fields.fields,
				command: tt.fields.command,
			}

			if err := a.prettyPrint(&b); (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if b.String() != tt.wantOutput {
				t.Errorf("wrong output = %v, wantOutput %v", b.String(), tt.wantOutput)
			}
		})
	}
}
