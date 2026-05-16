package diagnose

import "time"

type Band string

const (
	BandOK       Band = "OK"
	BandWarn     Band = "WARN"
	BandCritical Band = "CRITICAL"
)

type Severity string

const (
	SeverityInfo     Severity = "INFO"
	SeverityWarn     Severity = "WARN"
	SeverityCritical Severity = "CRIT"
)

type Checker interface {
	Check(Inventory) CategoryResult
}

type CategoryResult struct {
	Name          string
	Weight        float64
	CoverageScore int
	QualityScore  int
	Score         int
	Summary       string
	Findings      []Finding
}

type Finding struct {
	Pillar         string    `json:"pillar"`
	Severity       Severity  `json:"severity"`
	Title          string    `json:"title"`
	Recommendation string    `json:"recommendation"`
	Evidence       []Snippet `json:"evidence,omitempty"`
}

type Snippet struct {
	Path      string `json:"path"`
	StartLine int    `json:"startLine"`
	EndLine   int    `json:"endLine"`
	Content   string `json:"content"`
}

type Report struct {
	ProjectSlug  string           `json:"projectSlug"`
	Path         string           `json:"path"`
	ScannedAt    time.Time        `json:"scannedAt"`
	Elapsed      time.Duration    `json:"elapsed"`
	FileCount    int              `json:"fileCount"`
	GlobalScore  int              `json:"globalScore"`
	Band         Band             `json:"band"`
	Categories   []CategoryResult `json:"categories"`
	Improvements []Finding        `json:"improvements"`
}

func BandForScore(score int) Band {
	switch {
	case score >= 85:
		return BandOK
	case score >= 60:
		return BandWarn
	default:
		return BandCritical
	}
}
