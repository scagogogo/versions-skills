package cli

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// captureStdout 将 os.Stdout 重定向到 pipe，返回 restore 与读取函数。
// 调用者必须在结束时显式调用 restore()。
func captureStdout(t *testing.T) (restore func(), getOutput func() string) {
	t.Helper()
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("os.Pipe: %v", err)
	}
	os.Stdout = w

	done := make(chan struct{})
	var buf bytes.Buffer
	go func() {
		_, _ = io.Copy(&buf, r)
		close(done)
	}()

	restore = func() {
		os.Stdout = old
		_ = w.Close()
		<-done
	}
	getOutput = func() string {
		return buf.String()
	}
	return
}

// withFormat 在测试期间临时设置 format，并在结束时还原。
func withFormat(t *testing.T, f string) {
	t.Helper()
	oldF := format
	format = f
	t.Cleanup(func() { format = oldF })
}

// withQuiet 在测试期间临时设置 quiet，并在结束时还原。
func withQuiet(t *testing.T, q bool) {
	t.Helper()
	oldQ := quiet
	quiet = q
	t.Cleanup(func() { quiet = oldQ })
}

// resetRootArgs 在测试结束还原 rootCmd 的 args，避免污染其它测试。
func resetRootArgs(t *testing.T) {
	t.Helper()
	t.Cleanup(func() {
		rootCmd.SetArgs([]string{})
	})
}

// resetCmdFlags 将所有命令级 flag 变量重置为默认值。
// pflag 在重复 Parse 之间不会重置已设的 flag 值，因此跨测试需手动还原。
func resetCmdFlags() {
	// root
	format = "json"
	quiet = false
	// parse
	parseDelimiters = ""
	// sort
	sortDesc = false
	sortFromFile = ""
	// group
	groupID = ""
	groupFromFile = ""
	// filter
	filterStable = false
	filterPrerelease = false
	filterMajor = ""
	filterMinor = ""
	filterPatch = ""
	filterPrefix = ""
	filterSuffix = ""
	filterConstraint = ""
	filterConstraintType = "set"
	filterFromFile = ""
	// constraint
	constraintType = "set"
	// range
	rangeIncludeStart = true
	rangeIncludeEnd = false
	rangeFromFile = ""
	// build
	buildPrefix = ""
	buildMajor = ""
	buildMinor = ""
	buildPatch = ""
	buildSuffix = ""
	buildNumbers = ""
	// bump
	bumpMajor = false
	bumpMinor = false
	bumpPatch = false
	// check
	checkPrerelease = false
	checkStable = false
	checkDev = false
	checkAlpha = false
	checkBeta = false
	checkRC = false
	checkSnapshot = false
	checkMilestone = false
	checkNightly = false
	checkFinal = false
	checkGA = false
	checkPre = false
	checkRelease = false
	checkSP = false
	checkPost = false
	checkZero = false
	checkNewer = ""
	checkOlder = ""
	checkEqual = ""
	checkBetweenLow = ""
	checkBetweenHigh = ""
	// count
	countStable = false
	countPrerelease = false
	countMajor = ""
	countMinor = ""
	countPatch = ""
	countFromFile = ""
	// fileio
	writeFromFile = ""
	// minmax
	minmaxFromFile = ""
	// partition
	partitionStable = false
	partitionPrerelease = false
	partitionFromFile = ""
	// sort_strings
	sortStringsDesc = false
	sortStringsFromFile = ""
	// visualize
	visualizeMaxItems = 0
	visualizeGroups = false
	visualizeFromFile = ""
	// group_extra
	groupExtraFromFile = ""
	groupExtraGroupID = ""
	groupExtraVersion = ""
}
