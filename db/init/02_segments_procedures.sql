\c usersegmentappdb

CREATE OR REPLACE PROCEDURE DeleteUserSegments(userId integer, segmentName text, deletedAt date) AS 
$$
    BEGIN
        IF EXISTS (SELECT FROM users_segments as us WHERE us.user_id = userId AND us.segment_name = segmentName AND deleted_at IS NULL)
            THEN UPDATE users_segments SET deleted_at = deletedAt WHERE users_segments.segment_name = segmentName AND user_id = userId;             
        END IF;
    END;
$$
LANGUAGE plpgsql;

CREATE OR REPLACE PROCEDURE AddUserSegments(userId integer, segmentName text, addedAt date) AS 
$$
    BEGIN
        IF EXISTS (SELECT FROM segments)
        THEN
            BEGIN
                IF NOT EXISTS (SELECT FROM users_segments)
                    THEN INSERT INTO users_segments (user_id, segment_name, added_at) VALUES (userId,  segmentName, addedAt);
                ELSE
                    BEGIN
                        IF NOT EXISTS (SELECT FROM users_segments as us WHERE us.user_id = userId AND us.segment_name = segmentName AND us.deleted_at IS NULL)
                            THEN INSERT INTO users_segments (user_id, segment_name, added_at) VALUES (userId,  segmentName, addedAt);
                        END IF;
                    END;
                END IF;
            END;
        END IF;
        
    END;
$$
LANGUAGE plpgsql;