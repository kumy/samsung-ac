package proto

type counter struct {
	Counter byte
}

type Counter interface {
	SetCurrent(c byte)
	GetCurrent() byte
	Next() byte
}

func NewCounter(c byte) Counter {
	return &counter{
		Counter: c,
	}
}

func NewCounterFromBytes(b byte) Counter {
	return &counter{
		Counter: b,
	}
}

func (c *counter) Next() byte {
	x := c.GetCurrent()
	c.Counter++
	return x
}

func (c *counter) GetCurrent() byte {
	return c.Counter
}

func (c *counter) SetCurrent(counter byte) {
	c.Counter = counter
}
