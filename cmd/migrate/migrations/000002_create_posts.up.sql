CREATE TABLE IF NOT EXISTS posts
(
    id bigserial PRIMARY KEY,
    content citext NOT NULL,
    title citext NOT NULL,
    user_id bigint NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
