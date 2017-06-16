CREATE DATABASE IF NOT EXISTS sample;
USE sample;
DROP TABLE IF EXISTS user;
CREATE TABLE IF NOT EXISTS user (
  id    		BIGINT unsigned PRIMARY KEY,
  username      VARCHAR(255) NOT NULL,
  email  		VARCHAR(255) NOT NULL,
  password		VARCHAR(255) NOT NULL,
  permission	VARCHAR(255) NOT NULL,
  confirmed		BOOLEAN default false,
  is_active		BOOLEAN default true,
  is_anonymous  BOOLEAN default false,
  created_at	DATETIME,
  updated_at	DATETIME,
  closed_at		DATETIME
) ENGINE=innodb;