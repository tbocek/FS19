#!/bin/sh

iptables -t filter -N MINIUPNPD
iptables -t nat -N MINIUPNPD
iptables -t nat -N MINIUPNPD-POSTROUTING

iptables -t filter -A FORWARD -j MINIUPNPD
iptables -t nat -A PREROUTING -j MINIUPNPD
iptables -t nat -A POSTROUTING -j MINIUPNPD-POSTROUTING

miniupnpd -N -i nat_wan -a nat_lan -d > /tmp/upnp.log 2>&1 &

#unr: nc -l -p 8192
#glb: nc 10.0.2.17 4096
