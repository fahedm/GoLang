DROP TABLE IF EXISTS album;
/*Delete (drop) a table called album. 
Executing this command first makes it easier for you to re-run the script later 
if you want to start over with the table.*/

CREATE TABLE album(
    id  INT AUTO_INCREMENT NOT NULL,
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

/*
Intall mysql in system
start a terminal for mysql
create a new Database recordings, use recordings
run this file using source /path/to/create-tables.sql
query to check if record created or not > select * from album;
*/