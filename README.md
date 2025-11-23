Package redovalnica provides utilities for managing students and their grades.

The package allows you to initialize students, assign grades, calculate averages,
and print grade reports. It is designed as a simple utility for working with
redovalnica (grade books) using Go maps to store students by their vpisna številka.

Example usage:

```go
studenti := make(map[string]redovalnica.Student)

var s redovalnica.Student
redovalnica.InitStudent(&s, "Ana", "Novak")
studenti["63230001"] = s

redovalnica.DodajOceno(studenti, "63230001", 9)
redovalnica.DodajOceno(studenti, "63230001", 10)

redovalnica.IzpisRedovalnice(studenti)
redovalnica.IzpisiKoncniUspeh(studenti)
```

The grade book supports the following features:

-   Initializing new students with empty grade lists
-   Adding grades in range 0–10
-   Computing averages with at least six grades
-   Printing detailed and final grade summaries
