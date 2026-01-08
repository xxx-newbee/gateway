package utils

func GenerateToken(userId int64, username string) (string, error) {
	// 这里简化处理，实际项目中应生成 JWT 或其他形式的 token
	return "mocked-token-for-" + username, nil
}