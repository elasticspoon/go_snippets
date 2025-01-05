package gh_history

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"golang.org/x/exp/rand"
)

func Run() {
	r, _ := git.PlainOpen(".")
	w, _ := r.Worktree()
	_ = w

	startDate := time.Date(2024, time.July, 20, 1, 1, 1, 1, time.Local)
	endDate := time.Now()

	fmt.Println()

	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		switch d.Weekday() {
		case time.Saturday, time.Sunday:
			buildFakeCommit(w, d, rand.Intn(10)+2)
		default:
			if rand.Intn(100) < 30 {
				buildFakeCommit(w, d, rand.Intn(20))
			}
		}
	}
}

func buildFakeCommit(w *git.Worktree, d time.Time, count int) {
	filepath := "fake-history"

	for i := 0; i < count; i++ {
		newTime := d.Add(time.Duration(i) * time.Hour)
		msg := fmt.Sprintf("fake history %s", newTime)
		_ = os.WriteFile(filepath, []byte(msg), 0644)
		_, _ = w.Add(filepath)

		w.Commit(msg, &git.CommitOptions{
			Author: &object.Signature{
				Name:  "Yuri Bocharov",
				Email: "quesadillaman@gmail.com",
				When:  newTime,
			},
		})

	}
}
