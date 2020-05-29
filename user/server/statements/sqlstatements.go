package statements

type usrSql string

// Create update delete statements
const (
	SqlInsert usrSql = `insert into appuser (firstname , lastname, validfrom, validthru, active, pwd) 
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING id, firstname , lastname, validfrom, validthru, active, pwd, name`

	SqlUpdate usrSql = ` update appuser set 
						 firstname = $1,
						 lastname = $2,
						 validfrom = $3,
						 validthru = $4,
						 active = $5,
						 pwd = $6
						where id = $7
						RETURNING id, firstname , lastname, validfrom, validthru, active, pwd, name`

	SqlDelete usrSql = "delete from appuser where id = $1"
)

// Select statements
const (
	SqlSelectById usrSql = `select 
								id, firstname , lastname, validfrom, validthru, active, pwd, name 
							from appuser 
							where id = $1`

	SqlSelectAll usrSql = ` select 
								id, firstname , lastname, validfrom, validthru, active, pwd, name 
							from appuser`
)

//test statement
const (
	TestStatement usrSql = `select 1`
)

func (p usrSql) String() string {
	return string(p)
}
