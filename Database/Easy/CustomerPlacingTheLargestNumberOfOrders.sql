/* Write your T-SQL query statement below */
-- First Option
SELECT customer_number
FROM orders
GROUP BY customer_number
HAVING COUNT(order_number) = (
    SELECT MAX(ord)
    FROM (
        SELECT COUNT(order_number) AS 'ord'
        FROM orders
        GROUP BY customer_number
    ) t1
)

-- Second Option
SELECT TOP 1 customer_number
FROM orders
GROUP BY customer_number
ORDER BY COUNT(order_number) DESC
