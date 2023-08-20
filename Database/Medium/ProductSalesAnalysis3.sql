/* Write your T-SQL query statement below */
-- First Option
SELECT product_id, [year] AS first_year, quantity, price
FROM (
    SELECT
        *,
        DENSE_RANK() OVER (PARTITION BY product_id ORDER BY product_id, [year]) AS first_year
    FROM sales
) a
WHERE a.first_year = 1;

-- # Write your MySQL query statement below
-- Second Option
SELECT s.product_id, s.year AS first_year, s.quantity, s.price
FROM sales s
    INNER JOIN (
        SELECT product_id, MIN(year) 'first_year'
        FROM sales
        GROUP BY product_id
    ) m ON s.product_id = m.product_id AND s.year = m.first_year;