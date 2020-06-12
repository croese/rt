package main

import (
	"fmt"

	"github.com/croese/raytrace"
)

type projectile struct {
	position raytrace.Tuple4
	velocity raytrace.Tuple4
}

type environment struct {
	gravity raytrace.Tuple4
	wind    raytrace.Tuple4
}

func main() {
	p := projectile{
		position: raytrace.NewPoint(0, 1, 0),
		velocity: raytrace.NewVector(1, 1, 0).Normalize(),
	}

	e := environment{
		gravity: raytrace.NewVector(0, -0.1, 0),
		wind:    raytrace.NewVector(-0.01, 0, 0),
	}

	tickCount := 0
	for p.position.Y() > 0 {
		p = tick(e, p)
		tickCount++
		fmt.Printf("after tick %d, projectile = %v\n", tickCount, p)
	}
	fmt.Printf("the projectile took %d ticks to reach the ground\n", tickCount)
}

func tick(env environment, proj projectile) projectile {
	position := proj.position.Plus(proj.velocity)
	velocity := proj.velocity.Plus(env.gravity).Plus(env.wind)
	return projectile{position, velocity}
}
