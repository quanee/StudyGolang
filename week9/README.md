### 1. 总结几种 socket 粘包的解包方式: fix length/delimiter based/length field based frame decoder。尝试举例其应用

- #### fix length
	发送方每次发送不超过缓冲区大小的固定长度的数据, 接受方按固定长度区接受数据

- #### delimiter based
	发送方在数据包中添加特定的分隔符用来标记数据包边界

- #### length field based
	发送方在消息数据包头添加包长度信息

### 2. 实现一个从 socket connection 中解码出 goim 协议的解码器。


```golang
const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
)

func main() {
	buf := ReadTCP()
	if len(buf) < 16 {
		panic("buf len < 16.")
		return
	}
	decoderHead(buf[:_rawHeaderSize])
	fmt.Println("body: ", buf[_rawHeaderSize:])
}


func Decoder(data []byte) {
	packSize := Int32(data[:4])
	fmt.Printf("pack size: %v\n", packetLen)

	header := Int16(data[4:6])
	fmt.Printf("header size: %v\n", headerLen)

	version := Int16(data[6:8])
	fmt.Printf("version: %v\n", version)

	opSize := Int32(data[8:12])
	fmt.Printf("operation: %v\n", operation)

	seqSize := Int32(data[12:16])
	fmt.Printf("sequence: %v\n", sequence)
}

func Int16(b []byte) int16 {
	return int16(b[1]) | int16(b[0])<<8
}

func Int32(b []byte) int32 {
	return int32(b[3]) | int32(b[2])<<8 | int32(b[1])<<16 | int32(b[0])<<24
}
```
