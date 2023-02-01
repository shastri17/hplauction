  DROP Database IF EXISTS hpl_auction;
  CREATE DATABASE hpl_auction;
  USE hpl_auction

  CREATE TABLE `team` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(100) NOT NULL,
    `password` varchar(100) NOT NULL,
    `team_name`  varchar(100),
    `logo` varchar(100),
    `purse_amount` int(11),
    `max_bid_amount` int(11),
    `total_players` int,
    `owners_name` varchar(255),
    `icon1` varchar(100),
    `icon2` varchar(100),
    `is_admin` bool default false,
    `token` varchar(255),
    PRIMARY KEY (`id`)
  ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

  CREATE TABLE `player` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100),
    `nick_name` varchar(100),
    `skill_area`  varchar(255),
    `batting_hand` varchar(100),
    `bowling_hand` varchar(100),
    `mobile_number` BIGINT,
    `whatsapp_number` BIGINT,
    `previously_played_teams` varchar(255),
    `image` varchar(100),
    `bid_amount` int(11),
    `is_sold` bool default false,
    `team_id` int default null,
    `team_name` varchar(100) null ,
    PRIMARY KEY (`id`),
    CONSTRAINT `player_teams` FOREIGN KEY (`team_id`) REFERENCES `team` (`id`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT  CHARSET=utf8;

