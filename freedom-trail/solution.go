package freedomtrail

import "math"

func findRotateSteps(ring string, key string) int {
	var minSteps func(freedomTrailRing, int) int

	cache := make(map[[2]int]int)

	minSteps = func(ftr freedomTrailRing, keyPos int) int {
		if keyPos == len(key) {
			return 0
		}

		cacheKey := [2]int{ftr.Pos(), keyPos}
		if cached, ok := cache[cacheKey]; ok {
			return cached
		}

		result := math.MaxInt32

		for _, nextPos := range ftr.CharPos(key[keyPos]) {
			stepsToSpellChar := ftr.MinSteps(nextPos) + 1
			remainingSteps := minSteps(ftr.WithPos(nextPos), keyPos+1)

			result = min(result, stepsToSpellChar+remainingSteps)
		}

		cache[cacheKey] = result

		return result
	}

	var ftr freedomTrailRing
	ftr.Init(ring)

	return minSteps(ftr, 0)
}

type freedomTrailRing struct {
	charPos   map[byte][]int
	pos, size int
}

func (ftr *freedomTrailRing) Init(ring string) {
	ftr.pos = 0
	ftr.size = len(ring)

	ftr.charPos = make(map[byte][]int)
	for i := 0; i < len(ring); i++ {
		ftr.charPos[ring[i]] = append(ftr.charPos[ring[i]], i)
	}
}

func (ftr *freedomTrailRing) Pos() int {
	return ftr.pos
}

func (ftr *freedomTrailRing) WithPos(pos int) freedomTrailRing {
	return freedomTrailRing{ftr.charPos, pos, ftr.size}
}

func (ftr *freedomTrailRing) CharPos(char byte) []int {
	return ftr.charPos[char]
}

func (ftr *freedomTrailRing) MinSteps(target int) int {
	var stepsBetween int
	if ftr.pos < target {
		stepsBetween = target - ftr.pos
	} else {
		stepsBetween = ftr.pos - target
	}
	return min(stepsBetween, ftr.size-stepsBetween)
}
