# alvarium-example
An example application integrated with the Alvarium SDK.

The application will demonstrate the following:
* Instantiation of the Alvarium SDK

* Use of the BootstrapHandler pattern for graceful shutdown

* Use of the Create() method for creating a random piece of data

* Use of the Mutate() method to link the original data to a new piece of data

* Use of the Transit() method to show handling of data in transit from one service to another

LogLevel is set via config to `debug` by default for a high level of visibility in logging.

## Running the application

Having built the application, you have two modes in which to run it

This project includes an [MQTT docker-compose](https://github.com/project-alvarium/scripts/docker/mqtt-docker-compose.yml) file.
If you do not already have MQTT running, you can start up a docker container by using a terminal window to locate the above 
directory in your local file system. You can then use the following command line.

`docker-compose -f mqtt-docker-compose.yml up`

To start the application, execute the `make run_mqtt` command line.

Please review the config file relevant to this scenario at `./cmd/res/config-mqtt.json` and make any necessary changes for your
specific environment.