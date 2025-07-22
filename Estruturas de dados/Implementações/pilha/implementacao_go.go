package main
import "fmt"

// Implementação da estrutura de dados pilha na linguagem de programação go 
// implementação voltada para dados do tipo int

type pilha [] int

// push, pop, peek, clear, isEmpty
func isEmpty(x pilha) bool{
    if len(x) == 0{
        return true
    }
    return false
}
func clear(x pilha) pilha{
    var aux pilha
    x = aux
    return x
}
func peek(x pilha) int{
    return x[len(x)-1]
}
func push(x pilha, y int) pilha{
    x = append(x, y)
    return x
}
func pop(x pilha) pilha{
    var aux pilha
    aux = x[:len(x)-1]
    x = aux
    return x
}


func main() {
    var p pilha
  // Teste: pilha vazia
	fmt.Println("Inicialmente, está vazia?", isEmpty(p)) // true

	// Teste: push
	p = push(p, 10)
	p = push(p, 20)
	p = push(p, 30)
	fmt.Println("Após inserções, pilha:", p) // [10 20 30]
	fmt.Println("Topo da pilha:", peek(p))   // 30
	fmt.Println("Está vazia agora?", isEmpty(p)) // false

	// Teste: pop
	p = pop(p)
	fmt.Println("Após pop, pilha:", p) // [10 20]
	fmt.Println("Topo agora:", peek(p)) // 20

	// Teste: clear
	p = clear(p)
	fmt.Println("Após clear, pilha:", p)       // []
	fmt.Println("Está vazia após clear?", isEmpty(p)) // true
}