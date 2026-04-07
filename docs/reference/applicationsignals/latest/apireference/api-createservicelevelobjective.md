# CreateServiceLevelObjective

Creates a service level objective (SLO), which can help you ensure that your critical business operations are
meeting customer expectations. Use SLOs to set and track specific target levels for the
reliability and availability of your applications and services. SLOs use service level indicators (SLIs) to
calculate whether the application is performing at the level that you want.

Create an SLO to set a target for a service or operation’s availability or latency. CloudWatch
measures this target frequently you can find whether it has been breached.

The target performance quality that is defined for an SLO is the _attainment goal_.

You can set SLO targets for your applications that are discovered by Application Signals, using critical metrics such as latency and availability.
You can also set SLOs against any CloudWatch metric or math expression that produces a time series.

###### Note

You can't create an SLO for a service operation that was discovered by Application Signals until after that operation has reported standard
metrics to Application Signals.

When you create an SLO, you specify whether it is a _period-based SLO_
or a _request-based SLO_. Each type of SLO has a different way of evaluating
your application's performance against its attainment goal.

- A _period-based SLO_ uses defined _periods_ of time within
a specified total time interval. For each period of time, Application Signals determines whether the
application met its goal. The attainment rate is calculated as the `number of good periods/number of total periods`.

For example, for a period-based SLO, meeting an attainment goal of 99.9% means that within your interval, your application must meet its
performance goal during at least 99.9% of the
time periods.

- A _request-based SLO_ doesn't use pre-defined periods of time. Instead,
the SLO measures `number of good requests/number of total requests` during the interval. At any time, you can find the ratio of
good requests to total requests for the interval up to the time stamp that you specify, and measure that ratio against the goal set in your SLO.

After you have created an SLO, you can retrieve error budget reports for it.
An _error budget_ is the amount of time or amount of requests that your application can be non-compliant
with the SLO's goal, and still have your application meet the goal.

- For a period-based SLO, the error budget starts at a number defined by the highest number of periods that can fail to meet the threshold,
while still meeting the overall goal. The _remaining error budget_ decreases with every failed period
that is recorded. The error budget within one interval can never increase.

For example, an SLO with a threshold that 99.95% of requests must be completed under 2000ms every month
translates to an error budget of 21.9 minutes of downtime per month.

- For a request-based SLO, the remaining error budget is dynamic and can increase or decrease, depending on
the ratio of good requests to total requests.

For more information about SLOs, see [Service level objectives (SLOs)](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-servicelevelobjectives.md).

When you perform a `CreateServiceLevelObjective` operation, Application Signals creates the _AWSServiceRoleForCloudWatchApplicationSignals_ service-linked role,
if it doesn't already exist in your account. This service-
linked role has the following permissions:

- `xray:GetServiceGraph`

- `logs:StartQuery`

- `logs:GetQueryResults`

- `cloudwatch:GetMetricData`

- `cloudwatch:ListMetrics`

- `tag:GetResources`

- `autoscaling:DescribeAutoScalingGroups`

## Request Syntax

```nohighlight

POST /slo HTTP/1.1
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
   "Name": "string",
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
   },
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[BurnRateConfigurations](#API_CreateServiceLevelObjective_RequestSyntax)**

Use this array to create _burn rates_ for this SLO. Each
burn rate is a metric that indicates how fast the service is consuming the error budget, relative to the attainment goal of the SLO.

Type: Array of [BurnRateConfiguration](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_BurnRateConfiguration.html) objects

Array Members: Minimum number of 0 items. Maximum number of 10 items.

Required: No

**[Description](#API_CreateServiceLevelObjective_RequestSyntax)**

An optional description for this SLO.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**[Goal](#API_CreateServiceLevelObjective_RequestSyntax)**

This structure contains the attributes that determine the goal of the SLO.

Type: [Goal](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_Goal.html) object

Required: No

**[Name](#API_CreateServiceLevelObjective_RequestSyntax)**

A name for this SLO.

Type: String

Pattern: `[0-9A-Za-z][-._0-9A-Za-z ]{0,126}[0-9A-Za-z]`

Required: Yes

**[RequestBasedSliConfig](#API_CreateServiceLevelObjective_RequestSyntax)**

If this SLO is a request-based SLO, this structure defines the information about what performance metric this SLO will monitor.

You can't specify both `RequestBasedSliConfig` and `SliConfig` in the same operation.

Type: [RequestBasedServiceLevelIndicatorConfig](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_RequestBasedServiceLevelIndicatorConfig.html) object

Required: No

**[SliConfig](#API_CreateServiceLevelObjective_RequestSyntax)**

If this SLO is a period-based SLO, this structure defines the information about what performance metric this SLO will monitor.

You can't specify both `RequestBasedSliConfig` and `SliConfig` in the same operation.

Type: [ServiceLevelIndicatorConfig](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_ServiceLevelIndicatorConfig.html) object

Required: No

**[Tags](#API_CreateServiceLevelObjective_RequestSyntax)**

A list of key-value pairs to associate with the SLO. You can associate as many as 50 tags with an SLO.
To be able to associate tags with the SLO when you create the SLO, you must
have the `cloudwatch:TagResource` permission.

Tags can help you organize and categorize your resources. You can also use them to scope user
permissions by granting a user
permission to access or change only resources with certain tag values.

Type: Array of [Tag](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_Tag.html) objects

Array Members: Minimum number of 0 items. Maximum number of 200 items.

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

**[Slo](#API_CreateServiceLevelObjective_ResponseSyntax)**

A structure that contains information about the SLO that you just created.

Type: [ServiceLevelObjective](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/API_ServiceLevelObjective.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/applicationsignals/latest/APIReference/CommonErrors.html).

**AccessDeniedException**

You don't have sufficient permissions to perform this action.

HTTP Status Code: 403

**ConflictException**

This operation attempted to create a resource that already exists.

HTTP Status Code: 409

**ServiceQuotaExceededException**

This request exceeds a service quota.

HTTP Status Code: 402

**ThrottlingException**

The request was throttled because of quota limits.

HTTP Status Code: 429

**ValidationException**

The resource is not valid.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/application-signals-2024-04-15/CreateServiceLevelObjective)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/application-signals-2024-04-15/CreateServiceLevelObjective)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BatchUpdateExclusionWindows

DeleteGroupingConfiguration
