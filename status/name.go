package status

//Status http://ecomgx17.ecomtoday.com/edi/EDI_4010/el306.htm
const (
	Draft		Code = "1"
	Deleted 	Code = "3"
	Verify		Code = "4"
	Cancelled	Code = "C"
	Rejected	Code = "U"
	MutuallyDefined Code = "Z"
	Approved 	Code = "11"
	Reissue		Code = "18"
	Disapprove	Code = "21"
	Hold		Code = "24"
	Adjusted	Code = "25"
	Paid		Code = "29"
	Complete	Code = "51"
	Assigned	Code = "76"
	Active		Code = "AE"
	Inactive 	Code = "IA"
	Discontinue	Code = "DT"
	Issue		Code = "IS"
	Processing	Code = "IN"
	InProgress	Code = "IT"
	Paying		Code = "PJ"
	Open		Code = "OP"
	Confirmed	Code = "CF"
	Converted	Code = "CV"
	Returned	Code = "RU"
	Pending		Code = "SU"
	Transferred	Code = "R6"
)