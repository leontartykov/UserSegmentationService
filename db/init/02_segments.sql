CREATE OR REPLACE PROCEDURE DeleteUserSegments(userId integer, segment_name text, deletedAt date) AS 
$$
    BEGIN
        IF EXISTS (SELECT FROM userssegments JOIN users ON userssegments.userName = users.nickname WHERE users.id = userId AND userssegments.segmentName = segment_name AND deleted_at IS NULL)
            THEN UPDATE userssegments SET deleted_at = deletedAt WHERE userssegments.segmentName = segment_name AND userName = (select nickname from users where users.id = userId);             
        END IF;
    END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE AddUserSegments(userId integer, segment_name text, addedAt date) AS 
$$
    BEGIN
        IF NOT EXISTS (SELECT FROM userssegments JOIN users ON userssegments.userName = users.nickname WHERE users.id = userId AND userssegments.segmentName = segment_name AND deleted_at IS NULL)
            THEN INSERT INTO userssegments (userName, segmentName, added_at) VALUES ((SELECT users.nickname FROM users
                     WHERE users.id = userId),  segment_name, addedAt);
        END IF;
    END;
$$
LANGUAGE plpgsql;