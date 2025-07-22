// Exercício 1 : conversão decimal-->binário
// A única funções built-in que usarei será a função len e as presentes no pacote padrão (fmt)
package main
import "fmt"


func push(vet []int, x int) []int{
    vet = append(vet, x)
    return vet
}

func pop(vet []int) []int{
    var aux []int
    aux = vet[:len(vet)-1]
    vet = aux
    return vet
}

func peek(vet []int)int{
    return vet[len(vet)-1]
}

func converter(x int) string{
    mapa:= map[int]string{
    0: "0",
    1: "1",
}
    var restos []int
    // pega um numero, enquanto o numero for maior ou igual a 1, armazene na pilha
    for i:= x; i>=1; i= i/2{
        r := i%2
        restos = push(restos, r)
    }
    // desempilhar a pilha e concatenar na string
    var result string
    n := len(restos)
    for i:=0; i<n; i++{
        result+= mapa[peek(restos)]
        restos = pop(restos)
    }
    
    return result
}

func main() {
  
  var n int 
  fmt.Print("Insira um número em decimal: ")
  fmt.Scan(&n)
  fmt.Println("Convertido para binário: ",converter(n))
  
  
  
}