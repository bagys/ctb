package ctb

type TableAttrFunc func(t *Table)
type TableAttrFuncs []TableAttrFunc

func (f TableAttrFuncs) Apply(t *Table) {
	for _, fun := range f {
		fun(t)
	}
}

func WithPrefixTab(Prefix string) TableAttrFunc {
	return func(t *Table) {
		t.prefixTab = Prefix
	}
}

func WithPrefixContent(Prefix string) TableAttrFunc {
	return func(t *Table) {
		t.prefixContent = Prefix
	}
}

func WithPrefixDisable(Disable bool) TableAttrFunc {
	return func(t *Table) {
		t.prefixDisable = Disable
	}
}

func WithSpacing(spacing int) TableAttrFunc {
	return func(t *Table) {
		t.SetSpacing(spacing)
	}
}

func WithTab(tab []LineData) TableAttrFunc {
	return func(t *Table) {
		t.SetTab(tab)
	}
}

func WithDataOne(data []LineData) TableAttrFunc {
	return func(t *Table) {
		t.SetDataOne(data)
	}
}

func WithDataAll(data [][]LineData) TableAttrFunc {
	return func(t *Table) {
		t.SetDataAll(data)
	}
}
