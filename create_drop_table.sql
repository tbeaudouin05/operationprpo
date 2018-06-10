CREATE TABLE baa_application.operation.purchase_request (
  id_purchase_request int IDENTITY(1,1) PRIMARY KEY
  ,cost_center  VARCHAR(50)
  ,initiator  VARCHAR(50)
  ,pr_type  VARCHAR(50)
  ,cost_category  VARCHAR(50)
  ,invoice_number  VARCHAR(50)
  ,invoice_date  DATE
  ,vendor_name  VARCHAR(50)
  ,item_description  VARCHAR(50)
  ,unit_price  REAL
  ,vat_unit_price  REAL
  ,quantity  INTEGER
  ,payment_term  VARCHAR(50)
  ,payment_installment  VARCHAR(50)
  ,payment_center  VARCHAR(50)
  ,payment_type  VARCHAR(50)
  ,invoice_total AS unit_price*quantity
  ,vat_invoice_total AS vat_unit_price*quantity
  ,purchase_request_status VARCHAR(50)
);


CREATE TABLE baa_application.operation.user_access (
  email VARCHAR(50) PRIMARY KEY
  ,name VARCHAR(50) NOT NULL
  ,access VARCHAR(50) NOT NULL
);

INSERT INTO  baa_application.operation.user_access (
  email
  ,name
  ,access)
VALUES ('julien.lefebvre@bamilo.com', 'Julien Lefebvre', 'pr_admin')

DELETE FROM baa_application.operation.user_access
WHERE baa_application.operation.user_access.email = 'julien.lefebvre@bamilo.com';



DELETE FROM baa_application.operation.purchase_request;
DELETE FROM baa_application.operation.user_access;

DROP TABLE baa_application.operation.purchase_request;
DROP TABLE baa_application.operation.user_access;


