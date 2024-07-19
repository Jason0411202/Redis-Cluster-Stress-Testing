# 關於 redis EOF 的發生原因及解決方法探討
## redis EOF 錯誤簡介
* 當客戶端與 redis server 間建立連接後，因故導致連接異常終止，此時客戶端若使用了這個故障的連接進行操作時，便可能會出現 redis EOF 錯誤
* 以下整理可能的發生原因

### Connection pool 中的連接被 redis server 主動斷開
* 根據下列網路資料描述，go-redis 的 Connection pool 並沒有實現 Keepalive 機制
  * https://pandaychen.github.io/2020/02/22/A-REDIS-POOL-ANALYSIS/#0x05-%E4%B8%80%E4%BA%9B%E7%BB%86%E8%8A%82
* 在這種情況下，redis server 可能會斷開一些閒置過久的連接 (使用 --timeout 參數控制)；也就是說，這可能會導致 Connection pool 中的連接被 redis server 終止
* 在實際下指令時，可能會從 Connection pool 取到已經被 redis server 終止的連接，進而導致 redis EOF 錯誤
* 這種類型的錯誤，在這個 Issue 中也有提到
  * Issue 發起人將 redis sesrver 的 timeout 設定為 10 秒，但 idletimeout 參數卻仍為預設的 5 分鐘 (客戶端的空閒連接最多能保持 5 分鐘)，這導致了 Connection pool 中的連接可能已經被 redis server 終止，但仍然在 Connection pool 中
  * https://github.com/redis/go-redis/issues/1373

### redis server 的 Client Eviction 機制
* 根據 redis 官方文件，當所有客戶端的總記憶體使用量超過閾值，就會斷開客戶端連接 (優先斷開使用最多記憶體的客戶端)
  * https://redis.io/docs/latest/develop/reference/clients/
* 在這種情況下，也可能導致連接故障

## 解決 redis EOF 的流程
### 解法一： 透過 redis Hook 實現具一般性的 reconnect
* 以下的文章中提到，當操作時出現了 Redis EOF 錯誤，可以使用 reconnect 來解決
  * https://blog.51cto.com/u_16213459/8159401?fbclid=IwZXh0bgNhZW0CMTEAAR24CW_lSHO1Vs14b8wppnnycA5ou64hWzi_riv05Knnlnntq5mng-kZpIs_aem_Lmy2e8Xw3xjpu-4PSAnsBQ
* 不過上述的文章實作方式，需修改每一個操作，不夠具一般性
* 我試著提出了一種等價的解決方案，可以嘗試透過 redis Hook，將原本需要修改每一個操作的部分，改為在 Hook 中處理
* 此外，golang redis 的 reconnect 寫法可以參考以下文章
  * https://cloud.tencent.com/developer/article/2355690

### 解法二: 在應用層級上透過背景執行的 goroutine 實現 keepalive 心跳機制
* 可以參考以下文章
  * https://juejin.cn/post/6844903427097493517
* 這樣做能確保連接能被 redis server 視作 "非閒置"，進而避免被 redis server 斷開
* 缺點是不適合簡單的套用至使用 redis.NewClusterClient() 進行連接的情境中