package main

/*
*

49. 字母异位词分组
中等
相关标签
相关企业
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。

示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

提示：

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] 仅包含小写字母
*/
func main() {

}

func groupAnagrams(strs []string) [][]string {
	// 特殊情况
	var res = make([][]string, 0, len(strs))
	if len(strs) == 0 {
		return res
	}
	if len(strs) == 1 {
		res = append(res, strs)
		return res
	}
	// 正常情况
	flags := make(map[string]string)
	for k, v := range strs {
		if _, ok := flags[v]; ok {
			continue
		}
		temp := make([]string, 0)
		temp = append(temp, v)
		flags[v] = v
		for i := k + 1; i < len(strs); i++ {
			if len(v) != len(strs[i]) {
				continue
			}
			if checkAnagrans(v, strs[i]) {
				flags[strs[i]] = strs[i]
				temp = append(temp, strs[i])
			}
		}
		if len(temp) > 0 {
			res = append(res, temp)
		}
	}
	return res
}

func checkAnagrans(s1, s2 string) bool {
	temp := make(map[string]int)
	for _, v2 := range s2 {
		if _, ok := temp[string(v2)]; ok {
			temp[string(v2)] += 1
		} else {
			temp[string(v2)] = 1
		}
	}

	for _, v := range s1 {
		count, ok := temp[string(v)]
		if !ok {
			return false
		} else {
			if count == 0 {
				return false
			}
			temp[string(v)] -= 1
		}
	}

	return true
}
