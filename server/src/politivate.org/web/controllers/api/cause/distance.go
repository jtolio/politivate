package cause

import (
	"math"
	"time"
)

const (
	// We want to tell users that they can only complete a challenge once every
	// day, so that's what we'll message, but we only limit to once every 18 hours.
	// This way, if someone calls once late in the day and wants to call earlier
	// the next day, they can.
	MinChallengeInterval = time.Hour * 18

	EarthRadius = 6378137
)

func square(x float64) float64 { return x * x }

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	latr1 := lat1 * math.Pi / 180
	latr2 := lat2 * math.Pi / 180

	dlatr := (lat2 - lat1) * math.Pi / 180
	dlonr := (lon2 - lon1) * math.Pi / 180

	a := (square(math.Sin(dlatr/2)) +
		math.Cos(latr1)*math.Cos(latr2)*
			square(math.Sin(dlonr/2)))

	return EarthRadius * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}
