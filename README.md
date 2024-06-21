# Redis-Cluster-Stress-Testing
## Redis Cluster Environment Setup
The project directory will be organized as follows:
![Project Directory](readme_img/image.png)
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
```

### Start Redis Cluster
```shell
docker-compose up -d --build
```
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

## Run the Producer-Consumer Model
* Run `main.go` directly in the project directory

If you see log information similar to the following output, the program has run successfully:
![Log Output](readme_img/image1.png)

## References
1. https://pdai.tech/md/db/nosql-redis/db-redis-data-type-stream.html?source=post_page-----2a51f449343a--------------------------------
2. https://blog.yowko.com/docker-compose-redis-cluster/
3. https://www.yoyoask.com/?p=6051
4. https://blog.csdn.net/weixin_43798031/article/details/131322622
5. https://www.cnblogs.com/goldsunshine/p/17410148.html

# Redis-Cluster-Stress-Testing
## Redis cluster 環境配置
專案目錄將會如下
![alt text](readme_img/image.png)
### 於 redis 資料夾中新增 rediscluster.conf 檔案
範例
```shell
# ip
bind 0.0.0.0
# 啟用 cluster
cluster-enabled yes
# 指定 cluster config 檔案
cluster-config-file nodes.conf
# 指定 node 無法連線時間
cluster-node-timeout 5000
#設置主服務的連接密碼
masterauth 「自行設定的 redis 資料庫密碼」
#設置從服務的連接密碼
requirepass 「自行設定的 redis 資料庫密碼」
```

### 於專案根目錄中新增 .env 檔案
範例
```shell
ip=「你的設備 IP」
REDIS_PASSWORD=「自行設定的 redis 資料庫密碼，要與 rediscluster.conf 一致」
STREAM_NAME= 「用來交換訊息的 stream name」
CUSTOMER_GROUPNAME=「customer 的 group name」
```

### 啟動 Redis Cluster
```shell
docker-compose up -d --build
```

### 確認 redis cluster 是否正常運作
```shell
redis-cli -a 「自行設定的 redis 資料庫密碼」 -p 7000 cluster info
```

如果輸出類似以下資訊，代表 Redis Cluster 已經正常運作
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

## 執行 producer-consumer model
* 在專案目錄下直接執行 main.go 即可

如果看到類似以下輸出的 log 資訊，即為成功
![alt text](readme_img/image1.png)

## 實驗
### 持續送過程中把 master 砍掉會發生什麼事，以及觀察 failover 機制
1. 來到本專案根目錄下，先使用以下指令啟動 redis cluster
    ```shell
    docker-compose up -d --build
    ```
2. 接著透過以下指令啟動 producer-consumer model
    ```shell
    go run main.go
    ```
3. 透過以下指令觀察 redis cluster 狀態
    ```shell
    redis-cli -a "replace with your redis password" -p 7005 cluster info
    ```
    輸出:
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
    cluster_stats_messages_ping_sent:70
    cluster_stats_messages_pong_sent:72
    cluster_stats_messages_meet_sent:1
    cluster_stats_messages_sent:143
    cluster_stats_messages_ping_received:72
    cluster_stats_messages_pong_received:71
    cluster_stats_messages_received:143
    total_cluster_links_buffer_limit_exceeded:0
    ```
4. 透過以下指令觀察 redis cluster node 狀態
    ```shell
    redis-cli -a "replace with your redis password" -p 7005 cluster nodes
    ```
    輸出:
    ```
    461be0b4a4543e5d9acd00942dd9b93b882115b7 26.9.179.171:7000@17000 master - 0 1718974312455 1 connected 0-5460
    d2cc64d7259a4230b757e5b52c74f742b4bc9e03 26.9.179.171:7004@17004 slave 461be0b4a4543e5d9acd00942dd9b93b882115b7 0 1718974311000 1 connected
    c3f97ac46e687af0940c42e56ad5b1774cdcfa2d 26.9.179.171:7001@17001 master - 0 1718974312000 2 connected 5461-10922
    52b55e41072d08db5710971d62b5dbd23676286a 26.9.179.171:7002@17002 master - 0 1718974311000 3 connected 10923-16383
    d4103efba01a43ff73a0dde1a00962d9c49015cb 26.9.179.171:7003@17003 slave 52b55e41072d08db5710971d62b5dbd23676286a 0 1718974311554 3 connected
    49c048b897ef0a9fd8fc90326bb9c7a5945e325f 26.9.179.171:7005@17005 myself,slave c3f97ac46e687af0940c42e56ad5b1774cdcfa2d 0 1718974311000 2 connected
    ```
5. 手動砍掉一個 master node (:7000 node)，會發現數秒後， redis cluster 的 cluster_state 暫時轉成 fail，以至於觸發 main.go 中的 FATAL 錯誤
    ```
    CLUSTERDOWN The cluster is down
    ```
    緊接著 redis cluster 的 cluster_state 再次轉成 ok，重啟 main.go 後程式也順利執行。查看 redis cluster node 狀態
    輸出:
    ```
    461be0b4a4543e5d9acd00942dd9b93b882115b7 26.9.179.171:7000@17000 master,fail - 1718974349457 1718974347534 1 connected
    d2cc64d7259a4230b757e5b52c74f742b4bc9e03 26.9.179.171:7004@17004 master - 0 1718974388562 7 connected 0-5460
    c3f97ac46e687af0940c42e56ad5b1774cdcfa2d 26.9.179.171:7001@17001 master - 0 1718974388864 2 connected 5461-10922
    52b55e41072d08db5710971d62b5dbd23676286a 26.9.179.171:7002@17002 master - 0 1718974388000 3 connected 10923-16383
    d4103efba01a43ff73a0dde1a00962d9c49015cb 26.9.179.171:7003@17003 slave 52b55e41072d08db5710971d62b5dbd23676286a 0 1718974387000 3 connected
    49c048b897ef0a9fd8fc90326bb9c7a5945e325f 26.9.179.171:7005@17005 myself,slave c3f97ac46e687af0940c42e56ad5b1774cdcfa2d 0 1718974387000 2 connected
    ```
    觀察上述輸出，可以發現原本的 master 節點 (:7000 node) 已經被標記為 fail，並且原本的 slave node (:7004 node) 被提升成 master node
6. 接著再手動砍掉一個 master node (:7001 node)，跟砍掉 :7000 node 時相同，會發現數秒後， redis cluster 的 state 暫時轉成 fail，以至於觸發 main.go 中的 FATAL 錯誤
    ```
    CLUSTERDOWN The cluster is down
    ```
    緊接著 redis cluster 的 cluster_state 再次轉成 ok，重啟 main.go 後程式也順利執行。查看 redis cluster node 狀態
    輸出:
    ```
    461be0b4a4543e5d9acd00942dd9b93b882115b7 26.9.179.171:7000@17000 master,fail - 1718974349457 1718974347534 1 connected
    d2cc64d7259a4230b757e5b52c74f742b4bc9e03 26.9.179.171:7004@17004 master - 0 1718974477310 7 connected 0-5460
    c3f97ac46e687af0940c42e56ad5b1774cdcfa2d 26.9.179.171:7001@17001 master,fail - 1718974429262 1718974428555 2 connected
    52b55e41072d08db5710971d62b5dbd23676286a 26.9.179.171:7002@17002 master - 0 1718974476000 3 connected 10923-16383
    d4103efba01a43ff73a0dde1a00962d9c49015cb 26.9.179.171:7003@17003 slave 52b55e41072d08db5710971d62b5dbd23676286a 0 1718974476288 3 connected
    49c048b897ef0a9fd8fc90326bb9c7a5945e325f 26.9.179.171:7005@17005 myself,master - 0 1718974477000 8 connected 5461-10922
    ```
    觀察上述輸出，可以發現原本的 master 節點 (:7001 node) 已經被標記為 fail，並且原本的 slave node (:7005 node) 被提升成 master node
7. 重複上述步驟，砍掉 :7002 node，觀察 redis cluster node 狀態，這次在砍掉的瞬間，main.go 跳出了 EOF 的錯誤
    ```
    EOF
    ```
    緊接著 redis cluster 的 cluster_state 再次轉成 ok，重啟 main.go 後程式仍能順利執行。查看 redis cluster node 狀態
    輸出:
    ```
    461be0b4a4543e5d9acd00942dd9b93b882115b7 26.9.179.171:7000@17000 master,fail - 1718974349457 1718974347534 1 connected
    d2cc64d7259a4230b757e5b52c74f742b4bc9e03 26.9.179.171:7004@17004 master - 0 1718974584831 7 connected 0-5460
    c3f97ac46e687af0940c42e56ad5b1774cdcfa2d 26.9.179.171:7001@17001 master,fail - 1718974429262 1718974428555 2 connected
    52b55e41072d08db5710971d62b5dbd23676286a 26.9.179.171:7002@17002 master,fail - 1718974548970 1718974548688 3 disconnected
    d4103efba01a43ff73a0dde1a00962d9c49015cb 26.9.179.171:7003@17003 master - 0 1718974583825 9 connected 10923-16383
    49c048b897ef0a9fd8fc90326bb9c7a5945e325f 26.9.179.171:7005@17005 myself,master - 0 1718974584000 8 connected 5461-10922
    ```
    觀察上述輸出，可以發現原本的 master 節點 (:7002 node) 已經被標記為 fail，並且原本的 slave node (:7003 node) 被提升成 master node
8. 重複上述步驟，砍掉 :7003 node，觀察 redis cluster node 狀態，這次在砍掉的瞬間，main.go 還是跳出了 EOF 的錯誤
    ```
    EOF
    ```
    但這次 redis cluster 的 cluster_state 繼續維持 fail 狀態，重啟 main.go 後程式也連不上 redis cluster。查看 redis cluster node 狀態
    輸出:
    ```
    461be0b4a4543e5d9acd00942dd9b93b882115b7 26.9.179.171:7000@17000 master,fail - 1718974349457 1718974347534 1 connected
    d2cc64d7259a4230b757e5b52c74f742b4bc9e03 26.9.179.171:7004@17004 master - 0 1718974899205 7 connected 0-5460
    c3f97ac46e687af0940c42e56ad5b1774cdcfa2d 26.9.179.171:7001@17001 master,fail - 1718974429262 1718974428555 2 connected
    52b55e41072d08db5710971d62b5dbd23676286a 26.9.179.171:7002@17002 master,fail - 1718974548970 1718974548688 3 connected
    d4103efba01a43ff73a0dde1a00962d9c49015cb 26.9.179.171:7003@17003 master,fail - 1718974730388 1718974729485 9 connected 10923-16383
    49c048b897ef0a9fd8fc90326bb9c7a5945e325f 26.9.179.171:7005@17005 myself,master - 0 1718974898000 8 connected 5461-10922
    ```
    這應該是因為過半數 master node 失效時，redis cluster 將無法繼續提供服務

## 參考資料
1. https://pdai.tech/md/db/nosql-redis/db-redis-data-type-stream.html?source=post_page-----2a51f449343a--------------------------------
2. https://blog.yowko.com/docker-compose-redis-cluster/
3. https://www.yoyoask.com/?p=6051
4. https://blog.csdn.net/weixin_43798031/article/details/131322622
5. https://www.cnblogs.com/goldsunshine/p/17410148.html

## 待做
- [x] XAUTOCLAIM 邏輯
- [ ] 實驗: consumer拿掉，使 memory 漲超過 max memory，觀察發生什麼事
- [x] 實驗: 持續送過程中把 master 砍掉會發生什麼事，以及觀察 failover 機制
- [ ] 實驗: 分別關掉 ACK 以及 Auto claim，觀察掉資料的情況