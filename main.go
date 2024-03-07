package main

func main() {
	store := NewStore()
	api := newAPIServer(":3000", nil)
}
