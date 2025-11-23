package storage

var (
	insertUser               = "INSERT INTO users (id, username, is_active) VALUES ($1,$2, $3);"
	insertTeam               = "INSERT INTO teams (name, user_id) VALUES($1, $2);"
	selectTwoFreeActiveUsers = "SELECT id FROM users WHERE is_active LIMIT 2;"
	createPr                 = "INSERT INTO pr (id, name, author, opened) VALUES($1, $2);"
	addReviewers             = "INSERT INTO review (id, reviewer) VALUES ($1,$2);"
	selectUsersNotFromTeam   = `SELECT u.id 
    							FROM users u 
    							WHERE NOT EXISTS (
        							SELECT 1 FROM teams t 
        							WHERE t.user_id = u.id AND t.name = $1
    							) AND u.is_active = true;`
	selectReviwerData = `SELECT r.id, r.reviewer 
									FROM review r
									JOIN teams t ON r.reviewer = t.user_id
									WHERE t.name = $1;`
	dropUserFromReview = "DELETE FROM review WHERE id = $1 AND reviewer = $2;"
	selectStat         = `	SELECT  u.id AS user_id,  COUNT(r.id) AS review_count
							FROM  users u
							LEFT JOIN review r ON u.id = r.reviewer
							LEFT JOIN  pr ON r.id = pr.id AND pr.opened = true
							WHERE u.id > $1  
							GROUP BY u.id
							ORDER BY u.id  LIMIT $2;`

	mergePr = `UPDATE pr 
				SET opened = false 
				WHERE id = $1;`
	setUserActive = `	UPDATE users 
						SET is_active = $1 
						WHERE id = $2;`
	getAllUsersQuery = `
        SELECT id, username, is_active 
        FROM users`

	getAllTeamsQuery = `
        SELECT name, user_id 
        FROM teams`

	getAllPrsQuery = `
        SELECT id, name, author, opened 
        FROM pr`

	getAllReviewsQuery = `
        SELECT id, reviewer 
        FROM review`
)
