CREATE TABLE tweets (
  tweet_id BIGINT UNIQUE PRIMARY KEY,
  tweet_text TEXT,
  tweet_timestamp TIMESTAMP,
  created_dt TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
);