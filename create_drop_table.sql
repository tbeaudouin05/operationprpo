CREATE TABLE baa_application.operation.purchase_request (
  id_purchase_request INT IDENTITY(1,1) PRIMARY KEY
  ,gfk_cost_center  VARCHAR(9) FOREIGN KEY REFERENCES baa_application.operation.func(gid_function) -- global foreign key: unique accross all divisions
  ,initiator  VARCHAR(50)
  ,pr_type  VARCHAR(50)
  ,fk_cost_category  INT FOREIGN KEY REFERENCES baa_application.operation.pr_cost_category(id_cost_category)
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

CREATE TABLE baa_application.operation.pr_cost_category (
  id_cost_category INT IDENTITY(1,1) PRIMARY KEY
  ,name VARCHAR(100) NOT NULL
  ,name_fa NVARCHAR(100) NOT NULL
);



CREATE TABLE baa_application.operation.pr_user (
  id_user INT IDENTITY(1,1) PRIMARY KEY
  ,email VARCHAR(50) NOT NULL UNIQUE
  ,name VARCHAR(50) NOT NULL
  ,access VARCHAR(50) NOT NULL
);

INSERT INTO  baa_application.operation.pr_user (
  email
  ,name
  ,access)
VALUES ('julien.lefebvre@bamilo.com', 'Julien Lefebvre', 'pr_admin')

CREATE TABLE baa_application.operation.pr_department_access (
  id_department_access INT IDENTITY(1,1) PRIMARY KEY
  ,fk_user INT FOREIGN KEY REFERENCES baa_application.operation.pr_user(id_user)
  ,gfk_department VARCHAR(6) FOREIGN KEY REFERENCES baa_application.operation.department(gid_department)
)

CREATE TABLE baa_application.operation.pr_location_access (
  id_location_access INT IDENTITY(1,1) PRIMARY KEY
  ,fk_user INT FOREIGN KEY REFERENCES baa_application.operation.pr_user(id_user)
  ,gfk_location VARCHAR(4) FOREIGN KEY REFERENCES baa_application.operation.location(gid_location)
)

  CREATE TABLE baa_application.operation.pr_division_access (
  id_division_access INT IDENTITY(1,1) PRIMARY KEY
  ,fk_user INT FOREIGN KEY REFERENCES baa_application.operation.pr_user(id_user) 
  ,fk_division CHAR(2) FOREIGN KEY REFERENCES baa_application.finance.division(id_division)

);
  INSERT INTO  baa_application.operation.pr_division_access (
  fk_user
  ,fk_division)
VALUES (3, '69')


-- NB: global ID are needed when IDs are spread accross different schemas e.g. operation.location, marketing.location etc. this allows more flexibility in defining Finance codes
-- the goal is to make each department independent in defining there codes while enabling Finance to easily get an aggregated overview of all departments 

-- division defines the division so can only belong to Finance
CREATE TABLE baa_application.finance.division (
  id_division CHAR(2) PRIMARY KEY -- no need of global ID since there is only one division table: finance.division
  ,name VARCHAR(100) UNIQUE
  ,tag CHAR(3)
  ,tag_code CHAR(2)
);

-- if we take location in a very loose meaning (ie. could be "Warehouse east" or "Warehouse second floor") then let us make it belong to the division to enable more choices
CREATE TABLE baa_application.operation.location (
  gid_location AS CONCAT(fk_division,location_code) PERSISTED PRIMARY KEY -- global ID: unique accross all divisions
  ,fk_division CHAR(2) FOREIGN KEY REFERENCES baa_application.finance.division(id_division)
  ,location_code CHAR(2) NOT NULL
  ,name VARCHAR(100)
  ,tag CHAR(3)
  ,tag_code CHAR(2)
);

ALTER TABLE baa_application.operation.location
  ADD name_fa NVARCHAR(100);


-- department = team hence belongs to the division itself
CREATE TABLE baa_application.operation.department (
  gid_department AS CONCAT(gfk_location,department_code) PERSISTED PRIMARY KEY -- global ID: unique accross all divisions
  ,gfk_location VARCHAR(4) FOREIGN KEY REFERENCES baa_application.operation.location(gid_location) -- global foreign key: unique accross all divisions
  ,department_code CHAR(2) NOT NULL
  ,name VARCHAR(100)
  ,tag CHAR(3)
  ,tag_code CHAR(2)
);

ALTER TABLE baa_application.operation.department
  ADD name_fa NVARCHAR(100);

-- function = sub_team hence belongs to the division itself
CREATE TABLE baa_application.operation.func (
  gid_function AS CONCAT(gfk_department,function_code) PERSISTED PRIMARY KEY -- global ID: unique accross all divisions
  ,gfk_department VARCHAR(6) FOREIGN KEY REFERENCES baa_application.operation.department(gid_department) -- global foreign key: unique accross all divisions
  ,function_code CHAR(3) NOT NULL
  ,name VARCHAR(100)
  ,tag CHAR(4)
  ,tag_code CHAR(3)
);

ALTER TABLE baa_application.operation.func
  ADD name_fa NVARCHAR(100);

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
    ,CONCAT(fu.name, ' | ', de.name, ' | ', lo.name) cost_center_name
    ,de.gid_department gfk_department -- global foreign key: unique accross all divisions
    ,CONCAT(de.name, ' | ', lo.name) department_name
    ,lo.gid_location gfk_location -- global foreign key: unique accross all divisions
    ,lo.name location_name
    ,di.id_division fk_division 
    ,di.name division_name

    
  FROM baa_application.finance.division di
  JOIN baa_application.operation.location lo
  ON lo.fk_division = di.id_division
  JOIN baa_application.operation.department de
  ON de.gfk_location = lo.gid_location
  JOIN baa_application.operation.func fu
  ON  fu.gfk_department = de.gid_department


INSERT INTO  baa_application.operation.pr_user (
  email
  ,name
  ,access
  ,gfk_department)
VALUES ('thomas.beaudouin@bamilo.com', 'Thomas Beaudouin', 'admin', '690301')

DELETE FROM baa_application.operation.user_access
WHERE baa_application.operation.user_access.email = 'julien.lefebvre@bamilo.com';

INSERT INTO baa_application.finance.division (
  id_division
  ,name
  ,tag
  ,tag_code
  ) VALUES ('69','Operations','OPE','69')

INSERT INTO baa_application.operation.location (
  fk_division
  ,location_code
  ,name
  ,tag
  ,tag_code
  ) VALUES ('69','02','Customer Service','CUS','02')

  INSERT INTO baa_application.operation.department (
  gfk_location
  ,department_code
  ,name
  ,tag
  ,tag_code
  ) VALUES ('6901','01','FBB','FBB','01')

    DELETE FROM baa_application.operation.department
    WHERE gid_department = @p1

  INSERT INTO baa_application.operation.func (
  gfk_department
  ,function_code
  ,name
  ,tag
  ,tag_code
  ) VALUES ('690301','016','Consignment Inbound','CONI','016')

DELETE FROM baa_application.operation.purchase_request;
DELETE FROM baa_application.operation.user_access;

DROP TABLE baa_application.operation.purchase_request;
DROP TABLE baa_application.operation.user_access;
DROP TABLE baa_application.operation.func;


