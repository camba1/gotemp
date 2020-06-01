package statements

type prmSql string

var MaxRowsToFetch = 200

// Create update delete statements
const (
	SqlInsert prmSql = `insert into promotion (name, description, validfrom, validthru, customerid, active, approvalstatus, prevapprovalstatus) 
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
						RETURNING id, name, description, validfrom, validthru, customerid, active, approvalstatus,  prevapprovalstatus `

	SqlUpdate prmSql = ` update promotion set 
						 name = $1,
						 description = $2,
						 validfrom = $3,
						 validthru = $4,
						 customerid = $5,
						 active = $6,
						 approvalstatus = $7,
						 prevapprovalstatus = $8
						where id = $9
						RETURNING id, name, description, validfrom, validthru, customerid, active, approvalstatus,  prevapprovalstatus`

	SqlDelete prmSql = "delete from promotion where id = $1"
)

// Select statements
const (
	SqlSelectById prmSql = `select 
								id, name, description, validfrom, validthru, 
								customerid, active, approvalstatus,  prevapprovalstatus 
							from promotion 
							where id = $1`

	SqlSelectAll prmSql = ` select 
								id, name, description, validfrom, validthru, 
								customerid, active, approvalstatus,  prevapprovalstatus 
							from promotion %s
							FETCH FIRST %d ROWS ONLY`
)

func (p prmSql) String() string {
	return string(p)
}
