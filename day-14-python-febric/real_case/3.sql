WITH expenses_sum AS (
    SELECT
    'Expenses' AS category,
    COUNT(*) AS total_transactions,
    SUM(amount) AS total_amount
    FROM expenses
    WHERE dt >= '2024-03-01' AND dt < '2024-04-01'
),

income_sum AS (
    SELECT
    'Income' AS category,
    COUNT(*) AS total_transactions,
    SUM(amount) AS total_amount
    FROM income
    WHERE dt >= '2024-03-01' AND dt < '2024-04-01'
)
SELECT * FROM expenses_sum
UNION ALL
SELECT * FROM income_sum;