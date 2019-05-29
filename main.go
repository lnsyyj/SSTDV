package main

import (
	"github.com/lnsyyj/SSTDV/common"
	"os"
	"bufio"
	"fmt"
	"flag"
)

func main() {
	var (
		visualizationType = flag.String("visualizationType", "", "Currently supports parsing: vdbench")
		logPath	= flag.String("logPath", "", "Specify the log file to be parsed, absolute or relative path")
		mysqlHostIP = flag.String("mysqlHostIP", "", "Specify mysql IP address")
		mysqlDatabase = flag.String("mysqlDatabase", "", "Specify the mysql database name")
		mysqlTableName = flag.String("mysqlTableName", "", "Specify mysql table name")
		mysqlUserName = flag.String("mysqlUserName", "", "Specify mysql username")
		mysqlUserPassword = flag.String("mysqlUserPassword", "", "Specify mysql password")
	)
	flag.Parse()
	if flag.NArg() == 0 {
		flag.PrintDefaults()
		return
	}
	fmt.Println(visualizationType, logPath, mysqlDatabase, mysqlTableName, mysqlHostIP, mysqlUserName, mysqlUserPassword)

	file, err := os.Open(*logPath)
	lineInfo := ""
	defer file.Close()
	common.Check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lineInfo = scanner.Text()
		fmt.Println(lineInfo)
	}
	err = scanner.Err()
	common.Check(err)
}