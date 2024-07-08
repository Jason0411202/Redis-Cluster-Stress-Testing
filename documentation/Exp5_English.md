# The Impact of Different Redis Memory Limits on Performance
* This experiment will use Azure's Standard_E4ds_v4 virtual machines for testing. The memory limits for each Redis server will be set to 0.5GB, 1GB, 2GB, and 4GB respectively.
* In this experiment, I will try to control memory consumption until it is nearly exhausted to conduct stress tests.

## Read/Write Performance
* In this experiment, I will let the producer and consumer send and receive a total of 100,000 messages to test their performance.

0.5GB: 17.380386881s, 17.359720639s, 17.314774037s
1GB: 19.304909801s, 18.975422097s, 18.706215438s
2GB: 19.583243025s, 19.458951985s, 18.961349243s
4GB: 18.633168363s, 18.436542117s, 18.157488882s

Conclusion: From the experiment results, it can be seen that even with a memory limit difference of up to 8 times, the difference in timing results is almost less than 10%. Therefore, changing the Redis server's memory limit does not significantly affect read/write performance.

## Failover Time
* In this experiment, I will terminate a master node while the producer and consumer are sending and receiving messages to observe the failover time of the program.

0.5GB: 93.820713ms, 92.823194ms, 87.89061ms
1GB: 128.576252ms, 282.872117ms, 99.022911ms
2GB: 20.80214ms, 113.59096ms, 28.045194ms
4GB: 1.975791485s, 1.30987155s, 2.596017418s

Conclusion: From the experiment results, it can be seen that when the memory limit is increased, the failover time for 0.5GB, 1GB, and 2GB does not show significant differences. However, when the memory limit is increased to 4GB, the failover time significantly increases (suspected to be due to nearly exhausting the device memory). To detect differences, it may be necessary to apply other types of stress to Redis besides nearly exhausting the memory, thus extending the Redis failover time to reveal potential differences.

## Summary
This experiment was not able to measure a linear relationship as expected. It is possible that the impact of Redis server memory limits on performance is non-linear, or it could be that insufficient stress was applied to Redis, making it difficult to detect differences.
