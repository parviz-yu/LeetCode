/* Write your T-SQL query statement below */
SELECT COUNT(DISTINCT customer_id) AS rich_count
FROM store
WHERE amount > 500