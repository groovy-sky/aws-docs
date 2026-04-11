# CreateControl

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Creates a new custom control in AWS Audit Manager.

## Request Syntax

```nohighlight

POST /controls HTTP/1.1
Content-type: application/json

{
   "actionPlanInstructions": "string",
   "actionPlanTitle": "string",
   "controlMappingSources": [
      {
         "sourceDescription": "string",
         "sourceFrequency": "string",
         "sourceKeyword": {
            "keywordInputType": "string",
            "keywordValue": "string"
         },
         "sourceName": "string",
         "sourceSetUpOption": "string",
         "sourceType": "string",
         "troubleshootingText": "string"
      }
   ],
   "description": "string",
   "name": "string",
   "tags": {
      "string" : "string"
   },
   "testingInformation": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[actionPlanInstructions](#API_CreateControl_RequestSyntax)**

The recommended actions to carry out if the control isn't fulfilled.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**[actionPlanTitle](#API_CreateControl_RequestSyntax)**

The title of the action plan for remediating the control.

Type: String

Length Constraints: Maximum length of 300.

Pattern: `^[\w\W\s\S]*$`

Required: No

**[controlMappingSources](#API_CreateControl_RequestSyntax)**

The data mapping sources for the control.

Type: Array of [CreateControlMappingSource](api-createcontrolmappingsource.md) objects

Array Members: Minimum number of 1 item.

Required: Yes

**[description](#API_CreateControl_RequestSyntax)**

The description of the control.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**[name](#API_CreateControl_RequestSyntax)**

The name of the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: Yes

**[tags](#API_CreateControl_RequestSyntax)**

The tags that are associated with the control.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Value Pattern: `.{0,255}`

Required: No

**[testingInformation](#API_CreateControl_RequestSyntax)**

The steps to follow to determine if the control is satisfied.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "control": {
      "actionPlanInstructions": "string",
      "actionPlanTitle": "string",
      "arn": "string",
      "controlMappingSources": [
         {
            "sourceDescription": "string",
            "sourceFrequency": "string",
            "sourceId": "string",
            "sourceKeyword": {
               "keywordInputType": "string",
               "keywordValue": "string"
            },
            "sourceName": "string",
            "sourceSetUpOption": "string",
            "sourceType": "string",
            "troubleshootingText": "string"
         }
      ],
      "controlSources": "string",
      "createdAt": number,
      "createdBy": "string",
      "description": "string",
      "id": "string",
      "lastUpdatedAt": number,
      "lastUpdatedBy": "string",
      "name": "string",
      "state": "string",
      "tags": {
         "string" : "string"
      },
      "testingInformation": "string",
      "type": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[control](#API_CreateControl_ResponseSyntax)**

The new control that the `CreateControl` API returned.

Type: [Control](api-control.md) object

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

**ResourceNotFoundException**

The resource that's specified in the request can't be found.

**resourceId**

The unique identifier for the resource.

**resourceType**

The type of resource that's affected by the error.

HTTP Status Code: 404

**ServiceQuotaExceededException**

You've reached your account quota for this resource type. To perform the requested
action, delete some existing resources or [request a quota increase](../../../../general/latest/gr/aws-service-limits.md) from
the Service Quotas console. For a list of Audit Manager service quotas, see [Quotas and\
restrictions for AWS Audit Manager](../../../../services/audit-manager/latest/userguide/service-quotas.md).

HTTP Status Code: 402

**ValidationException**

The request has invalid or missing parameters.

**fields**

The fields that caused the error, if applicable.

**reason**

The reason the request failed validation.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/createcontrol.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/createcontrol.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateAssessmentReport

DeleteAssessment

All content copied from https://docs.aws.amazon.com/.
