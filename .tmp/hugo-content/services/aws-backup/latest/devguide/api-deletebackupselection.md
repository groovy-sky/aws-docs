# DeleteBackupSelection

Deletes the resource selection associated with a backup plan that is specified by the
`SelectionId`.

## Request Syntax

```nohighlight

DELETE /backup/plans/backupPlanId/selections/selectionId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[backupPlanId](#API_DeleteBackupSelection_RequestSyntax)**

Uniquely identifies a backup plan.

Required: Yes

**[selectionId](#API_DeleteBackupSelection_RequestSyntax)**

Uniquely identifies the body of a request to assign a set of resources to a backup
plan.

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

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/deletebackupselection.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/deletebackupselection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteBackupPlan

DeleteBackupVault

All content copied from https://docs.aws.amazon.com/.
