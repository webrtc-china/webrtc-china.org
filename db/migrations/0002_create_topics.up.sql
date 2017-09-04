CREATE TABLE IF NOT EXISTS topics (
        topic_id                                SERIAL PRIMARY KEY NOT NULL,
        user_id                                 VARCHAR NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        title                                   VARCHAR NOT NULL,
        content                                 VARCHAR NOT NULL, 
        node                                    VARCHAR NOT NULL,
        created_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
        updated_at                              TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW() 
);

CREATE INDEX ON topics (user_id);
CREATE INDEX ON topics (title);
CREATE INDEX ON topics (node);