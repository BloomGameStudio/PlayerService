// Package publicModels provides public data models for the player service.
// It includes various models representing different entities such as players, positions, levels, states, etc.
// The package depends on the `gorm.io/gorm` package for database operations.
//
// Dependencies:
//  - gorm.io/gorm
//
// The package consists of the following files:
//   - generic.go: Contains the Vector3 struct, which serves as a generic model for representing three-dimensional vectors.
//   - level.go: Defines the Level struct, which holds the public fields for the Level model. It includes the LevelID field
//                and a IsValid() method for validating the level.
//   - player.go: Defines the Player struct, which holds the public fields for the Player model. It includes fields such as
//                Name, Layer, ENS, Active, Level, Transform, and an array of State. Additionally, it includes a IsValid()
//                method for validating the player.
//   - position.go: Defines the Position struct, which includes the Vector3 struct and methods such as IsValid() for
//                  validating the position.
//   - rotation.go: Defines the Rotation struct, which includes the Vector3 struct, the W field, and a IsValid() method
//                  for validating the rotation.
//   - scale.go: Defines the Scale struct, which includes the Vector3 struct and a IsValid() method for validating the scale.
//   - state.go: Defines the State struct, which holds the public fields for the State model. It includes the StateID and
//               Value fields, as well as a IsValid() method for validating the state.
//   - transform.go: Defines the Transform struct, which includes the Position, Rotation, and Scale structs. It also
//                   includes a IsValid() method for validating the transform.
//
// For more information on each model and its associated methods, please refer to their individual documentation.
package publicModels