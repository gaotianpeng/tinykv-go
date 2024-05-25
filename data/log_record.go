package data

// LogRecordPos 数据内存索引，主要描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 // 文件id，表示将数据存储到了哪个文件中
	Offset int64  // 偏移，表示将数据存储琶了数据文件中的哪个位置
}