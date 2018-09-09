CREATE TABLE IF NOT EXISTS odrders {
  id CHAR(27) PRIMARY KEY,
  create_at TIMESTAMP WITH TIME ZONE NOT NULL,
  account_id CHAR(27) NOT NULL,
  total_price MONEY NOT NULL
};

CREATE TABLE IF NOT EXISTS order_products {
  order_id CHAR(27) PEFERENCES orders (id) ON DELETE CASCADE,
  product_id CHAR(27),
  quantity INT NOT NULL,
  PRIMARY KEY (product_id, order_id)
};