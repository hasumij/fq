package postgres11

import (
	"github.com/wader/fq/format/postgres/common"
	"github.com/wader/fq/pkg/decode"
	"github.com/wader/fq/pkg/scalar"
)

// type = struct ControlFileData {
/*    0      |     8 */ // uint64 system_identifier;
/*    8      |     4 */ // uint32 pg_control_version;
/*   12      |     4 */ // uint32 catalog_version_no;
/*   16      |     4 */ // DBState state;
/* XXX  4-byte hole */
/*   24      |     8 */ // pg_time_t time;
/*   32      |     8 */ // XLogRecPtr checkPoint;
/*   40      |    80 */ // CheckPoint checkPointCopy;
/*  120      |     8 */ // XLogRecPtr unloggedLSN;
/*  128      |     8 */ // XLogRecPtr minRecoveryPoint;
/*  136      |     4 */ // TimeLineID minRecoveryPointTLI;
/* XXX  4-byte hole */
/*  144      |     8 */ // XLogRecPtr backupStartPoint;
/*  152      |     8 */ // XLogRecPtr backupEndPoint;
/*  160      |     1 */ // _Bool backupEndRequired;
/* XXX  3-byte hole */
/*  164      |     4 */ // int wal_level;
/*  168      |     1 */ // _Bool wal_log_hints;
/* XXX  3-byte hole */
/*  172      |     4 */ // int MaxConnections;
/*  176      |     4 */ // int max_worker_processes;
/*  180      |     4 */ // int max_prepared_xacts;
/*  184      |     4 */ // int max_locks_per_xact;
/*  188      |     1 */ // _Bool track_commit_timestamp;
/* XXX  3-byte hole */
/*  192      |     4 */ // uint32 maxAlign;
/* XXX  4-byte hole */
/*  200      |     8 */ // double floatFormat;
/*  208      |     4 */ // uint32 blcksz;
/*  212      |     4 */ // uint32 relseg_size;
/*  216      |     4 */ // uint32 xlog_blcksz;
/*  220      |     4 */ // uint32 xlog_seg_size;
/*  224      |     4 */ // uint32 nameDataLen;
/*  228      |     4 */ // uint32 indexMaxKeys;
/*  232      |     4 */ // uint32 toast_max_chunk_size;
/*  236      |     4 */ // uint32 loblksize;
/*  240      |     1 */ // _Bool float4ByVal;
/*  241      |     1 */ // _Bool float8ByVal;
/* XXX  2-byte hole */
/*  244      |     4 */ // uint32 data_checksum_version;
/*  248      |    32 */ // char mock_authentication_nonce[32];
/*  280      |     4 */ // pg_crc32c crc;
/* XXX  4-byte padding */
//
/* total size (bytes):  288 */

// type = struct CheckPoint {
/*    0      |     8 */ // XLogRecPtr redo;
/*    8      |     4 */ // TimeLineID ThisTimeLineID;
/*   12      |     4 */ // TimeLineID PrevTimeLineID;
/*   16      |     1 */ // _Bool fullPageWrites;
/* XXX  3-byte hole */
/*   20      |     4 */ // uint32 nextXidEpoch;
/*   24      |     4 */ // TransactionId nextXid;
/*   28      |     4 */ // Oid nextOid;
/*   32      |     4 */ // MultiXactId nextMulti;
/*   36      |     4 */ // MultiXactOffset nextMultiOffset;
/*   40      |     4 */ // TransactionId oldestXid;
/*   44      |     4 */ // Oid oldestXidDB;
/*   48      |     4 */ // MultiXactId oldestMulti;
/*   52      |     4 */ // Oid oldestMultiDB;
/*   56      |     8 */ // pg_time_t time;
/*   64      |     4 */ // TransactionId oldestCommitTsXid;
/*   68      |     4 */ // TransactionId newestCommitTsXid;
/*   72      |     4 */ // TransactionId oldestActiveXid;
/* XXX  4-byte padding */
//
/* total size (bytes):   80 */
//
func DecodePgControl(d *decode.D, in any) any {
	d.SeekAbs(0)
	/*    0      |     8 */ // uint64 system_identifier;
	/*    8      |     4 */ // uint32 pg_control_version;
	/*   12      |     4 */ // uint32 catalog_version_no;
	/*   16      |     4 */ // DBState state;
	/* XXX  4-byte hole  */
	d.FieldU64("system_identifier")
	d.FieldU32("pg_control_version")
	d.FieldU32("catalog_version_no")
	d.FieldU32("state", common.DBState)
	d.FieldU32("hole0")

	/*   24      |     8 */ // pg_time_t time;
	/*   32      |     8 */ // XLogRecPtr checkPoint;
	/*   40      |    80 */ // CheckPoint checkPointCopy;
	d.FieldS64("time", common.TimeMapper)
	d.FieldU64("checkPoint", common.XLogRecPtrMapper)
	d.FieldStruct("checkPointCopy", func(d *decode.D) {
		/*    0      |     8 */ // XLogRecPtr redo;
		/*    8      |     4 */ // TimeLineID ThisTimeLineID;
		/*   12      |     4 */ // TimeLineID PrevTimeLineID;
		/*   16      |     1 */ // _Bool fullPageWrites;
		/* XXX  3-byte hole */
		d.FieldU64("redo", common.XLogRecPtrMapper)
		d.FieldU32("ThisTimeLineID")
		d.FieldU32("PrevTimeLineID")
		d.FieldU8("fullPageWrites")
		d.FieldU24("hole1")

		/*   20      |     4 */ // uint32 nextXidEpoch;
		/*   24      |     4 */ // TransactionId nextXid;
		/*   28      |     4 */ // Oid nextOid;
		/*   32      |     4 */ // MultiXactId nextMulti;
		/*   36      |     4 */ // MultiXactOffset nextMultiOffset;
		/*   40      |     4 */ // TransactionId oldestXid;
		/*   44      |     4 */ // Oid oldestXidDB;
		/*   48      |     4 */ // MultiXactId oldestMulti;
		/*   52      |     4 */ // Oid oldestMultiDB;
		/*   56      |     8 */ // pg_time_t time;
		/*   64      |     4 */ // TransactionId oldestCommitTsXid;
		/*   68      |     4 */ // TransactionId newestCommitTsXid;
		/*   72      |     4 */ // TransactionId oldestActiveXid;
		/* XXX  4-byte padding */
		d.FieldU32("nextXidEpoch")
		d.FieldU32("nextXid")
		d.FieldU32("nextOid")
		d.FieldU32("nextMulti")
		d.FieldU32("nextMultiOffset")
		d.FieldU32("oldestXid")
		d.FieldU32("oldestXidDB")
		d.FieldU32("oldestMulti")
		d.FieldU32("oldestMultiDB")
		d.FieldS64("time", common.TimeMapper)
		d.FieldU32("oldestCommitTsXid")
		d.FieldU32("newestCommitTsXid")
		d.FieldU32("oldestActiveXid")
		d.FieldU32("padding0")
	})

	/*  120      |     8 */ // XLogRecPtr unloggedLSN;
	/*  128      |     8 */ // XLogRecPtr minRecoveryPoint;
	/*  136      |     4 */ // TimeLineID minRecoveryPointTLI;
	/* XXX  4-byte hole */
	d.FieldU64("unloggedLSN", common.LocPtrMapper)
	d.FieldU64("minRecoveryPoint", common.LocPtrMapper)
	d.FieldU32("minRecoveryPointTLI")
	d.FieldU32("hole2")

	/*  144      |     8 */ // XLogRecPtr backupStartPoint;
	/*  152      |     8 */ // XLogRecPtr backupEndPoint;
	/*  160      |     1 */ // _Bool backupEndRequired;
	/* XXX  3-byte hole */
	d.FieldU64("backupStartPoint", common.LocPtrMapper)
	d.FieldU64("backupEndPoint", common.LocPtrMapper)
	d.FieldU8("backupEndRequired")
	d.FieldU24("hole3")

	/*  164      |     4 */ // int wal_level;
	/*  168      |     1 */ // _Bool wal_log_hints;
	/* XXX  3-byte hole */
	d.FieldS32("wal_level", common.WalLevel)
	d.FieldU8("wal_log_hints")
	d.FieldU24("hole4")

	/*  172      |     4 */ // int MaxConnections;
	/*  176      |     4 */ // int max_worker_processes;
	/*  180      |     4 */ // int max_prepared_xacts;
	/*  184      |     4 */ // int max_locks_per_xact;
	/*  188      |     1 */ // _Bool track_commit_timestamp;
	/* XXX  3-byte hole  */
	d.FieldS32("MaxConnections")
	d.FieldS32("max_worker_processes")
	d.FieldS32("max_prepared_xacts")
	d.FieldS32("max_locks_per_xact")
	d.FieldU8("track_commit_timestamp")
	d.FieldU24("hole5")

	/*  192      |     4 */ // uint32 maxAlign;
	/* XXX  4-byte hole */
	d.FieldU32("maxAlign")
	d.FieldU32("hole6")

	/*  200      |     8 */ // double floatFormat;
	/*  208      |     4 */ // uint32 blcksz;
	/*  212      |     4 */ // uint32 relseg_size;
	/*  216      |     4 */ // uint32 xlog_blcksz;
	/*  220      |     4 */ // uint32 xlog_seg_size;
	/*  224      |     4 */ // uint32 nameDataLen;
	/*  228      |     4 */ // uint32 indexMaxKeys;
	/*  232      |     4 */ // uint32 toast_max_chunk_size;
	/*  236      |     4 */ // uint32 loblksize;
	/*  240      |     1 */ // _Bool float4ByVal;
	/*  241      |     1 */ // _Bool float8ByVal;
	/* XXX  2-byte hole */
	d.FieldF64("floatFormat")
	d.FieldU32("blcksz")
	d.FieldU32("relseg_size")
	d.FieldU32("xlog_blcksz")
	d.FieldU32("xlog_seg_size")
	d.FieldU32("nameDataLen")
	d.FieldU32("indexMaxKeys")
	d.FieldU32("toast_max_chunk_size")
	d.FieldU32("loblksize")
	d.FieldU8("float4ByVal")
	d.FieldU8("float8ByVal")
	d.FieldU16("hole7")

	/*  252      |     4 */ // uint32 data_checksum_version;
	/*  256      |    32 */ // char mock_authentication_nonce[32];
	/*  288      |     4 */ // pg_crc32c crc;
	/* XXX  4-byte padding  */
	d.FieldU32("data_checksum_version")
	d.FieldRawLen("mock_authentication_nonce", 32*8, scalar.RawHex)
	d.FieldU32("crc")
	d.FieldU32("padding1")
	/* total size (bytes):  288 */

	d.AssertPosBytes(288)
	d.FieldRawLen("unused", d.BitsLeft())

	return nil
}
