/* Write your T-SQL query statement below */
SELECT [name]
FROM salesperson
WHERE sales_id NOT IN (
    SELECT o.sales_id
    FROM orders o
        INNER JOIN company c ON o.com_id = c.com_id
    WHERE c.name = 'RED'
);
