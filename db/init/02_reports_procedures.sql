\c usersegmentappdb

CREATE OR REPLACE FUNCTION get_report_about_users_segments(time_date date) RETURNS TABLE (
		id integer,
		segment text,
        action text,
        date_time date
	) 
AS
$BODY$
    BEGIN
        RETURN QUERY
            WITH added_segs AS (SELECT user_id, segment_name, added_at FROM users_segments WHERE added_at >= time_date), 
                deleted_segs AS (SELECT user_id, segment_name, deleted_at FROM users_segments WHERE deleted_at >= time_date) 
                
            SELECT user_id, segment_name, 'add' AS action, added_at AS date_time 
                FROM added_segs 
            UNION 
            SELECT user_id, segment_name, 'delete' AS action, deleted_at AS date_time
                FROM deleted_segs
            ORDER BY user_id, action;
    END;
$BODY$
LANGUAGE plpgsql;