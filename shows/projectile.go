package shows

import (
	"fmt"

	"github.com/zealsham/raytrace/tuple"
)

type projectile struct {
	position tuple.Tuple
	velocity tuple.Tuple
}

type enviroment struct {
	gravity tuple.Tuple
	wind    tuple.Tuple
}

func tick(env enviroment, proj projectile) projectile {
	pos := tuple.AddTuple(proj.position, proj.velocity)
	velocity := tuple.AddTuple(proj.velocity, tuple.AddTuple(env.gravity, env.wind))
	return projectile{pos, velocity}
}

func Run() {
	p := projectile{tuple.Point(0, 1, 0), tuple.NormalizeTuple(tuple.Vector(1, 1, 0))}
	e := enviroment{tuple.Vector(0, -0.1, 0), tuple.Vector(-0.01, 0, 0)}
	for i := 0; p.position.Y >= 0; i++ {
		fmt.Printf("Position after tick %d: is x=%4.2f , y=%4.2f, z=%4.2f \n", i, p.position.X, p.position.Y, p.position.Z)

		p = tick(e, p)
	}

}
