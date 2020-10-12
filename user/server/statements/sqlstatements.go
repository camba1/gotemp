package statements

type usrSql string

//MaxRowsToFetch is the maximum number of row that a get all sql statement should return
var MaxRowsToFetch = 200

// Create update delete statements
const (
	SqlInsert usrSql = `INSERT into appuser (firstname , lastname, validfrom, validthru, active, pwd, email, company) 
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
						RETURNING id, firstname , lastname, validfrom, validthru, active, pwd, name, email, company, createdate, updatedate`

	SqlUpdate usrSql = ` UPDATE appuser SET 
						 firstname = $1,
						 lastname = $2,
						 validfrom = $3,
						 validthru = $4,
						 active = $5,
						 pwd = $6,
						 email = $7,
						 company = $8,
						 updatedate = $9
						WHERE id = $10
						RETURNING id, firstname , lastname, validfrom, validthru, active, pwd, name, email, company, createdate, updatedate`

	SqlDelete usrSql = "DELETE FROM appuser WHERE id = $1"
)

// Select statements
const (
	SqlSelectById usrSql = `SELECT 
								id, firstname , lastname, validfrom, validthru, active, pwd, name, email, company, createdate, updatedate 
							FROM appuser 
							WHERE id = $1`

	SqlSelectAll usrSql = ` SELECT 
								id, firstname , lastname, validfrom, validthru, active, pwd, name, email, company, createdate, updatedate 
							FROM appuser %s
							FETCH FIRST %d ROWS ONLY`
)

//test statement
const (
	TestStatement usrSql = `select 1`
)

func (p usrSql) String() string {
	return string(p)
}
