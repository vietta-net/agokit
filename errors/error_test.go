package errors

import (
	"errors"
	"google.golang.org/grpc/codes"
	"testing"
	require "github.com/stretchr/testify/assert"

)

func TestError(t *testing.T) {
	t.Run("WrapErr", func(tt *testing.T) {
		e := WrapErr(errors.New("is it borked?"), "it is borked")
		require.Equal(tt, "it is borked: is it borked?", e.Error())
	})

	t.Run("E", func(tt *testing.T) {
		e := E("it is borked", errors.New("is it borked?"), 404)
		require.Equal(tt, "it is borked: is it borked?", e.Error())
		require.Equal(tt, uint32(404), e.(*Error).Code)

		e = E( codes.OK)
		require.Equal(tt, uint32(0), e.(*Error).Code)

		var errs = make(map[string]string)
		errs["Name"] = "Name should not empty!"
		e = E( codes.OK, errs)
		require.Equal(tt, uint32(0), e.(*Error).Code)
		require.Equal(tt, 1, len(e.(*Error).NestedErrors))
	})
}

