package rest

import (
	"testing"
)

/*func TestNewRest(t *testing.T) {
	rt := NewRest("master", 9099)

	t.Log(rt)
}

func TestVersion(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Version())
}

func TestStatus(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Status())
}

func TestListNameSpace(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.ListNameSpace())
}

func TestDescribeNameSpace(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.DescribeNameSpace("default"))
}

func TestCreateNameSpace(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.CreateNameSpace("demo6"))
}

func TestAlterNameSpace(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.AlterNameSpace("demo6"))
}

func TestDropNameSpace(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.DropNameSpace("demo3"))
}

func TestListNameSpaceTables(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.ListNameSpaceTables("default"))
}

func TestList(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.List())
}

func TestDescribe(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Describe("test"))
}

func TestCreate(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Create("test120", []string{"qf", "bf", "cf"}))
}

func TestAlter(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Alter("test120", []string{"qf", "bf", "cf", "df"}))
}

func TestDrop(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Drop("test120"))
}*

func TestTableRegions(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.TableRegions("test"))
}

func TestGetTbRow(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Get("test", "row3", "", "", ""))

	t.Log(rt.Get("test", "row3", "cf:a", "", ""))

	t.Log(rt.Get("test", "row1", "cf:a", "1513489808867", ""))

	t.Log(rt.Get("test", "row1", "cf:a", "", "1"))
}

func TestScanner(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Scanner("test"))
}

func TestScan(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.Scan("test", "15143026402696bb742bc"))
}

func TestDeleteScan(t *testing.T) {
	rt := NewRest("master", 9099)
	t.Log(rt.DeleteScan("test", "15143026402696bb742bc"))
}*/

func TestPut(t *testing.T) {
	rt := NewRest("slave1", 9099)
	t.Log(rt.Put("test", "row9", "cf:a", "value9"))
}
