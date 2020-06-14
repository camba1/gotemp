package statements

//auditSql: String based type that relates service specific sql statements
type auditSql string

//MaxRowsToFetch: Maximum number of rows to return from the database in one select
var MaxRowsToFetch = 200

// Create, update, delete statements
const (
	SqlInsert auditSql = `INSERT into audit (actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, objectDetail)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
						RETURNING actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, recordedtime`

	SqlDelete auditSql = "DELETE FROM audit WHERE objectName = $1 AND objectId = $2 AND actiontime <= $3 and actiontime >= $4"
)

// Select statements
const (
	SqlSelectById auditSql = `SELECT
								actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, recordedtime, id, objectDetail
							FROM audit
							WHERE id = $1`

	//SqlSelectById auditSql = `SELECT
	//							actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, recordedtime, id
	//						FROM audit
	//						WHERE id = $1`

	SqlSelectAll auditSql = ` SELECT 
								actiontime, topic, service, actionFunc, actionType, objectName, objectId, performedBy, recordedtime, id
							FROM audit %s
							FETCH FIRST %d ROWS ONLY`
)

//test statement
const (
	TestStatement auditSql = `select 1`
)

//String: Return sql statement
func (p auditSql) String() string {
	return string(p)
}
