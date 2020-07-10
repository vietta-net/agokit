package status

type Code string

func (c Code) String() string {
	return string(c)
}

func (c Code) Equal(s string) bool {
	return string(c) == s
}

func (c Code) Name() string {
	switch c {
	case "1":
		return "Draft"
	case "3":
		return "Deleted"
	case "4":
		return "Verify"
	case "C":
		return "Cancelled"
	case "U":
		return "Rejected"
	case "11":
		return "Approved"
	case "18":
		return "Reissue"
	case "21":
		return "Disapprove"
	case "24":
		return "Hold"
	case "25":
		return "Adjusted"
	case "29":
		return "Paid"
	case "51":
		return "Complete"
	case "76":
		return "Assigned"
	case "AE":
		return "Active"
	case  "IA":
		return "Inactive"
	case "DT":
		return "Discontinue"
	case "IS":
		return "Issue"
	case "IN":
		return "Processing"
	case "IT":
		return "In Progress"
	case "PJ":
		return "Paying"
	case "OP":
		return "Open"
	case "CF":
		return "Confirmed"
	case "CV":
		return "Converted"
	case "RU":
		return "Returned"
	case "SU":
		return "Pending"
	case "R6":
		return "Transferred"
	default:
		return "Mutually Defined"
	}
}

