package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type command struct {
	name string
	base string
}

type field struct {
	name   string
	base   string
	ranges [2]int
	values []int
}

func (f *field) parse() error {
	if len(f.base) == 0 {
		return errWrongArgs
	}

	if f.ranges[0] >= f.ranges[1] {
		return errors.New("wrong ranges")
	}

	//fixed
	if len(f.base) == 1 {
		switch f.base {
		case string(allSign):
			valuesCount := f.ranges[1] - f.ranges[0] + 1
			res := make([]int, 0, valuesCount)
			for i := f.ranges[0]; i <= f.ranges[1]; i++ {
				res = append(res, i)
			}
			f.values = res
			return nil
		default:
			digit, err := strconv.Atoi(f.base)
			if err != nil {
				return fmt.Errorf("unexpected value %s: %w", f.base, err)
			}
			if digit < f.ranges[0] || digit > f.ranges[1] {
				return fmt.Errorf("value not in range, value: %d, ranges: %v", digit, f.ranges)
			}
			f.values = []int{digit}
			return nil
		}
	}

	//range
	if f.base[1] == rangeSeparator {
		ranges := strings.Split(f.base, string(rangeSeparator))
		if len(ranges) != 2 {
			return fmt.Errorf("unexpected range value: %s", f.base)
		}
		lower, err := strconv.Atoi(ranges[0])
		if err != nil {
			return fmt.Errorf("unexpected value %s: %w", f.base, err)
		}
		upper, err := strconv.Atoi(ranges[1])
		if err != nil {
			return fmt.Errorf("unexpected value %s: %w", f.base, err)
		}
		if lower > upper {
			return fmt.Errorf("range first value greater than second:  %s", f.base)
		}
		if lower < f.ranges[0] || upper > f.ranges[1] {
			return fmt.Errorf("value not in range, value: %v, ranges: %v", f.base, f.ranges)
		}

		valuesCount := upper - lower + 1
		res := make([]int, 0, valuesCount)
		for i := lower; i <= upper; i++ {
			res = append(res, i)
		}
		f.values = res
		return nil
	}

	// steps
	if f.base[1] == stepsSeparator {
		steps := strings.Split(f.base, string(stepsSeparator))
		if len(steps) != 2 {
			return fmt.Errorf("unexpected steps value: %s", f.base)
		}

		step, err := strconv.Atoi(steps[1])
		if err != nil {
			return fmt.Errorf("unexpected steps value %s: %w", f.base, err)
		}

		if step < f.ranges[0] || step > f.ranges[1] {
			return fmt.Errorf("value not in range, value: %v, ranges: %v", f.base, f.ranges)
		}

		valuesCount := (f.ranges[1] - f.ranges[0] + 1) / step
		res := make([]int, 0, valuesCount)
		for i := f.ranges[0]; i <= f.ranges[1]; {
			res = append(res, i)
			i += step
		}
		f.values = res
		return nil
	}

	//csv
	csv := strings.Split(f.base, string(commaSeparator))
	if len(csv) < 2 {
		return fmt.Errorf("unexpected comma separated values: %s", f.base)
	}
	valuesCount := len(csv)
	res := make([]int, 0, valuesCount)

	for _, v := range csv {
		num, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("unexpected value %s: %w", f.base, err)
		}

		res = append(res, num)
	}

	sort.Ints(res)
	if res[0] < f.ranges[0] || res[len(res)-1] > f.ranges[1] {
		return fmt.Errorf("value not in range, value: %v, ranges: %v", f.base, f.ranges)
	}

	f.values = res
	return nil
}
