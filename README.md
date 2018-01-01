# go-hbase-rest
### hbase rest api接口链接管理【golang语言版】
---
### go get github.com/qizexi/go-hbase-rest
---
		关于hbase的rest接口的详细信息可以到官网查看[http://hbase.apache.org/book.html#_rest]
		测试环境：ubuntu16.04+hadoop2.4.0+hbase1.2.6+jdk1.8.0_141
		联系作者：qizexi@163.com
		快速体验：
		package main
		import (
			"fmt"
			"github.com/qizexi/dhbase/rest"
		)
		func main() {
			//初始化rest请求,ip地址和端口
			rt := rest.NewRest("localhost", 9099)
			//查看版本信息
			fmt.Println(rt.Version())
			//创建一个表:表名，列簇名
			//fmt.Println(rt.Create("mydemo", []string{"cf"}))
			//列出所有表的信息
			fmt.Println(rt.List())
			//修改表的信息:表名，列簇名
			//fmt.Println(rt.Alter("mydemo", []string{"cf"}))
			//删除表:表名
			//fmt.Println(rt.Drop("mydemo"))
			//添加一条记录
			//fmt.Println(rt.Put("mydemo", "row1", "cf:a", "I am first value"))
			//获取一条记录：表名，行键值，列值，时间戳，版本号
			fmt.Println(rt.Get("mydemo", "row1", "", "", ""))
			//遍历表的记录
			vid, err := rt.Scanner("mydemo")
			if vid != "" && err == nil {
				for {
					rs, err := rt.Scan("mydemo", vid)
					fmt.Println("scan", rs, err)
					if err != nil {
						break
					}
				}
			}
		}

### 更多信息请看rest.go源代码，还可以发发邮件做进一步交流
---
