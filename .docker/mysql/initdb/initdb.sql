CREATE DATABASE IF NOT EXISTS codepay;
CREATE DATABASE IF NOT EXISTS codepay_test;

USE codepay;

CREATE TABLE IF NOT EXISTS transactions
(
  id            VARCHAR(255) NOT NULL,
  account_id    VARCHAR(255) NOT NULL,
  amount        double       NOT NULL,
  status        VARCHAR(255) NOT NULL,
  error_message VARCHAR(255) NOT NULL,
  created_at    DATETIME     NOT NULL,
  updated_at    DATETIME     NOT NULL
);
