-- +migrate Up
CREATE TABLE IF NOT EXISTS user_table(
    user_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name varchar(255),
    second_name varchar(255)
);
-- +migrate Down
DROP TABLE IF EXISTS user_table;