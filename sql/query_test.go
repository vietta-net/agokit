package sql_test
import (
	"crypto/md5"
	"github.com/guregu/null"
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
	ID string `gorm:"primary_key;column:id;type:char;size:36;" json:"id"`
	Username string `gorm:"column:username;type:varchar;size:30;" json:"name"`
	Password string `gorm:"column:password;type:varchar;size:32;" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
	DeletedAt null.Time `gorm:"column:deleted_at;type:datetime;" json:"deleted_at"`
}

func TestBuildWhereDateRange(t *testing.T) {
	dataSource := "dev:passw0rd@tcp(127.0.0.1:3306)/dev?parseTime=true&charset=utf8mb4"
	db, err := gorm.Open("mysql", dataSource)
	assert.Nil(t, err)
	user := &User{
		ID: uuid.NewV1().String(),
		Username: "pntn",
		Password: fmt.Sprintf("%x", md5.Sum([]byte("pass"))),
	}

	db.FirstOrCreate(&user, user)


	format := "2006-01-02"
	from, _ := time.Parse(format, "2010-07-13")

	to := user.CreatedAt.Add(time.Hour * 24)

	query := pb.Query{
		Dates: []*pb.DateRange{
			{
				From: &timestamppb.Timestamp{
					Seconds: from.Unix(),
					Nanos:   0,
				},
				To: &timestamppb.Timestamp{
					Seconds: to.Unix(),
					Nanos:   0,
				},
			},
			{
				From: &timestamppb.Timestamp{
					Seconds: from.Unix(),
					Nanos:   0,
				},
				To: &timestamppb.Timestamp{
					Seconds: to.Unix(),
					Nanos:   0,
				},
				Field: "updated_at",
			},
		},
	}

	resultOrm := db.Model(&user )
	timezone  := "Asia/Ho_Chi_Minh"

	resultOrm.Where("id = ?", user.ID)
	resultOrm, err = sql.BuildWhereDateRanges(resultOrm, query.Dates, timezone)
	assert.Nil(t, err)

	var results = []User{}
	resultOrm.Find(&results)

	assert.Equal(t, 1, len(results))

}