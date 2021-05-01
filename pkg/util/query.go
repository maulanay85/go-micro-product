package util

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func Pagination(pageString string, limitString string) (skip int32, limit int32) {
	page, _ := strconv.Atoi(pageString)
	tmp, _ := strconv.Atoi(limitString)
	var over int32
	if page < 1 {
		page = 1
	}
	if page == 1 {
		over = 0
	} else {
		over = int32((page - 1) * tmp)
	}
	return over, int32(tmp)
}

func Filter(query map[string]interface{}) bson.M {
	result := make(bson.M, len(query))
	for k, v := range query {
		if _, ok := v.(string); ok {
			result[k] = primitive.Regex{Pattern: fmt.Sprint(v), Options: "i"} //doesnt support case insensitive
		} else {
			result[k] = v
		}
	}
	return result
}

type Paging struct {
	Skip   int32
	Limit  int32
	Filter map[string]interface{}
}
