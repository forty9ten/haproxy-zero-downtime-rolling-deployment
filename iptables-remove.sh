#!/bin/sh
iptables -t nat -D OUTPUT -p tcp --dport 5910 -j REDIRECT --to-port 4910
iptables -t nat -D PREROUTING -p tcp --dport 5910 -j REDIRECT --to-port 4910
