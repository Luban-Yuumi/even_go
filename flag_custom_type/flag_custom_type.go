//实现一个解析并格式化命令行输入的时间集合的例子

package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
	"time"
)

type interval []time.Duration

var inter interval

//实现String接口
func (a *interval) String() string {
	return fmt.Sprintf("%v", *a)
}

//实现Set接口,Set接口决定了如何解析flag的值
func (a *interval) Set(value string) error {
	if len(*a) > 0 {
		return errors.New("error!")
	}
	for _, dt := range strings.Split(value, ",") {
		duration, err := time.ParseDuration(dt)
		if err != nil {
			return err
		}
		*a = append(*a, duration)
	}
	return nil
}

func init() {
	flag.Var(&inter, "time_str", "input time_str return duration array")
}

func main() {
	flag.Parse()
	fmt.Println(inter)
}
