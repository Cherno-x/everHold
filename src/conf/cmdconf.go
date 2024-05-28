package conf

type RUNCMD struct {
	MethodValue  string
	AddValue     bool
	NameValue    string
	PayloadValue string
	ModuleValue  string
}

func NewRunCMD(module string, method string, add bool, name string, payload string) *RUNCMD {
	return &RUNCMD{
		ModuleValue:  module,
		MethodValue:  method,
		AddValue:     add,
		NameValue:    name,
		PayloadValue: payload,
	}
}
