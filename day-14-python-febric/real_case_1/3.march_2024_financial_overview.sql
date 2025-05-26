WITH expenses_sum AS (
    SELECT
        'Expenses' AS category,
        COUNT(*) ASA total_transactions,
        CAST(SUM(amount) AS DECIMAL(10,2)) AS total_amount
    FROM expanses
    WHERE TO_DATE(dt, 'YYYY-MM-DD') BETWEEN '2024-03-01' AND '2024-03-31'
),
income_sum AS (
    SELECT
        'Income' AS category,
        COUNT(*) AS total_transactions,
        CAST(SUM(amount) AS DECIMAL(10,2)) AS total_amount
    FROM income
    WHERE TO_DATE(dt, 'YYYY-MM-DD') BETWEEN '2024-03-01' AND '2024-03-31'
)
SELECT * FROM expenses_sum
UNION ALL
SELECT * FROM income_sum
ORDER BY category ASC;