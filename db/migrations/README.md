
sql-migrate has the following subcommands:
new : Prepares a new migration file
up : Applies the DB migrations to the latest version
down : Rolls back the previous migration
redo : down and up. This is useful when you are working on the new features and find out you need to fix the migration you’ve already done.
status : Shows the migration histories including current one
skip : Marks the current status of DB as the latest version without actually running them. (see: https://github.com/rubenv/sql-migrate/pull/100)
Getting Started
1/ Download
Download the package with the following command.
$ go get -v github.com/rubenv/sql-migrate/...
This installs the sql-migrate executable file at the same time.
2/ Create a config file
Save the following content in the dbconfig.yml file.
development:
  dialect: sqlite3
  datasource: report.db
  table: migrations
  dir: migrations/sqlite3
test:
  dialect: mysql
  datasource: root:pass@tcp(localhost:3306)/report
  dir: migrations
As you imagine, you can use different dialect for different environments (e.g. postgress for production.) The migration tool creates the migration history table automatically with the default name gorp_migrations. You can rename it with the table parameter.
Create the migrations/sqlite3 directory if needed.
3/ Run new subcommand
$ sql-migrate new -env="development" create_user_table
This creates a new migration file under the directory you specified in the config named with a version number prefix.
4/ Edit the SQL file
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id      INTEGER      AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255)                NOT NULL,
    PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE users;
Up is for the actual migration, and Down is for the rollback migration.
5/ Run the migration
The command supports dryrun option. You can try it if you want to check the result without actually running it.
$ sql-migrate up -env="development" -dryrun
==> Would apply migration 20190212230933-create_user_table.sql (up)
CREATE TABLE IF NOT EXISTS users (
    id      INTEGER      AUTO_INCREMENT NOT NULL,
    name    VARCHAR(255)                NOT NULL,
    PRIMARY KEY (id)
);
And run without -dryrun option:
$ sql-migrate up -env="development"
Applied 1 migration
You can run status subcommand to see the status of the applied migration.
Rolling back
The down subcommand also supports the dryrun option.
$ sql-migrate down -env="development" -dryrun
==> Would apply migration 20190212230933-create_user_table.sql (down)
DROP TABLE users;
When you actually run it, you can see the table is removed.
References
https://github.com/rubenv/sql-migrate