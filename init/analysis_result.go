package init

import (
	"fmt"

	"github.com/LanceLRQ/deer-common/constants"
	commonStructs "github.com/LanceLRQ/deer-common/structs"
)

func AnalysisResult(caseName string, result *commonStructs.JudgeResult) error {
	name, ok := constants.FlagMeansMap[result.JudgeResult]
	if !ok {
		name = "Unknown"
	}
	if result.JudgeResult != 0 {
		//return error name
		return nil
	}
	fmt.Println(name)

	//return ac
	return nil
}
