ALTER TABLE model
  ADD CONSTRAINT model_unique_name_4_brand UNIQUE(`name`, brand_id);