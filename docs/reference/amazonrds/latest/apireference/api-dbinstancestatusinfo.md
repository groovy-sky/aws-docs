# DBInstanceStatusInfo

Provides a list of status information for a DB instance.

## Contents

###### Note

In the following list, the required parameters are described first.

**Message**

Details of the error if there is an error for the instance. If the instance isn't in an error state, this value is blank.

Type: String

Required: No

**Normal**

Indicates whether the instance is operating normally (TRUE) or is in an error state (FALSE).

Type: Boolean

Required: No

**Status**

The status of the DB instance. For a StatusType of read replica, the values can be
replicating, replication stop point set, replication stop point reached, error, stopped,
or terminated.

Type: String

Required: No

**StatusType**

This value is currently "read replication."

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbinstancestatusinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbinstancestatusinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbinstancestatusinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBInstanceRole

DBMajorEngineVersion

All content copied from https://docs.aws.amazon.com/.
