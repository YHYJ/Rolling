/*
File: version.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-21 11:42:27

Description: 由程序子命令 version 执行
*/

package general

import (
	"fmt"
	"strconv"
	"time"
)

const (
	Name    string = "Rolling"                 // 程序名
	Version string = "v0.5.6"                  // 程序版本
	Project string = "github.com/yhyj/rolling" // 项目地址
)

var (
	GitCommitHash string = "Unknown" // Git 提交 Hash
	BuildTime     string = "Unknown" // 编译时间
	BuildBy       string = "Unknown" // 编译者
)

// ProgramInfo 返回程序信息
//
// 参数：
//   - only: 是否只返回程序版本
//
// 返回：
//   - 程序信息
func ProgramInfo(only bool) string {
	programInfo := fmt.Sprintf("%s\n", Version)
	if !only {
		BuildTimeConverted := "Unknown"
		timestamp, err := strconv.ParseInt(BuildTime, 10, 64)
		if err == nil {
			BuildTimeConverted = time.Unix(timestamp, 0).String()
		}
		programInfo = fmt.Sprintf("%s %s - Build rev %s\nBuilt on: %s\nBuilt by: %s\n", Name, Version, GitCommitHash, BuildTimeConverted, BuildBy)
	}
	return programInfo
}
