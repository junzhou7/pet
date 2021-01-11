package utils

// AllLike ...
func AllLike(condition string) string {
	return "%" + condition + "%"
}

// LeftLike ..
func LeftLike(condition string) string {
	return "%" + condition
}

// RightLike ..
func RightLike(condition string) string {
	return condition + "%"
}
