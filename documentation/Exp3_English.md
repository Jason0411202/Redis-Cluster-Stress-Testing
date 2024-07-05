# Impact of Auto Claim on Data Loss
1. To simulate scenarios where data loss may occur, the ConsumingMessage part in main.go was uncommented, allowing the consumer to occasionally crash after reading messages without processing them (not ACKed). Additionally, adjusting the number of messages sent and the triggering time of Auto Claim can also affect this scenario.
   ![alt text](readme_img/image-1.png)
2. Navigate to the root directory of this project and start the Redis cluster with the following command:
    ```shell
    docker-compose up -d --build
    ```
3. Then start the producer-consumer model and observe the log output with the following command:
    ```shell
    go run main.go
    ```
4. Initially, both the producer and consumer operate normally. The producer continuously sends messages to the Redis stream, and the consumer retrieves messages from the stream. However, after some time, the consumer triggers the crash mechanism after reading messages without processing them. The producer continues to send messages until all are sent. Finally, if a message is detected to have been read but not ACKed for a period, the Auto Claim mechanism automatically reclaims the message for processing. Without the Auto Claim mechanism, these unprocessed messages could be lost.
   ![alt text](readme_img/image-11.png)