package difficult

import "fmt"

type quesType struct {
	QuestionId string   `json:"id"`
	Filename   string   `json:"filename"`
	Question   []string `json:"question"`
}

// 补充 生成带答案的问题 功能的相关代码
func GetQuestion() quesType {
	var question quesType

	return question
}

// 补充 验证选择是否问题 功能的相关代码。
//
// 【输入】
// 函数的输入参数为前端请求中携带的相关信息，可自行选择需要的参数
//
// @id：任务id（取决于GetQuestion中是否返回QuestionId）
// @filename：文件名
// @chosenId：选择序号（eg: 3）
// @chosenText：选择选项（eg: "D"）
// @chosenName：选择文本（eg: "火影忍者"）
//
// 【输出】
// 若选择正确，返回True，否则返回False
func Verify(id string, filename string, chosenId int, chosenText string, chosenName string) bool {
	fmt.Println(id, filename, chosenId, chosenText, chosenName)

	return false
}
