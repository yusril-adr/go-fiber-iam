CREATE TABLE user_tokens (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP DEFAULT NULL,

    CONSTRAINT fk_bookings_user
        FOREIGN KEY (user_id) REFERENCES users(id)
);
