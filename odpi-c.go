package godror

/*
#include <stdlib.h>
#include "dpiImpl.h"
*/
import "C"
import (
	"unsafe"
)

//-----------------------------------------------------------------------------
// Connection Methods (dpiConn)
//-----------------------------------------------------------------------------

// add a reference to a connection
func dpiConn_addRef(conn *C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_addRef(conn); return nil },
		"msg", "dpiConn_addRef",
		"conn", conn,
	)
	return ret
}

// begin a distributed transaction
// DEPRECATED: use dpiConn_tpcBegin() instead
func dpiConn_beginDistribTrans(conn *C.dpiConn, formatId C.long,
	globalTransactionId *C.char, globalTransactionIdLength C.uint32_t,
	branchQualifier *C.char, branchQualifierLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_beginDistribTrans(conn, formatId, globalTransactionId, globalTransactionIdLength, branchQualifier, branchQualifierLength)
		return nil
	},
		"msg", "dpiConn_beginDistribTrans",
		"conn", conn,
		"formatId", formatId,
		"globalTransactionId", globalTransactionId,
		"globalTransactionIdLength", globalTransactionIdLength,
		"branchQualifier", branchQualifier,
		"branchQualifierLength", branchQualifierLength,
	)
	return ret
}

// break execution of the statement running on the connection
func dpiConn_breakExecution(conn *C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_breakExecution(conn); return nil },
		"msg", "dpiConn_breakExecution",
		"conn", conn,
	)
	return ret
}

// change the password for the specified user
func dpiConn_changePassword(conn *C.dpiConn, userName *C.char,
	userNameLength C.uint32_t, oldPassword *C.char,
	oldPasswordLength C.uint32_t, newPassword *C.char,
	newPasswordLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_changePassword(conn, userName, userNameLength, oldPassword, oldPasswordLength, newPassword, newPasswordLength)
		return nil
	},
		"msg", "dpiConn_changePassword",
		"conn", conn,
		"userName", userName,
		"userNameLength", userNameLength,
		"oldPassword", oldPassword,
		"oldPasswordLength", oldPasswordLength,
		"newPassword", newPassword,
		"newPasswordLength", newPasswordLength,
	)
	return ret
}

// close the connection now, not when the reference count reaches zero
func dpiConn_close(conn *C.dpiConn, mode C.dpiConnCloseMode, tag *C.char, tagLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_close(conn, mode, tag, tagLength); return nil },
		"msg", "dpiConn_close",
		"conn", conn,
		"mode", mode,
		"tag", tag,
		"tagLength", tagLength,
	)
	return ret
}

// commits the current active transaction
func dpiConn_commit(conn *C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_commit(conn); return nil },
		"msg", "dpiConn_commit",
		"conn", conn,
	)
	return ret
}

// create a connection and return a reference to it
func dpiConn_create(context *C.dpiContext, userName *C.char,
	userNameLength C.uint32_t, password *C.char, passwordLength C.uint32_t,
	connectString *C.char, connectStringLength C.uint32_t,
	commonParams *C.dpiCommonCreateParams,
	createParams *C.dpiConnCreateParams, conn **C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_create(context, userName, userNameLength, password, passwordLength, connectString, connectStringLength, commonParams, createParams, conn)
		return nil
	},
		"msg", "dpiConn_create",
		"context", context,
		"userName", userName,
		"userNameLength", userNameLength,
		"password", password,
		"passwordLength", passwordLength,
		"connectString", connectString,
		"connectStringLength", connectStringLength,
		"commonParams", commonParams,
		"createParams", createParams,
		"conn", conn,
	)
	return ret
}

// dequeue a message from a queue
func dpiConn_deqObject(conn *C.dpiConn, queueName *C.char,
	queueNameLength C.uint32_t, options *C.dpiDeqOptions, props *C.dpiMsgProps,
	payload *C.dpiObject,
	msgId **C.char, msgIdLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_deqObject(conn, queueName, queueNameLength, options, props, payload, msgId, msgIdLength)
		return nil
	},
		"msg", "dpiConn_deqObject",
		"conn", conn,
		"queueName", queueName,
		"queueNameLength", queueNameLength,
		"options", options,
		"props", props,
		"payload", payload,
		"msgId", msgId,
		"msgIdLength", msgIdLength,
	)
	return ret
}

// enqueue a message to a queue
func dpiConn_enqObject(conn *C.dpiConn, queueName *C.char,
	queueNameLength C.uint32_t, options *C.dpiEnqOptions, props *C.dpiMsgProps,
	payload *C.dpiObject,
	msgId **C.char, msgIdLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_enqObject(conn, queueName, queueNameLength, options, props, payload, msgId, msgIdLength)
		return nil
	},
		"msg", "dpiConn_enqObject",
		"conn", conn,
		"queueName", queueName,
		"queueNameLength", queueNameLength,
		"options", options,
		"props", props,
		"payload", payload,
		"msgId", msgId,
		"msgIdLength", msgIdLength,
	)
	return ret
}

// get call timeout in place for round-trips with this connection
func dpiConn_getCallTimeout(conn *C.dpiConn, value *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getCallTimeout(conn, value); return nil },
		"msg", "dpiConn_getCallTimeout",
		"conn", conn,
		"value", value,
	)
	return ret
}

// get current schema associated with the connection
func dpiConn_getCurrentSchema(conn *C.dpiConn,
	value **C.char,
	valueLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getCurrentSchema(conn, value, valueLength); return nil },
		"msg", "dpiConn_getCurrentSchema",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// get edition associated with the connection
func dpiConn_getEdition(conn *C.dpiConn,

	value **C.char,
	valueLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getEdition(conn, value, valueLength); return nil },
		"msg", "dpiConn_getEdition",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// return the encoding information used by the connection
func dpiConn_getEncodingInfo(conn *C.dpiConn, info *C.dpiEncodingInfo) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getEncodingInfo(conn, info); return nil },
		"msg", "dpiConn_getEncodingInfo",
		"conn", conn,
		"info", info,
	)
	return ret
}

// get external name associated with the connection
func dpiConn_getExternalName(conn *C.dpiConn,
	value **C.char,
	valueLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getExternalName(conn, value, valueLength); return nil },
		"msg", "dpiConn_getExternalName",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// get the OCI service context handle associated with the connection
func dpiConn_getHandle(conn *C.dpiConn, handle *unsafe.Pointer) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getHandle(conn, handle); return nil },
		"msg", "dpiConn_getHandle",
		"conn", conn,
		"handle", handle,
	)
	return ret
}

// get internal name associated with the connection
func dpiConn_getInternalName(conn *C.dpiConn,
	value **C.char,
	valueLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getInternalName(conn, value, valueLength); return nil },
		"msg", "dpiConn_getInternalName",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// get the health of a connection
func dpiConn_getIsHealthy(conn *C.dpiConn, isHealthy *C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getIsHealthy(conn, isHealthy); return nil },
		"msg", "dpiConn_getIsHealthy",
		"conn", conn,
		"isHealthy", isHealthy,
	)
	return ret
}

// get logical transaction id associated with the connection
func dpiConn_getLTXID(conn *C.dpiConn,
	value **C.char,
	valueLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getLTXID(conn, value, valueLength); return nil },
		"msg", "dpiConn_getLTXID",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// create a new object type and return it for subsequent object creation
func dpiConn_getObjectType(conn *C.dpiConn, name *C.char,
	nameLength C.uint32_t, objType **C.dpiObjectType) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getObjectType(conn, name, nameLength, objType); return nil },
		"msg", "dpiConn_getObjectType",
		"conn", conn,
		"name", name,
		"nameLength", nameLength,
		"objType", objType,
	)
	return ret
}

// generic method for getting an OCI connection attribute
// WARNING: use only as directed by Oracle
func dpiConn_getOciAttr(conn *C.dpiConn, handleType C.uint32_t,
	attribute C.uint32_t, value *C.dpiDataBuffer, valueLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getOciAttr(conn, handleType, attribute, value, valueLength); return nil },
		"msg", "dpiConn_getOciAttr",
		"conn", conn,
		"handleType", handleType,
		"attribute", attribute,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// return information about the server version in use
func dpiConn_getServerVersion(conn *C.dpiConn,
	releaseString **C.char, releaseStringLength *C.uint32_t,
	versionInfo *C.dpiVersionInfo) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_getServerVersion(conn, releaseString, releaseStringLength, versionInfo)
		return nil
	},
		"msg", "dpiConn_getServerVersion",
		"conn", conn,
		"releaseString", releaseString,
		"releaseStringLength", releaseStringLength,
		"versionInfo", versionInfo,
	)
	return ret
}

// get SODA interface object
func dpiConn_getSodaDb(conn *C.dpiConn, db **C.dpiSodaDb) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getSodaDb(conn, db); return nil },
		"msg", "dpiConn_getSodaDb",
		"conn", conn,
		"db", db,
	)
	return ret
}

// return the statement cache size
func dpiConn_getStmtCacheSize(conn *C.dpiConn, cacheSize *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_getStmtCacheSize(conn, cacheSize); return nil },
		"msg", "dpiConn_getStmtCacheSize",
		"conn", conn,
		"cacheSize", cacheSize,
	)
	return ret
}

// create a new dequeue options object and return it
func dpiConn_newDeqOptions(conn *C.dpiConn, options **C.dpiDeqOptions) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_newDeqOptions(conn, options); return nil },
		"msg", "dpiConn_newDeqOptions",
		"conn", conn,
		"options", options,
	)
	return ret
}

// create a new enqueue options object and return it
func dpiConn_newEnqOptions(conn *C.dpiConn, options **C.dpiEnqOptions) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_newEnqOptions(conn, options); return nil },
		"msg", "dpiConn_newEnqOptions",
		"conn", conn,
		"options", options,
	)
	return ret
}

// create a new message properties object and return it
func dpiConn_newMsgProps(conn *C.dpiConn, props **C.dpiMsgProps) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_newMsgProps(conn, props); return nil },
		"msg", "dpiConn_newMsgProps",
		"conn", conn,
		"props", props,
	)
	return ret
}

// create a new AQ queue
func dpiConn_newQueue(conn *C.dpiConn, name *C.char,
	nameLength C.uint32_t, payloadType *C.dpiObjectType, queue **C.dpiQueue) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_newQueue(conn, name, nameLength, payloadType, queue); return nil },
		"msg", "dpiConn_newQueue",
		"conn", conn,
		"name", name,
		"nameLength", nameLength,
		"payloadType", payloadType,
		"queue", queue,
	)
	return ret
}

// create a new temporary LOB
func dpiConn_newTempLob(conn *C.dpiConn, lobType C.dpiOracleTypeNum,
	lob **C.dpiLob) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_newTempLob(conn, lobType, lob); return nil },
		"msg", "dpiConn_newTempLob",
		"conn", conn,
		"lobType", lobType,
		"lob", lob,
	)
	return ret
}

// create a new variable and return it for subsequent binding/defining
func dpiConn_newVar(conn *C.dpiConn, oracleTypeNum C.dpiOracleTypeNum,
	nativeTypeNum C.dpiNativeTypeNum, maxArraySize C.uint32_t, size C.uint32_t,
	sizeIsBytes C.int, isArray C.int, objType *C.dpiObjectType, var_ **C.dpiVar, data **C.dpiData) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_newVar(conn, oracleTypeNum, nativeTypeNum, maxArraySize, size, sizeIsBytes, isArray, objType, var_, data)
		return nil
	},
		"msg", "dpiConn_newVar",
		"conn", conn,
		"oracleTypeNum", oracleTypeNum,
		"nativeTypeNum", nativeTypeNum,
		"maxArraySize", maxArraySize,
		"size", size,
		"sizeIsBytes", sizeIsBytes,
		"isArray", isArray,
		"objType", objType,
		"var", var_,
		"data", data,
	)
	return ret
}

// ping the connection to see if it is still alive
func dpiConn_ping(conn *C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_ping(conn); return nil },
		"msg", "dpiConn_ping",
		"conn", conn,
	)
	return ret
}

// prepare a distributed transaction for commit
// DEPRECATED: use dpiConn_tpcPrepare() instead
func dpiConn_prepareDistribTrans(conn *C.dpiConn, commitNeeded *C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_prepareDistribTrans(conn, commitNeeded); return nil },
		"msg", "dpiConn_prepareDistribTrans",
		"conn", conn,
		"commitNeeded", commitNeeded,
	)
	return ret
}

// prepare a statement and return it for subsequent execution/fetching
func dpiConn_prepareStmt(conn *C.dpiConn, scrollable C.int, sql *C.char, sqlLength C.uint32_t, tag *C.char, tagLength C.uint32_t, stmt **C.dpiStmt) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiConn_prepareStmt(conn, scrollable, sql, sqlLength, tag, tagLength, stmt)
		return nil
	},
		"msg", "dpiConn_prepareStmt",
		"conn", conn,
		"scrollable", scrollable,
		"sql", sql,
		"sqlLength", sqlLength,
		"tag", tag,
		"tagLength", tagLength,
		"stmt", stmt,
	)
	return ret
}

// release a reference to the connection
func dpiConn_release(conn *C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_release(conn); return nil },
		"msg", "dpiConn_release",
		"conn", conn,
	)
	return ret
}

// rolls back the current active transaction
func dpiConn_rollback(conn *C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_rollback(conn); return nil },
		"msg", "dpiConn_rollback",
		"conn", conn,
	)
	return ret
}

// set action associated with the connection
func dpiConn_setAction(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setAction(conn, value, valueLength); return nil },
		"msg", "dpiConn_setAction",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set call timeout for subsequent round-trips with this connection
func dpiConn_setCallTimeout(conn *C.dpiConn, value C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setCallTimeout(conn, value); return nil },
		"msg", "dpiConn_setCallTimeout",
		"conn", conn,
		"value", value,
	)
	return ret
}

// set client identifier associated with the connection
func dpiConn_setClientIdentifier(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setClientIdentifier(conn, value, valueLength); return nil },
		"msg", "dpiConn_setClientIdentifier",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set client info associated with the connection
func dpiConn_setClientInfo(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setClientInfo(conn, value, valueLength); return nil },
		"msg", "dpiConn_setClientInfo",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set current schema associated with the connection
func dpiConn_setCurrentSchema(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setCurrentSchema(conn, value, valueLength); return nil },
		"msg", "dpiConn_setCurrentSchema",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set database operation associated with the connection
func dpiConn_setDbOp(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setDbOp(conn, value, valueLength); return nil },
		"msg", "dpiConn_setDbOp",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set execution context id associated with the connection
func dpiConn_setEcontextId(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setEcontextId(conn, value, valueLength); return nil },
		"msg", "dpiConn_setEcontextId",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set external name associated with the connection
func dpiConn_setExternalName(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setExternalName(conn, value, valueLength); return nil },
		"msg", "dpiConn_setExternalName",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set internal name associated with the connection
func dpiConn_setInternalName(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setInternalName(conn, value, valueLength); return nil },
		"msg", "dpiConn_setInternalName",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set module associated with the connection
func dpiConn_setModule(conn *C.dpiConn, value *C.char,
	valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setModule(conn, value, valueLength); return nil },
		"msg", "dpiConn_setModule",
		"conn", conn,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// generic method for setting an OCI connection attribute
// WARNING: use only as directed by Oracle
func dpiConn_setOciAttr(conn *C.dpiConn, handleType C.uint32_t,
	attribute C.uint32_t, value unsafe.Pointer, valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setOciAttr(conn, handleType, attribute, value, valueLength); return nil },
		"msg", "dpiConn_setOciAttr",
		"conn", conn,
		"handleType", handleType,
		"attribute", attribute,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set the statement cache size
func dpiConn_setStmtCacheSize(conn *C.dpiConn, cacheSize C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_setStmtCacheSize(conn, cacheSize); return nil },
		"msg", "dpiConn_setStmtCacheSize",
		"conn", conn,
		"cacheSize", cacheSize,
	)
	return ret
}

// shutdown the database
func dpiConn_shutdownDatabase(conn *C.dpiConn, mode C.dpiShutdownMode) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_shutdownDatabase(conn, mode); return nil },
		"msg", "dpiConn_shutdownDatabase",
		"conn", conn,
		"mode", mode,
	)
	return ret
}

// startup the database
func dpiConn_startupDatabase(conn *C.dpiConn, mode C.dpiStartupMode) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_startupDatabase(conn, mode); return nil },
		"msg", "dpiConn_startupDatabase",
		"conn", conn,
		"mode", mode,
	)
	return ret
}

// startup the database with a PFILE
func dpiConn_startupDatabaseWithPfile(conn *C.dpiConn,
	pfile *C.char, pfileLength C.uint32_t, mode C.dpiStartupMode) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_startupDatabaseWithPfile(conn, pfile, pfileLength, mode); return nil },
		"msg", "dpiConn_startupDatabaseWithPfile",
		"conn", conn,
		"pfile", pfile,
		"pfileLength", pfileLength,
		"mode", mode,
	)
	return ret
}

// subscribe to events in the database
func dpiConn_subscribe(conn *C.dpiConn, params *C.dpiSubscrCreateParams,
	subscr **C.dpiSubscr) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_subscribe(conn, params, subscr); return nil },
		"msg", "dpiConn_subscribe",
		"conn", conn,
		"params", params,
		"subscr", subscr,
	)
	return ret
}

// begin a TPC (two-phase commit) transaction
func dpiConn_tpcBegin(conn *C.dpiConn, xid *C.dpiXid,
	transactionTimeout C.uint32_t, flags C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_tpcBegin(conn, xid, transactionTimeout, flags); return nil },
		"msg", "dpiConn_tpcBegin",
		"conn", conn,
		"xid", xid,
		"transactionTimeout", transactionTimeout,
		"flags", flags,
	)
	return ret
}

// commit a TPC (two-phase commit) transaction
func dpiConn_tpcCommit(conn *C.dpiConn, xid *C.dpiXid, onePhase C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_tpcCommit(conn, xid, onePhase); return nil },
		"msg", "dpiConn_tpcCommit",
		"conn", conn,
		"xid", xid,
		"onePhase", onePhase,
	)
	return ret
}

// end (detach from) a TPC (two-phase commit) transaction
func dpiConn_tpcEnd(conn *C.dpiConn, xid *C.dpiXid, flags C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_tpcEnd(conn, xid, flags); return nil },
		"msg", "dpiConn_tpcEnd",
		"conn", conn,
		"xid", xid,
		"flags", flags,
	)
	return ret
}

// forget a TPC (two-phase commit) transaction
func dpiConn_tpcForget(conn *C.dpiConn, xid *C.dpiXid) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_tpcForget(conn, xid); return nil },
		"msg", "dpiConn_tpcForget",
		"conn", conn,
		"xid", xid,
	)
	return ret
}

// prepare a TPC (two-phase commit) transaction for commit
func dpiConn_tpcPrepare(conn *C.dpiConn, xid *C.dpiXid,
	commitNeeded *C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_tpcPrepare(conn, xid, commitNeeded); return nil },
		"msg", "dpiConn_tpcPrepare",
		"conn", conn,
		"xid", xid,
		"commitNeeded", commitNeeded,
	)
	return ret
}

// rollback a TPC (two-phase commit) transaction
func dpiConn_tpcRollback(conn *C.dpiConn, xid *C.dpiXid) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_tpcRollback(conn, xid); return nil },
		"msg", "dpiConn_tpcRollback",
		"conn", conn,
		"xid", xid,
	)
	return ret
}

// unsubscribe from events in the database
func dpiConn_unsubscribe(conn *C.dpiConn, subscr *C.dpiSubscr) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiConn_unsubscribe(conn, subscr); return nil },
		"msg", "dpiConn_unsubscribe",
		"conn", conn,
		"subscr", subscr,
	)
	return ret
}

//-----------------------------------------------------------------------------
// Session Pools Methods (dpiPool)
//-----------------------------------------------------------------------------

// acquire a connection from the pool and return it
func dpiPool_acquireConnection(pool *C.dpiPool, userName *C.char,
	userNameLength C.uint32_t, password *C.char, passwordLength C.uint32_t,
	createParams *C.dpiConnCreateParams, conn **C.dpiConn) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiPool_acquireConnection(pool, userName, userNameLength, password, passwordLength, createParams, conn)
		return nil
	},
		"msg", "dpiPool_acquireConnection",
		"pool", pool,
		"userName", userName,
		"userNameLength", userNameLength,
		"password", password,
		"passwordLength", passwordLength,
		"createParams", createParams,
		"conn", conn,
	)
	return ret
}

// add a reference to a pool
func dpiPool_addRef(pool *C.dpiPool) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_addRef(pool); return nil },
		"msg", "dpiPool_addRef",
		"pool", pool,
	)
	return ret
}

// destroy the pool now, not when its reference count reaches zero
func dpiPool_close(pool *C.dpiPool, closeMode C.dpiPoolCloseMode) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_close(pool, closeMode); return nil },
		"msg", "dpiPool_close",
		"pool", pool,
		"closeMode", closeMode,
	)
	return ret
}

// create a session pool and return it
func dpiPool_create(context *C.dpiContext, userName *C.char,
	userNameLength C.uint32_t, password *C.char, passwordLength C.uint32_t,
	connectString *C.char, connectStringLength C.uint32_t,
	commonParams *C.dpiCommonCreateParams,
	createParams *C.dpiPoolCreateParams, pool **C.dpiPool) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiPool_create(context, userName, userNameLength, password, passwordLength, connectString, connectStringLength, commonParams, createParams, pool)
		return nil
	},
		"msg", "dpiPool_create",
		"context", context,
		"userName", userName,
		"userNameLength", userNameLength,
		"password", password,
		"passwordLength", passwordLength,
		"connectString", connectString,
		"connectStringLength", connectStringLength,
		"commonParams", commonParams,
		"createParams", createParams,
		"pool", pool,
	)
	return ret
}

// get the pool's busy count
func dpiPool_getBusyCount(pool *C.dpiPool, value *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getBusyCount(pool, value); return nil },
		"msg", "dpiPool_getBusyCount",
		"pool", pool,
		"value", value,
	)
	return ret
}

// return the encoding information used by the session pool
func dpiPool_getEncodingInfo(pool *C.dpiPool, info *C.dpiEncodingInfo) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getEncodingInfo(pool, info); return nil },
		"msg", "dpiPool_getEncodingInfo",
		"pool", pool,
		"info", info,
	)
	return ret
}

// get the pool's "get" mode
func dpiPool_getGetMode(pool *C.dpiPool, value *C.dpiPoolGetMode) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getGetMode(pool, value); return nil },
		"msg", "dpiPool_getGetMode",
		"pool", pool,
		"value", value,
	)
	return ret
}

// get the pool's maximum lifetime session
func dpiPool_getMaxLifetimeSession(pool *C.dpiPool, value *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getMaxLifetimeSession(pool, value); return nil },
		"msg", "dpiPool_getMaxLifetimeSession",
		"pool", pool,
		"value", value,
	)
	return ret
}

// get the pool's maximum sessions per shard
func dpiPool_getMaxSessionsPerShard(pool *C.dpiPool, value *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getMaxSessionsPerShard(pool, value); return nil },
		"msg", "dpiPool_getMaxSessionsPerShard",
		"pool", pool,
		"value", value,
	)
	return ret
}

// get the pool's open count
func dpiPool_getOpenCount(pool *C.dpiPool, value *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getOpenCount(pool, value); return nil },
		"msg", "dpiPool_getOpenCount",
		"pool", pool,
		"value", value,
	)
	return ret
}

// return whether the SODA metadata cache is enabled or not
func dpiPool_getSodaMetadataCache(pool *C.dpiPool, enabled *C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getSodaMetadataCache(pool, enabled); return nil },
		"msg", "dpiPool_getSodaMetadataCache",
		"pool", pool,
		"enabled", enabled,
	)
	return ret
}

// return the statement cache size
func dpiPool_getStmtCacheSize(pool *C.dpiPool, cacheSize *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getStmtCacheSize(pool, cacheSize); return nil },
		"msg", "dpiPool_getStmtCacheSize",
		"pool", pool,
		"cacheSize", cacheSize,
	)
	return ret
}

// get the pool's timeout value
func dpiPool_getTimeout(pool *C.dpiPool, value *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getTimeout(pool, value); return nil },
		"msg", "dpiPool_getTimeout",
		"pool", pool,
		"value", value,
	)
	return ret
}

// get the pool's wait timeout value
func dpiPool_getWaitTimeout(pool *C.dpiPool, value *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getWaitTimeout(pool, value); return nil },
		"msg", "dpiPool_getWaitTimeout",
		"pool", pool,
		"value", value,
	)
	return ret
}

// get the pool-ping-interval
func dpiPool_getPingInterval(pool *C.dpiPool, value *C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_getPingInterval(pool, value); return nil },
		"msg", "dpiPool_getPingInterval",
		"pool", pool,
		"value", value,
	)
	return ret
}

// release a reference to the pool
func dpiPool_release(pool *C.dpiPool) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_release(pool); return nil },
		"msg", "dpiPool_release",
		"pool", pool,
	)
	return ret
}

// set the pool's "get" mode
func dpiPool_setGetMode(pool *C.dpiPool, value C.dpiPoolGetMode) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setGetMode(pool, value); return nil },
		"msg", "dpiPool_setGetMode",
		"pool", pool,
		"value", value,
	)
	return ret
}

// set the pool's maximum lifetime session
func dpiPool_setMaxLifetimeSession(pool *C.dpiPool, value C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setMaxLifetimeSession(pool, value); return nil },
		"msg", "dpiPool_setMaxLifetimeSession",
		"pool", pool,
		"value", value,
	)
	return ret
}

// set the pool's maximum sessions per shard
func dpiPool_setMaxSessionsPerShard(pool *C.dpiPool, value C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setMaxSessionsPerShard(pool, value); return nil },
		"msg", "dpiPool_setMaxSessionsPerShard",
		"pool", pool,
		"value", value,
	)
	return ret
}

// set whether the SODA metadata cache is enabled or not
func dpiPool_setSodaMetadataCache(pool *C.dpiPool, enabled C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setSodaMetadataCache(pool, enabled); return nil },
		"msg", "dpiPool_setSodaMetadataCache",
		"pool", pool,
		"enabled", enabled,
	)
	return ret
}

// set the statement cache size
func dpiPool_setStmtCacheSize(pool *C.dpiPool, cacheSize C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setStmtCacheSize(pool, cacheSize); return nil },
		"msg", "dpiPool_setStmtCacheSize",
		"pool", pool,
		"cacheSize", cacheSize,
	)
	return ret
}

// set the pool's timeout value
func dpiPool_setTimeout(pool *C.dpiPool, value C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setTimeout(pool, value); return nil },
		"msg", "dpiPool_setTimeout",
		"pool", pool,
		"value", value,
	)
	return ret
}

// set the pool's wait timeout value
func dpiPool_setWaitTimeout(pool *C.dpiPool, value C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setWaitTimeout(pool, value); return nil },
		"msg", "dpiPool_setWaitTimeout",
		"pool", pool,
		"value", value,
	)
	return ret
}

// set the pool-ping-interval value
func dpiPool_setPingInterval(pool *C.dpiPool, value C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiPool_setPingInterval(pool, value); return nil },
		"msg", "dpiPool_setPingInterval",
		"pool", pool,
		"value", value,
	)
	return ret
}

//-----------------------------------------------------------------------------
// Statement Methods (dpiStmt)
//-----------------------------------------------------------------------------

// add a reference to a statement
func dpiStmt_addRef(stmt *C.dpiStmt) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_addRef(stmt); return nil },
		"msg", "dpiStmt_addRef",
		"stmt", stmt,
	)
	return ret
}

// bind a variable to the statement using the given name
func dpiStmt_bindByName(stmt *C.dpiStmt, name *C.char,
	nameLength C.uint32_t, var_ *C.dpiVar) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_bindByName(stmt, name, nameLength, var_); return nil },
		"msg", "dpiStmt_bindByName",
		"stmt", stmt,
		"name", name,
		"nameLength", nameLength,
		"var", var_,
	)
	return ret
}

// bind a variable to the statement at the given position
// positions are determined by the order in which names are introduced
func dpiStmt_bindByPos(stmt *C.dpiStmt, pos C.uint32_t, var_ *C.dpiVar) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_bindByPos(stmt, pos, var_); return nil },
		"msg", "dpiStmt_bindByPos",
		"stmt", stmt,
		"pos", pos,
		"var", var_,
	)
	return ret
}

// bind a value to the statement using the given name
// this creates the variable by looking at the type and then binds it
func dpiStmt_bindValueByName(stmt *C.dpiStmt, name *C.char,
	nameLength C.uint32_t, nativeTypeNum C.dpiNativeTypeNum, data *C.dpiData) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_bindValueByName(stmt, name, nameLength, nativeTypeNum, data); return nil },
		"msg", "dpiStmt_bindValueByName",
		"stmt", stmt,
		"name", name,
		"nameLength", nameLength,
		"nativeTypeNum", nativeTypeNum,
		"data", data,
	)
	return ret
}

// bind a value to the statement at the given position
// this creates the variable by looking at the type and then binds it
func dpiStmt_bindValueByPos(stmt *C.dpiStmt, pos C.uint32_t,
	nativeTypeNum C.dpiNativeTypeNum, data *C.dpiData) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_bindValueByPos(stmt, pos, nativeTypeNum, data); return nil },
		"msg", "dpiStmt_bindValueByPos",
		"stmt", stmt,
		"pos", pos,
		"nativeTypeNum", nativeTypeNum,
		"data", data,
	)
	return ret
}

// close the statement now, not when its reference count reaches zero
func dpiStmt_close(stmt *C.dpiStmt, tag *C.char,
	tagLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_close(stmt, tag, tagLength); return nil },
		"msg", "dpiStmt_close",
		"stmt", stmt,
		"tag", tag,
		"tagLength", tagLength,
	)
	return ret
}

// define a variable to accept the data for the specified column (1 based)
func dpiStmt_define(stmt *C.dpiStmt, pos C.uint32_t, var_ *C.dpiVar) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_define(stmt, pos, var_); return nil },
		"msg", "dpiStmt_define",
		"stmt", stmt,
		"pos", pos,
		"var", var_,
	)
	return ret
}

// define type of data to use for the specified column (1 based)
func dpiStmt_defineValue(stmt *C.dpiStmt, pos C.uint32_t,
	oracleTypeNum C.dpiOracleTypeNum, nativeTypeNum C.dpiNativeTypeNum,
	size C.uint32_t, sizeIsBytes C.int, objType *C.dpiObjectType) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiStmt_defineValue(stmt, pos, oracleTypeNum, nativeTypeNum, size, sizeIsBytes, objType)
		return nil
	},
		"msg", "dpiStmt_defineValue",
		"stmt", stmt,
		"pos", pos,
		"oracleTypeNum", oracleTypeNum,
		"nativeTypeNum", nativeTypeNum,
		"size", size,
		"sizeIsBytes", sizeIsBytes,
		"objType", objType,
	)
	return ret
}

// execute the statement and return the number of query columns
// zero implies the statement is not a query
func dpiStmt_execute(stmt *C.dpiStmt, mode C.dpiExecMode,
	numQueryColumns *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_execute(stmt, mode, numQueryColumns); return nil },
		"msg", "dpiStmt_execute",
		"stmt", stmt,
		"mode", mode,
		"numQueryColumns", numQueryColumns,
	)
	return ret
}

// execute the statement multiple times (queries not supported)
func dpiStmt_executeMany(stmt *C.dpiStmt, mode C.dpiExecMode,
	numIters C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_executeMany(stmt, mode, numIters); return nil },
		"msg", "dpiStmt_executeMany",
		"stmt", stmt,
		"mode", mode,
		"numIters", numIters,
	)
	return ret
}

// fetch a single row and return the index into the defined variables
// this will internally perform any execute and array fetch as needed
func dpiStmt_fetch(stmt *C.dpiStmt, found *C.int,
	bufferRowIndex *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_fetch(stmt, found, bufferRowIndex); return nil },
		"msg", "dpiStmt_fetch",
		"stmt", stmt,
		"found", found,
		"bufferRowIndex", bufferRowIndex,
	)
	return ret
}

// return the number of rows that are available in the defined variables
// up to the maximum specified; this will internally perform execute/array
// fetch only if no rows are available in the defined variables and there are
// more rows available to fetch
func dpiStmt_fetchRows(stmt *C.dpiStmt, maxRows C.uint32_t,
	bufferRowIndex *C.uint32_t, numRowsFetched *C.uint32_t, moreRows *C.int) C.int {
	var ret C.int
	_ = ExecInTrace(func() error {
		ret = C.dpiStmt_fetchRows(stmt, maxRows, bufferRowIndex, numRowsFetched, moreRows)
		return nil
	},
		"msg", "dpiStmt_fetchRows",
		"stmt", stmt,
		"maxRows", maxRows,
		"bufferRowIndex", bufferRowIndex,
		"numRowsFetched", numRowsFetched,
		"moreRows", moreRows,
	)
	return ret
}

// get the number of batch errors that took place in the previous execution
func dpiStmt_getBatchErrorCount(stmt *C.dpiStmt, count *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getBatchErrorCount(stmt, count); return nil },
		"msg", "dpiStmt_getBatchErrorCount",
		"stmt", stmt,
		"count", count,
	)
	return ret
}

// get the batch errors that took place in the previous execution
func dpiStmt_getBatchErrors(stmt *C.dpiStmt, numErrors C.uint32_t,
	errors *C.dpiErrorInfo) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getBatchErrors(stmt, numErrors, errors); return nil },
		"msg", "dpiStmt_getBatchErrors",
		"stmt", stmt,
		"numErrors", numErrors,
		"errors", errors,
	)
	return ret
}

// get the number of bind variables that are in the prepared statement
func dpiStmt_getBindCount(stmt *C.dpiStmt, count *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getBindCount(stmt, count); return nil },
		"msg", "dpiStmt_getBindCount",
		"stmt", stmt,
		"count", count,
	)
	return ret
}

// get the names of the bind variables that are in the prepared statement
func dpiStmt_getBindNames(stmt *C.dpiStmt, numBindNames *C.uint32_t,
	bindNames **C.char, bindNameLengths *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getBindNames(stmt, numBindNames, bindNames, bindNameLengths); return nil },
		"msg", "dpiStmt_getBindNames",
		"stmt", stmt,
		"numBindNames", numBindNames,
		"bindNames", bindNames,
		"bindNameLengths", bindNameLengths,
	)
	return ret
}

// get the number of rows to (internally) fetch at one time
func dpiStmt_getFetchArraySize(stmt *C.dpiStmt, arraySize *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getFetchArraySize(stmt, arraySize); return nil },
		"msg", "dpiStmt_getFetchArraySize",
		"stmt", stmt,
		"arraySize", arraySize,
	)
	return ret
}

// get next implicit result from previous execution; NULL if no more exist
func dpiStmt_getImplicitResult(stmt *C.dpiStmt,
	implicitResult **C.dpiStmt) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getImplicitResult(stmt, implicitResult); return nil },
		"msg", "dpiStmt_getImplicitResult",
		"stmt", stmt,
		"implicitResult", implicitResult,
	)
	return ret
}

// return information about the statement
func dpiStmt_getInfo(stmt *C.dpiStmt, info *C.dpiStmtInfo) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getInfo(stmt, info); return nil },
		"msg", "dpiStmt_getInfo",
		"stmt", stmt,
		"info", info,
	)
	return ret
}

// get the rowid of the last row affected by a DML statement
func dpiStmt_getLastRowid(stmt *C.dpiStmt, rowid **C.dpiRowid) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getLastRowid(stmt, rowid); return nil },
		"msg", "dpiStmt_getLastRowid",
		"stmt", stmt,
		"rowid", rowid,
	)
	return ret
}

// get the number of query columns (zero implies the statement is not a query)
func dpiStmt_getNumQueryColumns(stmt *C.dpiStmt,
	numQueryColumns *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getNumQueryColumns(stmt, numQueryColumns); return nil },
		"msg", "dpiStmt_getNumQueryColumns",
		"stmt", stmt,
		"numQueryColumns", numQueryColumns,
	)
	return ret
}

// generic method for getting an OCI statement attribute
// WARNING: use only as directed by Oracle
func dpiStmt_getOciAttr(stmt *C.dpiStmt, attribute C.uint32_t,
	value *C.dpiDataBuffer, valueLength *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getOciAttr(stmt, attribute, value, valueLength); return nil },
		"msg", "dpiStmt_getOciAttr",
		"stmt", stmt,
		"attribute", attribute,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// return the number of rows that are prefetched by the Oracle Client library
func dpiStmt_getPrefetchRows(stmt *C.dpiStmt, numRows *C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getPrefetchRows(stmt, numRows); return nil },
		"msg", "dpiStmt_getPrefetchRows",
		"stmt", stmt,
		"numRows", numRows,
	)
	return ret
}

// return metadata about the column at the specified position (1 based)
func dpiStmt_getQueryInfo(stmt *C.dpiStmt, pos C.uint32_t,
	info *C.dpiQueryInfo) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getQueryInfo(stmt, pos, info); return nil },
		"msg", "dpiStmt_getQueryInfo",
		"stmt", stmt,
		"pos", pos,
		"info", info,
	)
	return ret
}

// get the value for the specified column of the current row fetched
func dpiStmt_getQueryValue(stmt *C.dpiStmt, pos C.uint32_t,
	nativeTypeNum *C.dpiNativeTypeNum, data **C.dpiData) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getQueryValue(stmt, pos, nativeTypeNum, data); return nil },
		"msg", "dpiStmt_getQueryValue",
		"stmt", stmt,
		"pos", pos,
		"nativeTypeNum", nativeTypeNum,
		"data", data,
	)
	return ret
}

// get the row count for the statement
// for queries, this is the number of rows that have been fetched so far
// for non-queries, this is the number of rows affected by the last execution
func dpiStmt_getRowCount(stmt *C.dpiStmt, count *C.uint64_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getRowCount(stmt, count); return nil },
		"msg", "dpiStmt_getRowCount",
		"stmt", stmt,
		"count", count,
	)
	return ret
}

// get the number of rows affected for each DML operation just executed
// using the mode DPI_MODE_EXEC_ARRAY_DML_ROWCOUNTS
func dpiStmt_getRowCounts(stmt *C.dpiStmt, numRowCounts *C.uint32_t,
	rowCounts **C.uint64_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getRowCounts(stmt, numRowCounts, rowCounts); return nil },
		"msg", "dpiStmt_getRowCounts",
		"stmt", stmt,
		"numRowCounts", numRowCounts,
		"rowCounts", rowCounts,
	)
	return ret
}

// get subscription query id for continuous query notification
func dpiStmt_getSubscrQueryId(stmt *C.dpiStmt, queryId *C.uint64_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_getSubscrQueryId(stmt, queryId); return nil },
		"msg", "dpiStmt_getSubscrQueryId",
		"stmt", stmt,
		"queryId", queryId,
	)
	return ret
}

// release a reference to the statement
func dpiStmt_release(stmt *C.dpiStmt) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_release(stmt); return nil },
		"msg", "dpiStmt_release",
		"stmt", stmt,
	)
	return ret
}

// scroll the statement to the desired row
// this is only valid for scrollable statements
func dpiStmt_scroll(stmt *C.dpiStmt, mode C.dpiFetchMode, offset C.int32_t,
	rowCountOffset C.int32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_scroll(stmt, mode, offset, rowCountOffset); return nil },
		"msg", "dpiStmt_scroll",
		"stmt", stmt,
		"mode", mode,
		"offset", offset,
		"rowCountOffset", rowCountOffset,
	)
	return ret
}

// set the number of rows to (internally) fetch at one time
func dpiStmt_setFetchArraySize(stmt *C.dpiStmt, arraySize C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_setFetchArraySize(stmt, arraySize); return nil },
		"msg", "dpiStmt_setFetchArraySize",
		"stmt", stmt,
		"arraySize", arraySize,
	)
	return ret
}

// generic method for setting an OCI statement attribute
// WARNING: use only as directed by Oracle
func dpiStmt_setOciAttr(stmt *C.dpiStmt, attribute C.uint32_t,
	value unsafe.Pointer, valueLength C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_setOciAttr(stmt, attribute, value, valueLength); return nil },
		"msg", "dpiStmt_setOciAttr",
		"stmt", stmt,
		"attribute", attribute,
		"value", value,
		"valueLength", valueLength,
	)
	return ret
}

// set the number of rows that are prefetched by the Oracle Client library
func dpiStmt_setPrefetchRows(stmt *C.dpiStmt,
	numRows C.uint32_t) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_setPrefetchRows(stmt, numRows); return nil },
		"msg", "dpiStmt_setPrefetchRows",
		"stmt", stmt,
		"numRows", numRows,
	)
	return ret
}

// set the flag to exclude the current SQL statement from the statement
// cache
func dpiStmt_deleteFromCache(stmt *C.dpiStmt) C.int {
	var ret C.int
	_ = ExecInTrace(func() error { ret = C.dpiStmt_deleteFromCache(stmt); return nil },
		"msg", "dpiStmt_deleteFromCache",
		"stmt", stmt,
	)
	return ret
}
