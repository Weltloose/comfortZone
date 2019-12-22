package model

func AddComment(user, content string) bool {
	return createComment(user, content)
}

func FetchComment(curLen int) Comments {
	return queryComment(curLen)
}
