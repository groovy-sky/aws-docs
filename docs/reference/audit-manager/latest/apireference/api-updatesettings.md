# UpdateSettings

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Updates Audit Manager settings for the current account.

## Request Syntax

```nohighlight

PUT /settings HTTP/1.1
Content-type: application/json

{
   "defaultAssessmentReportsDestination": {
      "destination": "string",
      "destinationType": "string"
   },
   "defaultExportDestination": {
      "destination": "string",
      "destinationType": "string"
   },
   "defaultProcessOwners": [
      {
         "roleArn": "string",
         "roleType": "string"
      }
   ],
   "deregistrationPolicy": {
      "deleteResources": "string"
   },
   "evidenceFinderEnabled": boolean,
   "kmsKey": "string",
   "snsTopic": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[defaultAssessmentReportsDestination](#API_UpdateSettings_RequestSyntax)**

The default S3 destination bucket for storing assessment reports.

Type: [AssessmentReportsDestination](api-assessmentreportsdestination.md) object

Required: No

**[defaultExportDestination](#API_UpdateSettings_RequestSyntax)**

The default S3 destination bucket for storing evidence finder exports.

Type: [DefaultExportDestination](api-defaultexportdestination.md) object

Required: No

**[defaultProcessOwners](#API_UpdateSettings_RequestSyntax)**

A list of the default audit owners.

Type: Array of [Role](api-role.md) objects

Required: No

**[deregistrationPolicy](#API_UpdateSettings_RequestSyntax)**

The deregistration policy for your Audit Manager data. You can
use this attribute to determine how your data is handled when you deregister Audit Manager.

Type: [DeregistrationPolicy](api-deregistrationpolicy.md) object

Required: No

**[evidenceFinderEnabled](#API_UpdateSettings_RequestSyntax)**

Specifies whether the evidence finder feature is enabled. Change this attribute to
enable or disable evidence finder.

###### Important

When you use this attribute to disable evidence finder, Audit Manager deletes the
event data store that’s used to query your evidence data. As a result, you can’t
re-enable evidence finder and use the feature again. Your only alternative is to [deregister](api-deregisteraccount.md) and then [re-register](api-registeraccount.md)
Audit Manager.

Type: Boolean

Required: No

**[kmsKey](#API_UpdateSettings_RequestSyntax)**

The AWS KMS key details.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 2048.

Pattern: `^arn:.*:kms:.*|DEFAULT`

Required: No

**[snsTopic](#API_UpdateSettings_RequestSyntax)**

The Amazon Simple Notification Service (Amazon SNS) topic that AWS Audit Manager sends
notifications to.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 2048.

Pattern: `^arn:.*:sns:.*|NONE`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "settings": {
      "defaultAssessmentReportsDestination": {
         "destination": "string",
         "destinationType": "string"
      },
      "defaultExportDestination": {
         "destination": "string",
         "destinationType": "string"
      },
      "defaultProcessOwners": [
         {
            "roleArn": "string",
            "roleType": "string"
         }
      ],
      "deregistrationPolicy": {
         "deleteResources": "string"
      },
      "evidenceFinderEnablement": {
         "backfillStatus": "string",
         "enablementStatus": "string",
         "error": "string",
         "eventDataStoreArn": "string"
      },
      "isAwsOrgEnabled": boolean,
      "kmsKey": "string",
      "snsTopic": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[settings](#API_UpdateSettings_ResponseSyntax)**

The current list of settings.

Type: [Settings](api-settings.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

Your account isn't registered with AWS Audit Manager. Check the delegated
administrator setup on the Audit Manager settings page, and try again.

HTTP Status Code: 403

**InternalServerException**

An internal service error occurred during the processing of your request. Try again
later.

HTTP Status Code: 500

**ValidationException**

The request has invalid or missing parameters.

**fields**

The fields that caused the error, if applicable.

**reason**

The reason the request failed validation.

HTTP Status Code: 400

## Examples

### Enabling evidence finder

This is an example response for the `GetSettings` API operation, where
the `evidenceFinderEnabled` parameter was used to enable evidence
finder.

This response returns the following `evidenceFinderEnablement`
data:

- `enablementStatus` shows the current status of evidence finder.
In this case, `ENABLE_IN_PROGRESS` indicates that you requested to
enable evidence finder, and an event data store is being created to support
evidence finder queries.

- `backfillStatus` shows the current status of the evidence data
backfill. In this case, `NOT_STARTED` indicates that the backfill
hasn’t started yet.

#### Sample Response

```json

{
    "settings": {
        "isAwsOrgEnabled": false,
        "snsTopic": "arn:aws:sns:us-east-1:111122223333:my-assessment-topic",
        "defaultAssessmentReportsDestination": {
            "destinationType": "S3",
            "destination": "s3://my-assessment-report-destination"
        },
        "defaultProcessOwners": [
            {
                "roleType": "PROCESS_OWNER",
                "roleArn": "arn:aws:iam::111122223333:role/Administrator"
            }
        ],
        "kmsKey": "DEFAULT",
        "evidenceFinderEnablement": {
            "enablementStatus": "ENABLE_IN_PROGRESS",
            "backfillStatus": "NOT_STARTED"
        }
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/updatesettings.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/updatesettings.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateControl

ValidateAssessmentReportIntegrity

All content copied from https://docs.aws.amazon.com/.
