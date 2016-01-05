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
}

var routes = append(mona, mssql...)
