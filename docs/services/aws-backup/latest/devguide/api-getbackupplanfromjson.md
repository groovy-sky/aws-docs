# GetBackupPlanFromJSON

Returns a valid JSON document specifying a backup plan or an error.

## Request Syntax

```nohighlight

POST /backup/template/json/toPlan HTTP/1.1
Content-type: application/json

{
   "BackupPlanTemplateJson": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[BackupPlanTemplateJson](#API_GetBackupPlanFromJSON_RequestSyntax)**

A customer-supplied backup plan document in JSON format.

Type: String

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "BackupPlan": {
      "AdvancedBackupSettings": [
         {
            "BackupOptions": {
               "string" : "string"
            },
            "ResourceType": "string"
         }
      ],
      "BackupPlanName": "string",
      "Rules": [
         {
            "CompletionWindowMinutes": number,
            "CopyActions": [
               {
                  "DestinationBackupVaultArn": "string",
                  "Lifecycle": {
                     "DeleteAfterDays": number,
                     "DeleteAfterEvent": "string",
                     "MoveToColdStorageAfterDays": number,
                     "OptInToArchiveForSupportedResources": boolean
                  }
               }
            ],
            "EnableContinuousBackup": boolean,
            "IndexActions": [
               {
                  "ResourceTypes": [ "string" ]
               }
            ],
            "Lifecycle": {
               "DeleteAfterDays": number,
               "DeleteAfterEvent": "string",
               "MoveToColdStorageAfterDays": number,
               "OptInToArchiveForSupportedResources": boolean
            },
            "RecoveryPointTags": {
               "string" : "string"
            },
            "RuleId": "string",
            "RuleName": "string",
            "ScanActions": [
               {
                  "MalwareScanner": "string",
                  "ScanMode": "string"
               }
            ],
            "ScheduleExpression": "string",
            "ScheduleExpressionTimezone": "string",
            "StartWindowMinutes": number,
            "TargetBackupVaultName": "string",
            "TargetLogicallyAirGappedBackupVaultArn": "string"
         }
      ],
      "ScanSettings": [
         {
            "MalwareScanner": "string",
            "ResourceTypes": [ "string" ],
            "ScannerRoleArn": "string"
         }
      ]
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[BackupPlan](#API_GetBackupPlanFromJSON_ResponseSyntax)**

Specifies the body of a backup plan. Includes a `BackupPlanName` and one or
more sets of `Rules`.

Type: [BackupPlan](api-backupplan.md) object

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

**LimitExceededException**

A limit in the request has been exceeded; for example, a maximum number of items allowed
in a request.

**Context**

**Type**

HTTP Status Code: 400

**MissingParameterValueException**

Indicates that a required parameter is missing.

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

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/getbackupplanfromjson.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/getbackupplanfromjson.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBackupPlan

GetBackupPlanFromTemplate

All content copied from https://docs.aws.amazon.com/.
