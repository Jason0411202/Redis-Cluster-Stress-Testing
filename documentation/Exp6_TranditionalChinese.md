# 關於 redis connection pool timeout 的發生原因及解決方法探討
* 本篇將著重於探討在 producer-consumer model 中，發生 connection pool timeout 錯誤的情況

## connection pool timeout 錯誤簡介
* 根據 go redis 的官方文檔 (https://redis.uptrace.dev/guide/go-redis-debugging.html#connection-pool-size)，當在 golang 中使用 redis.NewClient() 連接至 redis 資料庫時，go-redis 會自動配置適當的 PoolSize (預設為 10*可用 CPU 數，可自訂)
* 當超過 Options.PoolTimeout 的時間後，pool 中仍然沒有可用的連接時，便會出現 connection pool timeout 的錯誤

## producer-consumer model 什麼時候會發生 connection pool timeout
* 當以下條件滿足時，便很可能會發生 connection pool timeout 的錯誤
    1. PoolSize 過小
    2. consumer 在使用 XReadGroup() 時，Block 參數設定為 0 (代表阻塞直到新訊息到達)
    3. 有多個 consumer 平行執行

當上述條件滿足，此時發生 connection pool timeout 的原因可能為，XReadGroup() 在阻塞時，是會佔用 connections 的，當有多個 consumer 平行執行，且各自阻塞時，可能會導致所有可用的 connections 都被佔用，進而導致 connection pool timeout 的錯誤

## producer-consumer model 避免發生 connection pool timeout 的推薦解決方案
1. 手動配置 PoolSize，確保 PoolSize 足夠大
2. consumer 在使用 XReadGroup() 時，不要將 Block 參數設定為 0，可以改成 1*time.Second 或其他適當的阻塞時間
3. 避免 consumer 的數量過多