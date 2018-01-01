package rest

/*
 *集群版本结构信息
 */
type ClusterVersion string

/*
 *集群状态相关结构信息
 */
//节点的区块信息
type Region struct {
	CurrentCompactedKVs    string `xml:"currentCompactedKVs,attr"`
	MemstoreSizeMB         string `xml:"memstoreSizeMB,attr"`
	Name                   string `xml:"name,attr"`
	ReadRequestsCount      string `xml:"readRequestsCount,attr"`
	RootIndexSizeKB        string `xml:"rootIndexSizeKB,attr"`
	StorefileIndexSizeMB   string `xml:"storefileIndexSizeMB,attr"`
	StorefileSizeMB        string `xml:"storefileSizeMB,attr"`
	Storefiles             string `xml:"storefiles,attr"`
	Stores                 string `xml:"stores,attr"`
	TotalCompactingKVs     string `xml:"totalCompactingKVs,attr"`
	TotalStaticBloomSizeKB string `xml:"totalStaticBloomSizeKB,attr"`
	TotalStaticIndexSizeKB string `xml:"totalStaticIndexSizeKB,attr"`
	WriteRequestsCount     string `xml:"writeRequestsCount,attr"`
}

//节点信息
type Node struct {
	HeapSizeMB    string `xml:"heapSizeMB,attr"`
	MaxHeapSizeMB string `xml:"maxHeapSizeMB,attr"`
	Name          string `xml:"name,attr"`
	Requests      string `xml:"requests,attr"`
	StartCode     string `xml:"startCode,attr"`

	Region []Region `xml:"Region"`
}

//当机的机器
type DeadNodes struct {
	Node []Node `xml:"Node"`
}

//活跃的机器
type LiveNodes struct {
	Node []Node `xml:"Node"`
}

//集群状态信息
type ClusterStatus struct {
	Describe    string
	AverageLoad string    `xml:"averageLoad,attr"`
	Regions     string    `xml:"regions,attr"`
	Requests    string    `xml:"requests,attr"`
	DeadNodes   DeadNodes `xml:"DeadNodes"`
	LiveNodes   LiveNodes `xml:"LiveNodes"`
}

/*
 *集群命名空间结构信息
 */
type Namespaces struct {
	Namespace []string `xml:"Namespace"`
}
type Namespace struct {
	Name string
}

/*
 *集群数据表相关结构信息
 */
//数据表结构
type Table struct {
	Name string `xml:"name,attr"`
}

//数据表列表结构
type TableList struct {
	Table []Table `xml:"table"`
}

//表的列shema
type ColumnSchema struct {
	Name                string `xml:"name,attr"`
	BLOOMFILTER         string `xml:"BLOOMFILTER,attr"`
	VERSIONS            string `xml:"VERSIONS,attr"`
	IN_MEMORY           string `xml:"IN_MEMORY,attr"`
	KEEP_DELETED_CELLS  string `xml:"KEEP_DELETED_CELLS,attr"`
	DATA_BLOCK_ENCODING string `xml:"DATA_BLOCK_ENCODING,attr"`
	TTL                 string `xml:"TTL,attr"`
	COMPRESSION         string `xml:"COMPRESSION,attr"`
	MIN_VERSIONS        string `xml:"MIN_VERSIONS,attr"`
	BLOCKCACHE          string `xml:"BLOCKCACHE,attr"`
	BLOCKSIZE           string `xml:"BLOCKSIZE,attr"`
	REPLICATION_SCOPE   string `xml:"REPLICATION_SCOPE,attr"`
}

//表的schema
type TableSchema struct {
	ColumnSchema []ColumnSchema `xml:"ColumnSchema"`
}

//表的region
type TableRegion struct {
	EndKey   string `xml:"endKey,attr"`
	Id       string `xml:"id,attr"`
	Location string `xml:"location,attr"`
	Name     string `xml:"name,attr"`
	StartKey string `xml:"startKey,attr"`
}

//表的region列表
type TableRegions struct {
	Name   string        `xml:name,attr`
	Region []TableRegion `xml:"Region"`
}

//记录的cell
type Cell struct {
	Column    string `xml:"column,attr"`
	Timestamp string `xml:"timestamp,attr"`
	Value     string `xml:",innerxml"`
}

//记录的row
type Row struct {
	Key  string `xml:"key,attr"`
	Cell []Cell `xml:"Cell"`
}

//一条记录信息
type CellSet struct {
	Row []Row `xml:"Row"`
}
