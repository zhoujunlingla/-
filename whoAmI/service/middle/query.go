package middle

type quesWithAType struct {
	Filename string   `json:"filename"`
	Question []string `json:"question"`
	Answer   int      `json:"answer"`
}

// 补充 生成带答案的问题 功能的相关代码
func GetQuestionWithA() quesWithAType {
	var question quesWithAType

	return question
}
