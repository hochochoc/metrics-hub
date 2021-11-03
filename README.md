### Prerequisites
- go1.16.4
- postgreSQL
### Local development
```
$ cp .env.example .env
```
- Create a database for metrics hub
- Update postgreSQL info in `.env`
```
$ go mod download
$ go run main.go
```
### APIs
1. `process_report`
- accept `POST` request when the process finishes
- Body 
```
{
	"server_name": "test5",
	"start_time": "2021-05-21T20:11:33Z",
	"end_time": "2021-05-21T21:01:53Z"
}
```
2. `process_statistics`
- Get mean, standard deviation of all processes' duration
- If the number of reports < 10, throw response with status 500, and error `not enough reports received`  
- Method: `GET`
- Response 
```
{
    "mean": 1664,
    "stddev": 76
}
```
3. `process_outliers`
- Return all server names (a set) for which a process was detected whose duration is more than three standard deviations greater or smaller than the mean of all process durations
- Method: `GET`
- Response
```
["server1", "server2"] 
```