package ds

type MaxPQ struct {
	pq   []string
	size int
}

func NewMaxPQ(maxN int) *MaxPQ {
	return &MaxPQ{
		pq:   make([]string, maxN+1),
		size: 0,
	}
}

func (p *MaxPQ) IsEmpty() bool {
	return p.size == 0
}

func (p *MaxPQ) GetSize() int {
	return p.size
}

func (p *MaxPQ) GetPQ() []string {
	return p.pq
}

func (p *MaxPQ) Insert(v string) {
	p.size++
	p.pq[p.size] = v
	p.swim(p.size)
}

func (p *MaxPQ) DeleteMax() string {
	max := p.pq[1]
	p.size--
	p.exchange(1, p.size)
	p.pq[p.size+1] = ""
	p.sink(1)
	return max
}

func (p *MaxPQ) less(i, j int) bool {
	return p.pq[i] < p.pq[j]
}

func (p *MaxPQ) exchange(i, j int) {
	t := p.pq[i]
	p.pq[i] = p.pq[j]
	p.pq[j] = t
}

func (p *MaxPQ) swim(k int) {
	var kk int = k
	for {
		if kk > 1 && p.less(kk/2, kk) {
			p.exchange(kk/2, kk)
			kk = kk / 2
		}
		break
	}
}

func (p *MaxPQ) sink(k int) {
	var kk int = k
	for {
		if 2*kk <= p.size {
			j := 2 * kk

			if j < p.size && p.less(j, j+1) {
				j++
			}

			if !p.less(kk, j) {
				break
			}

			p.exchange(kk, j)
			kk = j
		}
		break
	}
}
