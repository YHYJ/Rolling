/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-15 10:24:11

Description: 子命令 'view' 的实现
*/

package cli

import (
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"
	"github.com/yhyj/rolling/general"
)

// SystemInfo 获取系统信息
func SystemInfo() {
	// 检索的 Pacman 日志文件
	var fileName = "/var/log/pacman.log"

	// 获取系统安装时间和当前时间
	lineText := general.ReadFileLine(fileName, 1)
	startTimeStrTZ := strings.Split(strings.Split(lineText, "[")[1], "]")[0] // 2023-03-10T10:49:09+0800
	currentTimeStr := time.Now().Format("2006-01-02 15:04")

	// 获取初始和当前内核版本
	keyText := general.ReadFileKey(fileName, "installed linux ")
	firstKernel := strings.Replace(strings.Split(strings.Split(keyText, " (")[1], ")")[0], ".", "-", -1)
	unameArgs := []string{"-r"}
	latestKernel, err := general.RunCommandGetResult("uname", unameArgs)
	if err != nil {
		color.Error.Println(err)
	}

	// 计算系统安装天数
	local, _ := time.LoadLocation("Asia/Shanghai")
	startTime, _ := time.ParseInLocation("2006-01-02T15:04:05Z0700", startTimeStrTZ, local)
	startTimeStr := startTime.Format("2006-01-02 15:04")
	startTimeStamp := startTime.Unix()
	currentTime, _ := time.ParseInLocation("2006-01-02 15:04", currentTimeStr, local)
	currentTimeStamp := currentTime.Unix()
	systemDays := int((currentTimeStamp - startTimeStamp) / 86400)

	// 获取系统/内核更新相关数据
	systemUpdateCount := general.ReadFileCount(fileName, "starting full system upgrade")
	systemUpdateMean := float32(systemUpdateCount) / float32(systemDays)
	kernelUpdateCount := general.ReadFileCount(fileName, "upgraded linux ")
	kernelUpdateMean := float32(systemDays) / float32(kernelUpdateCount)

	// 从“系统使用时长”和“系统更新次数”中选出最大值
	max := func() int {
		if systemDays > systemUpdateCount {
			return systemDays
		}
		return systemUpdateCount
	}()
	length := len(strconv.Itoa(max))

	// 获取吉祥物
	repoArgs := []string{""}
	mascot, err := general.RunCommandGetResult("repo-elephant", repoArgs)
	if err != nil {
		color.Error.Println(err)
	}

	// 输出
	titleFormat := "%27v %-2v %-27v\n"
	dataFormat := "%23v %-2v %-3v %v\n"
	color.Printf(titleFormat, general.FgCyan("[", startTimeStr, "]"), "--", general.FgCyan("[", currentTimeStr, "]"))
	color.Printf(titleFormat, general.FgMagenta(firstKernel), "--", general.FgMagenta(latestKernel))
	color.Printf(dataFormat, general.PrimaryText("系统使用时长"), "--", general.FgYellow(color.Sprintf("%-*.2v", length, systemDays)), general.SecondaryText("天"))
	color.Printf(dataFormat, general.PrimaryText("系统更新次数"), "--", general.FgYellow(color.Sprintf("%-*.2v", length, systemUpdateCount)), general.SecondaryText("次"))
	color.Printf(dataFormat, general.PrimaryText("系统更新频率"), "--", general.FgYellow(color.Sprintf("%-*.2v", length, systemUpdateMean)), general.SecondaryText("次/天"))
	color.Printf(dataFormat, general.PrimaryText("内核更新次数"), "--", general.FgYellow(color.Sprintf("%-*.2v", length, kernelUpdateCount)), general.SecondaryText("次"))
	color.Printf(dataFormat, general.PrimaryText("内核更新频率"), "--", general.FgYellow(color.Sprintf("%-*.2v", length, kernelUpdateMean)), general.SecondaryText("天/次"))
	color.Println(general.SuccessText(mascot))
}
