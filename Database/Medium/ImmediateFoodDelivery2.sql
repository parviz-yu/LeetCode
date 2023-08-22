/* Write your T-SQL query statement below */
-- First Option
WITH first_order AS (
    SELECT order_date, customer_pref_delivery_date
    FROM (
        SELECT
            *,
            ROW_NUMBER() OVER(PARTITION BY customer_id ORDER BY order_date) AS r
        FROM delivery
    ) q
    WHERE r = 1
)

SELECT ROUND(COUNT(*) * 100.0 / (SELECT COUNT(*) FROM first_order), 2) AS immediate_percentage
FROM first_order
WHERE order_date = customer_pref_delivery_date;

-- # Write your MySQL query statement below
-- Second Option
SELECT 
    ROUND(AVG(order_date = customer_pref_delivery_date) * 100.0, 2) AS immediate_percentage
FROM delivery
WHERE (customer_id, order_date) IN (
    SELECT customer_id, MIN(order_date) AS first_order
    FROM delivery
    GROUP BY customer_id
);