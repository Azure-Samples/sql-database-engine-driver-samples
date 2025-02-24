<!--
---
page_type: sample
name: Go - Connect to SQL database engine
description: This sample uses the Go driver for SQL to connect to a database engine and perform a basic query.
urlFragment: sample
languages:
- go
- sql
- tsql
products:
- sql-server
- azure-sql-database
- azure-sql-edge
- azure-sql-managed-instance
- azure-sql-virtual-machines
- azure-sqlserver-stretchdb
- azure-sqlserver-vm
---
-->

# Go - Connect to SQL database engine

This sample uses the Go driver for SQL to connect to a database engine and perform a basic query.

## Prerequisites

- [Go 1.23 or newer](https://go.dev/)

## Run this sample

1. Navigate to this sample's directory

    ```shell
    cd ./001-connect/
    ```

1. Create a *.env* file with the following environment variable set.

    ```dotenv
    SQL_CREDENTIAL="<your-sql-connection-string>"
    ```

1. Run this sample using `go run`.

    ```shell
    go run .
    ```

1. Observe the output.
