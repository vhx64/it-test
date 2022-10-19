CREATE
    EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id         UUID         NOT NULL DEFAULT uuid_generate_v4(),
    user_name  varchar(500) NOT NULL,
    last_name  varchar(500) NOT NULL,
    first_name varchar(500) NOT NULL,
    password   varchar(60)  NOT NULL,
    email      varchar(500) NOT NULL,
    mobile     varchar(500) NOT NULL,
    aszf       boolean      NOT NULL,
    PRIMARY KEY (id)
);

----------- user index -----------
CREATE INDEX idx_user_id
    ON users (id);
CREATE unique index idx_unique_user_name
    on users (user_name);
CREATE unique index idx_unique_user_email
    on users (email);


---------- add admin user ------------ pass: test
INSERT INTO users (id, user_name, last_name, first_name, password, email, mobile, aszf)
VALUES ('c1a2dc4b-8709-47d8-ab45-fc3eb925833f', 'Admin', 'Admin', 'Admin',
        '$2a$08$4ncg/ibKDezk0lykPsD/COL4vsVGr0lKwWAp2o9iKlyinGkCFFsim', 'admin@admin.com', '06202694565', true);
