# Linux services checker

This service allows using the web interface to display systemd services running on a Linux system. The server side is written in Golang, the client side is in Vue JS.

To display only the services you need, there is a ___config.yml___ file in the project directory, in which you need to write down the names of your services, as well as the port number on which the service is launched.

The service allows you to:
1. View service name, current job status, systemctl and journalctl help;
2. Start and stop services by pressing a button;
3. Sort services by configuration file, by state of work. 

## View
<p align="center">
<img  src="./readme_assets/3.png" width="80%">
</p>

## Test project setup
For a test run, you can write from the directory
```
./service_checker
```

## Project setup
To run the project as a separate service, copy the ___service_checker.service___ to _/etc/systemd/system_
```
systemctl daemon-reload
systemctl start service_checker
```
