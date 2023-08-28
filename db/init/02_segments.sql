CREATE OR REPLACE PROCEDURE changeUserSegments(userId integer, segment_name text, addedAt date, deletedAt date) AS 
$$
    BEGIN
        IF EXISTS (SELECT FROM userssegments JOIN users ON userssegments.userName = users.nickname WHERE users.id = userId AND userssegments.segmentName = segment_name AND deleted_at IS NULL)
            THEN UPDATE userssegments SET deleted_at = deletedAt;
            ELSE 
                WITH UserNickname AS 
                    (SELECT users.nickname FROM users
                     WHERE users.id = userId) 
                INSERT INTO userssegments (userName, segmentName, added_at) VALUES ((SELECT users.nickname FROM users
                     WHERE users.id = userId),  segment_name, addedAt);
        END IF;
    END;
$$
LANGUAGE plpgsql;