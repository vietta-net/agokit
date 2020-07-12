package sql_test
import (
	"crypto/md5"
	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/vietta-net/agokit/pb"
	"github.com/vietta-net/agokit/sql"
	"testing"
	"github.com/stretchr/testify/assert"

	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/satori/go.uuid"
	"fmt"
)

type User struct {
	gorm.Model
	ID string `gorm:"primary_key;column:id;type:char;size:36;" json:"id"`
	Username string `gorm:"column:username;type:varchar;size:30;" json:"name"`
	Password string `gorm:"column:password;type:varchar;size:32;" json:"name"`
}

func TestBuildWhereDateRange(t *testing.T) {
	format := "2006-01-02"
	from, _ := time.Parse(format, "2010-07-13")
	to, _ := time.Parse(format, "2010-07-13")

	date := &pb.DateRange{
		From: &timestamppb.Timestamp{
			Seconds: from.Unix(),
			Nanos:   0,
		},
		To: &timestamppb.Timestamp{
			Seconds: to.Unix(),
			Nanos:   0,
		},
	}

	dataSource := "dev:passw0rd@tcp(127.0.0.1:3306)/dev?parseTime=true&charset=utf8mb4"
	db, err := gorm.Open("mysql", dataSource)
	assert.Nil(t, err)
	user := &User{
		ID: uuid.NewV1().String(),
		Username: "pntn",
		Password: fmt.Sprintf("%x", md5.Sum([]byte("pass"))),
	}

	db.FirstOrCreate(&user, user)

	resultOrm := db.Model(&user )
	timezone  := "Asia/Ho_Chi_Minh"
	resultOrm, err = sql.BuildWhereDateRange(resultOrm, date, timezone)
	assert.Nil(t, err)

}