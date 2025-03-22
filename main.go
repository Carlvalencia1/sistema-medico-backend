package main

func main() {
    dependencies := NewDependencies() // Llama directamente a la función NewDependencies
    if err := dependencies.Run(); err != nil {
        panic(err)
    }
}