package polymorphism

import "fmt"

type Income interface {
	calculate() int
	aisatsu() string
}

type kachou struct {
	name                    string
	workingYear, baseSalary int
}

type ippan struct {
	name                                 string
	baseSalary, workingYear, performance int
}

func (kachou kachou) calculate() int {
	return kachou.baseSalary + (kachou.workingYear * kachou.workingYear)
}
func (kachou kachou) aisatsu() string {
	return "私の名前は" + kachou.name + "です"

}

func (ippan ippan) calculate() int {
	return ippan.baseSalary + (1000 * ippan.workingYear * ippan.performance)
}

func (ippan ippan) aisatsu() string {
	return "私の名前は" + ippan.name + "です"
}

func calculateIncome(ic []Income) int {
	sum := 0
	for _, worker := range ic {
		sum += worker.calculate()
	}
	return sum
}

func calculateRoop(ic []Income) {
	for _, worker := range ic {
		fmt.Println("それぞれの給与トータルはこちらです", worker.calculate())
	}
}

func aisatsuRoop(ic []Income) {
	for _, worker := range ic {
		fmt.Println(worker.aisatsu())
	}
}

func Fuga02() {
	hanako := kachou{
		name:        "hanako",
		workingYear: 5,
		baseSalary:  11,
	}
	choko := ippan{
		name:        "choko",
		baseSalary:  60000,
		workingYear: 50,
	}
	motoko := ippan{
		name:        "motoko",
		baseSalary:  40000,
		workingYear: 25,
	}

	workers := []Income{hanako, choko, motoko} //yahoo, google型がincomeインターフェースを満たしているから渡せる
	calculateRoop(workers)
	fmt.Printf("Total income: %d\n", calculateIncome(workers))
	aisatsuRoop(workers)

}
func Output() {
	Fuga02()
	// fmt.Println(OutputName(dog{}))
	// fmt.Println(OutputName(cat{}))
}
