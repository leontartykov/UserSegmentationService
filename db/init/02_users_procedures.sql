\c usersegmentappdb

CREATE OR REPLACE FUNCTION get_users_without_current_segment(segmentName text) RETURNS TABLE (
		id_user integer
	) 
AS
$BODY$
    BEGIN
        RETURN QUERY
            SELECT DISTINCT user_id 
            FROM users_segments 
            WHERE segment_name = segmentName AND deleted_at IS NOT NULL 
            
            EXCEPT
            SELECT DISTINCT user_id
            FROM users_segments 
            WHERE segment_name = segmentName AND deleted_at IS NULL 
            GROUP BY user_id;
    END;
$BODY$
LANGUAGE plpgsql;