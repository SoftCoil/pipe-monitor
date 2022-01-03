package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stats struct {
	Name      string
	Size      int64
	Written   int64
	Buffered  int
	Read      int
	Format    string
	StartTime time.Time
	lastTime  time.Time
}

func (stats *Stats) Record(written int) {
	stats.Written += int64(written)
}

func (stats *Stats) WriteStats(force bool) {
	minDuration, _ := time.ParseDuration("2s")
	now := time.Now()

	if force || now.Sub(stats.lastTime) >= minDuration {
		stats.lastTime = now

		exteranlFormatString := stats.defaultFormatString()
		if len(stats.Format) > 0 {
			exteranlFormatString = stats.Format
		}
		statsString := stats.createStatsString(exteranlFormatString)
		fmt.Fprintln(os.Stderr, statsString)
	}
}

func (stats *Stats) createStatsString(externalFormatString string) string {
	replacer := strings.NewReplacer(stats.toNameValueArray()...)
	return replacer.Replace(externalFormatString)
}

func (stats *Stats) defaultFormatString() string {

	if len(stats.Format) > 0 {
		return stats.Format
	}

	fs := ""
	if len(stats.Name) > 0 {
		fs += "%name: "
	}

	fs += "Processed %written bytes"

	if stats.Size > 0 {
		fs += " of %size (%percent complete)"
	}

	fs += ". %buffered bytes buffered. Running %time"

	if stats.Size > 0 {
		fs += ", eta: %eta"
	}

	return fs
}

func (stats *Stats) toNameValueArray() []string {
	statsArray := []string{
		"%name", stats.Name,
		"%size", stats.sizeString(),
		"%time", stats.timeString(),
		"%eta", stats.etaString(),
		"%percent", stats.percentString(),
		"%written", strconv.FormatInt(stats.Written, 10),
		"%buffered", strconv.Itoa(stats.Buffered),
	}

	return statsArray
}

func (stats *Stats) sizeString() string {
	if stats.Size > 0 {
		return fmt.Sprintf("%d", stats.Size)
	}
	return ""
}

func (stats *Stats) percent() float64 {
	if stats.Size == 0 {
		return 0
	}

	return float64(100) * float64(stats.Written) / float64(stats.Size)
}

func (stats *Stats) percentString() string {
	if stats.Size > 0 {
		return fmt.Sprintf("%d%%", int(stats.percent()))
	}
	return ""
}

func (stats *Stats) timeString() string {
	now := time.Now()
	runningTime := now.Sub(stats.StartTime)

	return fmt.Sprint(runningTime.Round(time.Second))
}

func (stats *Stats) etaString() string {
	percent := stats.percent()

	if stats.Size == 0 || percent == 0 {
		return "<unknown>"
	}

	if percent == 100 {
		return "0s"
	}

	since := time.Since(stats.StartTime)

	millis := int64((float64(since.Milliseconds()) / percent) * (float64(100) - percent))

	eta := time.Duration(millis) * time.Millisecond

	return fmt.Sprint(eta.Round(time.Second))
}
