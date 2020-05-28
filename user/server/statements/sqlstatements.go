package statements

type usrSql string

// Create update delete statements
const (
	SqlInsert usrSql = `insert into user (firstname , lastname, validfrom, validthru, active, pwd) 
						VALUES ($1, $2, $3, $4, $5, $6)
						RETURNING id, firstname , lastname, validfrom, validthru, active, pwd, name`

	SqlUpdate usrSql = ` update user set 
						 firstname = $1,
						 lastname = $2,
						 validfrom = $3,
						 validthru = $4,
						 active = $5,
						 pwd = $6
						where id = $7
						RETURNING id, firstname , lastname, validfrom, validthru, active, pwd, name`

	SqlDelete usrSql = "delete from user where id = $1"
)

// Select statements
const (
	SqlSelectById usrSql = `select 
								id, firstname , lastname, validfrom, validthru, active, pwd, name 
							from user 
							where id = $1`

	SqlSelectAll usrSql = ` select 
								id, firstname , lastname, validfrom, validthru, active, pwd, name 
							from user`
)

//test statement
const (
	TestStatement usrSql = `select 1`
)

func (p usrSql) String() string {
	return string(p)
}
