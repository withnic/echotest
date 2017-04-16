
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Follow(id integer primary key, user_id integer, follow_id integer, created_at text, FOREIGN KEY(user_id) REFERENCES User(id), FOREIGN KEY(follow_id) REFERENCES User(id));
CREATE INDEX user_id_follow ON Follow(user_id);
INSERT INTO Follow VALUES(1, 1, 2, '2017-04-01 10:00:00');
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Follow;
