CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `pwd` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `soarconfig` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cfgname` varchar(255) NOT NULL DEFAULT '',
  `cfgonline` varchar(255) NOT NULL DEFAULT '',
  `cfgtest` varchar(255) NOT NULL DEFAULT '',
  `allowonlineastest` tinyint(4) NOT NULL DEFAULT '0',
  `sampling` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;