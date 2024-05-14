CREATE TABLE chats (
    id bigserial PRIMARY KEY,
    type varchar(15) NOT NULL,
    title varchar(127),
    description text,
    invite_hash varchar(16) UNIQUE,
    join_by_link_count integer DEFAULT 0 NOT NULL CHECK (join_by_link_count >= 0),
    is_private boolean DEFAULT FALSE NOT NULL,
    join_by_request boolean DEFAULT FALSE NOT NULL
);
