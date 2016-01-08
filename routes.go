package main

// This should be stored in a KV store like consul
var mona = Routes{
    Route{
        "Index",
        "GET",
        "/api/v1/",
        Index,
    },
}

var mssql = Routes{
    Route{
        "Index",
        "GET",
        "/api/mssql/v1/",
        mssqlIndex,
    },
    Route{
        "ListAllInstances",
        "GET",
        "/api/mssql/v1/ListAllInstances",
        mssqlAllInstances,
    },
    Route{
        "ListAllConnections",
        "GET",
        "/api/mssql/v1/ListAllConnections",
        mssqlAllConnections,
    },
    Route{
        "ListAllActiveConnections",
        "GET",
        "/api/mssql/v1/ListAllActiveConnections",
        mssqlAllActiveConnections,
    },
    Route{
        "checksqlservices",
        "GET",
        "/api/mssql/v1/checksqlservices",
        checkSQLServices,
    },
}

var routes = append(mona, mssql...)
