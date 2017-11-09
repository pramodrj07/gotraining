// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mathext

import (
	"math"
)

// CompleteK computes the complete elliptic integral of the 1st kind, 0≤m≤1.
// It returns math.NaN() if m is not in [0,1].
//
//	K(m) = \int_{0}^{\pi/2} 1 / {\sqrt{1-m{\sin^2\theta}}} d\theta
//
// See: http://dx.doi.org/10.1016/j.cam.2014.12.038 for the computation method.
func CompleteK(m float64) float64 {
	if m < 0 || 1 < m || math.IsNaN(m) {
		return math.NaN()
	}

	mc := 1 - m

	if mc > 0.592990 {
		t := 2.45694208987494165*mc - 1.45694208987494165
		p := 3703.75266375099019 + t*(5462.47093231923466+t*(2744.82029097576810+t*(543.839017382099411+t*(36.2381612593459565+t*0.393188651542789784))))
		q := 2077.94377067058435 + t*(3398.00069767755460+t*(1959.05960044399275+t*(472.794455487539279+t*(43.5464368440078942+t))))
		return p / q
	}
	if mc > 0.350756 {
		t := 4.12823963605439369*mc - 1.44800482178389491
		p := 4264.28203103974630 + t*(6341.90978213264024+t*(3214.59187442783167+t*(642.790566685354573+t*(43.2589626155454993+t*0.475223892294445943))))
		q := 2125.06914237062279 + t*(3479.95663350926514+t*(2006.03187933518870+t*(482.900172581418890+t*(44.1848041560412224+t))))
		return p / q
	}
	if mc > 0.206924 {
		t := 6.95255575949719117*mc - 1.43865064797819679
		p := 4870.25402224986382 + t*(7307.18826377416591+t*(3738.29369283392307+t*(754.928587580583704+t*(51.3609902253065926+t*0.571948962277566451))))
		q := 2172.51745704102287 + t*(3565.04737778032566+t*(2056.13612019430497+t*(493.962405117599400+t*(44.9026847057686146+t))))
		return p / q
	}
	if mc > 0.121734 {
		t := 11.7384669562155183*mc - 1.42897053644793990
		p := 5514.8512729127464 + t*(8350.4595896779631+t*(4313.60788246750934+t*(880.27903031894216+t*(60.598720224393536+t*0.68504458747933773))))
		q := 2218.41682813309737 + t*(3650.41829123846319+t*(2107.97379949034285+t*(505.74295207655096+t*(45.6911096775045314+t))))
		return p / q
	}
	if mc > 0.071412 {
		t := 19.8720241643813839*mc - 1.41910098962680339
		p := 6188.8743957372448 + t*(9459.3331440432847+t*(4935.41351498551527+t*(1018.21910476032105+t*(70.981049144472361+t*0.81599895108245948))))
		q := 2260.73112539748448 + t*(3732.66955095581621+t*(2159.68721749761492+t*(517.86964191812384+t*(46.5298955058476510+t))))
		return p / q
	}
	if mc > 0.041770 {
		t := 33.7359152553808785*mc - 1.40914918021725929
		p := 6879.5170681289562 + t*(10615.0836403687221+t*(5594.8381504799829+t*(1167.26108955935542+t*(82.452856129147838+t*0.96592719058503951))))
		q := 2296.88303450660439 + t*(3807.37745652028212+t*(2208.74949754945558+t*(529.79651353072921+t*(47.3844470709989137+t))))
		return p / q
	}
	if mc > 0.024360 {
		t := 57.4382538770821367*mc - 1.39919586444572085
		p := 7570.6827538712100 + t*(11792.9392624454532+t*(6279.2661370014890+t*(1325.01058966228180+t*(94.886883830605940+t*1.13537029594409690))))
		q := 2324.04824540459984 + t*(3869.56755306385732+t*(2252.22250562615338+t*(540.85752251676412+t*(48.2089280211559345+t))))
		return p / q
	}
	if mc > 0.014165 {
		t := 98.0872976949485042*mc - 1.38940657184894556
		p := 8247.2601660137746 + t*(12967.7060124572914+t*(6974.7495213178613+t*(1488.54008220335966+t*(108.098282908839979+t*1.32411616748380686))))
		q := 2340.47337508405427 + t*(3915.63324533769906+t*(2287.70677154700516+t*(550.45072377717361+t*(48.9575432570382154+t))))
		return p / q
	}
	if mc > 0.008213 {
		t := 168.010752688172043*mc - 1.37987231182795699
		p := 8894.2961573611293 + t*(14113.7038749808951+t*(7666.5611739483371+t*(1654.60731579994159+t*(121.863474964652041+t*1.53112170837206117))))
		q := 2344.88618943372377 + t*(3942.81065054556536+t*(2313.28396270968662+t*(558.07615380622169+t*(49.5906602613891184+t))))
		return p / q
	}
	if mc > 0 {
		t := 1.0 - 121.758188238159016*mc
		p := -math.Log(mc*0.0625) * (34813.4518336350547 + t*(235.767716637974271+t*0.199792723884069485)) / (69483.5736412906324 + t*(614.265044703187382+t))
		q := -mc * (9382.53386835986099 + t*(51.6478985993381223+t*0.00410754154682816898)) / (37327.7262507318317 + t*(408.017247271148538+t))
		return p + q
	}

	return math.Inf(1)
}

// CompleteE computes the complete elliptic integral of the 2nd kind, 0≤m≤1.
// It returns math.NaN() if m is not in [0,1].
//
//	E(m) = \int_{0}^{\pi/2} {\sqrt{1-m{\sin^2\theta}}} d\theta
//
// See: http://dx.doi.org/10.1016/j.cam.2014.12.038 for the computation method.
func CompleteE(m float64) float64 {
	if m < 0 || 1 < m || math.IsNaN(m) {
		return math.NaN()
	}

	mc := 1 - m

	if mc > 0.566638 {
		t := 2.30753965506897236*mc - 1.30753965506897236
		p := 19702.2363352671642 + t*(31904.1559574281609+t*(18177.1879313824040+t*(4362.94760768571862+t*(409.975559128654710+t*10.3244775335024885))))
		q := 14241.2135819448616 + t*(20909.9899599927367+t*(10266.4884503526076+t*(1934.86289070792954+t*(117.162100771599098+t))))
		return p / q
	}
	if mc > 0.315153 {
		t := 3.97638030101198879*mc - 1.25316818100483130
		p := 16317.0721393008221 + t*(26627.8852140835023+t*(15129.4009798463159+t*(3574.15857605556033+t*(326.113727011739428+t*7.93163724081373477))))
		q := 13047.1505096551210 + t*(19753.5762165922376+t*(9964.25173735060361+t*(1918.72232033637537+t*(117.670514069579649+t))))
		return p / q
	}
	if mc > 0.171355 {
		t := 6.95419964116329852*mc - 1.19163687951153702
		p := 13577.3850240991520 + t*(22545.4744699553993+t*(12871.9137872656293+t*(3000.74575264868572+t*(263.964361648520708+t*6.08522443139677663))))
		q := 11717.3306408059832 + t*(18431.1264424290258+t*(9619.40382323874064+t*(1904.06010727307491+t*(118.690522739531267+t))))
		return p / q
	}
	if mc > 0.090670 {
		t := 12.3938774245522712*mc - 1.12375286608415443
		p := 11307.9485341543712 + t*(19328.6173704569489+t*(11208.6068472959372+t*(2596.54874477084334+t*(219.253495956962613+t*4.66931143174036616))))
		q := 10307.6837501971393 + t*(16982.2450249024383+t*(9241.7604666150102+t*(1893.41905403040679+t*(120.498555754227847+t))))
		return p / q
	}
	if mc > 0.046453 {
		t := 22.6157360291290680*mc - 1.05056878576113260
		p := 9383.1490856819874 + t*(16718.9730458676860+t*(9977.2498973537718+t*(2323.49987246555537+t*(188.618148076418837+t*3.59313532204509922))))
		q := 8877.1964704758383 + t*(15450.0537230364062+t*(8840.2771293410661+t*(1889.13672102820913+t*(123.422125687316355+t))))
		return p / q
	}
	if mc > 0.022912 {
		t := 42.4790790535661187*mc - 0.973280659275306911
		p := 7719.1171817802054 + t*(14521.7363804934985+t*(9045.3996063894006+t*(2149.92068078627829+t*(169.386557799782496+t*2.78515570453129137))))
		q := 7479.7539074698012 + t*(13874.4978011497847+t*(8420.3848818926324+t*(1892.69753150329759+t*(127.802109608726363+t))))
		return p / q
	}
	if mc > 0.010809 {
		t := 82.6241427745187144*mc - 0.893084359249772784
		p := 6261.6095608987273 + t*(12593.0874916293982+t*(8304.3265605809870+t*(2048.68391263416822+t*(159.371262600702237+t*2.18867046462858104))))
		q := 6156.4532048239501 + t*(12283.8373999680518+t*(7979.7435857665227+t*(1903.60556312663537+t*(133.911640385965187+t))))
		return p / q
	}
	if mc > 0.004841 {
		t := 167.560321715817694*mc - 0.811159517426273458
		p := 4978.06146583586728 + t*(10831.7178150656694+t*(7664.6703673290453+t*(1995.66437151562090+t*(156.689647694892782+t*1.75859085945198570))))
		q := 4935.56743322938333 + t*(10694.5510113880077+t*(7506.8028283118051+t*(1918.38517009740321+t*(141.854303920116856+t))))
		return p / q
	}
	if mc > 0 {
		t := 1.0 - 206.568890725056806*mc
		p := -mc * math.Log(mc*0.0625) * (41566.6612602868736 + t*(154.034981522913482+t*0.0618072471798575991)) / (165964.442527585615 + t*(917.589668642251803+t))
		q := (132232.803956682877 + t*(353.375480007017643-t*1.40105837312528026)) / (132393.665743088043 + t*(192.112635228732532-t))
		return p + q
	}

	return 1
}