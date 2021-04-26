package hurricane

import "math"

/* Constants
SPH: Standard Project Hurricane
PMH: Probable Maximum Hurricane
Pressure profile equation
(p-Cp)/(Pw - Cp) = e ^ (-R/r)
Pw: Peripheral Pressure, pressure at edge of storm, should be a bit below MSLP
Cp: Central Pressure (P0 in paper)
Rmax: Radius of Maximum Winds (R in paper)
Fspeed: Forward speed of hurricane center (T in paper)
Dir: Track direction
Vgx: Maximum Gradient Winds
Rho0: Surface air density
r: distance (radius) from hurricane center
fcorr: Coriolis parameter, dependant on latitude
Vx: Observed maximum 10-m, 10-min winds over open water.  75% to 105% of Vgx.  Standard is 95%
For moving hurricane: Vx = 0.95 * Vgx + (1.5 * T ^ 0.63 * To ^ 0.37 * cos(beta)
Vpt: 10m, 10min winds at a point (V in paper)
*/

const Pw_SPH_kPa float64 = 100.8
const Pw_PMH_kPa float64 = 102.0
const Pw_SPH_inhg float64 = 29.77
const Pw_PMH_inhg float64 = 30.12
const Rho0_kPa float64 = 101.325  // Mean Sea Level Pressure
const KmToNmi float64 = 0.539957
const MpsToKts float64 = 1.94384
const KpaToInhg float64 = 0.2953
const MbToInhg float64 = 0.02953

func linearInterpolation(x float64, x1 float64, x2 float64, y1 float64, y2 float64) float64 {
	return (((y2 - y1) / (x2 - x1)) * (x - x1)) + y1
}

// radialDecay Calculates the radial decay factor for a given radius, between 0.0 and 1.0.
// When rMaxNmi < rNmi: NWS 23 pdf page 53, page 27, Figure 2.12, empirical fit.
// When rMaxNmi > rNmi: NWS 23 pdf page 54, page 28, Figure 2.13, empirical fit (logistic regression).
//
// rNmi: Point radius from center of storm in nautical miles
//
// rMaxNmi Radius of maximum winds in nautical miles
//
// return 0 <= radial decay <= 1
func radialDecay(rNmi float64, rMaxNmi float64) float64 {
	ret := 1.0

	if rMaxNmi < rNmi {
		// NWS 23 pdf page 53
		slope := (-0.051 * math.Log(rMaxNmi)) - 0.1757
		intercept := (0.4244 * math.Log(rMaxNmi)) + 0.7586
		ret = (slope * math.Log(rNmi)) + intercept
	}
	// Skip this else block as a concession for modeling time series, where everything within the max wind radius is
	//	expected to experience the max wind radius while the storm translates
	// else {

		// NWS 23 pdf page 54
		// ret = 1.01231578 / (1 + math.exp(-8.612066494 * ((r_nmi / float(rmax_nmi)) - 0.678031222)))
		// ret = 1
	// }

	// clamp radial decay between 0 and 1
	return math.Max(math.Min(ret, 1.0), 0.0)
}


// Calculate the coriolis factor for a given latitude
//
// latDeg: latitude in degrees
//
// return coriolis factor in hr^-1
func coriolisFrequency(latDeg float64) float64 {
	w := 2.0 * math.Pi / 24
	return 2.0 * w * math.Sin(latDeg * math.Pi / 180.0)
}