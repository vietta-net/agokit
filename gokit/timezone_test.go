package gokit_test

import (
	"github.com/vietta-net/agokit/gokit"
	"google.golang.org/grpc/metadata"
	"context"
	"testing"
	"github.com/stretchr/testify/assert"

)

var (
	ServerTimezone = "Asia/Ho_Chi_Minh"
)

func TestGetTimezoneFromContext(t *testing.T) {


	tz :=  "America/New_York"
	ctx := SetTimezone(tz)
	timezone := gokit.GetTimezoneFromContext(ctx)
	assert.Equal(t, tz, timezone)

	tz = ""
	ctx = SetTimezone(tz)
	timezone = gokit.GetTimezoneFromContext(ctx)
	assert.Equal(t, ServerTimezone, timezone)

}

func SetTimezone(timezine string)(ctx context.Context){
	ctx = context.Background()
	md := metadata.Pairs(
		"content-language", "vi",)
	if timezine !="" {
		md = metadata.Pairs("timezone", timezine)
	}
	ctx = metadata.NewOutgoingContext(ctx, md)

	// capital "Key" is illegal in HTTP/2.
	timezone, ok := md[gokit.TimezoneKey]
	tz := ServerTimezone
	if ok {
		tz = timezone[0]
	}

	ctx = context.WithValue(ctx, gokit.TimezoneKey, tz)
	return ctx
}