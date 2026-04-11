# ListReportPlans

Returns a list of your report plans. For detailed information about a single report
plan, use `DescribeReportPlan`.

## Request Syntax

```nohighlight

GET /audit/report-plans?MaxResults=MaxResults&NextToken=NextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[MaxResults](#API_ListReportPlans_RequestSyntax)**

The number of desired results from 1 to 1000. Optional. If unspecified, the query will
return 1 MB of data.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[NextToken](#API_ListReportPlans_RequestSyntax)**

An identifier that was returned from the previous call to this operation, which can be
used to return the next set of items in the list.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "NextToken": "string",
   "ReportPlans": [
      {
         "CreationTime": number,
         "DeploymentStatus": "string",
         "LastAttemptedExecutionTime": number,
         "LastSuccessfulExecutionTime": number,
         "ReportDeliveryChannel": {
            "Formats": [ "string" ],
            "S3BucketName": "string",
            "S3KeyPrefix": "string"
         },
         "ReportPlanArn": "string",
         "ReportPlanDescription": "string",
         "ReportPlanName": "string",
         "ReportSetting": {
            "Accounts": [ "string" ],
            "FrameworkArns": [ "string" ],
            "NumberOfFrameworks": number,
            "OrganizationUnits": [ "string" ],
            "Regions": [ "string" ],
            "ReportTemplate": "string"
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListReportPlans_ResponseSyntax)**

An identifier that was returned from the previous call to this operation, which can be
used to return the next set of items in the list.

Type: String

**[ReportPlans](#API_ListReportPlans_ResponseSyntax)**

The report plans with detailed information for each plan. This information
includes the Amazon Resource Name (ARN), report plan name, description, settings, delivery
channel, deployment status, creation time, and last times the report plan attempted to and
successfully ran.

Type: Array of [ReportPlan](api-reportplan.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValueException**

Indicates that something is wrong with a parameter's value. For example, the value is
out of range.

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

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/listreportplans.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/listreportplans.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/listreportplans.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/listreportplans.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/listreportplans.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/listreportplans.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/listreportplans.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/listreportplans.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/listreportplans.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/listreportplans.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListReportJobs

ListRestoreAccessBackupVaults

All content copied from https://docs.aws.amazon.com/.
