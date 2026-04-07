# StoreImageTaskResult

The information about the AMI store task, including the progress of the task.

## Contents

**amiId**

The ID of the AMI that is being stored.

Type: String

Required: No

**bucket**

The name of the Amazon S3 bucket that contains the stored AMI object.

Type: String

Required: No

**progressPercentage**

The progress of the task as a percentage.

Type: Integer

Required: No

**s3objectKey**

The name of the stored AMI object in the bucket.

Type: String

Required: No

**storeTaskFailureReason**

If the tasks fails, the reason for the failure is returned. If the task succeeds,
`null` is returned.

Type: String

Required: No

**storeTaskState**

The state of the store task ( `InProgress`, `Completed`, or
`Failed`).

Type: String

Required: No

**taskStartTime**

The time the task started.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/storeimagetaskresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/storeimagetaskresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/storeimagetaskresult.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StorageLocation

Subnet
