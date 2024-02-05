CREATE TABLE participations (
    id bigserial PRIMARY KEY,
    role varchar(15) NOT NULL,
    are_notifications_enabled boolean DEFAULT TRUE NOT NULL,
    chat_id integer NOT NULL,
    user_id integer NOT NULL,

    FOREIGN KEY (chat_id) REFERENCES chats(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);
