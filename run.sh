#!/bin/bash
cd src && go build -o main.out && ./main.out "$1"
