# 不同的 redis 記憶體上限限制對於性能的影響
* 本實驗將使用 Azure 提供的 Standard_E4ds_v4 虛擬機進行測試，並將每個 redis server 的記憶體上限限制分別設定為 0.5GB、1GB、2GB、4GB
* 在本實驗中，我會盡量控制 memory 快耗盡，以進行壓力測試

## read/write performance
* 在這個實驗中，我將會讓 producer 及 consumer 收發共 100000 條訊息，測試其性能

0.5GB: 19.638418574s, 19.873413649s, 19.916937318s

## failover time
* 在這個實驗中，我將會在 producer 及 consumer 收發訊息的時候，砍掉一個 master node，觀察程式的 failover 時間

0.5GB: 19.861812ms, 44.516484ms, 55.052168ms, 103.394765ms, 113.002525ms