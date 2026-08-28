package main

func Ex3() {
	people := []Person{}
	// people := make([]Person, 0, 10_000_000) // this is WAY more performant. Like 0.3 vs 1.2 seconds
	for i := 0; i < 10_000_000; i++ {
		people = append(people, MakePerson("Jane", "Doe", 25))
	}
	// default GOGC (100) = 1.25 s; 17 GCs
	// 80 = 1.23s; 19 GCs
	// 20 = 1.52s; 30 GCs
	// 1000 = 813ms; 4 GCs
	// off = 627ms; 0 GCs
}
