#! /bin/sh
sudo docker stop user-service
sudo docker run -d -it --name user-service -p 3000:3000 user-service 