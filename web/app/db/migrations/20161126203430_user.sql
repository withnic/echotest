
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE User(id integer primary key, email varchar(255) unique);
INSERT INTO User VALUES(1, "hogehoge@test.com");
INSERT INTO User VALUES(2, "hogefuga@test.com");
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE User;
