
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE Follow(id integer primary key, user_id integer, follow_id integer, created_at text, FOREIGN KEY(user_id) REFERENCES User(id), FOREIGN KEY(follow_id) REFERENCES User(id));

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE Follow;
