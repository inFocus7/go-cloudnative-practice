# Go Cloud Native Practice
_Following the "[Learning Cloud Native Go](https://learning-cloud-native-go.github.io/)" tutorial._
_W.I.P_

## Local Setup
0) 🐳 Make sure you have Docker on your system. If not, visit and go through Dockers' [Getting Started](https://docs.docker.com/get-started/) doc.
1) Clone this repository. `git clone https://github.com/ginFocus7/go-cloudnative-practice.git`
2) Run `docker-compose build` then `docker-compose up` on the root folder.
3) 🚀 The site should now be running on `localhost:8080` assuming no Docker setup errors.

**NOTE**: *Built Docker image is ~500 MB. To remove the image, on the command prompt run `docker images` to find the Image ID, then `docker rmi -f <image_id>`*

_Personal Note: I need to better understand how logging works_