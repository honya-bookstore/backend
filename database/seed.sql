INSERT INTO order_statuses (name) VALUES 
  ('pending'),
  ('processing'),
  ('shipping'),
  ('delivered'),
  ('cancelled');

INSERT INTO discount_types (name) VALUES 
  ('percentage'),
  ('fixed');

INSERT INTO discount_categories (name) VALUES 
  ('publisher'),
  ('title');

INSERT INTO discount_operators (name) VALUES 
  ('greater_than'),
  ('less_than'),
  ('equal_to'),
  ('in');

INSERT INTO payment_statuses (name) VALUES 
  ('pending'),
  ('completed'),
  ('failed');

INSERT INTO payment_providers (name) VALUES 
  ('cod'),
  ('vnpay'),
  ('momo'),
  ('zalopay');

INSERT INTO ticket_statuses (name) VALUES 
  ('open'),
  ('closed');
