package helper

import (
	"github.com/vietta-net/agokit/pb"
	"encoding/json"
)

func Convert(from interface{}, to interface{}) (err error) {
	jData, err := json.Marshal(from)
	json.Unmarshal(jData, &to)
	return err
}

func ToDateRange(data interface{}) (*pb.DateRange, error){
	date := pb.DateRange{}
	err := Convert(data,  &date)
	return &date,err
}


