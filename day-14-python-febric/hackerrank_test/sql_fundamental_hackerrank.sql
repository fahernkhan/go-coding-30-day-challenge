
-- ===============================
-- 📂 Template & Soal SQL Fundamental
-- ===============================

-- SOAL 1: Menampilkan semua kolom dari tabel 'employees'
SELECT * FROM employees;

-- SOAL 2: Menampilkan nama dan gaji karyawan dengan gaji > 5000
SELECT name, salary
FROM employees
WHERE salary > 5000;

-- SOAL 3: Menghitung rata-rata gaji semua karyawan
SELECT AVG(salary) AS average_salary
FROM employees;

-- SOAL 4: Menampilkan jumlah karyawan untuk setiap departemen
SELECT department_id, COUNT(*) AS total_employees
FROM employees
GROUP BY department_id;

-- SOAL 5: Menampilkan nama karyawan dan nama departemen-nya (JOIN)
SELECT e.name AS employee_name, d.name AS department_name
FROM employees e
JOIN departments d ON e.department_id = d.id;

-- SOAL 6: Menampilkan 3 karyawan dengan gaji tertinggi
SELECT name, salary
FROM employees
ORDER BY salary DESC
LIMIT 3;

-- SOAL 7: Menampilkan departemen yang tidak memiliki karyawan
SELECT d.name
FROM departments d
LEFT JOIN employees e ON d.id = e.department_id
WHERE e.id IS NULL;

-- SOAL 8: Menampilkan karyawan yang memiliki gaji di atas rata-rata
SELECT name, salary
FROM employees
WHERE salary > (
    SELECT AVG(salary) FROM employees
);
