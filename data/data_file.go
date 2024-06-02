package data

import (
	"errors"
	"tinykv-go/fio"
)

var (
	ErrInvalidCRC = errors.New("invalid crc value, log record maybe corrupted")
)

const DataFileNameSuffix = ".data"

type DataFile struct {
	FileId    uint32        // 文件id
	WriteOff  int64         // 文件写到了哪个位置
	IoManager fio.IOManager // io 读写流程
}

// 打开新的数据文件
func OpenDataFile(dirPath string, fileId uint32) (*DataFile, error) {
	return nil, nil
}

func (df *DataFile) ReadLogRecored(offset int64) (*LogRecord, int64, error) {
	return nil, 0, nil
}

func (df *DataFile) Write(buf []byte) error {
	return nil
}

func (df *DataFile) Sync() error {
	return df.IoManager.Sync()
}
