-- Version: 1.01
-- Description: Create table admins
CREATE TABLE admins
(
    id              UUID        NOT NULL,
    name            TEXT        NOT NULL,
    phone_number    TEXT UNIQUE NOT NULL,
    hashed_password TEXT        NOT NULL,
    active          BOOLEAN     NOT NULL,
    created_at      TIMESTAMP   NOT NULL,
    updated_at      TIMESTAMP   NOT NULL,

    PRIMARY KEY (id)
);

-- Version: 1.02
-- Description: Create table users
CREATE TABLE users
(
    id              UUID        NOT NULL,
    name            TEXT        NOT NULL,
    phone_number    TEXT UNIQUE NOT NULL,
    hashed_password TEXT        NOT NULL,
    active          BOOLEAN     NOT NULL,
    created_at      TIMESTAMP   NOT NULL,
    updated_at      TIMESTAMP   NOT NULL,

    PRIMARY KEY (id)
);

-- Version: 1.03
-- Description: Create table products
CREATE TABLE products
(
    id          UUID      NOT NULL,
    name        TEXT      NOT NULL,
    description TEXT      NOT NULL,
    image_id    UUID      NOT NULL,
    price       INT       NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP NOT NULL,

    PRIMARY KEY (id)
);