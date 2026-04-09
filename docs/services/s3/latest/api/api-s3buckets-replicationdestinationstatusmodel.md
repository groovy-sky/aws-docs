# ReplicationDestinationStatusModel

Contains status information for a replication destination, including the current replication state, last successful update, and any error messages.

## Contents

**destinationTableBucketArn**

The Amazon Resource Name (ARN) of the destination table bucket.

Type: String

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

**replicationStatus**

The current status of replication to this destination.

Type: String

Valid Values: `pending | completed | failed`

Required: Yes

**destinationTableArn**

The Amazon Resource Name (ARN) of the destination table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63}/table/[a-zA-Z0-9-_]{1,255})`

Required: No

**failureMessage**

If replication has failed, this field contains an error message describing the failure reason.

Type: String

Required: No

**lastSuccessfulReplicatedUpdate**

Information about the most recent successful replication update to this destination.

Type: [LastSuccessfulReplicatedUpdate](api-s3buckets-lastsuccessfulreplicatedupdate.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/replicationdestinationstatusmodel.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/replicationdestinationstatusmodel.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/replicationdestinationstatusmodel.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationDestination

ReplicationInformation

All content copied from https://docs.aws.amazon.com/.
