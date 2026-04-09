# GetControl

###### Important

AWS Audit Manager will no longer be open to new
customers starting April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](../../../../services/audit-manager/latest/userguide/audit-manager-availability-change.md).

Gets information about a specified control.

## Request Syntax

```nohighlight

GET /controls/controlId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[controlId](#API_GetControl_RequestSyntax)**

The identifier for the control.

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: Yes

## Request Body

The request does not have a request body.

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

**[control](#API_GetControl_ResponseSyntax)**

The details of the control that the `GetControl` API returned.

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

**ValidationException**

The request has invalid or missing parameters.

**fields**

The fields that caused the error, if applicable.

**reason**

The reason the request failed validation.

HTTP Status Code: 400

## Examples

### Getting information about a control

This example shows a sample response for the `GetControl` API
operation. This example returns details about a specific custom control. In this
case, the control has three automated data sources. Each data source is a specific
CloudTrail event ( `ec2_ModifyInstanceAttribute`,
`ec2_CreateImage`, and `ec2_RunInstances`).

#### Sample Response

```json

{
    "control": {
        "arn": "arn:aws:auditmanager:us-east-1:111122223333:control/a1b2c3d4-5678-90ab-cdef-example11111",
        "id": "a1b2c3d4-5678-90ab-cdef-example11111",
        "type": "Custom",
        "name": "1.0.1 - CloudTrail Instance Events_CUSTOM",
        "description": "CloudTrail eventName for instances",
        "testingInformation": "CloudTrail eventName:\n\n- CreateImage\n- RunInstances\n- ModifyInstanceAttribute ",
        "actionPlanTitle": "The title of my action plan",
        "actionPlanInstructions": "The details of my action plan instructions",
        "controlSources": "AWS CloudTrail",
        "controlMappingSources": [
            {
                "sourceId": "sourceid-5678-90ab-cdef-example11111",
                "sourceName": "Data source 1",
                "sourceSetUpOption": "System_Controls_Mapping",
                "sourceType": "AWS_Cloudtrail",
                "sourceKeyword": {
                    "keywordInputType": "SELECT_FROM_LIST",
                    "keywordValue": "ec2_ModifyInstanceAttribute"
                }
            },
            {
                "sourceId": "sourceid-5678-90ab-cdef-example22222",
                "sourceName": "Data source 2",
                "sourceSetUpOption": "System_Controls_Mapping",
                "sourceType": "AWS_Cloudtrail",
                "sourceKeyword": {
                    "keywordInputType": "SELECT_FROM_LIST",
                    "keywordValue": "ec2_CreateImage"
                }
            },
            {
                "sourceId": "sourceid-5678-90ab-cdef-example33333",
                "sourceName": "Data source 3",
                "sourceSetUpOption": "System_Controls_Mapping",
                "sourceType": "AWS_Cloudtrail",
                "sourceKeyword": {
                    "keywordInputType": "SELECT_FROM_LIST",
                    "keywordValue": "ec2_RunInstances"
                }
            }
        ],
        "createdAt": "2023-03-13T12:00:58.943000-07:00",
        "lastUpdatedAt": "2023-03-13T12:00:58.943000-07:00",
        "createdBy": "john-doe",
        "lastUpdatedBy": "jane-doe",
        "state": "ACTIVE",
        "tags": {
            "Department": "Legal"
        }
    }
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for Python](../../../../services/goto/boto3/auditmanager-2017-07-25/getcontrol.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/getcontrol.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetChangeLogs

GetDelegations

All content copied from https://docs.aws.amazon.com/.
