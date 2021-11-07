#!/bin/bash
cd src && go build -o main && ./main "$1" > output.csv