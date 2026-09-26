ALTER TABLE followers
    DROP CONSTRAINT fk_user;

ALTER TABLE followers
    DROP CONSTRAINT fk_follower;

ALTER TABLE followers
    ADD CONSTRAINT fk_user
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE;

ALTER TABLE followers
    ADD CONSTRAINT fk_follower
        FOREIGN KEY (follower_id)
            REFERENCES users (id)
            ON DELETE CASCADE;
