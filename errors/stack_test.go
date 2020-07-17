package errors

import (
	"testing"

	require "github.com/stretchr/testify/assert"
)

func TestUnMarshal(t *testing.T) {
	var s = &Stack{}
	s.Callers = callers()
	one, _ := s.Marshal()
	s.Unmarshal(one)
	two, err := s.Marshal()
	require.Equal(t, one, two)
	require.Nil(t, err)
}

func TestMarshal(t *testing.T) {
	var s = &Stack{}
	s.Callers = callers()
	data, err := s.Marshal()
	require.NotNil(t, data)
	require.Nil(t, err)

}
