package _0_regular_expression_matching

// 回溯
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	fMatch := s != "" && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (fMatch && isMatch(s[1:], p))
	} else {
		return fMatch && isMatch(s[1:], p[1:])
	}
}

// 动态规划 dp
func isMatchDP(s string, p string) bool {
	// TODO:
	return false
}
