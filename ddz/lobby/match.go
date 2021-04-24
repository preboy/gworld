package lobby

type Match struct {
	ID   uint32
	pids []string
}

func (self *Match) Init(pids []string) {
	self.pids = pids

}

func (self *Match) OnUpdate() {

}

func (self *Match) Over() bool {

	return false
}

func (self *Match) Exist(pid string) bool {
	for _, v := range self.pids {
		if v == pid {
			return true
		}
	}

	return false
}
