#!/bin/bash

svcName=${1}

if [ -d "api" ];then
    cd api && air
fi
