SELECT user_id, CONCAT(UPPER(LEFT(name, 1)), LOWER(RIGHT(name, LENGTH(name)-1))) AS name FROM users
ORDER BY user_id;

/* Write your T-SQL query statement below */
SELECT
  [user_id], 
  CONCAT(UPPER(SUBSTRING([name], 1, 1)), LOWER(SUBSTRING([name], 2, LEN([name])))) AS 'name'
FROM users
ORDER BY [user_id]