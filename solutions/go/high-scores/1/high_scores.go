package highscores

import "sort"

type HighScores struct{
    scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	safeScores := make([]int, len(scores))

    copy(safeScores, scores)

    return &HighScores{scores: safeScores}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
    out := make([]int, len(s.scores))
    copy(out, s.scores)
	return out
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
    if(len(s.scores) == 0){
        return 0
    }
	return s.scores[len(s.scores) - 1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	if(len(s.scores) == 0){
        return 0
    }
    
	best := s.scores[0]

    for _,v := range s.scores {
        if v > best {
            best = v
        }
    }

    return best
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	sorted := make([]int, len(s.scores))

    copy(sorted, s.scores)

    sort.Slice(sorted, func(i, j int) bool {
      return  sorted[i] > sorted[j]
    })

    if len(sorted) > 3 {
        return sorted[:3]
    }

    return sorted
}
