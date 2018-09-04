CREATE DATABASE IF NOT EXISTS test_db;

DROP TABLE IF EXISTS `test_db`.`sshconfigs`;
CREATE TABLE `test_db`.`sshconfigs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `hostname` varchar(253) DEFAULT NULL,
  `password` varchar(255) DEFAULT '',
  `username` varchar(32) DEFAULT NULL,
  `authkey` text,
  `proxy` int DEFAULT 0,
  `port` int DEFAULT 22,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
INSERT INTO test_db.sshconfigs(hostname, password, username, authkey, port) values('test-host.name', 'testdata', 'admin', '-----BEGIN RSA PRIVATE KEY-----test data-----END RSA PRIVATE KEY-----', 22);