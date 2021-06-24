package riscv

type Program interface {
	Name() string
	Machine() string
	Class() string
	Order() string
	Sections() []Section
	AddSection(s Section)
}

type program struct {
	name     string
	machine  string
	class    string
	order    string
	sections []Section
}

func NewProgram(name, machine, class, order string) Program {
	return &program{
		name:    name,
		machine: machine,
		class:   class,
		order:   order,
	}
}

func (p *program) Name() string {
	return p.name
}

func (p *program) Machine() string {
	return p.machine
}

func (p *program) Class() string {
	return p.class
}

func (p *program) Order() string {
	return p.order
}

func (p *program) Sections() []Section {
	return p.sections
}

func (p *program) AddSection(s Section) {
	p.sections = append(p.sections, s)
}
