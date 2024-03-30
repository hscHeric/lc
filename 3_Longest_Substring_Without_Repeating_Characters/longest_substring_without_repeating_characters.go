package longestsubstringwithoutrepeatingcharacters

import "strings"

func lengthOfLongestSubstring(s string) int {
	var ans int
	sub := make([]string, 0)
	for _, c := range s {
		if !strings.Contains(strings.Join(sub, ""), string(c)) {
			sub = append(sub, string(c))
			ans = max(ans, len(sub))
		} else {
			// Encontra o index do primeiro caractere repetido
			cut := strings.Index(strings.Join(sub, ""), string(c))
			sub = append(sub[cut+1:], string(c))
		}
	}
	return ans
}
