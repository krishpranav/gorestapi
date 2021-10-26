#!/bin/bash

echo "Go Rest Api Test"
go test
sleep 1

echo "Aboutapi Test"
cd api/aboutapi
go test
sleep 1
cd ../
