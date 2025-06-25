# 🚀 DevOps Reverse Proxy Project

This project demonstrates using Docker Compose and Nginx to reverse proxy two services: one in Golang and one in Python.

## 📦 Project Structure

.
├── docker-compose.yml
├── nginx/
│ ├── nginx.conf
│ └── Dockerfile
├── service_1/ # Golang service
│ └── Dockerfile
├── service_2/ # Python service
│ └── Dockerfile
└── README.md

Accessible from:
http://47.129.47.163/service1/hello
http://47.129.47.163/service2/hello
