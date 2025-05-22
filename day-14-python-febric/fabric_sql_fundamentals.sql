-- =======================================
-- 📘 FABRIC SQL FUNDAMENTALS - CUSTOMER DATA PIPELINE
-- =======================================

-- ---------------------------------------
-- 🧱 1. CREATE TABLE
-- ---------------------------------------

CREATE TABLE customers (
    customer_id INT,
    name VARCHAR(100),
    signup_date DATE,
    total_purchase INT,
    is_active BOOLEAN
);

-- ---------------------------------------
-- 🌱 2. INSERT DATA
-- ---------------------------------------

INSERT INTO customers (customer_id, name, signup_date, total_purchase, is_active)
VALUES 
(101, 'Ali', '2023-01-01', 200000, TRUE),
(102, 'Budi', '2023-01-03', 150000, TRUE),
(103, 'Citra', '2023-02-10', 320000, FALSE),
(104, 'Dewi', '2023-03-25', 270000, TRUE);

-- ---------------------------------------
-- 🔍 3. SELECT DASAR
-- ---------------------------------------

SELECT * FROM customers;

-- Menampilkan kolom tertentu
SELECT name, total_purchase FROM customers;

-- Filter data
SELECT * FROM customers
WHERE is_active = TRUE;

-- ---------------------------------------
-- 🧮 4. PENGGUNAAN FUNGSI
-- ---------------------------------------

-- Mengambil tahun dari tanggal
SELECT customer_id, EXTRACT(YEAR FROM signup_date) AS signup_year
FROM customers;

-- Klasifikasi berdasarkan total_purchase
SELECT *,
    CASE
        WHEN total_purchase <= 200000 THEN 'Bronze'
        WHEN total_purchase <= 300000 THEN 'Silver'
        ELSE 'Gold'
    END AS loyalty_tier
FROM customers;

-- ---------------------------------------
-- 📊 5. AGGREGATION
-- ---------------------------------------

-- Total pembelian semua pelanggan
SELECT SUM(total_purchase) AS total_all_purchase FROM customers;

-- Rata-rata pembelian
SELECT AVG(total_purchase) AS avg_purchase FROM customers;

-- Group by loyalty tier
SELECT 
    CASE
        WHEN total_purchase <= 200000 THEN 'Bronze'
        WHEN total_purchase <= 300000 THEN 'Silver'
        ELSE 'Gold'
    END AS loyalty_tier,
    COUNT(*) AS customer_count,
    AVG(total_purchase) AS avg_purchase
FROM customers
GROUP BY loyalty_tier;

-- ---------------------------------------
-- 🔄 6. UPDATE DATA
-- ---------------------------------------

-- Update status aktif pelanggan
UPDATE customers
SET is_active = FALSE
WHERE customer_id = 102;

-- ---------------------------------------
-- 🗑️ 7. DELETE DATA
-- ---------------------------------------

-- Hapus pelanggan yang tidak aktif
DELETE FROM customers
WHERE is_active = FALSE;

-- ---------------------------------------
-- 🧹 8. DROP TABLE
-- ---------------------------------------

-- Menghapus tabel jika sudah tidak diperlukan
DROP TABLE IF EXISTS customers;
