package features

type Env struct {
	Gravity Tuple
	Wind    Tuple
}

type Proj struct {
	Position Tuple
	Velocity Tuple
}

func Tick(env Env, proj Proj) Proj {
	pos := Add(proj.Position, proj.Velocity)
	vel := Add(Add(proj.Velocity, env.Gravity), env.Wind)

	ret := Proj{Position: pos, Velocity: vel}

	return ret
}

