package model

import (
	"errors"
	"fmt"
)

type Filter struct {
	OrderBy string `schema:"$orderBy"`
	Limit   int    `schema:"$top"`
	Offset  int    `schema:"$offset"`
}

var minLimit = 0
var maxLimit = 20
var minOffset = 0

func (f Filter) Validate() error {
	if f.Limit < minLimit || f.Limit > maxLimit {
		return fmt.Errorf("Limit must be between %d and %d", minLimit, maxLimit)
	}
	if f.Offset < minOffset {
		return errors.New("Offset must be greater or equal to zero")
	}

	return nil
}
