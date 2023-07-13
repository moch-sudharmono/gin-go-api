DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);


DROP TABLE IF EXISTS basicauth;
CREATE TABLE basicauth (
  id          INT AUTO_INCREMENT NOT NULL,
  user        VARCHAR(128) NOT NULL,
  scrt        VARCHAR(255) NOT NULL,
  active      TINYINT(1) NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`)
);

INSERT INTO basicauth
  (user, scrt, active)
VALUES
  ('root', 'pass', 1),
  ('user', 'pass', 0);