# RemoveRoleFromDBInstance

Disassociates an AWS Identity and Access Management (IAM) role from a DB instance.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**DBInstanceIdentifier**

The name of the DB instance to disassociate the IAM role from.

Type: String

Required: Yes

**FeatureName**

The name of the feature for the DB instance that the IAM role is to be disassociated from.
For information about supported feature names, see `DBEngineVersion`.

Type: String

Required: Yes

**RoleArn**

The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB instance,
for example, `arn:aws:iam::123456789012:role/AccessRole`.

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DBInstanceNotFound**

`DBInstanceIdentifier` doesn't refer to an existing DB instance.

HTTP Status Code: 404

**DBInstanceRoleNotFound**

The specified `RoleArn` value doesn't match the specified feature for
the DB instance.

HTTP Status Code: 404

**InvalidDBInstanceState**

The DB instance isn't in a valid state.

HTTP Status Code: 400

## Examples

### Example

This example illustrates one usage of RemoveRoleFromDBInstance.

#### Sample Request

```

https://rds.us-east-1.amazonaws.com/
    ?Action=RemoveRoleFromDBInstance
    &DBInstanceIdentifier=sample-instance
    &RoleArn=arn%3Aaws%3Aiam%3A%3A123456789012%3Arole%2Fsample-role
    &FeatureName=s3Import

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for Python](../../../../services/goto/boto3/rds-2014-10-31/removerolefromdbinstance.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/removerolefromdbinstance.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemoveRoleFromDBCluster

RemoveSourceIdentifierFromSubscription

All content copied from https://docs.aws.amazon.com/.
