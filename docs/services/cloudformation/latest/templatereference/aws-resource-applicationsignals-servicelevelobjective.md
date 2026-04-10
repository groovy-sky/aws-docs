This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::ServiceLevelObjective

Creates or updates a service level objective (SLO), which can help you ensure that your critical business operations are
meeting customer expectations. Use SLOs to set and track specific target levels for the
reliability and availability of your applications and services. SLOs use service level indicators (SLIs) to
calculate whether the application is performing at the level that you want.

Create an SLO to set a target for a service operation, or service dependency's availability or latency. CloudWatch
measures this target frequently you can find whether it has been breached.

The target performance quality that is defined for an SLO is the _attainment goal_. An
attainment goal is the percentage of time or requests that the SLI is expected to meet the threshold over each time interval.
For example, an attainment goal of 99.9% means that within your interval, you are targeting 99.9% of the
periods to be in healthy state.

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

When you call this operation, Application Signals creates the _AWSServiceRoleForCloudWatchApplicationSignals_ service-linked role,
if it doesn't already exist in your account. This service-
linked role has the following permissions:

- `xray:GetServiceGraph`

- `logs:StartQuery`

- `logs:GetQueryResults`

- `cloudwatch:GetMetricData`

- `cloudwatch:ListMetrics`

- `tag:GetResources`

- `autoscaling:DescribeAutoScalingGroups`

You can easily set SLO targets for your applications, and their dependencies, that are discovered by Application Signals, using critical metrics such as latency and availability.
You can also set SLOs against any CloudWatch metric or math expression that produces a time series.

###### Note

You can't create an SLO for a service operation that was discovered by Application Signals until after that operation has reported standard
metrics to Application Signals.

You cannot change from a period-based SLO to a request-based SLO, or change from a request-based SLO to a period-based SLO.

For more information about SLOs, see [Service level objectives (SLOs)](../../../amazoncloudwatch/latest/monitoring/cloudwatch-servicelevelobjectives.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApplicationSignals::ServiceLevelObjective",
  "Properties" : {
      "BurnRateConfigurations" : [ BurnRateConfiguration, ... ],
      "Description" : String,
      "ExclusionWindows" : [ ExclusionWindow, ... ],
      "Goal" : Goal,
      "Name" : String,
      "RequestBasedSli" : RequestBasedSli,
      "Sli" : Sli,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApplicationSignals::ServiceLevelObjective
Properties:
  BurnRateConfigurations:
    - BurnRateConfiguration
  Description: String
  ExclusionWindows:
    - ExclusionWindow
  Goal:
    Goal
  Name: String
  RequestBasedSli:
    RequestBasedSli
  Sli:
    Sli
  Tags:
    - Tag

```

## Properties

`BurnRateConfigurations`

Each object in this array defines the length of the look-back window used to calculate one burn rate metric
for this SLO. The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO.

_Required_: No

_Type_: Array of [BurnRateConfiguration](aws-properties-applicationsignals-servicelevelobjective-burnrateconfiguration.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

An optional description for this SLO.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExclusionWindows`

The time window to be excluded from the SLO performance metrics.

_Required_: No

_Type_: Array of [ExclusionWindow](aws-properties-applicationsignals-servicelevelobjective-exclusionwindow.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Goal`

This structure contains the attributes that determine the goal of an SLO. This includes
the time period for evaluation and the attainment threshold.

_Required_: No

_Type_: [Goal](aws-properties-applicationsignals-servicelevelobjective-goal.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for this SLO.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z][-._0-9A-Za-z ]{0,126}[0-9A-Za-z]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RequestBasedSli`

A structure containing information about the performance metric that this SLO monitors, if this is a request-based SLO.

_Required_: No

_Type_: [RequestBasedSli](aws-properties-applicationsignals-servicelevelobjective-requestbasedsli.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sli`

A structure containing information about the performance metric that this SLO monitors, if this is a period-based SLO.

_Required_: No

_Type_: [Sli](aws-properties-applicationsignals-servicelevelobjective-sli.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs to associate with the SLO. You can associate as many as 50 tags with an SLO. To be able to associate tags with the SLO when you create the SLO, you must have the cloudwatch:TagResource permission.

Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.

_Required_: No

_Type_: Array of [Tag](aws-properties-applicationsignals-servicelevelobjective-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the SLO. For example, `arn:aws:cloudwatch:us-west-1:123456789012:slo:my-slo-name`

### Fn::GetAtt

`Arn`

The ARN of this SLO.

`CreatedTime`

The date and time that this SLO was created.

`EvaluationType`

Displays whether this is a period-based SLO or a request-based SLO.

`LastUpdatedTime`

The time that this SLO was most recently updated.

## Examples

- [A period-based SLO on service dependency latency](#aws-resource-applicationsignals-servicelevelobjective--examples--A_period-based_SLO_on_service_dependency_latency)

- [A request-based SLO on service dependency latency](#aws-resource-applicationsignals-servicelevelobjective--examples--A_request-based_SLO_on_service_dependency_latency)

- [A period-based SLO on service operation latency](#aws-resource-applicationsignals-servicelevelobjective--examples--A_period-based_SLO_on_service_operation_latency)

- [A period-based SLO on service operation availability](#aws-resource-applicationsignals-servicelevelobjective--examples--A_period-based_SLO_on_service_operation_availability)

- [A period-based SLO on CloudWatch metrics](#aws-resource-applicationsignals-servicelevelobjective--examples--A_period-based_SLO_on_CloudWatch_metrics)

- [A request-based SLO on CloudWatch metrics](#aws-resource-applicationsignals-servicelevelobjective--examples--A_request-based_SLO_on_CloudWatch_metrics)

### A period-based SLO on service dependency latency

The following example creates a period-based SLO on service dependency latency.

#### YAML

```yaml

Resources:
  ServiceLevelObjective:
    Type: AWS::ApplicationSignals::ServiceLevelObjective
    Properties:
      Name: !Join ["-", ["CanaryTestSlo", !Select [2, !Split ['/', !Ref AWS::StackId]]]]
      Description: "Sample period-based latency slo with dependency config."
      Sli:
        SliMetric:
          KeyAttributes:
            Environment: "eks:pet-clinic-slo-canary/default"
            Name: "TestServiceForCloudFormation-1"
            Type: "Service"
          OperationName: "Operation-1"
          MetricType: "LATENCY"
          Statistic: "p99"
          PeriodSeconds: 60
          DependencyConfig:
            DependencyKeyAttributes:
              Identifier: "sample_kinesis_stream"
              ResourceType: "AWS::Kinesis::Stream"
              Type: "AWS::Resource"
            DependencyOperationName: "Operation-1"
        MetricThreshold: 2.5
        ComparisonOperator: "LessThanOrEqualTo"
      Goal:
        Interval:
          RollingInterval:
            DurationUnit: "DAY"
            Duration: 1
        AttainmentGoal: 90.0
        WarningThreshold: 80.0
```

### A request-based SLO on service dependency latency

The following example creates a request-based SLO on service dependency availability.

#### YAML

```yaml

Resources:
  ServiceLevelObjective:
    Type: AWS::ApplicationSignals::ServiceLevelObjective
    Properties:
      Name: !Join ["-", ["CanaryTestSlo", !Select [2, !Split ['/', !Ref AWS::StackId]]]]
      Description: "Sample request-based availability slo with dependency config."
      RequestBasedSli:
        RequestBasedSliMetric:
          KeyAttributes:
            Environment: "eks:pet-clinic-slo-canary/default"
            Name: "TestServiceForCloudFormation-1"
            Type: "Service"
          OperationName: "Operation-1"
          MetricType: "AVAILABILITY"
          DependencyConfig:
            DependencyKeyAttributes:
              Identifier: "sample_kinesis_stream"
              ResourceType: "AWS::Kinesis::Stream"
              Type: "AWS::Resource"
            DependencyOperationName: "Operation-1"
      Goal:
        Interval:
          RollingInterval:
            DurationUnit: "DAY"
            Duration: 1
        AttainmentGoal: 90.0
        WarningThreshold: 80.0
```

### A period-based SLO on service operation latency

The following example creates a period-based SLO on service operation latency.

#### YAML

```yaml

Resources:
  ServiceLevelObjective:
    Type: AWS::ApplicationSignals::ServiceLevelObjective
    Properties:
      Name: !Join ["-", ["CanaryTestSlo", !Select [2, !Split ['/', !Ref AWS::StackId]]]]
      Description: "Canary test custom-metric slo description."
      Sli:
        SliMetric:
          KeyAttributes:
            Environment: "eks:pet-clinic-slo-canary/default"
            Name: "TestServiceForCloudFormation-1"
            Type: "Service"
          OperationName: "Operation-1"
          MetricType: "LATENCY"
          Statistic: "p99"
          PeriodSeconds: 60
        MetricThreshold: 2.5
        ComparisonOperator: "LessThanOrEqualTo"
      Goal:
        Interval:
          RollingInterval:
            DurationUnit: "DAY"
            Duration: 1
        AttainmentGoal: 90.0
        WarningThreshold: 80.0
```

### A period-based SLO on service operation availability

The following example creates a period-based SLO on service operation availability.

#### YAML

```yaml

Resources:
  ServiceLevelObjective:
    Type: AWS::ApplicationSignals::ServiceLevelObjective
    Properties:
      Name: !Join ["-", ["CanaryTestSlo", !Select [2, !Split ['/', !Ref AWS::StackId]]]]
      Description: "Canary test custom-metric slo description."
      Sli:
        SliMetric:
          KeyAttributes:
            Environment: "eks:pet-clinic-slo-canary/default"
            Name: "TestServiceForCloudFormation-1"
            Type: "Service"
          OperationName: "Operation-1"
          MetricType: "AVAILABILITY"
          PeriodSeconds: 60
        MetricThreshold: 2.5
        ComparisonOperator: "LessThanOrEqualTo"
      Goal:
        Interval:
          RollingInterval:
            DurationUnit: "DAY"
            Duration: 1
        AttainmentGoal: 90.0
        WarningThreshold: 80.0
```

### A period-based SLO on CloudWatch metrics

The following example creates a period-based SLO on CloudWatch metrics.

#### YAML

```yaml

Resources:
  ServiceLevelObjective:
    Type: AWS::ApplicationSignals::ServiceLevelObjective
    Properties:
      Name: !Join ["-", ["CanaryTestSlo", !Select [2, !Split ['/', !Ref AWS::StackId]]]]
      Description: "Canary test custom-metric slo description."
      Sli:
        SliMetric:
          MetricDataQueries:
            - MetricStat:
                Period: 60
                Metric:
                  MetricName: "Latency"
                  Dimensions:
                    - Value: "Operation-1"
                      Name: "Operation"
                    - Value: "TestServiceForCloudFormation-1"
                      Name: "Service"
                    - Value: "eks:pet-clinic-slo-canary/default"
                      Name: "Environment"
                  Namespace: "ApplicationSignals"
                Stat: "p99"
                Unit: "Milliseconds"
              Id: "latency"
              ReturnData: false
            - Id: "latencyfinal"
              ReturnData: true
              Expression: "latency * 1"
        MetricThreshold: 2.5
        ComparisonOperator: "LessThanOrEqualTo"
      Goal:
        Interval:
          RollingInterval:
            DurationUnit: "DAY"
            Duration: 1
        AttainmentGoal: 90.0
        WarningThreshold: 80.0
```

### A request-based SLO on CloudWatch metrics

The following example creates a request-based SLO on CloudWatch metrics.

#### YAML

```yaml

Resources:
  ServiceLevelObjective:
    Type: AWS::ApplicationSignals::ServiceLevelObjective
    Properties:
      Name: !Join ["-", ["CanaryTestSlo", !Select [2, !Split ['/', !Ref AWS::StackId]]]]
      Description: "Canary test request-based slo description."
      RequestBasedSli:
        RequestBasedSliMetric:
          TotalRequestCountMetric:
            - MetricStat:
                Period: 60
                Metric:
                  MetricName: "Latency"
                  Dimensions:
                    - Value: "Operation-1"
                      Name: "Operation"
                    - Value: "TestServiceForCloudFormation-1"
                      Name: "Service"
                    - Value: "eks:pet-clinic-slo-canary/default"
                      Name: "Environment"
                  Namespace: "ApplicationSignals"
                Stat: "SampleCount"
                Unit: "Milliseconds"
              Id: "totalCount"
              ReturnData: false
            - Id: "totalCountFinal"
              ReturnData: true
              Expression: "totalCount * 1"
          MonitoredRequestCountMetric:
            GoodCountMetric:
              - MetricStat:
                  Period: 60
                  Metric:
                    MetricName: "Latency"
                    Dimensions:
                      - Value: "Operation-1"
                        Name: "Operation"
                      - Value: "TestServiceForCloudFormation-1"
                        Name: "Service"
                      - Value: "eks:pet-clinic-slo-canary/default"
                        Name: "Environment"
                    Namespace: "ApplicationSignals"
                  Stat: "SampleCount"
                  Unit: "Milliseconds"
                Id: "goodCount"
                ReturnData: false
              - Id: "goodCountFinal"
                ReturnData: true
                Expression: "goodCount - 1"
      Goal:
        Interval:
          RollingInterval:
            DurationUnit: "DAY"
            Duration: 1
        AttainmentGoal: 90.0
        WarningThreshold: 80.0
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GroupingAttributeDefinition

BurnRateConfiguration

All content copied from https://docs.aws.amazon.com/.
