/* Write your T-SQL query statement below */
SELECT
  p.product_id,
  ROUND(SUM(p.price * u.units) * 1.0 / SUM(u.units), 2) AS 'average_price'
FROM unitssold u
LEFT JOIN prices p
  ON p.product_id = u.product_id
  AND u.purchase_date BETWEEN p.start_date AND p.end_date
GROUP BY p.product_id