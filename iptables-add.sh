#!/bin/sh
iptables -t nat -A OUTPUT -p tcp --dport 5910 -j REDIRECT --to-port 4910
iptables -t nat -A PREROUTING -p tcp --dport 5910 -j REDIRECT --to-port 4910

iptables -t nat -A OUTPUT -p tcp --dport 5911 -j REDIRECT --to-port 4911
iptables -t nat -A PREROUTING -p tcp --dport 5911 -j REDIRECT --to-port 4911
