SELECT product_id FROM products WHERE low_fats = 'Y'
UNION
SELECT product_id FROM products WHERE recyclable = 'Y';