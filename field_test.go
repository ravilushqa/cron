package main

import (
	"reflect"
	"testing"
)

func TestField_parse(t *testing.T) {
	tests := []struct {
		name      string
		field     field
		wantErr   bool
		wantValue []int
	}{
		{
			name: "parse fixed correct #1",
			field: field{
				name:   "test",
				base:   "*",
				ranges: [2]int{0, 5},
				values: nil,
			},
			wantErr:   false,
			wantValue: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "parse fixed correct #2",
			field: field{
				name:   "test",
				base:   "1",
				ranges: [2]int{0, 5},
				values: nil,
			},
			wantErr:   false,
			wantValue: []int{1},
		},
		{
			name: "parse fixed incorrect range #1",
			field: field{
				name:   "test",
				base:   "6",
				ranges: [2]int{0, 5},
				values: nil,
			},
			wantErr:   true,
			wantValue: nil,
		},
		{
			name: "parse range correct #1",
			field: field{
				name:   "test",
				base:   "1-3",
				ranges: [2]int{0, 5},
				values: nil,
			},
			wantErr:   false,
			wantValue: []int{1, 2, 3},
		},
		{
			name: "parse step correct #1",
			field: field{
				name:   "test",
				base:   "*/15",
				ranges: [2]int{0, 59},
				values: nil,
			},
			wantErr:   false,
			wantValue: []int{0, 15, 30, 45},
		},
		{
			name: "parse csv correct #1",
			field: field{
				name:   "test",
				base:   "1,15",
				ranges: [2]int{1, 31},
				values: nil,
			},
			wantErr:   false,
			wantValue: []int{1, 15},
		},
		{
			name: "parse csv incorrect range #1",
			field: field{
				name:   "test",
				base:   "1,15,32",
				ranges: [2]int{1, 31},
				values: nil,
			},
			wantErr:   true,
			wantValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &field{
				name:   tt.field.name,
				base:   tt.field.base,
				ranges: tt.field.ranges,
				values: tt.field.values,
			}
			if err := f.parse(); (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(f.values, tt.wantValue) {
				t.Errorf("wrong value = %v, wantValue %v", tt.field.values, tt.wantValue)
			}
		})
	}
}
