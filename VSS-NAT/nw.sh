#!/bin/sh

# create 2 namespaces: unreachable and nat
ip netns add unr 
ip netns add nat

# setup virtual interfaces
ip link add nat_wan type veth peer name nat_real
ip link add nat_lan type veth peer name unr_lan

ip link set nat_wan netns nat
ip link set nat_lan netns nat
ip link set unr_lan netns unr

# the external interface
ip address add 10.0.2.16/24 dev nat_real
ip link set nat_real up

# the two router interfaces
ip netns exec nat ip address add 10.0.2.17/24 dev nat_wan
ip netns exec nat ip link set nat_wan up
ip netns exec nat ip address add 172.20.0.1/24 dev nat_lan
ip netns exec nat ip link set nat_lan up

# the unreachable peer
ip netns exec unr ip address add 172.20.0.2/24 dev unr_lan
ip netns exec unr ip link set unr_lan up

#ip netns exec unr ifconfig lo 127.0.0.1 up
#ip netns exec nat ifconfig lo 127.0.0.1 up

ip route add 10.0.2.17 dev nat_real
ip netns exec unr route add default gw 172.20.0.1
ip netns exec nat route add default gw 10.0.2.16

# enable packet forwarding
ip netns exec nat sh -c "echo 1 > /proc/sys/net/ipv4/ip_forward"
echo 1 > /proc/sys/net/ipv4/ip_forward
echo 1 > /proc/sys/net/ipv4/conf/all/proxy_arp

# setup iptables
ip netns exec nat iptables -t nat -A POSTROUTING -o nat_wan -j MASQUERADE
ip netns exec nat iptables -A FORWARD -i nat_wan -o nat_lan -m state --state RELATED,ESTABLISHED -j ACCEPT
ip netns exec nat iptables -A FORWARD -i nat_lan -o nat_wan -j ACCEPT

#https://gist.github.com/dpino/6c0dca1742093346461e11aa8f608a99
#https://github.com/tomp2p/TomP2P/blob/c9589a967f6b2b23d14f8d5fd49d2957f67d0a41/nat/src/test/resources/nat-net.sh#L131

#test with:
#unr: nc -u -p 5000 10.0.2.15 4000
#unr: nc -l -u -p 5000
#glb: nc -u -s 10.0.2.15 -p 4000 10.0.2.17 5000

#watch conntrack -L
