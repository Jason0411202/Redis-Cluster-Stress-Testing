# 開啟實驗七中的 redis Hook 對於性能的影響
* 本實驗將使用 Azure 提供的 Standard_E4ds_v4 虛擬機進行測試

## read/write performance
* 在這個實驗中，我將會讓 producer 及 consumer 收發共 300000 條訊息，測試其性能

* 未開啟 redis Hook: 53.860710054s, 55.209615959s, 54.701627418s
* 開啟 redis Hook: 52.344460477s, 52.925375122s, 

## 總結
這次的實驗並沒辦法如預料測出線性關係，有可能是 redis server 的記憶體上限限制對性能的影響是非線性的，也不排除是未能給 redis 足夠的壓力，導致無法測出差異
