WITH date_range AS (
    SELECT 
        MIN(ORDER_DATE) AS start_date,
        MIN(ORDER_DATE) + INTERVAL '10 year' AS end_date
    FROM ORDERS
),
max_price AS (
    SELECT MAX(PRICE) AS max_price
    FROM ORDERS
    WHERE ORDER_DATE BETWEEN (SELECT start_date FROM date_range) AND (SELECT end_date FROM date_range)
)
SELECT 
    c.NAME,
    o.PRICE
FROM CUSTOMERS c
JOIN ORDERS o ON c.ORDER_ID = o.ID
WHERE 
    o.ORDER_DATE BETWEEN (SELECT start_date FROM date_range) AND (SELECT end_date FROM date_range)
    AND o.PRICE = (SELECT max_price FROM max_price)
ORDER BY o.PRICE DESC;