package main

import (
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

// Checks slice of int if it contains int i, return True of False
func ContainsInt(slice []int, i int) bool {
	res := false
	for _, num := range slice {
		if num == i {
			res = true
			break
		}
	}
	return res
}
func ContainsInt32(slice []int32, i int32) bool {
	res := false
	for _, num := range slice {
		if num == i {
			res = true
			break
		}
	}
	return res
}

func int32ToInt4(i int32) pgtype.Int4 {
	var p pgtype.Int4
	p.Int32 = i
	p.Valid = true
	return p
}

func strToInt32(s string) (int, error) {
	return strconv.Atoi(s)
}
func strToInt64(s string) (int64, error) {
	i, e := strToInt32(s)
	return int64(i), e
}
