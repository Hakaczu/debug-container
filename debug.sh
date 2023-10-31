#!/bin/sh
hostname
echo "=== List files ==="
ls -la /
ls -la /mnt
ls 

echo "=== List network ==="
netstat -plnt
ifconfig

echo "=== Goint Sleep ==="

sleep infinity