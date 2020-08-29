package main

type BytePoolCap struct {
	c    chan []byte
	w    int
	wcap int
}

func NewBytePoolCap(maxSize, width, capWidth int) (bp *BytePoolCap) {
	return &BytePoolCap{
		c:    make(chan []byte, maxSize),
		w:    width,
		wcap: capWidth,
	}
}

func (bp *BytePoolCap) Get() (b []byte) {
	select {
	case b = <-bp.c:
	default:
		if bp.wcap > 0 {
			b = make([]byte, bp.w, bp.wcap)
		} else {
			b = make([]byte, bp.w)
		}
	}
	return
}

func (bp *BytePoolCap) Put(b []byte) {
	select {
	case bp.c <- b:
	default:

	}
}

func main() {
	bp := NewBytePoolCap(500, 1024, 1024)
	buf := bp.Get()
	defer bp.Put(buf)
}
