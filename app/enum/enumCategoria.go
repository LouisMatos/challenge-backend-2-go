package enum

type Categoria int

// Declare related constants for each direction starting with index 1
const (
	Alimentacao Categoria = iota + 1 // EnumIndex = 1
	Saude                            // EnumIndex = 2
	Moradia                          // EnumIndex = 3
	Transporte                       // EnumIndex = 4
	Educação                         // EnumIndex = 5
	Lazer                            // EnumIndex = 6
	Imprevistos                      // EnumIndex = 7
	Outras                           // EnumIndex = 8
)

// String - Creating common behavior - give the type a String function
func (c Categoria) String() string {
	return [...]string{"Alimentação", "Saúde", "Moradia", "Transporte", "Educação", "Lazer", "Imprevistos", "Outras"}[c-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (c Categoria) EnumIndex() int {
	return int(c)
}
