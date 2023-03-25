package fruitMap

var fruits = map[string]string{
	"maca":     "É uma fruta que cresce em arvore",
	"banana":   "É uma fruta amarela",
	"morango":  "O morango é uma fruta vermelha, cuja origem é a Europa",
	"melancia": "A melancia é uma fruta rasteira, originária da África",
}

func FindFruit(fruit string) string {
	fruit, ok := fruits[fruit]

	if ok {
		return fruit
	}
	return "Fruta não encontrada"
}
