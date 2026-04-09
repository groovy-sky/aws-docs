# DeleteRecoveryPoint

Deletes the recovery point specified by a recovery point ID.

If the recovery point ID belongs to a continuous backup, calling this endpoint deletes
the existing continuous backup and stops future continuous backup.

When an IAM role's permissions are insufficient to call this API, the service sends back
an HTTP 200 response with an empty HTTP body, but the recovery point is not deleted.
Instead, it enters an `EXPIRED` state.

`EXPIRED` recovery points can be deleted with this API once the IAM role
has the `iam:CreateServiceLinkedRole` action. To learn more about adding this role, see
[Troubleshooting manual deletions](deleting-backups.md#deleting-backups-troubleshooting).

If the user or role is deleted or the permission within the role is removed,
the deletion will not be successful and will enter an `EXPIRED` state.

## Request Syntax

```nohighlight

DELETE /backup-vaults/backupVaultName/recovery-points/recoveryPointArn HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[backupVaultName](#API_DeleteRecoveryPoint_RequestSyntax)**

The name of a logical container where backups are stored. Backup vaults are identified
by names that are unique to the account used to create them and the AWS
Region where they are created.

Pattern: `^[a-zA-Z0-9\-\_]{2,50}$`

Required: Yes

**[recoveryPointArn](#API_DeleteRecoveryPoint_RequestSyntax)**

An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for example,
`arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45`.

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValueException**

Indicates that something is wrong with a parameter's value. For example, the value is
out of range.

**Context**

**Type**

HTTP Status Code: 400

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
parameter is of the wrong type.

**Context**

**Type**

HTTP Status Code: 400

**InvalidResourceStateException**

AWS Backup is already performing an action on this recovery point. It can't
perform the action you requested until the first action finishes. Try again later.

**Context**

**Type**

HTTP Status Code: 400

**MissingParameterValueException**

Indicates that a required parameter is missing.

**Context**

**Type**

HTTP Status Code: 400

**ResourceNotFoundException**

A resource that is required for the action doesn't exist.

**Context**

**Type**

HTTP Status Code: 400

**ServiceUnavailableException**

The request failed due to a temporary failure of the server.

**Context**

**Type**

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/deleterecoverypoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/deleterecoverypoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteFramework

DeleteReportPlan

All content copied from https://docs.aws.amazon.com/.
