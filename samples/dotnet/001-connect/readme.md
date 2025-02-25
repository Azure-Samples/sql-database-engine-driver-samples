<!--
---
page_type: sample
name: .NET - Connect to SQL database engine
description: This sample uses the .NET driver for SQL to connect to a database engine and perform a basic query.
urlFragment: sample
languages:
- csharp
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

# .NET - Connect to SQL database engine

This sample uses the .NET driver for SQL to connect to a database engine and perform a basic query.

## Prerequisites

- [.NET 9.0](https://dotnet.microsoft.com/download)

## Run this sample

1. Navigate to this sample's directory

    ```shell
    cd ./001-connect/
    ```

1. Set the `SQL_CREDENTIAL` environment variable to your current connection string using `dotnet user-secrets set`

    ```shell
    dotnet user-secrets set "SQL_CREDENTIAL" "<your-sql-connection-string>"
    ```

1. Run this sample using `dotnet run`.

    ```shell
    dotnet run
    ```

1. Observe the output.
