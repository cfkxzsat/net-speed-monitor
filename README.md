# net-speed-monitor
CLI net speed monitor program for Linux

## Solution specifically for linux
Linux kernel sysfs
```/sys/class/net/eth0/statistics/rx_packets:```收到的数据包数据

```/sys/class/net/eth0/statistics/tx_packets:```传输的数据包数量

```/sys/class/net/eth0/statistics/rx_bytes:```接收的字节数

```/sys/class/net/eth0/statistics/tx_bytes:```传输的字节数

```/sys/class/net/eth0/statistics/rx_dropped:```收包时丢弃的数据包

```/sys/class/net/eth0/statistics/tx_dropped:```发包时丢弃的数据包
