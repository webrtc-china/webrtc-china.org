CREATE TABLE IF NOT EXISTS replies (
        id                                      SERIAL PRIMARY KEY NOT NULL,
        user_id                                 VARCHAR NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        topic_id                                VARCHAR NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
        content                                 VARCHAR NOT NULL, 
        created_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        updated_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW() 
);

CREATE INDEX ON replies (topic_id);