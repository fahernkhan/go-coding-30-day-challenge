SELECT
    a.email,
    -- q1_income
    CAST(COALESCE(SUM(CASE 
        WHEN EXTRACT(QUARTER FROM TO_DATE(r.dt, 'YYYY-MM-DD')) = 1 
        THEN r.amount 
    END), 0) AS DECIMAL(10,2)) AS q1_income,

    -- q2_income
    CAST(COALESCE(SUM(CASE 
        WHEN EXTRACT(QUARTER FROM TO_DATE(r.dt, 'YYYY-MM-DD')) = 2 
        THEN r.amount 
    END), 0) AS DECIMAL(10,2)) AS q2_income,

    -- q3_income
    CAST(COALESCE(SUM(CASE 
        WHEN EXTRACT(QUARTER FROM TO_DATE(r.dt, 'YYYY-MM-DD')) = 3 
        THEN r.amount 
    END), 0) AS DECIMAL(10,2)) AS q3_income,

    -- q4_income
    CAST(COALESCE(SUM(CASE 
        WHEN EXTRACT(QUARTER FROM TO_DATE(r.dt, 'YYYY-MM-DD')) = 4 
        HEN r.amount 
    END), 0) AS DECIMAL(10,2)) AS q4_income,

    -- total_yearly_income
    CAST(COALESCE(SUM(r.amount), ) AS DECIMAL(10,2)) AS total_yearly_income
FROM account a
INNER JOIN reports r ON a.id = r.account_id
WHERE TO_DATE(r.dt, 'YYYY-MM-DD') BETWEEN '2023-01-01' AND '2023-12-31'
GROUP BY a.id, a.email
ORDER BY a.email ASC;