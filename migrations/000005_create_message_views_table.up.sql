CREATE TABLE message_views (
    message_id integer REFERENCES messages(id) ON UPDATE CASCADE ON DELETE CASCADE,
    user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,

    CONSTRAINT message_views_pk PRIMARY KEY (message_id, user_id)
);
