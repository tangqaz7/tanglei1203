package main

import (
	"MassX/cmd"
	"MassX/dao"
)

func main() {
	dao.MysqlInit()
	cmd.Entrance()
}
