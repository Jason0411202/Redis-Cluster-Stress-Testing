# Remove consumer, let memory exceed max memory, and observe what happens
1. To speed up the experiment process, modify the docker compose file to limit Redis's max memory to 3MB. This size allows Redis to start successfully but will quickly be exhausted by the producer's messages.
    ```yaml
    entrypoint: [redis-server, /etc/redis/rediscluster.conf, --port,"7000", --cluster-announce-ip,"${ip}", --appendonly, yes, "--maxmemory", "3mb", "--maxmemory-policy", "allkeys-lru"]
    ```
2. Additionally, modify the program in `main.go` by commenting out the consumer part, leaving only the producer.
    ```go
	go Producer(log) // start producer
	AutoClaim(log)   // start auto claim, auto claim will claim messages that have been idle for 300 seconds
	//Consumer(log)     // start consumer
    ```
3. Navigate to the root directory of this project and start the Redis cluster with the following command:
    ```shell
    docker-compose up -d --build
    ```
4. Then start the producer-consumer model and observe the log output with the following command:
    ```shell
    go run main.go
    ```
5. The producer initially works normally, continuously sending messages to the Redis stream.
   ![alt text](readme_img/image-2.png)
6. After sending about 9000 messages, the AutoClaim part encounters some unexpected errors, but the producer continues to work normally.
   ![alt text](readme_img/image-3.png)
7. When the producer sends about 20000 messages, Redis's memory is completely exhausted, and even the producer triggers OOM (Out of Memory) errors while sending messages. After exhausting the retry count, the program stops running.
   ![alt text](readme_img/image-4.png)

## Summary
When Redis runs out of memory, a series of unexpected situations and errors may occur. It is recommended to handle these potential situations specifically in the program and to prevent Redis memory from being completely exhausted as much as possible.
