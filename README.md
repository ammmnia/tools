<h3 align="center" style="border-bottom: none">
      ⭐️  FlyIM tools.  ⭐️ <br>
<h3>


<p align=center>
<a href="https://pkg.go.dev/github.com/ammmnia/tools"><img src="https://img.shields.io/badge/Language-Go-blue.svg"></a>
</p>


<p align="center">
    <a href="./README.md"><b>English</b></a> •
    <a href="./README_zh-CN.md"><b>中文</b></a>
</p>

-----

## Project Toolkit Documentation

This project includes a series of tools and libraries applicable to FlyIM, along with some other project-provided tool supports, aimed at supporting efficient solution development. Below is an introduction to the functionalities of each module:

## a2r

- `api2rpc.go`: A tool for converting API to RPC, used for converting HTTP API requests into RPC calls.

## apiresp

- `format.go`, `gin.go`, `http.go`, `resp.go`: Handles the formatting, encapsulation, and sending of API responses, supporting different web frameworks.

## checker

- `check.go`: Provides service health checks and dependency verification functions.

## config

- `config.go`, `config_parser.go`, `config_source.go`, `manager.go`, `path.go`: Configuration management module, supporting the parsing, loading, and dynamic updating of configurations.
- `validation`: Provides tools and libraries for configuration validation.

## db

- `mongo`, `pagination`, `redis`, `tx.go`: Database operation-related tools, including support for MongoDB, Redis, and transaction management.

## discovery

- `discovery_register.go`: Service discovery and registration functions.
- `zookeeper`: Service discovery implementation based on Zookeeper.

## env

- `env.go`, `env_test.go`: Environment variable management tools, including loading and parsing environment variables.

## errs

- `code.go`, `coderr.go`, `error.go`, `predefine.go`, `relation.go`: Error code management and custom error types.

## field

- `file.go`, `path.go`: Utilities for file operations and path generation.
- Related test files.

## log

- `color.go`, `encoder.go`, `logger.go`, `sql_logger.go`, `zap.go`, `zk_logger.go`: Log management module, supporting multiple log formats and outputs.

> [!IMPORTANT]
> For more information about OpenIM log, please read [https://github.com/ammmnia/im-server/blob/main/docs/contrib/logging.md](https://github.com/ammmnia/im-server/blob/main/docs/contrib/logging.md)

## mcontext

- `ctx.go`: Context management tool, used for passing request-related information between middleware and services.

## mq

- `kafka`: Support for message queues based on Kafka.

## mw

- `gin.go`, `intercept_chain.go`, `rpc_client_interceptor.go`, `rpc_server_interceptor.go`: Middleware and interceptors, used for preprocessing and postprocessing of requests.
- `specialerror`: Special error handling module.

## tokenverify

- `jwt_token.go`, `jwt_token_test.go`: JWT token verification and testing.

## utils

The utils contain multiple utility libraries, such as `encoding`, `encrypt`, `httputil`, `jsonutil`, `network`, `splitter`, `stringutil`, `timeutil`: Providing various common functionalities, such as encryption, encoding, network operations, etc.

#### encoding

- `base64.go` & `base64_test.go`: Provides utility functions for Base64 encoding and decoding, and their unit tests.

#### encrypt

- `encryption.go` & `encryption_test.go`: Contains functionalities for encryption and decryption, supporting common encryption algorithms, and related unit tests.

#### goassist

- `jsonutils.go` & `jsonutils_test.go`: Provides utility functions for handling JSON data, such as parsing and generating JSON, and related unit tests.

#### httputil

- `http_client.go` & `http_client_test.go`: Encapsulates HTTP client operations, providing convenient methods for sending HTTP requests, and their unit tests.

#### jsonutil

- `interface.go`, `json.go` & `json_test.go`: Focuses on JSON data handling, including more advanced JSON operations and customized JSON parsing methods, and their unit tests.

#### network

- `ip.go` & `ip_test.go`: Provides network-related utility functions, such as parsing and validating IP addresses, and related unit tests.

#### splitter

- `splitter.go` & `splitter_test.go`: Provides tools for splitting strings, supporting various splitting strategies and complex splitting scenarios, and their unit tests.

#### stringutil

- `strings.go` & `strings_test.go`: Contains a series of utility functions for string operations, such as modifying, searching, comparing strings, and their unit tests.

#### timeutil

- `time_format.go` & `time_format_test.go`: Provides utility functions related to time, including parsing and formatting time formats, and related unit tests.

## version

- `base.go`, `doc.go`, `types.go`, `version.go`: Version management tool, used for defining and managing project version information.
