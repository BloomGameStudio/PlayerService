package models

//should be an enum
/*
Grounded,
    Airborne,
    Waterborne,
    Mounted,

    //Movement
    Idle,
    Walking,
    Running,

    //Grounded/Airborne
    Dashing,
    Blinking,
    Shrinking,

    //Grounded
    Crouched,

    //Airborne
    Jumping,
    Gliding,
    Falling,
    Pound,

    //Water
    WaterSurface,
    Underwater,
    FallingIntoWater,
    Swimming,
    WaterborneIdle,
    WaterRotatingAngleUp,
    WaterRotatingAngleDown,

    WaterDive,
    SwimStroke,
  
    //Remaining & General
    Defending,

    AttackingMeleeQuick,
    AttackingProjectileFar,
    AttackingProjectileNear
*/
type State struct {
	Grounded, Airborn, Waterborn bool
}

func (s State) CanEnter() bool {

	return true
}
