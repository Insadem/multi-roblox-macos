package fspath

import "sync"

type Path struct {
	owp  string
	path func() (string, error)
	mut  sync.Mutex
}

func New(path func() (string, error)) *Path {
	return &Path{
		path: path,
	}
}

// Set overwrites path if s is not empty
func (p *Path) Set(s string) {
	p.mut.Lock()
	defer p.mut.Unlock()

	p.owp = s
}

func (p *Path) Get() (string, error) {
	p.mut.Lock()
	defer p.mut.Unlock()

	if p.owp != "" {
		return p.owp, nil
	}

	return p.path()
}
