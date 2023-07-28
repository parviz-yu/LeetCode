/* Write your T-SQL query statement below */
-- First Option
SELECT
ROUND(SUM(CASE
  WHEN order_date = customer_pref_delivery_date THEN 1
  ELSE 0
END) * 100.0 / COUNT(*), 2) AS 'immediate_percentage'
FROM delivery

-- Second Option
SELECT
  ROUND(SUM(CASE
    WHEN DATEDIFF(DAY, order_date, customer_pref_delivery_date) = 0 THEN 1
    ELSE 0
  END) * 100.0 / COUNT(*), 2) AS 'immediate_percentage'
FROM delivery