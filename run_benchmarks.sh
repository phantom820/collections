#! /usr/bin/bash
echo "-- RUNNING SET BENCHMARKS --"
go test -bench=./... -benchmem -benchtime=5x github.com/phantom820/collections/benchmarks/sets > sets.csv
echo "-- COMPLETED SET BENCHMARKS OUTPUT : sets.csv --"
echo "-- RUNNING LIST BENCHMARKS --"
go test -bench=./... -benchmem -benchtime=5x github.com/phantom820/collections/benchmarks/lists > lists.csv
echo "-- COMPLETED LIST BENCHMARKS OUTPUT : lists.csv --"
echo "-- RUNNING QUEUE BENCHMARKS --"
go test -bench=./... -benchmem -benchtime=5x github.com/phantom820/collections/benchmarks/queues > queues.csv
echo "-- COMPLETED QUEUE BENCHMARKS OUTPUT : queues.csv --"
