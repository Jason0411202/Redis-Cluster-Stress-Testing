# Observe what happens when a master is killed during continuous sending and the failover mechanism
1. Navigate to the root directory of this project and start the Redis cluster with the following command:
    ```shell
    docker-compose up -d --build
    ```
2. Then start the producer-consumer model with the following command:
    ```shell
    go run main.go
    ```
3. Check the Redis cluster status with the following command:
    ```shell
    redis-cli -a "replace with your redis password" -p 7005 cluster info
    ```
    Output:
    ```
    cluster_state:ok
    cluster_slots_assigned:16384
    cluster_slots_ok:16384
    cluster_slots_pfail:0
    cluster_slots_fail:0
    cluster_known_nodes:6
    cluster_size:3
    cluster_current_epoch:6
    cluster_my_epoch:3
    cluster_stats_messages_ping_sent:19
    cluster_stats_messages_pong_sent:17
    cluster_stats_messages_meet_sent:1
    cluster_stats_messages_sent:37
    cluster_stats_messages_ping_received:17
    cluster_stats_messages_pong_received:20
    cluster_stats_messages_received:37
    total_cluster_links_buffer_limit_exceeded:0
    ```
4. Check the Redis cluster node status with the following command:
    ```shell
    redis-cli -a "replace with your redis password" -p 7000 cluster nodes
    ```
    Output:
    ```
    9e1b9f1122acf92f85c0986d37872294a691d55b 26.9.179.171:7005@17005 slave 230310e5c4ef0c4e544c2f22e71a5cf74c7ce0a8 0 1719308437906 3 connected
    e3e0441634d8a6cd332d0698165d83dac6064dd9 26.9.179.171:7003@17003 slave b4ac2b113e50c5bc20c9b422b8c5ffc6e9248189 0 1719308437000 1 connected
    b4ac2b113e50c5bc20c9b422b8c5ffc6e9248189 26.9.179.171:7000@17000 myself,master - 0 1719308437000 1 connected 0-5460
    a62ff34550f00412fbb2024e0ca40c4f274ef5c6 26.9.179.171:7004@17004 slave 69b216af950a9ae48ac7a44200dfa87b9f0f67aa 0 1719308438911 2 connected
    69b216af950a9ae48ac7a44200dfa87b9f0f67aa 26.9.179.171:7001@17001 master - 0 1719308438509 2 connected 5461-10922
    230310e5c4ef0c4e544c2f22e71a5cf74c7ce0a8 26.9.179.171:7002@17002 master - 0 1719308437505 3 connected 10923-16383
    ```
5. In Docker, manually restart a master node (e.g., the :7000 node). You'll find that the "master node restart time" is shorter than the "time for Redis cluster failover mechanism to determine that the node has failed." This means the master node restarts in time before being marked as failed and continues to operate without being switched to a slave. The producer-consumer model also continues to operate normally.
    ![alt text](readme_img/image-7.png)
6. To simulate triggering a failover, manually pause the master node (:7000). After a few seconds, the Redis cluster goes temporarily offline, causing the producer-consumer model to keep retrying. After about 0.1 seconds, the Redis cluster resumes normal operation, and the producer-consumer model continues to operate normally.
    ![alt text](readme_img/image-8.png)
    ![alt text](readme_img/image-9.png)
7. Finally, manually restart the previously paused node (:7000). You'll find that the node has been switched to a slave without affecting the operation of the producer-consumer model.
    ![alt text](readme_img/image-10.png)