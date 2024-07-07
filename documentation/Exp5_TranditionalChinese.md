# 不同的 redis 記憶體上限限制對於性能的影響
* 本實驗將使用 Azure 提供的 Standard_E4ds_v4 虛擬機進行測試，並將每個 redis server 的記憶體上限限制分別設定為 0.5GB、1GB、2GB、4GB
* 在本實驗中，我會盡量控制 memory 快耗盡，以進行壓力測試

## read/write performance
* 在這個實驗中，我將會讓 producer 及 consumer 收發共 100000 條訊息，測試其性能

0.5GB: 17.380386881s, 17.359720639s, 17.314774037s
1GB: 23.06953716s, 22.098108313s, 21.81327895s, 19.304909801s
2GB: 23.188280602s, 23.004992405s, 23.48166519s
4GB: 18.633168363s, 18.436542117s, 18.157488882s


## failover time
* 在這個實驗中，我將會在 producer 及 consumer 收發訊息的時候，砍掉一個 master node，觀察程式的 failover 時間

0.5GB: 20.07971ms, 33.415115ms, 92.885029ms