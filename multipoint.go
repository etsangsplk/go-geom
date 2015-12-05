package geom

type MultiPoint struct {
	geom1
}

var _ T = &MultiPoint{}

func NewMultiPoint(layout Layout) *MultiPoint {
	return NewMultiPointFlat(layout, nil)
}

func NewMultiPointFlat(layout Layout, flatCoords []float64) *MultiPoint {
	mp := new(MultiPoint)
	mp.layout = layout
	mp.stride = layout.Stride()
	mp.flatCoords = flatCoords
	return mp
}

func (mp *MultiPoint) Clone() *MultiPoint {
	flatCoords := make([]float64, len(mp.flatCoords))
	copy(flatCoords, mp.flatCoords)
	return NewMultiPointFlat(mp.layout, flatCoords)
}

func (mp *MultiPoint) MustSetCoords(coords1 [][]float64) *MultiPoint {
	if err := mp.SetCoords(coords1); err != nil {
		panic(err)
	}
	return mp
}

func (mp *MultiPoint) NumPoints() int {
	return mp.NumCoords()
}

func (mp *MultiPoint) Point(i int) *Point {
	return NewPointFlat(mp.layout, mp.Coord(i))
}

func (mp *MultiPoint) Push(p *Point) error {
	if p.layout != mp.layout {
		return ErrLayoutMismatch{Got: p.layout, Want: mp.layout}
	}
	mp.flatCoords = append(mp.flatCoords, p.flatCoords...)
	return nil
}
