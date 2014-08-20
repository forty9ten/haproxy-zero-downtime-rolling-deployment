#!/bin/sh

# enable
echo 1 > /proc/sys/net/ipv4/ip_forward

# reste
iptables -F
iptables -t nat -F
