# Redis-Cluster-Stress-Testing
如果您想閱讀繁體中文版的安裝指南跟實驗結果，請點選 [此連結](documentation/Tranditional_install.md)
## Redis Cluster Environment Setup
The project directory will be organized as follows:
![Project Directory](documentation/readme_img/image.png)
### Create `rediscluster.conf` File in the Redis Directory

Example configuration:

```shell
# IP
bind 0.0.0.0
# Enable cluster
cluster-enabled yes
# Specify cluster config file
cluster-config-file nodes.conf
# Specify node timeout
cluster-node-timeout 5000
# Set master connection password
masterauth "your_redis_password"
# Set replica connection password
requirepass "your_redis_password"
```

### Create .env File in the Project Root Directory
Example configuration:
```shell
ip="your_device_ip"
REDIS_PASSWORD="your_redis_password"  # Must match the password in rediscluster.conf
STREAM_NAME="your_stream_name"
CUSTOMER_GROUPNAME="your_customer_groupname"
MaxEntries="maximum number of messages to be preloaded to control memory usage"
```

### Start Redis Cluster and the Producer-Consumer Model
```shell
docker-compose up -d --build
```

If you see log information similar to the following output, the program has run successfully:
![Log Output](documentation/readme_img/image1.png)

### Check if Redis Cluster is Running Properly
```shell
redis-cli -a "your_redis_password" -p 7000 cluster info
```

If the output is similar to the following, the Redis Cluster is operating correctly:
```
cluster_state:ok
cluster_slots_assigned:16384
cluster_slots_ok:16384
cluster_slots_pfail:0
cluster_slots_fail:0
cluster_known_nodes:6
cluster_size:3
cluster_current_epoch:6
cluster_my_epoch:2
cluster_stats_messages_ping_sent:63
cluster_stats_messages_pong_sent:69
cluster_stats_messages_meet_sent:4
cluster_stats_messages_sent:136
cluster_stats_messages_ping_received:68
cluster_stats_messages_pong_received:67
cluster_stats_messages_meet_received:1
cluster_stats_messages_received:136
```

### Check Redis Cluster Node Status
```shell
redis-cli -a "your_redis_password" -p 7000 cluster nodes
```

If the output is similar to the following, the Redis Cluster is operating correctly:
```shell
7e793eb8a2134550d9a285713f52ddaf8cc89189 26.9.179.171:7003@17003 slave 7651df59f610c9619ea8ecef840737344a762e12 0 1719301574086 3 connected       
a1c228d29fba6e1e0360f0caf5863f33688a93d0 26.9.179.171:7004@17004 master - 0 1719301573000 7 connected 0-5460
890b48ca513fd553e3319505e442c751cd25a1ea 26.9.179.171:7000@17000 myself,slave a1c228d29fba6e1e0360f0caf5863f33688a93d0 0 1719301573000 7 connected
cc4a188cf6ba936c097776f1a1fe203395f710bb 26.9.179.171:7001@17001 slave eb672df8d3073c0327084123bda8f022216b239e 0 1719301574488 8 connected       
eb672df8d3073c0327084123bda8f022216b239e 26.9.179.171:7005@17005 master - 0 1719301572982 8 connected 5461-10922
7651df59f610c9619ea8ecef840737344a762e12 26.9.179.171:7002@17002 master - 0 1719301573485 3 connected 10923-16383
```

## Experiment
- [Remove consumer, let memory exceed max memory, and observe what happens](documentation/Exp1_English.md)
- [Observe what happens when a master is killed during continuous sending and the failover mechanism](documentation/Exp2_English.md)
- [Impact of Auto Claim on Data Loss](documentation/Exp3_English.md)
- [Impact of AOF (Append-Only File) on the Experiment Structure](documentation/Exp4_English.md)

## References
1. https://pdai.tech/md/db/nosql-redis/db-redis-data-type-stream.html?source=post_page-----2a51f449343a--------------------------------
2. https://blog.yowko.com/docker-compose-redis-cluster/
3. https://www.yoyoask.com/?p=6051
4. https://blog.csdn.net/weixin_43798031/article/details/131322622
5. https://www.cnblogs.com/goldsunshine/p/17410148.html