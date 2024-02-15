# GO_LinkedList

This package implements a basic singly linked list in Go.

## Usage

```go
package main

import (
    "fmt"
    "github.com/your-username/linkedlist"
)

func main() {
    // Create a new linked list
    list := linkedlist.Make()

    // Add elements to the list
    list.Push(1)
    list.Push(2)
    list.Push(3)

    // Print the list
    list.Print()

    // Remove an element from the list
    list.Remove(2)

    // Print the modified list
    list.Print()

    // Check if a value belongs to the list
    fmt.Println(list.Belongs(1)) // Output: true
    fmt.Println(list.Belongs(5)) // Output: false

    // Get an element from the list by index
    val, err := list.Get(1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Value at index 1:", val) // Output: Value at index 1: 3
    }
}
```

## Methods

- `Make() LinkedList`: Creates a new linked list.
- `Push(value int64)`: Adds a new node with the given value to the beginning of the list.
- `Put(after int64, value int64)`: Inserts a new node with the given value after the node with the specified value.
- `Append(value int64)`: Adds a new node with the given value to the end of the list.
- `Pop() int64`: Removes and returns the value of the node at the beginning of the list.
- `Remove(value int64) bool`: Removes the first occurrence of the node with the specified value from the list.
- `Get(index int64) (int64, error)`: Returns the value of the node at the specified index in the list.
- `Belongs(value int64) bool`: Checks if the list contains a node with the specified value.
- `Print()`: Prints the values of all nodes in the list.
