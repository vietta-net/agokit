
-- +migrate Up
CREATE TABLE users (
  id char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  username varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  password varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  created_at datetime DEFAULT NULL,
  updated_at datetime DEFAULT NULL,
  deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


ALTER TABLE users
  ADD PRIMARY KEY (id),
  ADD UNIQUE KEY username (username);

-- +migrate Down
