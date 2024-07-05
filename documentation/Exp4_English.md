# Impact of AOF (Append-Only File) on the Experiment Structure
* Experiment conducted by modifying the --appendonly parameter in docker-compose.yaml:
  * Without AOF: 12.6642449s, 12.7502929s, 12.6436454s, average 12.6860611s
  * With AOF: 13.1419384s, 12.8154547s, 13.0021116s, average 12.9865016s
* It is observed that enabling AOF does slightly impact performance.