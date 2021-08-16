package task

type Detail struct {
	value string
}

func NewDetail(value string) (detail *Detail) {
	detail = new(Detail)
	detail.value = value
	return
}

func (Detail) Undefined() (detail *Detail) {
	detail = new(Detail)
	detail.value = ""
	return
}

func (d *Detail) IsUndefined() bool {
	return d.value == ""
}
