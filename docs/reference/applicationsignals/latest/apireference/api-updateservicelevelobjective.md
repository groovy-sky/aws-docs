# UpdateServiceLevelObjective

Updates an existing service level objective (SLO). If you omit parameters, the previous values
of those parameters are retained.

You cannot change from a period-based SLO to a request-based SLO,
or change from a request-based SLO to a period-based SLO.

## Request Syntax

```nohighlight

PATCH /slo/Id HTTP/1.1
Content-type: application/json

{
   "BurnRateConfigurations": [
      {
         "LookBackWindowMinutes": number
      }
   ],
   "Description": "string",
   "Goal": {
      "AttainmentGoal": number,
      "Interval": { ... },
      "WarningThreshold": number
   },
   "RequestBasedSliConfig": {
      "ComparisonOperator": "string",
      "MetricThreshold": number,
      "RequestBasedSliMetricConfig": {
         "DependencyConfig": {
            "DependencyKeyAttributes": {
               "string" : "string"
            },
            "DependencyOperationName": "string"
         },
         "KeyAttributes": {
            "string" : "string"
         },
         "MetricName": "string",
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
   "SliConfig": {
      "ComparisonOperator": "string",
      "MetricThreshold": number,
      "SliMetricConfig": {
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
         "MetricName": "string",
         "MetricSource": {
            "MetricSourceAttributes": {
               "string" : "string"
            },
            "MetricSourceKeyAttributes": {
               "string" : "string"
            }
         },
         "MetricType": "string",
         "OperationName": "string",
         "PeriodSeconds": number,
         "Statistic": "string"
      }
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[Id](#API_UpdateServiceLevelObjective_RequestSyntax)**

The Amazon Resource Name (ARN) or name of the service level objective that you want to update.

Pattern: `[0-9A-Za-z][-._0-9A-Za-z ]{0,126}[0-9A-Za-z]$|^arn:(aws|aws-us-gov):application-signals:[^:]*:[^:]*:slo/[0-9A-Za-z][-._0-9A-Za-z ]{0,126}[0-9A-Za-z]`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[BurnRateConfigurations](#API_UpdateServiceLevelObjective_RequestSyntax)**

Use this array to create _burn rates_ for this SLO. Each
burn rate is a metric that indicates how fast the service is consuming the error budget, relative to the attainment goal of the SLO.

Type: Array of [BurnRateConfiguration](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_BurnRateConfiguration.html) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: No

**[Description](#API_UpdateServiceLevelObjective_RequestSyntax)**

An optional description for the SLO.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**[Goal](#API_UpdateServiceLevelObjective_RequestSyntax)**

A structure that contains the attributes that determine the goal of the SLO. This includes
the time period for evaluation and the attainment threshold.

Type: [Goal](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_Goal.html) object

Required: No

**[RequestBasedSliConfig](#API_UpdateServiceLevelObjective_RequestSyntax)**

If this SLO is a request-based SLO, this structure defines the information about what performance metric this SLO will monitor.

You can't specify both `SliConfig` and `RequestBasedSliConfig` in the same operation.

Type: [RequestBasedServiceLevelIndicatorConfig](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_RequestBasedServiceLevelIndicatorConfig.html) object

Required: No

**[SliConfig](#API_UpdateServiceLevelObjective_RequestSyntax)**

If this SLO is a period-based SLO, this structure defines the information about what performance metric this SLO will monitor.

Type: [ServiceLevelIndicatorConfig](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_ServiceLevelIndicatorConfig.html) object

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "Slo": {
      "Arn": "string",
      "BurnRateConfigurations": [
         {
            "LookBackWindowMinutes": number
         }
      ],
      "CreatedTime": number,
      "Description": "string",
      "EvaluationType": "string",
      "Goal": {
         "AttainmentGoal": number,
         "Interval": { ... },
         "WarningThreshold": number
      },
      "LastUpdatedTime": number,
      "MetricSourceType": "string",
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
      }
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Slo](#API_UpdateServiceLevelObjective_ResponseSyntax)**

A structure that contains information about the SLO that you just updated.

Type: [ServiceLevelObjective](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_ServiceLevelObjective.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/CommonErrors.html).

**ResourceNotFoundException**

Resource not found.

**ResourceId**

Can't find the resource id.

**ResourceType**

The resource type is not valid.

HTTP Status Code: 404

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

**ValidationException**

The resource is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/application-signals-2024-04-15/UpdateServiceLevelObjective)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/application-signals-2024-04-15/UpdateServiceLevelObjective)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UntagResource

Data Types
