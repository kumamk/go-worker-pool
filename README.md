## Description

Worker Pool is a pattern to achieve concurrency using fixed number of workers to execute multiple amount of tasks on a queue. This is basic implementation of the worker pool using Go ecosystem via goroutines and channel.

## Project Details

* No of goroutines taken is 5
* No of jobs to be done is 50
* Simple Worker Pool is implemented utilizing channel as queue and goroutine as worker.

## Pre Requirements

* GO installed in the system and any supported IDE

## Test

```bash
# run program
$ go run main.go

# output
2023/12/02 15:42:04 Worker 5 finished the job
2023/12/02 15:42:04 Job 47 finished with response 470
2023/12/02 15:42:04 Worker 2 finished the job
2023/12/02 15:42:04 Job 46 finished with response 460
2023/12/02 15:42:04 Worker 3 finished the job
2023/12/02 15:42:04 Job 48 finished with response 480
2023/12/02 15:42:04 Worker 4 finished the job
2023/12/02 15:42:04 Worker 1 finished the job
2023/12/02 15:42:04 Job 50 finished with response 500
2023/12/02 15:42:04 Job 49 finished with response 490
```