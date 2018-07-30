# Drop all tables
DROP TABLE IF EXISTS model_2_zip_count;
DROP TABLE IF EXISTS model;
DROP TABLE IF EXISTS brand;
DROP TABLE IF EXISTS zip_code;

# Zip codes
create table zip_code (
  id int(11) unsigned not null auto_increment,
  zip_code int(11) unsigned not null,

  CONSTRAINT `PK_zip` PRIMARY KEY (id),
  KEY (zip_code)
);

# Makes
create table brand (
  id int(11) unsigned not null auto_increment,
  name varchar(255) not null,

  CONSTRAINT `PK_brand` PRIMARY KEY (id),
  key (name)
);

# Models
create table model (
  id int(11) unsigned not null auto_increment,
  name varchar(255) not null,
  brand_id int(11) unsigned not null,


  CONSTRAINT `PK_model` PRIMARY KEY (id),
  CONSTRAINT `model_fkibfk_brand` FOREIGN KEY (brand_id) REFERENCES `brand`(id),

  KEY (name)
);

# Statistics table
create table model_2_zip_count (
  zip_code_id int(11) unsigned not null,
  model_id int(11) unsigned not null,
  total_count int(11) unsigned not null,

  CONSTRAINT `m2zc_ibfk_zip`
    FOREIGN KEY (zip_code_id)
    REFERENCES zip_code(id),

  CONSTRAINT `m2zc_ibfk_model`
    FOREIGN KEY (model_id)
    REFERENCES model(id)
)
