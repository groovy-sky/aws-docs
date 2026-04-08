# BatchGetServiceLevelObjectiveBudgetReport

Use this operation to retrieve one or more _service level objective (SLO) budget reports_.

An _error budget_ is the amount of time or requests in an unhealthy state that your service can
accumulate during an interval before your overall SLO budget health is breached and the SLO is considered to be
unmet. For example, an SLO with a threshold of 99.95% and a monthly interval
translates to an error budget of 21.9 minutes of
downtime in a 30-day month.

Budget reports include a health indicator, the attainment value, and
remaining budget.

For more information about SLO error budgets, see
[SLO concepts](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-servicelevelobjectives.md#CloudWatch-ServiceLevelObjectives-concepts).

## Request Syntax

```nohighlight

POST /budget-report HTTP/1.1
Content-type: application/json

{
   "SloIds": [ "string" ],
   "Timestamp": number
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[SloIds](#API_BatchGetServiceLevelObjectiveBudgetReport_RequestSyntax)**

An array containing the IDs of the service level objectives that you want to include in the report.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Required: Yes

**[Timestamp](#API_BatchGetServiceLevelObjectiveBudgetReport_RequestSyntax)**

The date and time that you want the report to be for. It is expressed as the number of milliseconds since Jan 1, 1970 00:00:00 UTC.

Type: Timestamp

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Errors": [
      {
         "Arn": "string",
         "ErrorCode": "string",
         "ErrorMessage": "string",
         "Name": "string"
      }
   ],
   "Reports": [
      {
         "Arn": "string",
         "Attainment": number,
         "BudgetRequestsRemaining": number,
         "BudgetSecondsRemaining": number,
         "BudgetStatus": "string",
         "EvaluationType": "string",
         "Goal": {
            "AttainmentGoal": number,
            "Interval": { ... },
            "WarningThreshold": number
         },
         "Name": "string",
         "RequestBasedSli": {
            "ComparisonOperator": "string",
            "MetricThreshold": number,
            "RequestBasedSliMetric": {
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
               "MetricType": "string",
               "MonitoredRequestCountMetric": { ... },
               "OperationName": "string",
               "TotalRequestCountMetric": [
                  {
                     "AccountId": "string",
                     "Expression": "string",
                     "Id": "string",
                     "Label": "string",
                     "MetricStat": {
                        "Metric": {
                           "Dimensions": [
                              {
                                 "Name": "string",
                                 "Value": "string"
                              }
                           ],
                           "MetricName": "string",
                           "Namespace": "string"
                        },
                        "Period": number,
                        "Stat": "string",
                        "Unit": "string"
                     },
                     "Period": number,
                     "ReturnData": boolean
                  }
               ]
            }
         },
         "Sli": {
            "ComparisonOperator": "string",
            "MetricThreshold": number,
            "SliMetric": {
               "DependencyConfig": {
                  "DependencyKeyAttributes": {
                     "string" : "string"
                  },
                  "DependencyOperationName": "string"
               },
               "KeyAttributes": {
                  "string" : "string"
               },
               "MetricDataQueries": [
                  {
                     "AccountId": "string",
                     "Expression": "string",
                     "Id": "string",
                     "Label": "string",
                     "MetricStat": {
                        "Metric": {
                           "Dimensions": [
                              {
                                 "Name": "string",
                                 "Value": "string"
                              }
                           ],
                           "MetricName": "string",
                           "Namespace": "string"
                        },
                        "Period": number,
                        "Stat": "string",
                        "Unit": "string"
                     },
                     "Period": number,
                     "ReturnData": boolean
                  }
               ],
               "MetricSource": {
                  "MetricSourceAttributes": {
                     "string" : "string"
                  },
                  "MetricSourceKeyAttributes": {
                     "string" : "string"
                  }
               },
               "MetricType": "string",
               "OperationName": "string"
            }
         },
         "TotalBudgetRequests": number,
         "TotalBudgetSeconds": number
      }
   ],
   "Timestamp": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Errors](#API_BatchGetServiceLevelObjectiveBudgetReport_ResponseSyntax)**

An array of structures, where each structure includes an error indicating that one
of the requests in the array was not valid.

Type: Array of [ServiceLevelObjectiveBudgetReportError](api-servicelevelobjectivebudgetreporterror.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

**[Reports](#API_BatchGetServiceLevelObjectiveBudgetReport_ResponseSyntax)**

An array of structures, where each structure is one budget report.

Type: Array of [ServiceLevelObjectiveBudgetReport](api-servicelevelobjectivebudgetreport.md) objects

Array Members: Minimum number of 0 items. Maximum number of 50 items.

**[Timestamp](#API_BatchGetServiceLevelObjectiveBudgetReport_ResponseSyntax)**

The date and time that the report is for. It is expressed as the number of milliseconds since Jan 1, 1970 00:00:00 UTC.

Type: Timestamp

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for Python](../../../../services/goto/boto3/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/application-signals-2024-04-15/batchgetservicelevelobjectivebudgetreport.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Actions

BatchUpdateExclusionWindows
