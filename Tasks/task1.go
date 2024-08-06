import "fmt"
func averageCalculator() uint16 {
	var name string
	var numSubjects int

	fmt.Println("What is your Full Name: ")
	fmt.Scanln(&name)

	fmt.Println("How many subjects have you taken: ")
	fmt.Scanln(&numSubjects)

	subjects := make([]string, numSubjects)
	grades := make(map[string]int)
	sum := 0

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scanln(&subjects[i])

		fmt.Printf("Enter the grade for %s: ", subjects[i])
		var grade int
		fmt.Scanln(&grade)

		for grade < 0 || grade > 100 {
			fmt.Println("Invalid grade! Please enter a grade between 0 and 100.")
			fmt.Printf("Enter the grade for %s: ", subjects[i])
			fmt.Scanln(&grade)
		}

		grades[subjects[i]] = grade
		sum += grade
	}

	average := uint16(sum / numSubjects)

	fmt.Printf("\nName: %s\n", name)
	fmt.Println("Subject Grades:")
	for _, subject := range subjects {
		fmt.Printf("%s: %d\n", subject, grades[subject])
	}
	fmt.Printf("Average Grade: %d\n", average)

	return average
}