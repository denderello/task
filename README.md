# Task

`task` is a simple Go library to execute tasks after each other. When one of the
tasks fails the queue will be walked backwards and each task gets the
opportunity to run rollback actions of the actions performed beforehand.
