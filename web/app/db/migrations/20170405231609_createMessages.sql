
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Message(id integer primary key, body varchar(255), user_id integer, created_at text, FOREIGN KEY(user_id) REFERENCES User(id));
INSERT INTO Message VALUES(1, "hello, world.", 1, '2017-04-01 10:00:00');
INSERT INTO Message VALUES(2, "Bye bye world.", 1,'2017-04-01 11:00:00');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Message;
