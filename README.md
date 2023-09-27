# PlayerService

The Bloom & LostLight Playerservice.

- [PlayerService](#playerservice)
  - [Quickstart Guide:](#quickstart-guide)
  - [Installation](#installation)
    - [**Native Bare Metel**](#native-bare-metel)
    - [**Using Docker \&\& DockerCompose**](#using-docker--dockercompose)
    - [**Docker-Compose**](#docker-compose)
    - [**Start the App and listen on port 1323**](#start-the-app-and-listen-on-port-1323)
    - [**Docker**](#docker)
  - [Contributors Guide](#contributors-guide)
  - [Endpoints](#endpoints)
    - [Rest Base Endpoints](#rest-base-endpoints)
    - [WebSocket Base Endpoints](#websocket-base-endpoints)
    - [Rest Endpoints](#rest-endpoints)
        - [CreatePlayer](#createplayer)
        - [GetPlayer](#getplayer)
        - [DeletePlayer](#deleteplayer)
    - [Websocket Endpoints](#websocket-endpoints)
        - [Player](#player)
        - [Position](#position)
        - [Rotation](#rotation)
        - [Scale](#scale)
        - [States](#states)
        - [Level](#level)
  - [Guides](#guides)
  - [How to Interact with the Player WebSocket](#how-to-interact-with-the-player-websocket)
  - [How to Interact with the Position WebSocket](#how-to-interact-with-the-position-websocket)
  - [Examples](#examples)
  - [Postman | OpenAPI](#postman--openapi)
  - [Rotation ressources](#rotation-ressources)

This Project keeps a [Changelog](CHANGELOG.md) in which all Versions and Notable Changes are Documentad.

## Quickstart Guide:

## Installation

### **Native Bare Metel**

        // Install Deps/Modules
        go mod download

        // Run the server on port 1323
        go run .

**Optional Live Code Reloading with Air**

Install Air via your preffered installation method: https://github.com/cosmtrek/air

        // Run the server on port 1323 with live code reloading
        air

### **Using Docker && DockerCompose**

---

### **Docker-Compose**

Chose your docker compose cli
Depending on what version you have or how you installed docker compose.

The examples will use the more wider used `docker-compose`

For more Information read: https://stackoverflow.com/questions/66514436/difference-between-docker-compose-and-docker-compose

The more wider used `docker-compose`.

        docker-compose <command>

The newer `docker compose`.

        docker compose <command>

### **Start the App and listen on port 1323**

Note: Depending on your system and context you may have to configure your image & container versions
view the official Docker Compose documentation on how Docker determines what and how it runs images & containers and how <docker-compose up> behaves.

TLDR: Docker and by extension Docker Compose will chose the latest container and if that does not exist the latest image to run your application.

**Run the latest version that was build from the branch main.**

Note: If you have build a later version or somehow else have a later version on your system a version that docker thinks is later than what was build from main it will most likely use that.
Which resulst in you not runing the version from main and not runing the intended version.
This will automatically be resolved for you if a new push to main happens.

        docker-compose up
        // CTRL + C to stop

**Run & Build the current state of the currently checkout out branch.**

Note: This will build a image and run and build a container which probably is a later version than the prebuild image built from the main branch.

        docker-compose up --build
        // CTRL + C to stop

        // If there are caching issues or some other problems or you want to be 100% sure that you run and have build the latest version of the current branch you can run:
        docker-compose up --build --force-recreate
        // This will recreate everything and might take longer.

---

### **Docker**

1.  Create a The docker volume for the database

        docker volume create playerservicevolume

2.  Run The Container

        // From Github Container Registry via Image
        // You can Replace the tag <main> at the end with whatever tag you want

                docker run --rm -p 1323:1323 -v playerservicevolume:/database ghcr.io/bloomgamestudio/playerservice:main

        // Build it yourself locally with build tag/name then run it

                docker build -t playerservice .

                docker run -p 1323:1323 -v playerservicevolume:/database playerservice


        // Build it yourself locally without tag/name then run it

                docker build .

                docker run -p 1323:1323 -v playerservicevolume:/database <Containername>

---

## Contributors Guide

1. Assign a Issue\*(Github Issue) to yourself and or clearly indicate to at least Balu if not everyone else involved that that you are now working on this task. If you lack permissions to assign a issue to yourself contact Balu or any other person with the needed permissions to assign the issue to you.
2. Continuesly Push your work to Github so that other people can follow the progress passivly. Ask for Help if you are stuck Notify Balu if you cant complete it. Ghosting is 0/10.
3. Test everything and make sure everything works as intended
4. Open a Pull Request. Ask Balu if you lack permissions.
5. Address comments on the PR if there are any.
6. Shake hands firmly.

Contributors shall work on delivering a complete contribution from start to end.

Contributing half finished and untested things is not ideal.

---

## Endpoints

Sending partial Data is Accepted and works on most Endpoints.
Indications will be made if partial Data is not supported for certain objects or Endpoints.

Sending Incorrect or Malformed JSON will always result in failure for the entire request.

### Rest Base Endpoints

Local Base Endpoint with Default Config:

- http://127.0.0.1:1323

Staging Base Endpoint:

- http://staging.player.bloomstudio.gg

Full Example:

- http://127.0.0.1:1323/player

### WebSocket Base Endpoints

Local Base Endpoint with Default Config:

- ws://localhost:1323/ws

Staging Base Endpoint:

- ws://staging.player.bloomstudio.gg/ws

Full Example:

- ws://staging.player.bloomstudio.gg/ws/position

### Rest Endpoints

##### CreatePlayer

`POST /player`

Creates a Player.

The primairy concern of the CreatePlayer endpoint is to handle the top level fields of the player.

E.g Name,UserID.
Not associated fields like Transform or States.
Use dedicated endpoints for associations and non top level fields if possible.

**Headers:** None

**Request Body:**

Expects a JSON serilized Player [publicModel](./publicModels/player.go) or a [model](./models/player.go) object in the body.

| Name | Type   | Mandatory | Info              |
| ---- | ------ | --------- | ----------------- |
| Name | STRING | YES       | Has to be unique. |

**Request Body Example With all Accepted Fields:**

```json
{
  "UserID": "33b7e1f3-6f8e-40b9-97dc-c54d9162vb05",
  "Name": "User1",
  "Layer": "layer1",
  "Ens": "",
  "Active": true,
  "Transform": {
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
  },
  "states": [
    {
      "id": 1,
      "value": 0.4
    },
    {
      "id": 2,
      "value": 0.1
    }
  ]
}
```

\*States are Unimplemented

**Response:**

```json
{
  "ID": 5,
  "CreatedAt": "2023-09-09T23:13:10.74373182+02:00",
  "UpdatedAt": "2023-09-09T23:13:10.74373182+02:00",
  "DeletedAt": null,
  "UserID": "65ceceb3-611a-4c5a-843d-cd4f060590e2",
  "name": "User1",
  "layer": "",
  "ens": "",
  "active": true,
  "transform": {
    "RotationID": 6,
    "ScaleID": 6,
    "PositionID": 6,
    "position": {
      "ID": 6,
      "CreatedAt": "2023-09-09T23:13:10.740599415+02:00",
      "UpdatedAt": "2023-09-09T23:13:10.740599415+02:00",
      "DeletedAt": null,
      "x": 1,
      "y": 2,
      "z": 3
    },
    "rotation": {
      "ID": 6,
      "CreatedAt": "2023-09-09T23:13:10.742062057+02:00",
      "UpdatedAt": "2023-09-09T23:13:10.742062057+02:00",
      "DeletedAt": null,
      "x": 4,
      "y": 5,
      "z": 6,
      "w": 0
    },
    "scale": {
      "ID": 6,
      "CreatedAt": "2023-09-09T23:13:10.742933738+02:00",
      "UpdatedAt": "2023-09-09T23:13:10.742933738+02:00",
      "DeletedAt": null,
      "x": 7,
      "y": 8,
      "z": 9
    }
  },
  "states": null
}
```

##### GetPlayer

`GET /player`

Get a list of Players.

**Headers:** None

**Query Parameters:**

| Name   | Type | Mandatory | Default | Info                                                                                                              |
| ------ | ---- | --------- | ------- | ----------------------------------------------------------------------------------------------------------------- |
| Active | BOOL | NO        | True    | Whether or not to include Inactive Players. Include Inactive Players with False. Does not exclude Active Players! |

**Request URL Example With all Accepted Query Params:**

```html
http://127.0.0.1:1323/player?active=true
```

**Response:**

```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-08-31T17:45:01.727297847+02:00",
    "UpdatedAt": "2023-09-03T13:29:14.59514667+02:00",
    "DeletedAt": null,
    "UserID": "d6ce83e8-afa8-4fcc-af13-ab1f2e26f9e3",
    "name": "User5",
    "layer": "",
    "ens": "",
    "active": true,
    "transform": {
      "RotationID": 1,
      "ScaleID": 1,
      "PositionID": 1,
      "position": {
        "ID": 1,
        "CreatedAt": "2023-08-31T17:45:01.724252373+02:00",
        "UpdatedAt": "2023-09-03T13:29:14.592187627+02:00",
        "DeletedAt": null,
        "x": 1,
        "y": 2,
        "z": 3
      },
      "rotation": {
        "ID": 1,
        "CreatedAt": "2023-08-31T17:45:01.725541515+02:00",
        "UpdatedAt": "2023-09-10T11:51:08.996298719+02:00",
        "DeletedAt": null,
        "x": 444,
        "y": 555,
        "z": 666,
        "w": 8
      },
      "scale": {
        "ID": 1,
        "CreatedAt": "2023-08-31T17:45:01.726407856+02:00",
        "UpdatedAt": "2023-09-03T13:29:14.594360829+02:00",
        "DeletedAt": null,
        "x": 7,
        "y": 8,
        "z": 9
      }
    },
    "states": [
      {
        "ID": 3,
        "CreatedAt": "2023-08-31T17:49:59.858256533+02:00",
        "UpdatedAt": "2023-08-31T17:49:59.858256533+02:00",
        "DeletedAt": null,
        "PlayerID": 1,
        "stateID": 0,
        "value": 0.4
      },
      {
        "ID": 12,
        "CreatedAt": "2023-09-03T13:29:14.596106641+02:00",
        "UpdatedAt": "2023-09-03T13:29:14.596106641+02:00",
        "DeletedAt": null,
        "PlayerID": 1,
        "stateID": 333,
        "value": 123456.991123123
      }
    ]
  }
]
```

##### DeletePlayer

`DELETE /player`

Soft Deletes a Player.

**Headers:** None

**Path Parameters:**

| Name | Type | Mandatory |
| ---- | ---- | --------- |
| ID   | INT  | YES       |

**Request URL Example With all Accepted Path Params:**

```html
http://127.0.0.1:1323/player/:ID
```

**Response:**

```html
Status: 200 OK
```

---

### Websocket Endpoints

---

Websocket Endpoints on first connect will send all relevant data\*¹.
After the first transmition it will only send objects that have changes since the last transmition.

\*¹Except Data that has to be activly included and would not be included normally by default.

E.g If by default inactive or soft deleted data|objects|rows wont be included
it also wont be included in the initial transfer unless queried|asked for by the client.

---

All Websocket Endpoints Accept the following:

**Query Parameters:**

| Name | Type | Mandatory | Default | Info                                                        |
| ---- | ---- | --------- | ------- | ----------------------------------------------------------- |
| rate | INT  | NO        | 1       | The Rate in Millisecond which the Server Updates the Client |

\*Capitalization matters

---

##### Player

`/player`

Handles Player Data.
Which encompases all Data Available. This can be thought of as the Base Layer of Information
Send and Receive Player Data.

**Headers:** None

**Request Body:**

Expects a JSON serilized Player [publicModel](./publicModels/player.go) or a [model](./models/player.go) object in the body.

| Name | Type   | Mandatory | Info              |
| ---- | ------ | --------- | ----------------- |
| Name | STRING | YES       | Has to be unique. |

**Request Body Example With all Accepted Fields:**

```json
{
  "ID": 4,
  "CreatedAt": "2023-09-09T22:12:18.668123608+02:00",
  "UpdatedAt": "2023-09-09T22:12:18.668123608+02:00",
  "DeletedAt": null,
  "UserID": "18a8bfd7-0f28-49be-aba6-13186fca7ec3",
  "name": "User12",
  "layer": "",
  "ens": "",
  "active": true,
  "transform": {
    "RotationID": 5,
    "ScaleID": 5,
    "PositionID": 5,
    "position": {
      "ID": 5,
      "CreatedAt": "2023-09-09T22:12:18.66237607+02:00",
      "UpdatedAt": "2023-09-09T22:12:18.66237607+02:00",
      "DeletedAt": null,
      "x": 1,
      "y": 2,
      "z": 3
    },
    "rotation": {
      "ID": 5,
      "CreatedAt": "2023-09-09T22:12:18.664758593+02:00",
      "UpdatedAt": "2023-09-09T22:12:18.664758593+02:00",
      "DeletedAt": null,
      "x": 4,
      "y": 5,
      "z": 6,
      "w": 0
    },
    "scale": {
      "ID": 5,
      "CreatedAt": "2023-09-09T22:12:18.666777896+02:00",
      "UpdatedAt": "2023-09-09T22:12:18.666777896+02:00",
      "DeletedAt": null,
      "x": 7,
      "y": 8,
      "z": 9
    }
  },
  "states": [
    {
      "id": 1,
      "value": 0.4
    },
    {
      "id": 2,
      "value": 0.1
    }
  ]
}
```

**Response:**

```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-08-31T17:45:01.727297847+02:00",
    "UpdatedAt": "2023-09-03T13:29:14.59514667+02:00",
    "DeletedAt": null,
    "UserID": "d6ce83e8-afa8-4fcc-af13-ab1f2e26f9e3",
    "name": "User5",
    "layer": "",
    "ens": "",
    "active": true,
    "transform": {
      "RotationID": 1,
      "ScaleID": 1,
      "PositionID": 1,
      "position": {
        "ID": 1,
        "CreatedAt": "2023-08-31T17:45:01.724252373+02:00",
        "UpdatedAt": "2023-09-10T22:28:41.658189968+02:00",
        "DeletedAt": null,
        "x": 11111,
        "y": 2,
        "z": 23
      },
      "rotation": {
        "ID": 1,
        "CreatedAt": "2023-08-31T17:45:01.725541515+02:00",
        "UpdatedAt": "2023-09-10T11:51:08.996298719+02:00",
        "DeletedAt": null,
        "x": 444,
        "y": 555,
        "z": 666,
        "w": 8
      },
      "scale": {
        "ID": 1,
        "CreatedAt": "2023-08-31T17:45:01.726407856+02:00",
        "UpdatedAt": "2023-09-03T13:29:14.594360829+02:00",
        "DeletedAt": null,
        "x": 7,
        "y": 8,
        "z": 9
      }
    },
    "states": [
      {
        "ID": 3,
        "CreatedAt": "2023-08-31T17:49:59.858256533+02:00",
        "UpdatedAt": "2023-08-31T17:49:59.858256533+02:00",
        "DeletedAt": null,
        "PlayerID": 1,
        "stateID": 0,
        "value": 0.4
      },
      {
        "ID": 4,
        "CreatedAt": "2023-08-31T17:49:59.858256533+02:00",
        "UpdatedAt": "2023-08-31T17:49:59.858256533+02:00",
        "DeletedAt": null,
        "PlayerID": 1,
        "stateID": 0,
        "value": 0.1
      },
      {
        "ID": 12,
        "CreatedAt": "2023-09-03T13:29:14.596106641+02:00",
        "UpdatedAt": "2023-09-03T13:29:14.596106641+02:00",
        "DeletedAt": null,
        "PlayerID": 1,
        "stateID": 333,
        "value": 123456.991123123
      }
    ]
  },
  {
    "ID": 2,
    "CreatedAt": "2023-08-31T17:46:17.440900559+02:00",
    "UpdatedAt": "2023-09-05T14:02:16.433866911+02:00",
    "DeletedAt": null,
    "UserID": "ea5f4920-c472-4cfb-ba34-ff504a4a19bb",
    "name": "User2",
    "layer": "",
    "ens": "",
    "active": true,
    "transform": {
      "RotationID": 2,
      "ScaleID": 2,
      "PositionID": 2,
      "position": {
        "ID": 2,
        "CreatedAt": "2023-08-31T17:46:17.436231232+02:00",
        "UpdatedAt": "2023-09-05T14:02:16.429573435+02:00",
        "DeletedAt": null,
        "x": 1123,
        "y": 2,
        "z": 345
      },
      "rotation": {
        "ID": 2,
        "CreatedAt": "2023-08-31T17:46:17.438476376+02:00",
        "UpdatedAt": "2023-09-10T11:52:54.218716343+02:00",
        "DeletedAt": null,
        "x": 444,
        "y": 555,
        "z": 666,
        "w": 8
      },
      "scale": {
        "ID": 2,
        "CreatedAt": "2023-08-31T17:46:17.439661988+02:00",
        "UpdatedAt": "2023-09-05T14:02:16.432585839+02:00",
        "DeletedAt": null,
        "x": 7,
        "y": 8,
        "z": 9
      }
    },
    "states": [
      {
        "ID": 13,
        "CreatedAt": "2023-09-03T13:33:28.807118497+02:00",
        "UpdatedAt": "2023-09-03T13:34:57.979531378+02:00",
        "DeletedAt": null,
        "PlayerID": 2,
        "stateID": 31,
        "value": 41.4435345
      },
      {
        "ID": 14,
        "CreatedAt": "2023-09-03T13:33:28.807118497+02:00",
        "UpdatedAt": "2023-09-03T13:34:57.979531378+02:00",
        "DeletedAt": null,
        "PlayerID": 2,
        "stateID": 41,
        "value": 51.991123123
      }
    ]
  },
  {
    "ID": 3,
    "CreatedAt": "2023-09-02T20:34:25.340807392+02:00",
    "UpdatedAt": "2023-09-02T20:34:25.340807392+02:00",
    "DeletedAt": null,
    "UserID": "5de16ba0-eaf1-466e-93cd-91ee2206852a",
    "name": "User2RUNDOT",
    "layer": "",
    "ens": "",
    "active": true,
    "transform": {
      "RotationID": 4,
      "ScaleID": 4,
      "PositionID": 4,
      "position": {
        "ID": 4,
        "CreatedAt": "2023-09-02T20:34:25.337078128+02:00",
        "UpdatedAt": "2023-09-02T20:34:25.337078128+02:00",
        "DeletedAt": null,
        "x": 1,
        "y": 2,
        "z": 3
      },
      "rotation": {
        "ID": 4,
        "CreatedAt": "2023-09-02T20:34:25.338551329+02:00",
        "UpdatedAt": "2023-09-02T20:34:25.338551329+02:00",
        "DeletedAt": null,
        "x": 4,
        "y": 5,
        "z": 6,
        "w": 0
      },
      "scale": {
        "ID": 4,
        "CreatedAt": "2023-09-02T20:34:25.339655251+02:00",
        "UpdatedAt": "2023-09-02T20:34:25.339655251+02:00",
        "DeletedAt": null,
        "x": 7,
        "y": 8,
        "z": 9
      }
    },
    "states": []
  },
  {
    "ID": 4,
    "CreatedAt": "2023-09-09T22:12:18.668123608+02:00",
    "UpdatedAt": "2023-09-09T22:14:46.68169814+02:00",
    "DeletedAt": null,
    "UserID": "18a8bfd7-0f28-49be-aba6-13186fca7ec3",
    "name": "User12",
    "layer": "",
    "ens": "",
    "active": true,
    "transform": {
      "RotationID": 5,
      "ScaleID": 5,
      "PositionID": 5,
      "position": {
        "ID": 5,
        "CreatedAt": "2023-09-09T22:12:18.66237607+02:00",
        "UpdatedAt": "2023-09-09T22:14:46.67448893+02:00",
        "DeletedAt": null,
        "x": 1,
        "y": 2,
        "z": 3
      },
      "rotation": {
        "ID": 5,
        "CreatedAt": "2023-09-09T22:12:18.664758593+02:00",
        "UpdatedAt": "2023-09-09T22:14:46.676704963+02:00",
        "DeletedAt": null,
        "x": 4,
        "y": 5,
        "z": 6,
        "w": 0
      },
      "scale": {
        "ID": 5,
        "CreatedAt": "2023-09-09T22:12:18.666777896+02:00",
        "UpdatedAt": "2023-09-09T22:14:46.677962475+02:00",
        "DeletedAt": null,
        "x": 7,
        "y": 8,
        "z": 9
      }
    },
    "states": [
      {
        "ID": 1,
        "CreatedAt": "2023-08-31T17:48:08.013886722+02:00",
        "UpdatedAt": "2023-09-09T22:14:46.680543299+02:00",
        "DeletedAt": null,
        "PlayerID": 4,
        "stateID": 0,
        "value": 0.4
      },
      {
        "ID": 2,
        "CreatedAt": "2023-08-31T17:48:08.013886722+02:00",
        "UpdatedAt": "2023-09-09T22:14:46.680543299+02:00",
        "DeletedAt": null,
        "PlayerID": 4,
        "stateID": 0,
        "value": 0.1
      }
    ]
  },
  {
    "ID": 5,
    "CreatedAt": "2023-09-09T23:13:10.74373182+02:00",
    "UpdatedAt": "2023-09-09T23:13:10.74373182+02:00",
    "DeletedAt": null,
    "UserID": "65ceceb3-611a-4c5a-843d-cd4f060590e2",
    "name": "User1",
    "layer": "",
    "ens": "",
    "active": true,
    "transform": {
      "RotationID": 6,
      "ScaleID": 6,
      "PositionID": 6,
      "position": {
        "ID": 6,
        "CreatedAt": "2023-09-09T23:13:10.740599415+02:00",
        "UpdatedAt": "2023-09-09T23:13:10.740599415+02:00",
        "DeletedAt": null,
        "x": 1,
        "y": 2,
        "z": 3
      },
      "rotation": {
        "ID": 6,
        "CreatedAt": "2023-09-09T23:13:10.742062057+02:00",
        "UpdatedAt": "2023-09-09T23:13:10.742062057+02:00",
        "DeletedAt": null,
        "x": 4,
        "y": 5,
        "z": 6,
        "w": 0
      },
      "scale": {
        "ID": 6,
        "CreatedAt": "2023-09-09T23:13:10.742933738+02:00",
        "UpdatedAt": "2023-09-09T23:13:10.742933738+02:00",
        "DeletedAt": null,
        "x": 7,
        "y": 8,
        "z": 9
      }
    },
    "states": []
  }
]
```

##### Position

`/position`

Handles Position Data. Send and Receive Position Data.

**Headers:** None

**Request Body:**

Expects a JSON serilized Position [publicModel](./publicModels/position.go) or a [model](./models/position.go) object in the body.

| Name | Type | Mandatory | Info              |
| ---- | ---- | --------- | ----------------- |
| ID   | INT  | YES       | Has to be unique. |

**Request Body Example With all Accepted Fields:**

```json
{
  "ID": 1,
  "x": 1,
  "y": 2,
  "z": 3
}
```

**Response:**

```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-08-31T17:45:01.724252373+02:00",
    "UpdatedAt": "2023-09-03T13:29:14.592187627+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 2,
    "CreatedAt": "2023-08-31T17:46:17.436231232+02:00",
    "UpdatedAt": "2023-09-05T14:02:16.429573435+02:00",
    "DeletedAt": null,
    "x": 1123,
    "y": 2,
    "z": 345
  },
  {
    "ID": 3,
    "CreatedAt": "2023-09-02T20:34:10.205426857+02:00",
    "UpdatedAt": "2023-09-02T20:34:10.205426857+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 4,
    "CreatedAt": "2023-09-02T20:34:25.337078128+02:00",
    "UpdatedAt": "2023-09-02T20:34:25.337078128+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 5,
    "CreatedAt": "2023-09-09T22:12:18.66237607+02:00",
    "UpdatedAt": "2023-09-09T22:14:46.67448893+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 6,
    "CreatedAt": "2023-09-09T23:13:10.740599415+02:00",
    "UpdatedAt": "2023-09-09T23:13:10.740599415+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  }
]
```

##### Rotation

`/rotation`

Handles Rotation Data. Send and Receive Rotation Data.

**Headers:** None

**Request Body:**

Expects a JSON serilized Rotation [publicModel](./publicModels/rotation.go) or a [model](./models/rotation.go) object in the body.

| Name | Type | Mandatory | Info              |
| ---- | ---- | --------- | ----------------- |
| ID   | INT  | YES       | Has to be unique. |

**Request Body Example With all Accepted Fields:**

```json
{
  "ID": 2,
  "x": 444,
  "y": 555,
  "z": 666,
  "w": 8
}
```

**Response:**

```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-08-31T17:45:01.725541515+02:00",
    "UpdatedAt": "2023-09-10T11:51:08.996298719+02:00",
    "DeletedAt": null,
    "x": 444,
    "y": 555,
    "z": 666,
    "w": 8
  },
  {
    "ID": 2,
    "CreatedAt": "2023-08-31T17:46:17.438476376+02:00",
    "UpdatedAt": "2023-09-10T11:52:54.218716343+02:00",
    "DeletedAt": null,
    "x": 444,
    "y": 555,
    "z": 666,
    "w": 8
  },
  {
    "ID": 3,
    "CreatedAt": "2023-09-02T20:34:10.206732079+02:00",
    "UpdatedAt": "2023-09-02T20:34:10.206732079+02:00",
    "DeletedAt": null,
    "x": 4,
    "y": 5,
    "z": 6,
    "w": 0
  },
  {
    "ID": 4,
    "CreatedAt": "2023-09-02T20:34:25.338551329+02:00",
    "UpdatedAt": "2023-09-02T20:34:25.338551329+02:00",
    "DeletedAt": null,
    "x": 4,
    "y": 5,
    "z": 6,
    "w": 0
  }
]
```

##### Scale

`/scale`

Handles Scale Data. Send and Receive Scale Data.

**Headers:** None

**Request Body:**

Expects a JSON serilized Scale [publicModel](./publicModels/scale.go) or a [model](./models/scale.go) object in the body.

| Name | Type | Mandatory | Info              |
| ---- | ---- | --------- | ----------------- |
| ID   | INT  | YES       | Has to be unique. |

**Request Body Example With all Accepted Fields:**

```json
{
  "ID": 1,
  "x": 1,
  "y": 2,
  "z": 3
}
```

**Response:**

```json
[
  {
    "ID": 1,
    "CreatedAt": "2023-08-31T17:45:01.724252373+02:00",
    "UpdatedAt": "2023-09-03T13:29:14.592187627+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 2,
    "CreatedAt": "2023-08-31T17:46:17.436231232+02:00",
    "UpdatedAt": "2023-09-05T14:02:16.429573435+02:00",
    "DeletedAt": null,
    "x": 1123,
    "y": 2,
    "z": 345
  },
  {
    "ID": 3,
    "CreatedAt": "2023-09-02T20:34:10.205426857+02:00",
    "UpdatedAt": "2023-09-02T20:34:10.205426857+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 4,
    "CreatedAt": "2023-09-02T20:34:25.337078128+02:00",
    "UpdatedAt": "2023-09-02T20:34:25.337078128+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 5,
    "CreatedAt": "2023-09-09T22:12:18.66237607+02:00",
    "UpdatedAt": "2023-09-09T22:14:46.67448893+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  },
  {
    "ID": 6,
    "CreatedAt": "2023-09-09T23:13:10.740599415+02:00",
    "UpdatedAt": "2023-09-09T23:13:10.740599415+02:00",
    "DeletedAt": null,
    "x": 1,
    "y": 2,
    "z": 3
  }
]
```

##### States

`/states`

Handles State Data. Send and Receive State Data.

**Headers:** None

**Request Body:**

Expects a JSON serilized State [publicModel](./publicModels/state.go) or a [model](./models/state.go) object in the body.

| Name | Type | Mandatory | Info              |
| ---- | ---- | --------- | ----------------- |
| ID   | INT  | YES       | Has to be unique. |

**Request Body Example With all Accepted Fields:**

```json
{
  "ID": 4,
  "CreatedAt": "2023-08-31T17:49:59.858256533+02:00",
  "UpdatedAt": "2023-08-31T17:49:59.858256533+02:00",
  "DeletedAt": null,
  "PlayerID": 1,
  "stateID": 0,
  "value": 0.1
}
```

**Response:**

```json
[
  {
    "ID": 3,
    "CreatedAt": "2023-08-31T17:49:59.858256533+02:00",
    "UpdatedAt": "2023-08-31T17:49:59.858256533+02:00",
    "DeletedAt": null,
    "PlayerID": 1,
    "stateID": 0,
    "value": 0.4
  },
  {
    "ID": 4,
    "CreatedAt": "2023-08-31T17:49:59.858256533+02:00",
    "UpdatedAt": "2023-08-31T17:49:59.858256533+02:00",
    "DeletedAt": null,
    "PlayerID": 1,
    "stateID": 0,
    "value": 0.1
  },
  {
    "ID": 7,
    "CreatedAt": "2023-08-31T17:54:09.475338837+02:00",
    "UpdatedAt": "2023-08-31T17:54:09.475338837+02:00",
    "DeletedAt": null,
    "PlayerID": 1,
    "stateID": 22,
    "value": 0.4435345
  },
  {
    "ID": 9,
    "CreatedAt": "2023-09-01T19:46:20.845375984+02:00",
    "UpdatedAt": "2023-09-01T19:46:20.845375984+02:00",
    "DeletedAt": null,
    "PlayerID": 1,
    "stateID": 223,
    "value": 99.4435345
  },
  {
    "ID": 12,
    "CreatedAt": "2023-09-03T13:29:14.596106641+02:00",
    "UpdatedAt": "2023-09-03T13:29:14.596106641+02:00",
    "DeletedAt": null,
    "PlayerID": 1,
    "stateID": 333,
    "value": 123456.991123123
  }
]
```

##### Level

`/level`

Handles Level Data. Send and Receive Level Data.

**Headers:** None

**Request Body:**

Expects a JSON serilized State [publicModel](./publicModels/level.go) or a [model](./models/level.go) object in the body.

| Name | Type | Mandatory | Info              |
| ---- | ---- | --------- | ----------------- |
| ID   | INT  | YES       | Has to be unique. |

**Request Body Example With all Accepted Fields:**

```json
{
  "ID": 9,
  "PlayerID": 9,
  "levelID": 99
}
```

**Response:**

```json
[
  {
    "ID": 1,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "2023-09-19T18:58:01.004352074+02:00",
    "DeletedAt": null,
    "PlayerID": 1,
    "levelID": 8
  },
  {
    "ID": 2,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "2023-09-17T21:13:22.907264347+02:00",
    "DeletedAt": null,
    "PlayerID": 1,
    "levelID": 99
  },
  {
    "ID": 3,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "2023-09-18T12:55:28.270739728+02:00",
    "DeletedAt": null,
    "PlayerID": 3,
    "levelID": 33
  },
  {
    "ID": 4,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "2023-09-18T12:59:50.210264594+02:00",
    "DeletedAt": null,
    "PlayerID": 3,
    "levelID": 33
  },
  {
    "ID": 5,
    "CreatedAt": "2023-09-18T13:23:07.32511678+02:00",
    "UpdatedAt": "2023-09-18T13:23:22.24989923+02:00",
    "DeletedAt": null,
    "PlayerID": 3,
    "levelID": 33
  },
  {
    "ID": 6,
    "CreatedAt": "2023-09-18T13:26:00.061928771+02:00",
    "UpdatedAt": "2023-09-18T13:26:19.385058543+02:00",
    "DeletedAt": null,
    "PlayerID": 32,
    "levelID": 323
  },
  {
    "ID": 9,
    "CreatedAt": "2023-09-19T18:57:03.427916819+02:00",
    "UpdatedAt": "2023-09-19T18:57:23.383696443+02:00",
    "DeletedAt": null,
    "PlayerID": 9,
    "levelID": 99
  }
]
```

---

## Guides

## How to Interact with the Player WebSocket

Assuming standard config and hosting locally.

1.  If not present create a player in the players table. The player needs to have a usable UserID(In Production) and or a Usable Name(In Debug/Dev)

    This can be be done either by:

    - Manually creating it in the db
    - Creating it throug the /player CreatePlayer endpoint

2.  The Websocket in production needs to derive a models.Player.UserID from a JWT. That UserID has to match with the UserID of the player in the database. In debug mode this is replaced by using the name of the player. That JWT can be created either by:

    - The Userservice
    - Manually

3.  Connect to the websocket with the JWT in the Header as shown below:

              // Token has been shortened for readability
              Authorization: Bearer eyJ_A

              ws://localhost:1323/ws/player

4.  Send and receive JSON Player objects from the websocket

## How to Interact with the Position WebSocket

1. Connect to ws://staging.player.bloomstudio.gg/ws/position
2. You will now recive a array of Positions
3. To update any position you can send a single position object back in its entirety
   This websocket is very similair to the player websocket

**How to setup private multiplayer**

1.  Create a Player in the DB. More Ressources than the below example are availabe either in the docs folder or in this Readme.

    - #### **How To Create a Player:**

      Send a `POST` request to http://staging.player.bloomstudio.gg/player with a adjusted Player Object.

      Player Object Example:

      Change "UserID" and "Name"

      ```json
      {
        "UserID": "33b7e1f3-6f8e-40b9-97dc-c54d9162vb05",
        "Name": "User1",
        "Layer": "layer1",
        "Transform": {
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
        },
        "states": [
          {
            "id": 1,
            "value": 0.4
          },
          {
            "id": 2,
            "value": 0.1
          }
        ]
      }
      ```

2.  Connect to the ws player websocket endpoint. Wss might also work. The following link is for reference and will connect to the remotely hosted staging player service and player endpoint

             ws://staging.player.bloomstudio.gg/ws/player

3.  You will now receive a list of player objects on that websocket. Example output can be found below.

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

4.  If you want to modify your player or any other players you can modify the list of player objects that you received
    and send the singel modified (Not the entire list) player object back within the same websocket your received it from to the service you received it from.
    Note: Sending the entire list of player objects back for update will fail and is currently not supported. You have to send a singel player object back.

    - #### **How to Update The Player object/s.**

      1. Modify and make the desired changes to the list of of Player objects that you have received in the previous step.
      2. Push a modified single player object(Not a list) to the [ws://staging.player.bloomstudio.gg/ws/player](ws://staging.player.bloomstudio.gg/ws/player) websocket
      3. Optional Confirm that the changes took place by looking at the next data push from the websocket.

## Examples

Examples can be found in the docs dir.

## Postman | OpenAPI

The Postman Collection: [PlayerService.postman_collection.json](docs/PlayerService.postman_collection.json)

The OpenAPI spec: [open_api.yml](docs/open_api.yml)

Can be found in docs dir.

The OpenAPI spec is generated from the postman collection in the same dir via:

https://joolfe.github.io/postman-to-openapi/

        p2o PlayerService.postman_collection.json -f open_api.yml

## Rotation ressources

- https://compsci290-s2016.github.io/CoursePage/Materials/EulerAnglesViz/index.html
- https://www.youtube.com/watch?v=2Cwa6hfn2K0
- https://docs.unity3d.com/ScriptReference/Transform-eulerAngles.html
- https://docs.unity3d.com/ScriptReference/Quaternion.html
