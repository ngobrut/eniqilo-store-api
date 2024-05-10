DROP INDEX IF EXISTS unique_user_phone_idx;
DROP INDEX IF EXISTS unique_customer_phone_idx;

DROP INDEX IF EXISTS idx_sku;
DROP INDEX IF EXISTS idx_name;
DROP INDEX IF EXISTS idx_invoices_customer_id;
DROP INDEX IF EXISTS idx_invoice_products_invoice_id;
DROP INDEX IF EXISTS idx_invoice_products_product_id;

DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS invoice_products;
DROP TABLE IF EXISTS invoices;
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS users;



