package gh_history

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func Run() {
	r, _ := git.PlainOpen(".")
	w, _ := r.Worktree()
	_ = w

	startDate := time.Date(2024, time.July, 20, 1, 1, 1, 1, time.Local)
	endDate := time.Now()

	for d := startDate; d.Before(endDate); d = d.AddDate(0, 0, 1) {
		switch d.Weekday() {
		case time.Saturday, time.Sunday:
			buildFakeCommit(w, d, randInt(10)+2)
		case time.Tuesday:
			buildFakeCommit(w, d, randInt(12)+1)
		default:
			if randInt(100) < 20 {
				buildFakeCommit(w, d, randInt(20))
			}
		}
	}
}

func randInt(max int) int {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println("Error generating random number:", err)
		return 0
	}
	return int(randomNumber.Int64())
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
