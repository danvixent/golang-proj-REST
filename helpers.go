package main

import (
	"sort"
	"strconv"
	"strings"
)

func process(price *string, time *string) error {
	var err error

	ref := *price
	_, err = strconv.Atoi(ref)
	if err != nil {
		return err
	}
	ref = strings.Replace(ref, "-", ".", 1)

	ref = *time
	_, err = strconv.Atoi(ref)
	if err != nil {
		return err
	} else {
		ref = ref + "mins"
		return nil
	}
}

func newID() int {
	sort.Slice(food, func(i, j int) bool {
		return food[i].ID < food[j].ID
	})
	lastid := food[len(food)-1].ID
	return lastid + 1
}
