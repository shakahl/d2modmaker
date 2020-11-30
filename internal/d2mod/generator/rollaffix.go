package generator

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores"
	"github.com/tlentz/d2modmaker/internal/d2fs/txts/propscores/propscorespartype"
	"github.com/tlentz/d2modmaker/internal/d2mod/d2items"
	"github.com/tlentz/d2modmaker/internal/util"
	"github.com/tlentz/d2modmaker/internal/weightrand"
)

const (
	// PropVariance controls the range between min & max that a randomly rolled prop/affix will have
	PropVariance = 0.2
)

// RollAffix Randomly roll a new Prop for a given item
func RollAffix(g *Generator, item d2items.Item, sbm float32, colIdx int, targetPropScore int, w *weightrand.Weights) *d2items.Affix {
	targetPropScore = util.Round32(float32(targetPropScore) * (1 / sbm)) // Adjust the targetPropScore up in order to compensate for sbm

	for rollcounter := 0; rollcounter < 500; rollcounter++ {
		scoreFileRowIndex := w.Generate()
		line := g.psi.RowLines[scoreFileRowIndex]
		//log.Printf("RollAffix: %s: ScoreMax=%d T:%d TargetScore=%d\n", l.Prop.Name, l.ScoreMax, l.PropParType, targetPropScore)
		if line.ScoreMax < (targetPropScore / 6) {
			//log.Print("<Reroll S>")
			continue
		}
		if !d2items.CheckIETypes(g.TypeTree, item.Types[0], line.Itypes, line.Etypes) {
			//log.Print("<Reroll T>")
			continue
		}
		if item.Lvl < line.MinLvl {
			//log.Print("<Reroll L>")
			continue
		}
		newa := d2items.NewAffixFromLine(line, colIdx, sbm)

		switch line.PropParType {
		case propscorespartype.R, propscorespartype.Rp, propscorespartype.Rt, propscorespartype.Smm, propscorespartype.C:
			//fmt.Printf("RollAffix: %d %d-%d -> ", targetPropScore, newa.P.Val.Min, newa.P.Val.Max)
			newa.P.Val.Min, newa.P.Val.Max = rollRange(newa.P.Val.Min, newa.P.Val.Max, line, item.Lvl, targetPropScore)
			newa.P.Min = strconv.Itoa(newa.P.Val.Min)
			newa.P.Max = strconv.Itoa(newa.P.Val.Max)
			//fmt.Printf("%s-%s\n", newa.P.Min, newa.P.Max)
		case propscorespartype.Lvl: // (pts or %)/lvl prop min & max are empty/ignored
			newa.P.Val.Par = rollMax(newa.P.Val.Par, line, item.Lvl, targetPropScore)
			newa.P.Par = strconv.Itoa(newa.P.Val.Par)
		case propscorespartype.Scl: // Skill, %chance, Level
			// OBC: I'm too lazy to roll both chance and level as that makes the targetscore calculation
			// a lot more complex, so roll either Chance or Level.
			if rand.Intn(4) >= 2 {
				// Roll Chance
				newa.P.Val.Min = rollMax(newa.P.Val.Min, line, item.Lvl, targetPropScore)
				newa.P.Min = strconv.Itoa(newa.P.Val.Min)
			} else {
				// Roll Level
				newa.P.Val.Max = rollMax(newa.P.Val.Max, line, item.Lvl, targetPropScore)
				newa.P.Max = strconv.Itoa(newa.P.Val.Max)
			}
		case propscorespartype.Sch: // Skill, #charges, Level
			// Don't touch # charges... roll for Level
			newa.P.Val.Max = rollMax(newa.P.Val.Max, line, item.Lvl, targetPropScore)
			newa.P.Max = strconv.Itoa(newa.P.Val.Max)
		case propscorespartype.S:
			// Nothing to adjust here
		default:
			log.Fatalf("calcPropScore: Unhandled Prop type %s:%d\n", newa.P.Name, line.PropParType)

		}
		return newa
	}
	log.Fatal("RollProp: Too many rolls (500).. couldn't find a a valid PropScore.txt line")
	panic(1)
}

// rollRange:  Calculate a min & max value +/- PropVariance centered around targetPropScore
func rollRange(min int, max int, line *propscores.Line, itemLvl int, targetPropScore int) (int, int) {
	newAvg := util.Interpolate(targetPropScore, targetPropScore, line.ScoreMin, line.ScoreMax, min, max)
	//fmt.Printf("rollRange: %d %d-%d -> %d %d\n", targetPropScore, line.ScoreMin, line.ScoreMax, min, max)
	variance := util.Round32(float32((max - min)) * (PropVariance / 2.0))
	newMax := newAvg + variance
	newMin := newAvg - variance

	newMin = util.MaxInt(newMin, min)
	newMax = util.MinInt(newMax, max)

	return newMax, newMin
}

// rollMax:  Roll for a maximum value within 20% (PropVariance = 0.2) of the targetPropScore
func rollMax(max int, line *propscores.Line, itemLvl int, targetPropScore int) int {
	if max < 0 {
		log.Fatalf("rollMax: upper limit cannot be negative: Prop:%s: %d", line.Prop.Name, max)
	}
	if targetPropScore > util.Round32(float32(line.ScoreMax)*(1.0-PropVariance)) {
		return max
	}
	newMax := util.Interpolate(targetPropScore, targetPropScore, line.ScoreMin, line.ScoreMax, 0, max)
	targetvariance := float32(newMax) * PropVariance
	newMax += util.Round32((rand.Float32() * targetvariance * 2) - targetvariance)
	if newMax > max {
		newMax = line.Prop.Val.Max
	}
	if newMax < 0 {
		// 0% str per level would have no effect at all  If score ends up being too high then the affix will ge rerolled
		newMax = 1
	}
	return newMax
}
