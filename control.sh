#!/bin/bash

chmod o+x backend
killall backend
nohup ./backend --config=env.production.yaml  >> storage/log/hade.log &