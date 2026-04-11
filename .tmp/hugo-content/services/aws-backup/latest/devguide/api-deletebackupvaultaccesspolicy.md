# DeleteBackupVaultAccessPolicy

Deletes the policy document that manages permissions on a backup vault.

## Request Syntax

```nohighlight

DELETE /backup-vaults/backupVaultName/access-policy HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[backupVaultName](#API_DeleteBackupVaultAccessPolicy_RequestSyntax)**

The name of a logical container where backups are stored. Backup vaults are identified
by names that are unique to the account used to create them and the AWS
Region where they are created. They consist of lowercase letters, numbers, and
hyphens.

Pattern: `^[a-zA-Z0-9\-\_]{2,50}$`

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

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/deletebackupvaultaccesspolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBackupVault

DeleteBackupVaultLockConfiguration

All content copied from https://docs.aws.amazon.com/.
