# UpdateRestoreTestingSelection

Updates the specified restore testing selection.

Most elements except the `RestoreTestingSelectionName`
can be updated with this request.

You can use either protected resource ARNs or conditions, but not both.

## Request Syntax

```nohighlight

PUT /restore-testing/plans/RestoreTestingPlanName/selections/RestoreTestingSelectionName HTTP/1.1
Content-type: application/json

{
   "RestoreTestingSelection": {
      "IamRoleArn": "string",
      "ProtectedResourceArns": [ "string" ],
      "ProtectedResourceConditions": {
         "StringEquals": [
            {
               "Key": "string",
               "Value": "string"
            }
         ],
         "StringNotEquals": [
            {
               "Key": "string",
               "Value": "string"
            }
         ]
      },
      "RestoreMetadataOverrides": {
         "string" : "string"
      },
      "ValidationWindowHours": number
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[RestoreTestingPlanName](#API_UpdateRestoreTestingSelection_RequestSyntax)**

The restore testing plan name is required to update the
indicated testing plan.

Required: Yes

**[RestoreTestingSelectionName](#API_UpdateRestoreTestingSelection_RequestSyntax)**

The required restore testing selection name of the restore
testing selection you wish to update.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[RestoreTestingSelection](#API_UpdateRestoreTestingSelection_RequestSyntax)**

To update your restore testing selection, you can use either
protected resource ARNs or conditions, but not both. That is, if your
selection has `ProtectedResourceArns`, requesting an update
with the parameter `ProtectedResourceConditions` will be
unsuccessful.

Type: [RestoreTestingSelectionForUpdate](api-restoretestingselectionforupdate.md) object

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "CreationTime": number,
   "RestoreTestingPlanArn": "string",
   "RestoreTestingPlanName": "string",
   "RestoreTestingSelectionName": "string",
   "UpdateTime": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[CreationTime](#API_UpdateRestoreTestingSelection_ResponseSyntax)**

The time the resource testing selection was
updated successfully.

Type: Timestamp

**[RestoreTestingPlanArn](#API_UpdateRestoreTestingSelection_ResponseSyntax)**

Unique string that is the name of the restore testing plan.

Type: String

**[RestoreTestingPlanName](#API_UpdateRestoreTestingSelection_ResponseSyntax)**

The restore testing plan with which the updated restore
testing selection is associated.

Type: String

**[RestoreTestingSelectionName](#API_UpdateRestoreTestingSelection_ResponseSyntax)**

The returned restore testing selection name.

Type: String

**[UpdateTime](#API_UpdateRestoreTestingSelection_ResponseSyntax)**

The time the update completed for the restore
testing selection.

Type: Timestamp

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

AWS Backup can't perform the action that you requested until it finishes
performing a previous action. Try again later.

**Context**

**Type**

HTTP Status Code: 400

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

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/updaterestoretestingselection.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/updaterestoretestingselection.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateRestoreTestingPlan

UpdateTieringConfiguration

All content copied from https://docs.aws.amazon.com/.
