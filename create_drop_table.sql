
ALTER TABLE baa_application.operation.purchase_request
ADD fk_vendor INT;

UPDATE baa_application.operation.purchase_request
SET fk_vendor = 1

ALTER TABLE baa_application.operation.purchase_request ALTER COLUMN item_description NVARCHAR(300);

ALTER TABLE baa_application.operation.purchase_request
ADD FOREIGN KEY (fk_vendor) REFERENCES baa_application.finance.vendor(id_vendor);

CREATE TABLE baa_application.operation.pr_cost_category (
  id_cost_category INT IDENTITY(1,1) PRIMARY KEY
  ,name VARCHAR(100) NOT NULL
  ,name_fa NVARCHAR(100) NOT NULL
);

ALTER TABLE baa_application.operation.pr_cost_category
ADD code BIGINT;

DELETE FROM baa_application.operation.pr_cost_category
WHERE id_cost_category = 30
 
ALTER TABLE baa_application.operation.pr_cost_category ALTER COLUMN code BIGINT NOT NULL;

ALTER TABLE baa_application.operation.pr_cost_category
ADD UNIQUE (code);

CREATE TABLE baa_application.finance.vendor (
  id_vendor INT IDENTITY(1,1) PRIMARY KEY
  ,name NVARCHAR(100) NOT NULL
  ,code BIGINT NOT NULL UNIQUE
);
                   


update baa_application.operation.pr_cost_category
set pr_cost_category.code = table1.code
from baa_application.operation.table1
inner join baa_application.operation.pr_cost_category on
pr_cost_category.name_fa = operation.table1.name_fa


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

  INSERT INTO  baa_application.operation.pr_department_access (
		fk_user
		,gfk_department)
	VALUES (@p1, @p2)

CREATE TABLE baa_application.operation.pr_location_access (
  id_location_access INT IDENTITY(1,1) PRIMARY KEY
  ,fk_user INT FOREIGN KEY REFERENCES baa_application.operation.pr_user(id_user)
  ,gfk_location VARCHAR(4) FOREIGN KEY REFERENCES baa_application.operation.location(gid_location)
)

  INSERT INTO  baa_application.operation.pr_location_access (
  fk_user
  ,gfk_location)
  VALUES (@p1, @p2)

  CREATE TABLE baa_application.operation.pr_division_access (
  id_division_access INT IDENTITY(1,1) PRIMARY KEY
  ,fk_user INT FOREIGN KEY REFERENCES baa_application.operation.pr_user(id_user) 
  ,fk_division CHAR(2) FOREIGN KEY REFERENCES baa_application.finance.division(id_division)

);
  INSERT INTO  baa_application.operation.pr_division_access (
  fk_user
  ,fk_division)
VALUES (@p1, @p2)


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
-- CREATE VIEW finance.cost_center_view AS operatio.cost_center_view UNION ALL marketing.cost_center_view etc.


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
  ,access)
VALUES ('thomas.beaudouin@bamilo.com', 'Thomas Beaudouin', 'admin')

DELETE FROM baa_application.operation.pr_user
WHERE baa_application.operation.pr_user.email = 'julien.lefebvre@bamilo.com';

UPDATE baa_application.operation.pr_user
SET access = 'cc_admin'
WHERE email = 'julien.lefebvre@bamilo.com'


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
  ) VALUES ('6910','01','FBB','FBB','01')

    DELETE FROM baa_application.operation.department
    WHERE gid_department = @p1

  INSERT INTO baa_application.operation.func (
  gfk_department
  ,function_code
  ,name
  ,tag
  ,tag_code
  ) VALUES ('691001','001','Receiving','RECE','001')

DELETE FROM baa_application.operation.purchase_request
WHERE initiator = 'Thomas Beaudouin'

DELETE FROM baa_application.operation.user_access;

DROP TABLE baa_application.operation.purchase_request;
DROP TABLE baa_application.operation.user_access;
DROP TABLE baa_application.operation.func;


