CREATE TABLE IF NOT EXISTS users(
        id                                      VARCHAR PRIMARY KEY CHECK (id ~ '^[0-9a-f]{24,24}$'),
        username                                VARCHAR NOT NULL UNIQUE,
        email                                   VARCHAR NOT NULL UNIQUE,
        encrypted_password                      VARCHAR NOT NULL,
        full_name                               VARCHAR,
        avatar_url                              VARCHAR,
        created_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        updated_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX ON users ((LOWER(username)));
CREATE UNIQUE INDEX ON users ((LOWER(email)));
