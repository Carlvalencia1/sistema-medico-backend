package main

//CORS GIN

func main() {
	dependencies := NewDependencies()
	if err := dependencies.Run(); err != nil {
		panic(err)
	}
}