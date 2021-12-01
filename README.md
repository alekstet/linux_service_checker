# Linux services checker

This service allows using the web interface to display systemd services running on a Linux system. The server side is written in Golang, the client side is in Vue JS.

To display only the services you need, there is a config.yml file in the project directory, in which you need to write down the names of your services, as well as the port number on which the service is launched.
