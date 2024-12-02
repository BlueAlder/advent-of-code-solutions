package defaultdict

type DefaultDict[T comparable, U any] struct {
	dict map[T]U
	dVal U
}

func NewDefaultDict[T comparable, U any](defaultVal U) *DefaultDict[T, U] {
	return &DefaultDict[T, U]{
		dict: make(map[T]U),
		dVal: defaultVal,
	}
}

func (d DefaultDict[T, U]) Add(k T, v U) {
	d.dict[k] = v
}

func (d DefaultDict[T, U]) Get(k T) U {
	v, ext := d.dict[k]
	if !ext {
		return d.dVal
	}
	return v
}

func (d DefaultDict[T, U]) Delete(k T) {
	delete(d.dict, k)
}

func (d DefaultDict[T, U]) Values() map[T]U {
	return d.dict
}
