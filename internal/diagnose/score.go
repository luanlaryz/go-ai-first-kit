package diagnose

import "math"

func CategoryScore(coverage, quality int) int {
	return clampScore(int(math.Round(0.6*float64(coverage) + 0.4*float64(quality))))
}

func CoverageScore(present, total int) int {
	if total <= 0 {
		return 100
	}
	return clampScore(int(math.Round(float64(present) / float64(total) * 100)))
}

func clampScore(score int) int {
	switch {
	case score < 0:
		return 0
	case score > 100:
		return 100
	default:
		return score
	}
}
