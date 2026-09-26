ALTER TABLE followers
    DROP CONSTRAINT fk_follower;

ALTER TABLE followers
    DROP CONSTRAINT fk_user;

DROP TABLE IF EXISTS followers;
