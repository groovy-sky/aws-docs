# ModifyActivityStream

Changes the audit policy state of a database activity stream to either locked (default) or unlocked. A locked policy is read-only,
whereas an unlocked policy is read/write. If your activity stream is started and locked, you can unlock it, customize your audit policy,
and then lock your activity stream. Restarting the activity stream isn't required. For more information, see [Modifying a database activity stream](../../../../services/amazonrds/latest/userguide/dbactivitystreams-modifying.md) in the
_Amazon RDS User Guide_.

This operation is supported for RDS for Oracle and Microsoft SQL Server.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**AuditPolicyState**

The audit policy state. When a policy is unlocked, it is read/write. When it is locked, it is
read-only. You can edit your audit policy only when the activity stream is unlocked or stopped.

Type: String

Valid Values: `locked | unlocked`

Required: No

**ResourceArn**

The Amazon Resource Name (ARN) of the RDS for Oracle or Microsoft SQL Server DB instance.
For example, `arn:aws:rds:us-east-1:12345667890:db:my-orcl-db`.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**EngineNativeAuditFieldsIncluded**

Indicates whether engine-native audit fields are included in the database activity stream.

Type: Boolean

**KinesisStreamName**

The name of the Amazon Kinesis data stream to be used for the database activity stream.

Type: String

**KmsKeyId**

The AWS KMS key identifier for encryption of messages in the database activity stream.

Type: String

**Mode**

The mode of the database activity stream.

Type: String

Valid Values: `sync | async`

**PolicyStatus**

The status of the modification to the policy state of the database activity stream.

Type: String

Valid Values: `locked | unlocked | locking-policy | unlocking-policy`

**Status**

The status of the modification to the database activity stream.

Type: String

Valid Values: `stopped | starting | started | stopping`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

**ResourceNotFoundFault**

The specified resource ID was not found.

HTTP Status Code: 404

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/modifyactivitystream.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/modifyactivitystream.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListTagsForResource

ModifyCertificates

All content copied from https://docs.aws.amazon.com/.
