# PlayerService

## Quickstart Guide:



## Installation
### **Native Bare Metel**

        // Install Deps/Modules
        go mod download

        // Run the server on port 1323
        go run main.go


**Optional Live Code Reloading with Air**

Install Air via your preffered installation method: https://github.com/cosmtrek/air

        // Run the server on port 1323 with live code reloading
        air


### **Using Docker && DockerCompose**
---
### **Docker**

        // With build tag/name

        docker build -t playerservice .

        docker run -p 1323:1323 playerservice


        // Without tag/name

        docker build .

        docker run -p 1323:1323 <Containername>
---
###  **Docker-Compose**
Chose your docker compose cli 
Depending on what version you have or how you installed docker compose.

The examples will use the more wider used `docker-compose`

For more Information read: https://stackoverflow.com/questions/66514436/difference-between-docker-compose-and-docker-compose

The more wider used `docker-compose`.
  
        
        docker-compose <command>
  
 
The newer `docker compose`.

        docker compose <command>
         

**Start the App and listen on port 1323**
   
        docker-compose up
        // CTRL + C to stop 
        

---

**Rotation ressources**
- https://compsci290-s2016.github.io/CoursePage/Materials/EulerAnglesViz/index.html
- https://www.youtube.com/watch?v=2Cwa6hfn2K0
- https://docs.unity3d.com/ScriptReference/Transform-eulerAngles.html
- https://docs.unity3d.com/ScriptReference/Quaternion.html


---
## How to Interact with the Player WebSocket (In Development)

Assuming standard config and hosting locally.


  1. If not present create a player in the players table. The player needs to have a usable UserID 
  
        This can be be done either by:

        - Manually creating it in the db
        - Creating it throug the /player CreatePlayer endpoint

  
  2.  The Websocket needs to derive a models.Player.UserID from a JWT. That UserID has to match with the UserID of the player in the database That JWT can be created either by:
   
        - The Userservice
        - Manually
   
  3.  Connect to the websocket with the JWT in the Header as shown below:
   
      
                // Token has been shortened for readability 
                Authorization: Bearer eyJ_A 

                ws://localhost:1323/ws/player

  4. Send and receive JSON Player objects from the websocket


---
## OpenAPI
The OpenAPI spec: [open_api.yml](docs/open_api.yml)

can be found in docs dir and is generated from the postman collection in the same dir via:

https://joolfe.github.io/postman-to-openapi/ 

        
        p2o PlayerService.postman_collection.json -f open_api.yml


## Guides

**How to setup private multiplayer**

1. Create a Player in the DB. More Ressources than the below example are availabe either in the docs folder or in this Readme.
   
   - **How To Create a Player:**

        Send a `POST` request to http://staging.player.bloomstudio.gg/player with a adjusted Player Object.

        Player Object Example:

        Change "UserID" and "Name" 
        ```json
        {
        "UserID": "33b7e1f3-6f8e-40b9-97dc-c54d9162vb05",
        "Name": "User1",
        "Layer": "layer1",
        "Position": {
                "x": 1,
                "y": 2,
                "z": 3
        },
        "Rotation": {
                "x": 4,
                "y": 5,
                "z": 6
        },
        "Scale": {
                "x": 7,
                "y": 8,
                "z": 9
        }
        }
        ```

2. Connect to the ws player websocket endpoint. Wss might also work. The following link is for reference and will connect to the remotely hosted staging player service and player endpoint 
             
             ws://staging.player.bloomstudio.gg/ws/player

3. You will now receive a list of player objects on that websocket. Example output can be found below.
   ```json

                
        [
                {
                        "ID": 1,
                        "CreatedAt": "2022-11-19T16:52:17.036734453+01:00",
                        "UpdatedAt": "2022-11-22T22:29:57.585125754+01:00",
                        "DeletedAt": null,
                        "UserID": "216f02a1-e252-4905-a300-69bc3aeb0cc1",
                        "name": "User1",
                        "layer": "",
                        "PositionID": 1,
                        "position": {
                                "ID": 1,
                                "CreatedAt": "2022-11-19T16:52:17.036154163+01:00",
                                "UpdatedAt": "2022-11-22T22:29:57.584197333+01:00",
                                "DeletedAt": null,
                                "x": 0.8777,
                                "y": 1.55555,
                                "z": 3.33333
                        },
                        "RotationID": 1,
                        "rotation": {
                                "ID": 1,
                                "CreatedAt": "2022-11-19T16:52:17.036448333+01:00",
                                "UpdatedAt": "2022-11-22T22:29:57.584519033+01:00",
                                "DeletedAt": null,
                                "x": 4,
                                "y": 5,
                                "z": 6,
                                "w": 0
                        },
                        "ScaleID": 1,
                        "scale": {
                                "ID": 1,
                                "CreatedAt": "2022-11-19T16:52:17.036644533+01:00",
                                "UpdatedAt": "2022-11-22T22:29:57.584636463+01:00",
                                "DeletedAt": null,
                                "x": 7,
                                "y": 5.444,
                                "z": 9.987
                        },
                        "ens": ""
                },
                {
                        "ID": 2,
                        "CreatedAt": "2023-08-01T14:40:44.986472893+02:00",
                        "UpdatedAt": "2023-08-01T14:40:44.986472893+02:00",
                        "DeletedAt": null,
                        "UserID": "735b2924-fa7f-4119-a0f2-8d51750c6e9e",
                        "name": "User3",
                        "layer": "",
                        "PositionID": 2,
                        "position": {
                                "ID": 2,
                                "CreatedAt": "2023-08-01T14:40:44.985785149+02:00",
                                "UpdatedAt": "2023-08-01T14:40:44.985785149+02:00",
                                "DeletedAt": null,
                                "x": 1,
                                "y": 2,
                                "z": 3
                        },
                        "RotationID": 2,
                        "rotation": {
                                "ID": 2,
                                "CreatedAt": "2023-08-01T14:40:44.986199897+02:00",
                                "UpdatedAt": "2023-08-01T14:40:44.986199897+02:00",
                                "DeletedAt": null,
                                "x": 4,
                                "y": 5,
                                "z": 6,
                                "w": 0
                        },
                        "ScaleID": 2,
                        "scale": {
                                "ID": 2,
                                "CreatedAt": "2023-08-01T14:40:44.986350052+02:00",
                                "UpdatedAt": "2023-08-01T14:40:44.986350052+02:00",
                                "DeletedAt": null,
                                "x": 7,
                                "y": 8,
                                "z": 9
                        },
                        "ens": ""
                }
        ]

    ```

4. If you want to modify your player or any other players you can modify the list of player objects that you received 
and send the singel modified (Not the entire list) player object back within the same websocket your received it from to the service you received it from.
Note: Sending the entire list of player objects back for update will fail and is currently not supported. You have to send a singel player object back.

   - **How to Update The Player object/s.**
    
        1. Modify and make the desired changes to the list of of Player objects that you have received in the previous step.
        2. Push a modified single player object(Not a list) to the [ws://staging.player.bloomstudio.gg/ws/player](ws://staging.player.bloomstudio.gg/ws/player) websocket 
        3. Optional Confirm that the changes took place by looking at the next data push from the websocket.

