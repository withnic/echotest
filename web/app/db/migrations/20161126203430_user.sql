
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE User(id integer primary key, email varchar(255) unique, passwd text);
--MyPassword
INSERT INTO User VALUES(1, "hogehoge@test.com", "$2a$10$kqfL5SRcJM222pJ5Zjt14eIH9ISaFShy41o.gvFb4MjmyzuQY22kO");
--MyPassword
INSERT INTO User VALUES(2, "hogefuga@test.com", "$2a$10$NmT3wf3vsolV3PLVL8KNnObueD5IeBfVmr2Pa1EH.Tgug2MLFEGJy");
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE User;
