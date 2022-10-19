DROP INDEX IF EXISTS idx_user_id;
DROP INDEX IF EXISTS idx_unique_user_name;
DROP INDEX IF EXISTS idx_unique_user_email;
DROP INDEX IF EXISTS idx_role_id;
DROP INDEX IF EXISTS idx_unique_role_name;
DROP INDEX IF EXISTS idx_user_role_id;
DROP INDEX IF EXISTS idx_user_role_role_id;
DROP INDEX IF EXISTS idx_user_role_user_id;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users_roles;