package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	prioridad int
	nombre    string
}

type TaskPQ []Task

func (self TaskPQ) Len() int { return len(self) }
func (self TaskPQ) Less(i, j int) bool {
	return self[i].prioridad < self[j].prioridad
}
func (self TaskPQ) Swap(i, j int)       { self[i], self[j] = self[j], self[i] }
func (self *TaskPQ) Push(x interface{}) { *self = append(*self, x.(Task)) }
func (self *TaskPQ) Pop() (popped interface{}) {
	popped = (*self)[len(*self)-1]
	*self = (*self)[:len(*self)-1]
	return
}

func main() {

	pq := &TaskPQ{
		{8, "Pedri"},
		{9, "Lewandoski"},
		{4, "Araujo"},
		{1, "Ter Stegen"}}

	heap.Init(pq)

	// encolar
	heap.Push(pq, Task{2, "Dest"})
	heap.Push(pq, Task{6, "Gavi"})

	fmt.Println("Football Club Barcelona inscripciones")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Ingrese el Dorsal del jugador a inscribir: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Print("ingrese el nombre del jugador: ")
	scanner.Scan()
	input2 := scanner.Text()
	fmt.Println("Nombre: ", input2, "\n Dorsal ", input)
	heap.Push(pq, Task{int(input), input2})
	fmt.Println("Jugadores Inscritos ordenados")

	for pq.Len() != 0 {
		fmt.Println(heap.Pop(pq))
	}

}
