package query_params

import (
	"net/http"
	"strconv"
)

func GetFilter(r *http.Request) (limit, offset int32) {
	var l int64 = 10
	var o int64 = 0

	lp := r.URL.Query().Get("limit")
	if len(lp) > 0 {
		val, err := strconv.ParseInt(lp, 10, 64)
		if err == nil {
			l = val
		}
	}

	op := r.URL.Query().Get("offset")
	if len(op) > 0 {
		val, err := strconv.ParseInt(op, 10, 64)
		if err == nil {
			o = val
		}
	}

	return int32(l), int32(o)
}

func GetId(name string, r *http.Request) (id int64) {
	var result int64 = 0

	p := r.URL.Query().Get(name)
	if len(p) > 0 {
		val, err := strconv.ParseInt(p, 10, 64)
		if err == nil {
			result = val
		}
	}

	return result
}
