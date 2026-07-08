package colorista

func (self *Colorista) Success(a any) string {
	return self.Apply(a, Bold, BrightGreen)
}

func (self *Colorista) Error(a any) string {
	return self.Apply(a, Bold, BrightRed)
}

func (self *Colorista) Warning(a any) string {
	return self.Apply(a, Bold, BrightYellow)
}

func (self *Colorista) Info(a any) string {
	return self.Apply(a, Bold, BrightBlue)
}

func (self *Colorista) Debug(a any) string {
	return self.Apply(a, BrightBlack)
}

func (self *Colorista) Title(a any) string {
	return self.Apply(a, Bold, Underline, BrightWhite)
}
