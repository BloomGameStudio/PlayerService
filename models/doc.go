// Package models provides data models for the player service.
// It includes various models representing different entities such as players, positions, levels, states etc.
// The package consists of the following files:
//   - coordinate.go: Defines the Coordinate struct for representing three-dimensional coordinates.
//   - generic.go: Contains the Vector3 struct, which serves as a generic model that includes a gorm.Model and
//                  publicModels.Vector3.
//   - grid.go: Defines the Grid struct
//   - level.go: Defines the Level struct, which includes a gorm.Model, PlayerID, and publicModels.Level. Additionally,
//               it includes a IsValid() method for validating the state of the level.
//   - orientation.go: Defines the Orientation struct, which represents the orientation of an entity in three-dimensional
//                     space using horizontal and vertical bearings.
//   - player.go: Defines the Player struct, which includes a gorm.Model, UserID, publicModels.Player, Transform,
//                an array of State, and Level. It also includes methods such as IsValid() and GetPosition() for validating
//                the player and retrieving its position.
//   - position.go: Defines the Position struct, which includes a gorm.Model and publicModels.Position. It also includes
//                  methods such as IsValid() and GetPosition() for validating the position and retrieving it.
//   - rotation.go: Defines the Rotation struct, which includes a gorm.Model and publicModels.Rotation. It also includes
//                  a IsValid() method for validating the rotation.
//   - scale.go: Defines the Scale struct, which includes a gorm.Model and publicModels.Scale. It also includes a
//               IsValid() method for validating the scale.
//   - state.go: Defines the State struct, which includes a gorm.Model, PlayerID, and publicModels.State. It also includes
//               a IsValid() method for validating the state.
//   - transform.go: Defines the Transform struct, which includes RotationID, ScaleID, PositionID, Position, Rotation, and
//                   Scale. Currently, validation methods are commented out in the file.
//
// For more information on each model and its associated methods, please refer to their individual documentation.
package models