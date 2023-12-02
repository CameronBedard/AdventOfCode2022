package solutions2022

import (
	"bufio"
	"fmt"
	"os"
)

var globalBest = 0

func day19p1() {
	// Open the file.
	file, err := os.Open("/Users/cameron.bedard/Documents/FunRepos/AdventOfCode2022/inputs2022/day19.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	result := 0

	// Create a scanner to read the file.
	scanner := bufio.NewScanner(file)
	// Read the file line by line.
	for scanner.Scan() {
		var id int
		oreRobot := oreRobot{}
		clayRobot := clayRobot{}
		obsRobot := obsRobot{}
		geodeRobot := geodeRobot{}

		fmt.Sscanf(scanner.Text(),
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&id, &oreRobot.oreCost, &clayRobot.oreCost, &obsRobot.oreCost, &obsRobot.clayCost, &geodeRobot.oreCost, &geodeRobot.obsCost)

		bp := blueprint{
			id,
			oreRobot,
			clayRobot,
			obsRobot,
			geodeRobot,
		}

		// ore, clay, obsidian, geodes
		/*robots := [4]int{1, 0, 0, 0}
		minerals := [4]int{0, 0, 0, 0}

		globalMax := 0
		blueprintMax := DFS(bp, robots, minerals, 24, &globalMax)
		total += bp.id * blueprintMax
		fmt.Println(total, globalMax, blueprintMax)*/

		result += bp.id * search(bp, 0, 0, 0, 24, 1, 0, 0, 0, 0)
		globalBest = 0
		fmt.Println(result)
	}

	fmt.Println("total:", result)

	// Check for any errors that occurred while reading the file.
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
}

type blueprint struct {
	id         int
	oreRobot   oreRobot
	clayRobot  clayRobot
	obsRobot   obsRobot
	geodeRobot geodeRobot
}

type oreRobot struct {
	oreCost int
}

type clayRobot struct {
	oreCost int
}

type obsRobot struct {
	oreCost, clayCost int
}

type geodeRobot struct {
	oreCost, obsCost int
}

func DFS(bp blueprint, robots [4]int, minerals [4]int, time int, globalMax *int) int {
	if time == 0 || *globalMax >= minerals[3]+sum(robots[3], robots[3]+time-1) {
		return 0
	}
	if robots[0] >= bp.geodeRobot.oreCost && robots[2] >= bp.geodeRobot.obsCost {
		return sum(robots[3], robots[3]+time-1)
	}

	oreLimitHit := robots[0] >= max(bp.geodeRobot.oreCost, max(bp.clayRobot.oreCost, bp.obsRobot.oreCost))
	clayLimitHit := robots[1] >= bp.obsRobot.clayCost
	obsLimitHit := robots[2] >= bp.geodeRobot.obsCost
	localMax := 0

	//spend our minute building + collecting
	if !oreLimitHit {
		localMax = max(localMax,
			DFS(
				bp,
				[4]int{robots[0], robots[1], robots[2], robots[3]},
				[4]int{minerals[0] + robots[0], minerals[1] + robots[1], minerals[2] + robots[2], minerals[3] + robots[3]},
				time-1,
				globalMax))
	}
	if !oreLimitHit && minerals[0] >= bp.oreRobot.oreCost {
		localMax = max(localMax,
			DFS(
				bp,
				[4]int{robots[0] + 1, robots[1], robots[2], robots[3]},
				[4]int{minerals[0] + robots[0] - bp.oreRobot.oreCost, minerals[1] + robots[1], minerals[2] + robots[2], minerals[3] + robots[3]},
				time-1,
				globalMax))
	}
	if !clayLimitHit && minerals[0] >= bp.clayRobot.oreCost {
		localMax = max(localMax,
			DFS(
				bp,
				[4]int{robots[0], robots[1] + 1, robots[2], robots[3]},
				[4]int{minerals[0] + robots[0] - bp.clayRobot.oreCost, minerals[1] + robots[1], minerals[2] + robots[2], minerals[3] + robots[3]},
				time-1,
				globalMax))
	}
	if !obsLimitHit && minerals[0] >= bp.obsRobot.oreCost && minerals[1] >= bp.obsRobot.clayCost {
		localMax = max(localMax,
			DFS(
				bp,
				[4]int{robots[0], robots[1], robots[2] + 1, robots[3]},
				[4]int{minerals[0] + robots[0] - bp.obsRobot.oreCost, minerals[1] + robots[1] - bp.obsRobot.clayCost, minerals[2] + robots[2], minerals[3] + robots[3]},
				time-1,
				globalMax))
	}
	if minerals[0] >= bp.geodeRobot.oreCost && minerals[2] >= bp.geodeRobot.obsCost {
		localMax = max(localMax,
			DFS(
				bp,
				[4]int{robots[0], robots[1], robots[2], robots[3] + 1},
				[4]int{minerals[0] + robots[0] - bp.geodeRobot.oreCost, minerals[1] + robots[1], minerals[2] + robots[2] - bp.geodeRobot.obsCost, minerals[3] + robots[3]},
				time-1,
				globalMax))
	}

	*globalMax = max(*globalMax, localMax)
	return localMax
}

func sum(from int, to int) int {
	return (to * (to + 1) / 2) - ((from - 1) * from / 2)
}

func search(bp blueprint, ore, clay, obs, time, oreRobots, clayRobots, obsRobots, geodeRobots, geodes int) int {
	if time == 0 || globalBest >= geodes+sum(geodeRobots, geodeRobots+time-1) {
		return 0
	}
	if oreRobots >= bp.geodeRobot.oreCost && obsRobots >= bp.geodeRobot.obsCost {
		return sum(geodeRobots, geodeRobots+time-1)
	}

	oreLimitHit := oreRobots >= max(bp.geodeRobot.oreCost, max(bp.clayRobot.oreCost, bp.obsRobot.oreCost))
	clayLimitHit := clayRobots >= bp.obsRobot.clayCost
	obsLimitHit := obsRobots >= bp.geodeRobot.obsCost
	best := 0

	if !oreLimitHit {
		best = max(
			best,
			geodeRobots+search(
				bp, ore+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.oreRobot.oreCost && !oreLimitHit {
		best = max(
			best,
			geodeRobots+search(
				bp, ore-bp.oreRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots+1, clayRobots, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.clayRobot.oreCost && !clayLimitHit {
		best = max(
			best, geodeRobots+search(
				bp, ore-bp.clayRobot.oreCost+oreRobots, clay+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots+1, obsRobots, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.obsRobot.oreCost && clay >= bp.obsRobot.clayCost && !obsLimitHit {
		best = max(
			best, geodeRobots+search(
				bp, ore-bp.obsRobot.oreCost+oreRobots, clay-bp.obsRobot.clayCost+clayRobots, obs+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots+1, geodeRobots, geodes+geodeRobots))
	}
	if ore >= bp.geodeRobot.oreCost && obs >= bp.geodeRobot.obsCost {
		best = max(
			best, geodeRobots+search(
				bp, ore-bp.geodeRobot.oreCost+oreRobots, clay+clayRobots, obs-bp.geodeRobot.obsCost+obsRobots,
				time-1, oreRobots, clayRobots, obsRobots, geodeRobots+1, geodes+geodeRobots))
	}

	globalBest = max(best, globalBest)
	return best
}
