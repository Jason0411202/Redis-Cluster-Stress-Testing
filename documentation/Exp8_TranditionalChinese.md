# 開啟實驗七中的 redis Hook 對於性能的影響
* 本實驗將使用 Azure 提供的 Standard_E4ds_v4 虛擬機進行測試

## read/write performance
* 在這個實驗中，我將會讓 producer 及 consumer 收發共 300000 條訊息，測試其性能
* 未開啟 redis Hook: 52.033732367s, 51.992560416s, 52.214875285s, 平均 52.080389356
* 開啟 redis Hook: 52.610379859s, 52.387042074s, 53.172067901s, 平均 52.723163278

## 總結
redis Hook 可能會對性能略有影響，不過在本次實驗中，影響並不大
