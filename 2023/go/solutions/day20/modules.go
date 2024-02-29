package day20

type ModuleData struct {
	label        string
	destinations []string
}

func (md ModuleData) GetDestinations() []string {
	return md.destinations
}
func (md ModuleData) GetLabel() string {
	return md.label
}

type Broadcast struct {
	ModuleData
}

func (b *Broadcast) receive(p Pulse) []Pulse {
	pulses := make([]Pulse, len(b.destinations))
	for i, dest := range b.destinations {
		pulses[i] = Pulse{src: b.label, dest: dest, high: p.high}
	}
	return pulses
}

type FlipFlop struct {
	status bool
	ModuleData
}

func (f *FlipFlop) receive(p Pulse) []Pulse {
	pulses := make([]Pulse, 0)
	if !p.high {
		f.status = !f.status
		for _, dest := range f.destinations {
			pulses = append(pulses, Pulse{src: f.label, dest: dest, high: f.status})
		}
	}
	return pulses
}

type Conjuction struct {
	inputLastValue map[string]bool
	ModuleData
}

func (c *Conjuction) receive(p Pulse) []Pulse {
	c.inputLastValue[p.src] = p.high
	pulses := make([]Pulse, 0)

	toSendHigh := false
	for _, v := range c.inputLastValue {
		if !v {
			toSendHigh = true
			break
		}
	}

	for _, dest := range c.destinations {
		pulses = append(pulses, Pulse{src: c.label, dest: dest, high: toSendHigh})
	}

	return pulses
}

type Pulse struct {
	src  string
	dest string
	high bool
}

type Module interface {
	receive(Pulse) []Pulse
	GetDestinations() []string
	GetLabel() string
}
