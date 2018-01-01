package rest

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/qizexi/qfunc"
)

type rest struct {
	Host string
	Port int
}

func NewRest(host string, port int) *rest {
	rt := new(rest)
	if host == "" {
		host = "127.0.0.1"
	}
	if port <= 0 {
		port = 8088
	}

	rt.Host = host
	rt.Port = port

	return rt
}

//获取集群的版本信息
func (rt *rest) Version() (ClusterVersion, error) {
	var ver ClusterVersion

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/version/cluster", rt.Port))
	if err != nil {
		return ver, err
	}

	err = xml.Unmarshal([]byte(rs), &ver)
	if err != nil {
		return ver, err
	}

	return ver, nil
}

//列出集群的状态
func (rt *rest) Status() (ClusterStatus, error) {
	var cls ClusterStatus

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/status/cluster", rt.Port))
	if err != nil {
		return cls, err
	}

	err = xml.Unmarshal([]byte(rs), &cls)
	if err != nil {
		return cls, err
	}

	//集群的描述
	cls.AverageLoad = fmt.Sprintf("%d live servers, %d dead servers, %s average load",
		len(cls.LiveNodes.Node), len(cls.DeadNodes.Node), cls.AverageLoad)

	return cls, nil
}

//列出所有的命名空间
func (rt *rest) ListNameSpace() (Namespaces, error) {
	var ns Namespaces

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/namespaces", rt.Port))
	if err != nil {
		return ns, err
	}

	err = xml.Unmarshal([]byte(rs), &ns)
	if err != nil {
		return ns, err
	}

	return ns, nil
}

//获取某个命名空间的信息
//nse 命名空间的名称
func (rt *rest) DescribeNameSpace(nse string) (Namespace, error) {
	var ns Namespace

	if nse == "" {
		return ns, errors.New("缺少命名空间名称.")
	}

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/namespaces/%s", rt.Port, nse))
	if err != nil {
		return ns, err
	}
	_ = rs
	ns.Name = nse

	return ns, nil
}

//创建一个命名空间
//nse 命名空间的名称
func (rt *rest) CreateNameSpace(nse string) error {
	if nse == "" {
		return errors.New("缺少命名空间名称.")
	}

	rq := NewRequest("", "text/plain")
	rs, err := rq.Post("http://"+rt.Host+fmt.Sprintf(":%d/namespaces/%s", rt.Port, nse), "")
	if err != nil {
		return err
	}
	if rs != "" {
		return errors.New(rs)
	}

	return nil
}

//更新一个命名空间
//nse 命名空间的名称
func (rt *rest) AlterNameSpace(nse string) error {
	if nse == "" {
		return errors.New("缺少命名空间名称.")
	}

	rq := NewRequest("", "text/plain")
	rs, err := rq.Put("http://"+rt.Host+fmt.Sprintf(":%d/namespaces/%s", rt.Port, nse), "")
	if err != nil {
		return err
	}
	if rs != "" {
		return errors.New(rs)
	}

	return nil
}

//删除一个命名空间
//nse 命名空间的名称
func (rt *rest) DropNameSpace(nse string) error {
	if nse == "" {
		return errors.New("miss nse param.")
	}

	rq := NewRequest("", "text/plain")
	rs, err := rq.Delete("http://" + rt.Host + fmt.Sprintf(":%d/namespaces/%s", rt.Port, nse))
	if err != nil {
		return err
	}
	if len(rs) > 0 {
		return errors.New(rs)
	}

	return nil
}

//获取某个命名空间相关表的信息
//nse 命名空间的名称
func (rt *rest) ListNameSpaceTables(nse string) (TableList, error) {
	var tb TableList

	if nse == "" {
		return tb, errors.New("缺少命名空间名称.")
	}

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/namespaces/%s/tables", rt.Port, nse))
	if err != nil {
		return tb, err
	}

	err = xml.Unmarshal([]byte(rs), &tb)
	if err != nil {
		return tb, err
	}

	return tb, nil
}

//列出所有非系统的数据表
func (rt *rest) List() (TableList, error) {
	var tb TableList

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d", rt.Port))
	if err != nil {
		return tb, err
	}

	err = xml.Unmarshal([]byte(rs), &tb)
	if err != nil {
		return tb, err
	}

	return tb, nil
}

//获取表的描述信息(schema)
func (rt *rest) Describe(tb string) (TableSchema, error) {
	var ts TableSchema

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/%s/schema", rt.Port, tb))
	if err != nil {
		return ts, err
	}

	err = xml.Unmarshal([]byte(rs), &ts)
	if err != nil {
		return ts, err
	}

	return ts, nil
}

//创建或修改表
//tb 表名
//cfs 列族名-数组
func (rt *rest) Create(tb string, cfs []string) error {
	if tb == "" {
		return errors.New("缺少表名参数.")
	}
	rq := NewRequest("", "text/xml")

	post := `<?xml version="1.0" encoding="UTF-8"?>`
	post += `<TableSchema name="` + tb + `">`
	if len(cfs) <= 0 {
		post += `<ColumnSchema name="cf" />`
	} else {
		for _, cf := range cfs {
			post += `<ColumnSchema name="` + cf + `" />`
		}
	}
	post += `</TableSchema>`

	rs, err := rq.Post("http://"+rt.Host+fmt.Sprintf(":%d/%s/schema", rt.Port, tb), post)
	if err != nil {
		return err
	}

	if rs != "" {
		return errors.New(rs)
	}

	return nil
}

//更新表的记录
//tb 表名
//cfs 列族名-数组
func (rt *rest) Alter(tb string, cfs []string) error {
	if tb == "" {
		return errors.New("缺少表名参数.")
	}
	rq := NewRequest("", "text/xml")

	post := `<?xml version="1.0" encoding="UTF-8"?>`
	post += `<TableSchema name="` + tb + `">`
	if len(cfs) <= 0 {
		post += `<ColumnSchema name="cf" KEEP_DELETED_CELLS="true" />`
	} else {
		for _, cf := range cfs {
			post += `<ColumnSchema name="` + cf + `" KEEP_DELETED_CELLS="true" />`
		}
	}
	post += `</TableSchema>`

	rs, err := rq.Put("http://"+rt.Host+fmt.Sprintf(":%d/%s/schema", rt.Port, tb), post)
	if err != nil {
		return err
	}

	if rs != "" {
		return errors.New(rs)
	}

	return nil
}

//删除表
//tb 表名
func (rt *rest) Drop(tb string) error {
	if tb == "" {
		return errors.New("缺少表名.")
	}

	rq := NewRequest("", "text/plain")
	rs, err := rq.Delete("http://" + rt.Host + fmt.Sprintf(":%d/%s/schema", rt.Port, tb))
	if err != nil {
		return err
	}
	if len(rs) > 0 {
		return errors.New(rs)
	}

	return nil
}

//获取表的region
//tb 表名
func (rt *rest) TableRegions(tb string) (TableRegions, error) {
	var tr TableRegions

	if tb == "" {
		return tr, errors.New("缺少表名.")
	}

	rq := NewRequest("", "")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/%s/regions", rt.Port, tb))
	if err != nil {
		return tr, err
	}

	err = xml.Unmarshal([]byte(rs), &tr)
	if err != nil {
		return tr, err
	}

	return tr, nil
}

//获取表的数据
//tb 表名
//row 行名
//cf 列名
//stamp 时间戳
//ver 版本信息
func (rt *rest) Get(tb string, row string, cf string, stamp string, ver string) (CellSet, error) {
	var cs CellSet

	if tb == "" {
		return cs, errors.New("缺少表名.")
	}

	if row == "" {
		return cs, errors.New("缺少列名.")
	}

	rq := NewRequest("", "")
	rs := ""
	var err error
	if cf == "" {
		rs, err = rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/%s/%s",
			rt.Port, tb, row))
	} else if stamp != "" || ver != "" {
		if ver == "" {
			rs, err = rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/%s/%s/%s/%s",
				rt.Port, tb, row, cf, stamp))
		} else {
			rs, err = rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/%s/%s/%s?v=%s",
				rt.Port, tb, row, cf, ver))
		}
	} else {
		rs, err = rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/%s/%s/%s",
			rt.Port, tb, row, cf))
	}

	if err != nil {
		return cs, err
	}

	err = xml.Unmarshal([]byte(rs), &cs)
	if err != nil {
		return cs, err
	}

	for k, v := range cs.Row {
		cs.Row[k].Key, _ = qfunc.Base64Decode(v.Key)
		for k1, v1 := range v.Cell {
			cs.Row[k].Cell[k1].Column, _ = qfunc.Base64Decode(v1.Column)
			cs.Row[k].Cell[k1].Value, _ = qfunc.Base64Decode(v1.Value)
		}
	}

	return cs, nil
}

//获取扫描表的id
//tb 表名
func (rt *rest) Scanner(tb string) (string, error) {
	if tb == "" {
		return "", errors.New("缺少表名称.")
	}

	rq := NewRequest("text/xml", "text/xml")
	rs, err := rq.Put("http://"+rt.Host+fmt.Sprintf(":%d/%s/scanner", rt.Port, tb),
		`<Scanner batch="1"/>`)
	if err != nil {
		return "", err
	}
	if rs != "" {
		return "", errors.New(rs)
	}

	surl := rq.RespHeader.Get("Location")
	if surl == "" {
		return "", errors.New("创建sanner失败，未知原因.")
	}
	sarr := strings.Split(surl, "/")
	sid := sarr[len(sarr)-1]

	return sid, nil
}

//通过扫描表的ID进行扫描
//tb 表名称
//sid scanner id
func (rt *rest) Scan(tb string, sid string) (CellSet, error) {
	var cs CellSet

	if tb == "" {
		return cs, errors.New("缺少表名称.")
	}
	if sid == "" {
		return cs, errors.New("缺少scanner id.")
	}

	rq := NewRequest("text/xml", "text/xml")
	rs, err := rq.Get("http://" + rt.Host + fmt.Sprintf(":%d/%s/scanner/%s", rt.Port, tb, sid))
	if err != nil {
		return cs, err
	}
	if rs == "" {
		return cs, errors.New("内容为空.")
	}

	err = xml.Unmarshal([]byte(rs), &cs)
	if err != nil {
		return cs, err
	}

	for k, v := range cs.Row {
		cs.Row[k].Key, _ = qfunc.Base64Decode(v.Key)
		for k1, v1 := range v.Cell {
			cs.Row[k].Cell[k1].Column, _ = qfunc.Base64Decode(v1.Column)
			cs.Row[k].Cell[k1].Value, _ = qfunc.Base64Decode(v1.Value)
		}
	}

	return cs, nil
}

//删除scanner
//tb 表名称
//sid scanner id
func (rt *rest) DeleteScan(tb string, sid string) error {
	if tb == "" {
		return errors.New("缺少表名称.")
	}
	if sid == "" {
		return errors.New("缺少scanner id.")
	}

	rq := NewRequest("text/xml", "text/xml")
	rs, err := rq.Delete("http://" + rt.Host + fmt.Sprintf(":%d/%s/scanner/%s", rt.Port, tb, sid))
	if err != nil {
		return err
	}
	if rs != "" {
		return errors.New(rs)
	}

	return nil
}

//添加记录
//tb 表名
//key 行键值
//cf 列名
//val 列值
func (rt *rest) Put(tb string, key string, cf string, val string) error {
	if tb == "" {
		return errors.New("缺少表名称.")
	}
	if key == "" {
		return errors.New("缺少行键值.")
	}
	if cf == "" {
		return errors.New("缺少行列簇名.")
	}
	if val == "" {
		return errors.New("缺少列值.")
	}

	body := `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>`
	var cs CellSet
	var row Row
	var cell Cell
	cell.Column, _ = qfunc.Base64Encode(cf)
	cell.Timestamp = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	cell.Value, _ = qfunc.Base64Encode(val)
	row.Key, _ = qfunc.Base64Encode(key)
	row.Cell = []Cell{cell}
	cs.Row = []Row{row}
	xstr, err := xml.Marshal(cs)
	if err != nil {
		return errors.New("xml格式不正确." + err.Error())
	}
	body += string(xstr)

	rq := NewRequest("text/xml", "text/xml")
	rs, err := rq.Put("http://"+rt.Host+fmt.Sprintf(":%d/%s/fakerow", rt.Port, tb), body)
	if err != nil {
		return err
	}
	if rs != "" {
		return errors.New(rs)
	}

	return nil
}
