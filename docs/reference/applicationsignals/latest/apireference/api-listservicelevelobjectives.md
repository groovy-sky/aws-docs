# ListServiceLevelObjectives

Returns a list of SLOs created in this account.

## Request Syntax

```nohighlight

POST /slos?IncludeLinkedAccounts=IncludeLinkedAccounts&MaxResults=MaxResults&NextToken=NextToken&OperationName=OperationName&SloOwnerAwsAccountId=SloOwnerAwsAccountId HTTP/1.1
Content-type: application/json

{
   "DependencyConfig": {
      "DependencyKeyAttributes": {
         "string" : "string"
      },
      "DependencyOperationName": "string"
   },
   "KeyAttributes": {
      "string" : "string"
   },
   "MetricSource": {
      "MetricSourceAttributes": {
         "string" : "string"
      },
      "MetricSourceKeyAttributes": {
         "string" : "string"
      }
   },
   "MetricSourceTypes": [ "string" ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[IncludeLinkedAccounts](#API_ListServiceLevelObjectives_RequestSyntax)**

If you are using this operation in a monitoring account, specify `true` to include SLO from source accounts in the returned data.

When you are monitoring an account, you can use AWS account ID in `KeyAttribute` filter for service source account and `SloOwnerawsaccountID` for SLO source account with `IncludeLinkedAccounts` to filter the returned data to only a single source account.

**[MaxResults](#API_ListServiceLevelObjectives_RequestSyntax)**

The maximum number of results to return in one operation. If you omit this
parameter, the default of 50 is used.

Valid Range: Minimum value of 1. Maximum value of 50.

**[NextToken](#API_ListServiceLevelObjectives_RequestSyntax)**

Include this value, if it was returned by the previous operation, to get the next set of service level objectives.

**[OperationName](#API_ListServiceLevelObjectives_RequestSyntax)**

The name of the operation that this SLO is associated with.

Length Constraints: Minimum length of 1. Maximum length of 255.

**[SloOwnerAwsAccountId](#API_ListServiceLevelObjectives_RequestSyntax)**

SLO's AWS account ID.

Pattern: `[0-9]{12}`

## Request Body

The request accepts the following data in JSON format.

**[DependencyConfig](#API_ListServiceLevelObjectives_RequestSyntax)**

Identifies the dependency using the `DependencyKeyAttributes` and `DependencyOperationName`.

Type: [DependencyConfig](api-dependencyconfig.md) object

Required: No

**[KeyAttributes](#API_ListServiceLevelObjectives_RequestSyntax)**

You can use this optional field to specify which services you want to retrieve SLO information for.

This is a string-to-string map. It can
include the following fields.

- `Type` designates the type of object this is.

- `ResourceType` specifies the type of the resource. This field is used only
when the value of the `Type` field is `Resource` or `AWS::Resource`.

- `Name` specifies the name of the object. This is used only if the value of the `Type` field
is `Service`, `RemoteService`, or `AWS::Service`.

- `Identifier` identifies the resource objects of this resource.
This is used only if the value of the `Type` field
is `Resource` or `AWS::Resource`.

- `Environment` specifies the location where this object is hosted, or what it belongs to.

Type: String to string map

Map Entries: Maximum number of 4 items.

Key Pattern: `[a-zA-Z]{1,50}`

Value Length Constraints: Minimum length of 1. Maximum length of 1024.

Value Pattern: `[ -~]*[!-~]+[ -~]*`

Required: No

**[MetricSource](#API_ListServiceLevelObjectives_RequestSyntax)**

Identifies the metric source to filter SLOs by.

Type: [MetricSource](api-metricsource.md) object

Required: No

**[MetricSourceTypes](#API_ListServiceLevelObjectives_RequestSyntax)**

Use this optional field to only include SLOs with the specified metric source types in the output. Supported types are:

- Service operation

- Service dependency

- Service

- CloudWatch metric

- AppMonitor

- Canary

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 3 items.

Valid Values: `ServiceOperation | CloudWatchMetric | ServiceDependency | AppMonitor | Canary | Service`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "NextToken": "string",
   "SloSummaries": [
      {
         "Arn": "string",
         "CreatedTime": number,
         "DependencyConfig": {
            "DependencyKeyAttributes": {
               "string" : "string"
            },
            "DependencyOperationName": "string"
         },
         "EvaluationType": "string",
         "KeyAttributes": {
            "string" : "string"
         },
         "MetricSource": {
            "MetricSourceAttributes": {
               "string" : "string"
            },
            "MetricSourceKeyAttributes": {
               "string" : "string"
            }
         },
         "MetricSourceType": "string",
         "Name": "string",
         "OperationName": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_ListServiceLevelObjectives_ResponseSyntax)**

Include this value in your next use of this API to get next set
of service level objectives.

Type: String

**[SloSummaries](#API_ListServiceLevelObjectives_ResponseSyntax)**

An array of structures, where each structure contains information about one SLO.

Type: Array of [ServiceLevelObjectiveSummary](api-servicelevelobjectivesummary.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

**ValidationException**

The resource is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/listservicelevelobjectives.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/listservicelevelobjectives.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ListServiceLevelObjectiveExclusionWindows

ListServiceOperations
