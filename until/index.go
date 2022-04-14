package until

//数字结构体
type GetNum struct {
	Target int
	Self   int
	Result int
}

//结构体初始化
func NewGetNum(a int, b int) *GetNum {
	return &GetNum{
		Target: a,
		Self:   b,
	}
}

//加法
func (s *GetNum) Add(new *GetNum) int {
	Sum := new.Target + new.Self
	return Sum
}

//减法
func (s *GetNum) Minus(new *GetNum) int {
	Sum := new.Target - new.Self
	return Sum
}

//乘法
func (s *GetNum) Take(new *GetNum) int {
	Sum := new.Target * new.Self
	return Sum
}

//除法
func (s *GetNum) Division(new *GetNum) int {
	Sum := new.Target / new.Self
	return Sum
}