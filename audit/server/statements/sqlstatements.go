package statements

type auditSql string

var MaxRowsToFetch = 200

// Create update delete statements
const (
	SqlInsert auditSql = `INSERT into audit (actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, objectDetail) 
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						RETURNING actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, recordedtime`

	SqlDelete auditSql = "DELETE FROM audit WHERE objectName = $1 AND objectId = $2 AND actiontime <= $3 and actiontime >= $4"
)

// Select statements
const (
	SqlSelectAll auditSql = ` SELECT 
								actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, recordedtime 
							FROM audit %s
							FETCH FIRST %d ROWS ONLY`
)

//test statement
const (
	TestStatement auditSql = `select 1`
)

func (p auditSql) String() string {
	return string(p)
}
