# DiningHallServerGO
This is the dining hall part of the Restaurant simulation of the first lab at the Network Programming course.
The kitchen part of the Restaurant simulation: https://github.com/GheorgheMorari/KitchenServerGO

# Docker stuff:
If you don't run linux, or don't have git bash, change the file type of the scripts to cmd.

run build_and_start_container.sh to build and start container

run start_server_from_container.sh to start or restart server

run build_docker_image.sh to build the image 

run remove_docker_stuff.sh to remove docker image and container 



# View in browser addresses:

localhost:7500/start  -to start sending one order every second

localhost:7500/stop  -to stop sending one order every second

localhost:7500/send  -to send one order and to display the response or error


# The dining hall system architecture:

![image](https://user-images.githubusercontent.com/53918731/133939450-7ce8bc35-0286-4d3d-951e-eb51d71869a2.png)

# The communication protocol:

Sending:

![image](https://user-images.githubusercontent.com/53918731/134770671-331833ae-bdf9-4983-95e4-1e213836c4f7.png)

Receiving:

![image](https://user-images.githubusercontent.com/53918731/133939490-04ea0dd2-96cd-4458-a31d-df68c66ca409.png)
