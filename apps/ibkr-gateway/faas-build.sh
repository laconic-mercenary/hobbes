#!/bin/sh

faas-cli build -f ibkr-gateway.yml

faas-cli push -f ibkr-gateway.yml
