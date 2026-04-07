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

Type: [LastSuccessfulReplicatedUpdate](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_LastSuccessfulReplicatedUpdate.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/ReplicationDestinationStatusModel)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/ReplicationDestinationStatusModel)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/ReplicationDestinationStatusModel)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReplicationDestination

ReplicationInformation
