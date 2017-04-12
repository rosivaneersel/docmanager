package alerts

type Alert struct {
	Title   string
	Class   string
	Message string
}

type alerts []Alert

func (a *alerts) New(title string, class string, message string) {
	*a = append(*a, Alert{title, class, message})
}

func (a *alerts) Get() []Alert {
	c := *a
	*a = nil
	return c
}

var Alerts alerts