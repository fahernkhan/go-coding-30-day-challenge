WITH report_2023 AS (
  SELECT
    a.email,
    r.amount,
    TO_DATE(r.dt, 'YYYY-MM-DD') AS transaction_date,
    EXTRACT(QUARTER FROM TO_DATE(r.dt, 'YYYY-MM-DD')) AS quarter
  FROM accounts a
  JOIN reports r
    ON a.id = r.account_id
  WHERE TO_DATE(r.dt, 'YYYY-MM-DD') BETWEEN '2023-01-01' AND '2023-12-31'
)
SELECT
  email,
  CAST(COALESCE(SUM(CASE WHEN quarter = 1 THEN amount END), 0) AS DECIMAL(10,2)) AS q1_income,
  CAST(COALESCE(SUM(CASE WHEN quarter = 2 THEN amount END), 0) AS DECIMAL(10,2)) AS q2_income,
  CAST(COALESCE(SUM(CASE WHEN quarter = 3 THEN amount END), 0) AS DECIMAL(10,2)) AS q3_income,
  CAST(COALESCE(SUM(CASE WHEN quarter = 4 THEN amount END), 0) AS DECIMAL(10,2)) AS q4_income,
  CAST(COALESCE(SUM(amount), 0) AS DECIMAL(10,2)) AS total_yearly_income
FROM report_2023
GROUP BY email
ORDER BY email ASC;