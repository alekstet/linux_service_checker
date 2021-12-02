# Linux services checker

This service allows using the web interface to display systemd services running on a Linux system. The server side is written in Golang, the client side is in Vue JS.

To display only the services you need, there is a ___config.yml___ file in the project directory, in which you need to write down the names of your services, as well as the port number on which the service is launched.

The service allows you to:
1. View service name, current job status, systemctl and journalctl help;
2. Start and stop services by pressing a button;
3. Sort services by configuration file, by state of work. 

## View
<p align="center">
<img  src="/readme_assets/1.PNG" width="80%">
</p>

## Test project setup
For a test run, you can write from the directory
```
npm install
npm run build
./service_checker
```

## Project setup
To run the project as a separate service, change path in ___service_checker.txt___ and 
copy the contents of the file, create a ___service_checker.service___ in _/etc/systemd/system_ and paste the content
```
systemctl daemon-reload
npm install
npm run build
systemctl start service_checker
```
