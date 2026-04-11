# DeleteSnapshotReturnCode

The snapshot ID and its deletion result code.

## Contents

**returnCode**

The result code from the snapshot deletion attempt. Possible values:

- `success` \- The snapshot was successfully deleted.

- `skipped` \- The snapshot was not deleted because it's associated with other
AMIs.

- `missing-permissions` \- The snapshot was not deleted because the role lacks
`DeleteSnapshot` permissions. For more information, see [How\
Amazon EBS works with IAM](../../../../services/ebs/latest/userguide/security-iam-service-with-iam.md).

- `internal-error` \- The snapshot was not deleted due to a server
error.

- `client-error` \- The snapshot was not deleted due to a client configuration
error.

For details about an error, check the `DeleteSnapshot` event in the CloudTrail
event history. For more information, see [View event history](../../../../services/awscloudtrail/latest/userguide/tutorial-event-history.md)
in the _AWS CloudTrail User Guide_.

Type: String

Valid Values: `success | skipped | missing-permissions | internal-error | client-error`

Required: No

**snapshotId**

The ID of the snapshot.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletesnapshotreturncode.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletesnapshotreturncode.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletesnapshotreturncode.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteQueuedReservedInstancesError

DeprecationTimeCondition
