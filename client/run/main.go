package run

import (
    "fmt"
    "github.com/LanceLRQ/deer-common/utils"
    "github.com/urfave/cli/v2"
    "os"
)


// 执行评测
func UserRunJudge(c *cli.Context) error {
    err := loadSystemConfiguration()
    if err != nil {
        return err
    }

    configFile, err := loadProblemConfiguration(c.Args().Get(0), c.String("work-dir"))
    if err != nil {
        return err
    }

    isBenchmarkMode := c.Int("benchmark") > 1
    if !isBenchmarkMode {
        // 普通的运行
        judgeResult, err := runUserJudge(c, configFile)
        if err != nil {
            return err
        }
        fmt.Println(utils.ObjectToJSONStringFormatted(judgeResult))
        os.Exit(judgeResult.JudgeResult)
    } else {
        // 基准测试
        err = runJudgeBenchmark(c, configFile)
        if err != nil {
            return err
        }
    }
    return nil
}