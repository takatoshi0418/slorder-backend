package controller

func getBetweenCondition(start string, end string) string {
	return "? BETWEEN " + start + " and IFNULL(" + end + ", '9999-12-31')"
}

type someoneUpdatedError struct{}
type noneMembersProject struct{}

func GetSomeoneUpdatedError() someoneUpdatedError {
	return someoneUpdatedError{}
}

func (someoneUpdatedError) Error() string {
	return "Someone has already updated this data"
}

func GetNoneMembersProject() noneMembersProject {
	return noneMembersProject{}
}

func (noneMembersProject) Error() string {
	return "Projects for which no members exist"
}
