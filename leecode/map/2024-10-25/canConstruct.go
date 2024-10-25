package main

func main() {

}
func canConstruct(ransomNote string, magazine string) bool {
	var flags = make(map[string]int32)
	for i := 0; i < len(magazine); i++ {
		_, ok := flags[string(magazine[i])]
		if ok {
			flags[string(magazine[i])] += 1
		} else {
			flags[string(magazine[i])] = 1
		}

	}

	for i := 0; i < len(ransomNote); i++ {
		key := string(ransomNote[i])
		temp, ok := flags[key]
		if ok {
			if temp == 0 {
				return false
			}
			flags[key] -= 1
		} else {
			return false
		}

	}
	return true
}
