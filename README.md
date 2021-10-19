# DiningHallServerGO

This is the dining hall part of the Restaurant simulation of the first lab at the Network Programming course. The
kitchen part of the Restaurant simulation: https://github.com/GheorgheMorari/KitchenServerGO

If you close one part of the simulation, the other part will likely throw an error, which will make it crash. Make sure to start the container again with the help of "start_server_from_image.sh" script

# Docker stuff:

If you don't run linux, or don't have git bash, change the file type of the scripts to cmd, or copy the commands from
the chosen script into your preferred cli.

run build_and_start_container.sh to build and start container

run start_server_from_container.sh to start or restart server

run build_docker_image.sh to build the image

run remove_docker_stuff.sh to remove docker image and container

# View in browser addresses:

http://localhost:7500/ -To view the status of the dining hall.

# Tunables:
There are multiple tunables in this simulation.

In _main.go_ there are the common tunables, such as the timeUnit.

There are also tunables for waiters, and tables, situated in _waiter.go_ and _table.go_ respectively.

# The dining hall system architecture:

![image](https://user-images.githubusercontent.com/53918731/133939450-7ce8bc35-0286-4d3d-951e-eb51d71869a2.png)

# The communication protocol:

Sending:

![image](https://user-images.githubusercontent.com/53918731/134770671-331833ae-bdf9-4983-95e4-1e213836c4f7.png)

Receiving:

![image](https://user-images.githubusercontent.com/53918731/133939490-04ea0dd2-96cd-4458-a31d-df68c66ca409.png)
