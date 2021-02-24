package utils

import (
	"fmt"
	"regexp"
	"time"
)

const logTimeLayout = `02/Jan/2006:15:04:05 -0700`

func LogTimeDiff(eachline string) (float64,error) {
	re2, err := regexp.Compile(`\[.*\]`) // error if regexp invalid
	if err != nil {
		return 0,fmt.Errorf("failed to compile regex 2: %v", err)
	}

	times := re2.FindAllString(eachline,1)

	if len(times) < 1 {
		return 0,fmt.Errorf("failed to find log time")
	}

	if len(times[0]) < 0 {
		return 0,fmt.Errorf("failed to find log time %v", err)
	}

	re3, err := regexp.Compile(`\[|\]`) // error if regexp invalid
	if err != nil {
		return 0,fmt.Errorf("failed to compile regex 3: %v", err)
	}

	logTimeClean := re3.ReplaceAllString(times[0], "")

	logTime,err := time.Parse(logTimeLayout,logTimeClean)
	if err != nil {
		return 0,fmt.Errorf("failed to parse time: %v", err)
	}

	diffTime := time.Since(logTime)

	return diffTime.Minutes(),nil
}
