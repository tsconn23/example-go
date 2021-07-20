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

At this time, the application requires a specific [version](https://github.com/iotaledger/streams/tree/9d469a09ee18c55f087821cb2ebf5de5715ca4f2/bindings/c) 
of the IOTA Streams C bindings. The compiled library of the bindings is ~80MB in size and thus is not checked in to Github.

This means you will need to build the C bindings yourself and copy the resulting *.so file into your go module cache where
the referenced version of the alvarium-sdk-go module is found. For example:

`~/go/pkg/mod/github.com/project-alvarium/alvarium-sdk-go@v0.0.0-20210720173148-76fd57ea3590/internal/iota/include`

The version of the alvarium-sdk-go will may differ in your local environment depending on which version you're referencing. If you
browse to that directory you should see a `channels.h` file. Just put the *.so file right next to that file in the same location.

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