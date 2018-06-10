CREATE TABLE baa_application.operation.purchase_request (
  id_purchase_request INT IDENTITY(1,1) PRIMARY KEY
  ,fk_cost_center  INT
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


-- NB: global ID are needed when IDs are spread accross different schemas e.g. operation.location, marketing.location etc. this allows more flexibility in defining Finance codes
-- the goal is to make each department independent in defining there codes while enabling Finance to easily get an aggregated overview of all departments 

-- division defines the division so can only belong to Finance
CREATE TABLE baa_application.finance.division (
  id_division CHAR(2) PRIMARY KEY -- no need of global ID since there is only one division table: finance.division
  ,name VARCHAR(100) UNIQUE
);

-- if we take location in a very loose meaning (ie. could be "Warehouse east" or "Warehouse second floor") then let us make it belong to the division to enable more choices
CREATE TABLE baa_application.operation.location (
  gid_location AS CONCAT(fk_division,location_code) PERSISTED PRIMARY KEY -- global ID: unique accross all divisions
  ,fk_division CHAR(2) FOREIGN KEY REFERENCES baa_application.finance.division(id_division)
  ,location_code CHAR(2)
  ,name VARCHAR(100) UNIQUE
);


-- department = team hence belongs to the division itself
CREATE TABLE baa_application.operation.department (
  gid_department AS CONCAT(gfk_location,department_code) PERSISTED PRIMARY KEY -- global ID: unique accross all divisions
  ,gfk_location VARCHAR(4) FOREIGN KEY REFERENCES baa_application.operation.location(gid_location) -- global foreign key: unique accross all divisions
  ,department_code CHAR(2)
  ,name VARCHAR(100) UNIQUE
);

-- function = sub_team hence belongs to the division itself
CREATE TABLE baa_application.operation.func (
  gid_function AS CONCAT(gfk_department,function_code) PERSISTED PRIMARY KEY -- global ID: unique accross all divisions
  ,gfk_department VARCHAR(6) FOREIGN KEY REFERENCES baa_application.operation.department(gid_department) -- global foreign key: unique accross all divisions
  ,function_code CHAR(9)
  ,name VARCHAR(100) UNIQUE
);

-- NB: given that gid_function is unique across all divisions THEN id_cost_center is unique across all divisions
-- view of cost_center across the whole organization
CREATE VIEW finance.cost_center_view AS
  SELECT 
    fu.gid_function id_cost_center -- global ID: unique accross all divisions
    ,CONCAT(fu.name, '-', di.name, '[',LEFT(lo.name,3),LEFT(de.name,3),']') cost_center_name
    ,di.id_division fk_division 
    ,di.name division_name
    ,lo.gid_location gfk_location -- global foreign key: unique accross all divisions
    ,lo.name location_name
    ,de.gid_department gfk_department -- global foreign key: unique accross all divisions
    ,de.name department_name
    ,fu.gid_function -- global ID: unique accross all divisions
    ,fu.name function_name
    
  FROM baa_application.finance.division di
  JOIN baa_application.operation.location lo
  ON lo.fk_division = di.id_division
  JOIN baa_application.operation.department de
  ON de.gfk_location = lo.gid_location
  JOIN baa_application.operation.func fu
  ON  fu.gfk_department = de.gid_department
  -- UNION ALL SELECT FROM ... JOIN marketing.location lo JOIN marketing.department

-- view of cost_center only for operation
CREATE VIEW operation.cost_center_view AS
  SELECT 
    fu.gid_function id_cost_center -- global ID: unique accross all divisions
    ,CONCAT(fu.name, '-', di.name, '[',LEFT(lo.name,3),LEFT(de.name,3),']') cost_center_name
    ,di.id_division fk_division 
    ,di.name division_name
    ,lo.gid_location gfk_location -- global foreign key: unique accross all divisions
    ,lo.name location_name
    ,de.gid_department gfk_department -- global foreign key: unique accross all divisions
    ,de.name department_name
    ,fu.gid_function -- global ID: unique accross all divisions
    ,fu.name function_name
    
  FROM baa_application.finance.division di
  JOIN baa_application.operation.location lo
  ON lo.fk_division = di.id_division
  JOIN baa_application.operation.department de
  ON de.gfk_location = lo.gid_location
  JOIN baa_application.operation.func fu
  ON  fu.gfk_department = de.gid_department




INSERT INTO  baa_application.operation.user_access (
  email
  ,name
  ,access)
VALUES ('julien.lefebvre@bamilo.com', 'Julien Lefebvre', 'pr_admin')

DELETE FROM baa_application.operation.user_access
WHERE baa_application.operation.user_access.email = 'julien.lefebvre@bamilo.com';

INSERT INTO baa_application.operation.location (
  fk_division
  ,location_code
  ,name
  ) VALUES ('02','07','Esfahan')

    INSERT INTO baa_application.operation.department (
  gfk_location
  ,department_code
  ,name
  ) VALUES ('0207','07','Esfahan')

  INSERT INTO baa_application.operation.func (
  gfk_department
  ,function_code
  ,name
  ) VALUES ('020707','07','Esfahan')

DELETE FROM baa_application.operation.purchase_request;
DELETE FROM baa_application.operation.user_access;

DROP TABLE baa_application.operation.purchase_request;
DROP TABLE baa_application.operation.user_access;
DROP TABLE baa_application.operation.func;


