CREATE TABLE memberships (
    chat_id integer REFERENCES chats(id) ON UPDATE CASCADE ON DELETE CASCADE,
    user_id integer REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    role varchar(15) NOT NULL,
    are_notifications_enabled boolean DEFAULT TRUE NOT NULL,

    CONSTRAINT memberships_pk PRIMARY KEY (chat_id, user_id)
);

