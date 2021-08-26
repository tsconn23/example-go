# alvarium-example
An example application integrated with the Alvarium SDK.

The application will demonstrate the following:
* Instantiation of the Alvarium SDK

* Use of the BootstrapHandler pattern for graceful shutdown

* Use of the Create() method for creating a random piece of data

* Use of the Mutate() method to link the original data to a new piece of data

* Use of the Transit() method to show handling of data in transit from one service to another

LogLevel is set via config to `debug` by default for a high level of visibility in logging.

# Build Notes

**NOTE: While generally applicable to all OS platforms, the specific instructions are relative to Linux (Ubuntu)**

This application has a dependency on the [alvarium-sdk-go](https://github.com/project-alvarium/alvarium-sdk-go) module 
and through that a dependency on the [IOTA Streams C bindings](https://github.com/iotaledger/streams/tree/develop/bindings/c).

The SDK contains a pre-built artifact of the [C bindings](https://github.com/project-alvarium/alvarium-sdk-go/blob/main/internal/iota/include/libiota_streams_c.so)
in its source tree that was built on Ubuntu 20.04. Obtaining the `alvarium-sdk-go` via `go get` will allow you to build the 
applications and run tests. However you will need to copy the shared library into a location your OS is aware of in order 
to load the library dynamically at runtime. For example, on Ubuntu 20.04 this location is `/usr/lib`.

Having done that, you will now be able to build the application using the `make build` command line.

## Running the application

Having built the application, you have two modes in which to run it

1.) With IOTA Streams

If you have a Tangle setup and a [Streams Author](https://github.com/project-alvarium/streams-author) endpoint, you can run
the application by invoking the `make run` command line.

Please review the config file relevant to this scenario at `./cmd/res/config.json` and make any necessary changes for your
specific environment.

2.) With MQTT

This project includes an [MQTT docker-compose](https://github.com/project-alvarium/scripts/docker/mqtt-docker-compose.yml) file.
If you do not already have MQTT running, you can start up a docker container by using a terminal window to locate the above 
directory in your local file system. You can then use the following command line.

`docker-compose -f mqtt-docker-compose.yml up`

To start the application, execute the `make run_mqtt` command line.

Please review the config file relevant to this scenario at `./cmd/res/config-mqtt.json` and make any necessary changes for your
specific environment.