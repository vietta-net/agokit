package sql

import (
	"agokit/helper"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"math"
	"strings"
)

//Dynamic filter with multiple field with single value
func BuildWhereFilter(resultOrm *gorm.DB, filter interface{}) ( *gorm.DB){
	var data = make(map[string]interface{})
	//Convert Struct to map
	helper.Convert(filter,  &data)

	var where = make(map[string]interface{})

	for field, value := range data {
		//Check field if it has value
		if value != "" {
			//Convert to lower case to map with db field name in lower case
			field = strings.ToLower(field)
			where[field] = value
		}
	}
	//Add where only it has conditions
	if len(where) > 0{
		resultOrm = resultOrm.Where(where)
	}

	return resultOrm
}

//Dynamic Search in multiple fields
func BuildWhereSearch(searchFields []string, keyword string) (string, []interface{}){
	var args   = []interface{}{}
	var whereString = ""
	//Nothing to do if searchFields and keyword are empty
	if len(searchFields) == 0 || keyword == "" {
		return whereString, args
	}

	//make where Or in (logic 1 OR logic 2 OR ...)
	var beginWhere = "("
	var endWhere 	= ")"
	var where  = make([]string, len(searchFields))

	//Search by contain a keyword
	key := "%"+ keyword + "%"
	for i, field := range searchFields {
		args = append(args, key)
		where[i] = fmt.Sprintf("%s LIKE ?", strings.ToLower(field) )
	}

	//Join a where string
	whereString = fmt.Sprintf("%s%s%s", beginWhere, strings.Join(where, " OR "),endWhere)
	return whereString, args
}

func GetPagination(resultOrm *gorm.DB, currentPage uint32, limit uint32) (interface{}, error){
	p := &Pagination{
		Count : 0,
		Pages: 0,
		Page: currentPage,
	}
	//Count Rows Total
	err := resultOrm.Count(&p.Count).Error
	if p.Count > 0{
		pageTotal := uint32(math.Ceil(float64(p.Count / limit)))
		if pageTotal < 1{
			p.Pages = 1
		}else{
			p.Pages = pageTotal
		}
	}
	return p, err
}