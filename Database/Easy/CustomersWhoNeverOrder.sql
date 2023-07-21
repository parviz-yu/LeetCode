-- First Option
SELECT name AS 'Customers'
FROM customers c
WHERE c.id NOT IN (SELECT customerId FROM orders);

-- Second Option
SELECT c.name AS 'Customers'
FROM Customers c
LEFT JOIN Orders o ON c.id = o.customerId
WHERE o.customerId IS Null;