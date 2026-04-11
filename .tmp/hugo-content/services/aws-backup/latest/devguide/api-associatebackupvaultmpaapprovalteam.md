# AssociateBackupVaultMpaApprovalTeam

Associates an MPA approval team with a backup vault.

## Request Syntax

```nohighlight

PUT /backup-vaults/backupVaultName/mpaApprovalTeam HTTP/1.1
Content-type: application/json

{
   "MpaApprovalTeamArn": "string",
   "RequesterComment": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[backupVaultName](#API_AssociateBackupVaultMpaApprovalTeam_RequestSyntax)**

The name of the backup vault to associate with the MPA approval team.

Pattern: `^[a-zA-Z0-9\-\_]{2,50}$`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[MpaApprovalTeamArn](#API_AssociateBackupVaultMpaApprovalTeam_RequestSyntax)**

The Amazon Resource Name (ARN) of the MPA approval team to associate with the backup vault.

Type: String

Required: Yes

**[RequesterComment](#API_AssociateBackupVaultMpaApprovalTeam_RequestSyntax)**

A comment provided by the requester explaining the association request.

Type: String

Required: No

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/associatebackupvaultmpaapprovalteam.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Backup

CancelLegalHold

All content copied from https://docs.aws.amazon.com/.
