package status

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, Draft.String(), "1")
}

func TestString(t *testing.T) {
	assert.Equal(t, Draft.Equal("1"), true)
}

func TestName(t *testing.T) {
	assert.Equal(t, "Draft"		, Code("1").Name())
	assert.Equal(t, "Deleted"		, Code("3").Name())
	assert.Equal(t, "Verify"		, Code("4").Name())
	assert.Equal(t, "Cancelled"	, Code("C").Name())
	assert.Equal(t, "Rejected"		, Code("U").Name())
	assert.Equal(t, "Approved" 	, Code("11").Name())
	assert.Equal(t, "Reissue"		, Code("18").Name())
	assert.Equal(t, "Disapprove"	, Code("21").Name())
	assert.Equal(t, "Hold"			, Code("24").Name())
	assert.Equal(t, "Adjusted"		, Code("25").Name())
	assert.Equal(t, "Paid"			, Code("29").Name())
	assert.Equal(t, "Complete"		, Code("51").Name())
	assert.Equal(t, "Assigned"		, Code("76").Name())
	assert.Equal(t, "Active"		, Code("AE").Name())
	assert.Equal(t, "Inactive" 	, Code("IA").Name())
	assert.Equal(t, "Discontinue"	, Code("DT").Name())
	assert.Equal(t, "Issue"		, Code("IS").Name())
	assert.Equal(t, "Processing"	, Code("IN").Name())
	assert.Equal(t, "In Progress"	, Code("IT").Name())
	assert.Equal(t, "Paying"		, Code("PJ").Name())
	assert.Equal(t, "Open"			, Code("OP").Name())
	assert.Equal(t, "Confirmed"	, Code("CF").Name())
	assert.Equal(t, "Converted"	, Code("CV").Name())
	assert.Equal(t, "Returned"		, Code("RU").Name())
	assert.Equal(t, "Pending"		, Code("SU").Name())
	assert.Equal(t, "Transferred"	, Code("R6").Name())
	assert.Equal(t, "Mutually Defined"	, Code("Z").Name())
}

func TestCode(t *testing.T) {
	assert.Equal(t, Draft		, Code("1"))
	assert.Equal(t, Deleted		, Code("3"))
	assert.Equal(t, Verify		, Code("4"))
	assert.Equal(t, Cancelled	, Code("C"))
	assert.Equal(t, Rejected	, Code("U"))
	assert.Equal(t, Approved 	, Code("11"))
	assert.Equal(t, Reissue		, Code("18"))
	assert.Equal(t, Disapprove	, Code("21"))
	assert.Equal(t, Hold		, Code("24"))
	assert.Equal(t, Adjusted	, Code("25"))
	assert.Equal(t, Paid		, Code("29"))
	assert.Equal(t, Complete	, Code("51"))
	assert.Equal(t, Assigned	, Code("76"))
	assert.Equal(t, Active		, Code("AE"))
	assert.Equal(t, Inactive 	, Code("IA"))
	assert.Equal(t, Discontinue	, Code("DT"))
	assert.Equal(t, Issue		, Code("IS"))
	assert.Equal(t, Processing	, Code("IN"))
	assert.Equal(t, InProgress	, Code("IT"))
	assert.Equal(t, Paying		, Code("PJ"))
	assert.Equal(t, Open		, Code("OP"))
	assert.Equal(t, Confirmed	, Code("CF"))
	assert.Equal(t, Converted	, Code("CV"))
	assert.Equal(t, Returned	, Code("RU"))
	assert.Equal(t, Pending		, Code("SU"))
	assert.Equal(t, Transferred	, Code("R6"))
	assert.Equal(t, MutuallyDefined	, Code("Z"))
}