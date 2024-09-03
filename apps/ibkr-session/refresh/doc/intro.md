## OVERVIEW

Makes a request to the ibkr-gateway app to see if the IBKR clientportal session is still open. 

The idea is to check if we need to login to the clientportal again. 

This is very important because a closed session means we cannot perform trades.

For more information about client portal sessions, see: https://interactivebrokers.github.io/cpwebapi/index.html#login

All interaction with the clientportal goes through the ibkr-gateway (in this project).

This application is meant to be executed like a polling job - i.e., some interval.

If the session is open, the following log is printed to stdout in this app.
```
{
  "timestamp" : "<timestamp>",
  "level" : "INFO",
  "thread" : "main",
  "ns" : "refresh.core",
  "line" : 15,
  "session" : true,    <<<< session is OK
  "column" : 3,
  "logger" : "refresh.core",
  "message" : "Response",
  "context" : "default"
}
```

If the session is not open, the following will be displayed
```
{
  "timestamp" : "<timestamp>",
  "level" : "INFO",
  "thread" : "main",
  "ns" : "refresh.core",
  "line" : 15,
  "session" : false,    <<<< session is NOT OK
  "column" : 3,
  "logger" : "refresh.core",
  "message" : "Response",
  "context" : "default"
}
```

