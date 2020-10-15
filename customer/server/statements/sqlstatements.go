package statements

type custSql string

// MaxRowsToFetch is the maximum number of row that a get all sql statement should return
var MaxRowsToFetch = 200

// Create update delete statements
//const (
//SqlInsert custSql = `insert into promotion (name, description, validfrom, validthru, customerid, active, approvalstatus, prevapprovalstatus)
//					VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
//					RETURNING id, name, description, validfrom, validthru, customerid, active, approvalstatus,  prevapprovalstatus `

//SqlUpdate custSql = ` update promotion set
//					 name = $1,
//					 description = $2,
//					 validfrom = $3,
//					 validthru = $4,
//					 customerid = $5,
//					 active = $6,
//					 approvalstatus = $7,
//					 prevapprovalstatus = $8
//					where id = $9
//					RETURNING id, name, description, validfrom, validthru, customerid, active, approvalstatus,  prevapprovalstatus`

//SqlDelete custSql = "delete from promotion where id = $1"
//)

// Select statements
const (
	//SqlSelectById custSql = `select
	//							id, name, description, validfrom, validthru,
	//							customerid, active, approvalstatus,  prevapprovalstatus
	//						from promotion
	//						where id = $1`

	//SqlSelectAll custSql = ` select
	//							id, name, description, validfrom, validthru,
	//							customerid, active, approvalstatus,  prevapprovalstatus
	//						from promotion %s
	//						FETCH FIRST %d ROWS ONLY`

	SqlSelectAll custSql = `FOR c IN customer %s
								LIMIT %d
								RETURN c`
)

func (p custSql) String() string {
	return string(p)
}
