CREATE TABLE users (
    -- id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    id bigserial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    email varchar(320) UNIQUE NOT NULL,
    password_hash varchar(255) NOT NULL,
    joined_at timestamp DEFAULT CURRENT_TIMESTAMP
);
