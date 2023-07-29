/* Write your T-SQL query statement below */
SELECT p.product_name, SUM(unit) AS 'unit'
FROM products p
    INNER JOIN orders o ON p.product_id = o.product_id
WHERE o.order_date BETWEEN '2020-02-01' AND '2020-02-29'
GROUP BY o.product_id, p.product_name
HAVING SUM(unit) > 99