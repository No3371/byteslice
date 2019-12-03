package byteslice

type ByteSlice struct {
	storage     []byte
	Payload     []byte
	left, right uint32
}

// ResetPayload resets the payload to the assigned memory segment assigned when created
func (c *ByteSlice) ResetPayload() {
	c.Payload = c.storage[c.left:c.right]
}

func (c *ByteSlice) SetPayloadSize(size uint32) {
	c.Payload = c.storage[:c.left+size]
}

func (c *ByteSlice) GetMaxPayloadSize() uint32 {
	return c.right - c.left
}

func CreatesSingleSlice(storage []byte) (bc *ByteSlice) {
	bc = &ByteSlice{
		storage: storage,
		left:    0,
		right:   uint32(len(storage)),
	}
	bc.ResetPayload()
	return bc
}

// CreateSlices create multiple ByteSlice using continuous memory
func CreateSlices(sizePerSlice uint32, slice uint32) (created []*ByteSlice) {
	storage := make([]byte, sizePerSlice*slice)
	created = make([]*ByteSlice, 0, slice)
	for i := uint32(0); i < slice; i++ {
		t := &ByteSlice{
			storage: storage,
			left:    sizePerSlice * i,
			right:   sizePerSlice * (i + 1),
		}
		t.ResetPayload()
		created = append(created, t)
	}
	return created
}
