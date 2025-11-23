package main

import (
	"context"
	"log"
	"math/rand/v2"
	"naloga-5/redovalnica"
	"os"

	"github.com/urfave/cli/v3"
)


func main() {
  cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "Upravljanje z redovalnico",
		Flags: []cli.Flag{
			&cli.Uint8Flag{
				Name:  "stOcen",
				Usage: "Minimalno stevilo ocen za izracun povprečja",
				Value: 6,
			},
			&cli.Uint8Flag{
				Name:  "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
			&cli.Uint8Flag{
				Name:  "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 1,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
  var studenti = make(map[string]redovalnica.Student)

  var s1 redovalnica.Student
  var s2 redovalnica.Student
  var s3 redovalnica.Student
  var s4 redovalnica.Student

  redovalnica.InitStudent(&s1, "Ana", "Novak")
  redovalnica.InitStudent(&s2, "Boris", "Kralj")
  redovalnica.InitStudent(&s3, "Janez", "Novak")
  redovalnica.InitStudent(&s4, "Neza", "Kraljica")

  studenti["63230001"] = s1
  studenti["63230002"] = s2
  studenti["63230003"] = s3
  studenti["63230004"] = s4

  for k := range studenti {
    stOcen := rand.IntN(4) + 5
    for range stOcen {
      o := rand.IntN(10) + 1
      redovalnica.DodajOceno(studenti, k, o)
    }
  }
  redovalnica.IzpisRedovalnice(studenti)
  redovalnica.IzpisiKoncniUspeh(studenti)
  return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}