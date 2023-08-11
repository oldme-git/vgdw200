package message

type VgDw200Out struct {
	cmdHeader [2]byte
	cmd       byte
	dataLen   [2]byte
	data      []byte
	checksum  byte
}

func New() *VgDw200Out {
	return &VgDw200Out{
		cmdHeader: [2]byte{0x55, 0xAA},
	}
}

func (v *VgDw200Out) Cmd(cmd byte) *VgDw200Out {
	v.cmd = cmd
	return v
}

func (v *VgDw200Out) Data(data []byte) *VgDw200Out {
	v.data = data
	return v
}

func (v *VgDw200Out) DataStr(data string) *VgDw200Out {
	v.data = []byte(data)
	return v
}

// Get 获取最终组合的报文
func (v *VgDw200Out) Get() []byte {
	var (
		msg     []byte
		dataLen = uint16(len(v.data))
	)
	v.dataLen = DataBytesBig(dataLen)

	msg = v.cmdHeader[0:2]
	msg = append(msg, v.cmd)
	msg = append(msg, v.dataLen[0:2]...)
	if dataLen > 0 {
		msg = append(msg, v.data...)
	}
	msg = append(msg, GetChecksum(msg))
	return msg
}
