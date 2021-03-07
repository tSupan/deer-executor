package cmd

import (
	"github.com/tSupan/deer-executor/init"
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
	err = init.AnalysisResult(result)
	if err != nil {
		return err
	}
	return nil
}
