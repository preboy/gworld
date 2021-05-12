package block

var (
	_block = NewBlock()
)

type Block struct {
	quit chan bool
}

// ----------------------------------------------------------------------------
// member

func NewBlock() *Block {
	return &Block{
		quit: make(chan bool),
	}
}

func (self *Block) Wait() {
	<-self.quit
}

func (self *Block) Signal() {
	close(self.quit)
}

// ----------------------------------------------------------------------------
// export

func Wait() {
	_block.Wait()
}

func Signal() {
	_block.Signal()
}
