- Make MySQL

```shell
docker run -v /path/to/db_data --name mysql_data busybox
docker run --volumes-from mysql_data --name mysql -e MYSQL_ROOT_PASSWORD=mysql -d -p 3306:3306 mysql
```

```musql
insert into testsecond.sshconfig(id, hostname, password, username, authkey, proxy, port) values(1, '192.168.0.2', 'password', 'user', '-----BEGIN RSA PRIVATE KEY-----
xxx ...
-----END RSA PRIVATE KEY-----', '0', 22);
insert into testsecond.sshconfig(id, hostname, password, username, authkey, proxy, port) values(2, '192.168.0.3', 'password', 'user', '-----BEGIN RSA PRIVATE KEY-----
ooo ...
-----END RSA PRIVATE KEY-----', '1', 22);
```

## ref

- [Golangでターミナル接続できるsshクライアントコマンドを作成する\(制御キー対応\) \| 俺的備忘録 〜なんかいろいろ〜](https://orebibou.com/2018/08/golang%e3%81%a7%e3%82%bf%e3%83%bc%e3%83%9f%e3%83%8a%e3%83%ab%e6%8e%a5%e7%b6%9a%e3%81%a7%e3%81%8d%e3%82%8bssh%e3%82%af%e3%83%a9%e3%82%a4%e3%82%a2%e3%83%b3%e3%83%88%e3%82%b3%e3%83%9e%e3%83%b3%e3%83%89/)
- [Golangでssh Proxy経由でのssh接続を行わせる\(多段プロキシ\) \| 俺的備忘録 〜なんかいろいろ〜](https://orebibou.com/2018/08/golang%e3%81%a7ssh-proxy%e7%b5%8c%e7%94%b1%e3%81%a7%e3%81%aessh%e6%8e%a5%e7%b6%9a%e3%82%92%e8%a1%8c%e3%82%8f%e3%81%9b%e3%82%8b%e5%a4%9a%e6%ae%b5proxy%e3%81%82%e3%82%8a/)

## TODO

- Dockerfile
- save remote host log