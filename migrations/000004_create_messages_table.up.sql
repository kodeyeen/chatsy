CREATE TABLE messages (
    id bigserial PRIMARY KEY,
    chat_id integer NOT NULL,
    sender_id integer,
    original_id integer,
    parent_id integer,
    text text NOT NULL,
    sent_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,

    FOREIGN KEY (chat_id) REFERENCES chats(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (sender_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (original_id) REFERENCES messages(id) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES messages(id) ON UPDATE CASCADE ON DELETE CASCADE
);
