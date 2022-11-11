# PlayerService

## Quickstart Guide:

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

        docker build -t playerservice.

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