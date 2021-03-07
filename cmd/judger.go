package cmd

import (
	"github.com/LanceLRQ/deer-common/constants"
	"github.com/LanceLRQ/deer-executor/v2/init"
)

func Judger() error {
	err := init.InitWorkRoot()
	if err != nil {
		return err
	}
	result, err := init.RunJudge("./data/problems/APlusB/problem.json", "./data/codes/APlusB/ac.c", "")
	if err != nil {
		return err
	}
	err = init.AnalysisResult("case 1", result, constants.JudgeFlagAC)
	if err != nil {
		return err
	}
	return nil
}
