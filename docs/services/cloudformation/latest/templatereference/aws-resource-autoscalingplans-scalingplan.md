This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScalingPlans::ScalingPlan

The `AWS::AutoScalingPlans::ScalingPlan` resource defines an AWS Auto Scaling scaling plan. A scaling plan is used to scale application resources to
size them appropriately to ensure that enough resource is available in the application at
peak times and to reduce allocated resource during periods of low utilization. The
following resources can be added to a scaling plan:

- Amazon EC2 Auto Scaling groups

- Amazon EC2 Spot Fleet requests

- Amazon ECS services

- Amazon DynamoDB tables and global secondary indexes

- Amazon Aurora Replicas

For more information, see the [Scaling Plans User Guide](../../../autoscaling/plans/userguide/what-is-a-scaling-plan.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AutoScalingPlans::ScalingPlan",
  "Properties" : {
      "ApplicationSource" : ApplicationSource,
      "ScalingInstructions" : [ ScalingInstruction, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::AutoScalingPlans::ScalingPlan
Properties:
  ApplicationSource:
    ApplicationSource
  ScalingInstructions:
    - ScalingInstruction

```

## Properties

`ApplicationSource`

A CloudFormation stack or a set of tags. You can create one scaling plan per application
source. The `ApplicationSource` property must be present to ensure
interoperability with the AWS Auto Scaling console.

_Required_: Yes

_Type_: [ApplicationSource](aws-properties-autoscalingplans-scalingplan-applicationsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingInstructions`

The scaling instructions.

_Required_: Yes

_Type_: Array of [ScalingInstruction](aws-properties-autoscalingplans-scalingplan-scalinginstruction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::AutoScalingPlans::ScalingPlan`
resource to the intrinsic `Ref` function, the function returns the Amazon
Resource Name (ARN) of the scaling plan. The format of the ARN is as follows:

`arn:aws:autoscaling:region:123456789012:scalingPlan:scalingPlanName/plan-name:scalingPlanVersion/plan-version`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

## Examples

### Scaling plan

The following example creates a scaling plan named `myScalingPlan` for
an existing Auto Scaling group ( [AWS::AutoScaling::AutoScalingGroup](../userguide/aws-properties-as-group.md)) whose name you specify when launching
the stack using this template. It specifies the `TagFilters` property as
the application source. You can specify any tag key and tag value you want without
affecting the stack, as long as the key-pair is unique for each scaling plan. This
can be any value you choose that helps you identify your scaling plan configuration.
However, if you also want to use the AWS Auto Scaling console to edit the
scaling plan, then the tag must match the tag you chose for the Auto Scaling
group.

The [ScalingInstructions](../userguide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.md) property includes information that's required to
enable predictive scaling and dynamic scaling. In this example, the predictive
scaling mode specifies `ForecastOnly`. In which case, AWS Auto Scaling generates forecasts with traffic predictions for the two days
ahead, but does not schedule scaling actions to match the forecast.

#### JSON

```json

{
  "AWSTemplateFormatVersion":"2010-09-09",
  "Parameters":{
    "myTagKey":{
      "Type":"String"
    },
    "myTagValue":{
      "Type":"String"
    },
    "myASGroup":{
      "Type":"String",
      "Description":"Name of the Auto Scaling group"
    },
    "ASGMinCapacity":{
      "Type":"Number"
    },
    "ASGMaxCapacity":{
      "Type":"Number"
    },
    "ASGTargetUtilization":{
      "Type":"Number",
      "Default":"50.0"
    },
    "ASGEstimatedInstanceWarmup":{
      "Type":"Number",
      "Default":"600"
    }
  },
  "Resources":{
    "myScalingPlan":{
      "Type":"AWS::AutoScalingPlans::ScalingPlan",
      "Properties":{
        "ApplicationSource":{
          "TagFilters":[
            {
              "Key":{
                "Ref":"myTagKey"
              },
              "Values":[{
                "Ref":"myTagValue"
              }]
            }
          ]
        },
        "ScalingInstructions":[
          {
            "MinCapacity":{
              "Ref":"ASGMinCapacity"
            },
            "MaxCapacity":{
              "Ref":"ASGMaxCapacity"
            },
            "ServiceNamespace":"autoscaling",
            "ScalableDimension":"autoscaling:autoScalingGroup:DesiredCapacity",
            "ResourceId":{
              "Fn::Join":[
                "/",
                [
                  "autoScalingGroup",
                  {
                    "Ref":"myASGroup"
                  }
                ]
              ]
            },
            "TargetTrackingConfigurations":[
              {
                "PredefinedScalingMetricSpecification":{
                  "PredefinedScalingMetricType":"ASGAverageCPUUtilization"
                },
                "TargetValue":{
                  "Ref":"ASGTargetUtilization"
                },
                "EstimatedInstanceWarmup":{
                  "Ref":"ASGEstimatedInstanceWarmup"
                }
              }
            ],
            "PredefinedLoadMetricSpecification":{
              "PredefinedLoadMetricType":"ASGTotalCPUUtilization"
            },
            "PredictiveScalingMode":"ForecastOnly",
            "PredictiveScalingMaxCapacityBehavior":"SetMaxCapacityAboveForecastCapacity",
            "PredictiveScalingMaxCapacityBuffer":25,
            "ScheduledActionBufferTime":600
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  myTagKey:
    Type: String
  myTagValue:
    Type: String
  myASGroup:
    Type: String
    Description: Name of the Auto Scaling group
  ASGMinCapacity:
    Type: Number
  ASGMaxCapacity:
    Type: Number
  ASGTargetUtilization:
    Type: Number
    Default: 50.0
  ASGEstimatedInstanceWarmup:
    Type: Number
    Default: 600
Resources:
  myScalingPlan:
    Type: AWS::AutoScalingPlans::ScalingPlan
    Properties:
      ApplicationSource:
        TagFilters:
          - Key: !Ref myTagKey
            Values:
              - !Ref myTagValue
      ScalingInstructions:
        - MinCapacity: !Ref ASGMinCapacity
          MaxCapacity: !Ref ASGMaxCapacity
          ServiceNamespace: autoscaling
          ScalableDimension: autoscaling:autoScalingGroup:DesiredCapacity
          ResourceId: !Join
            - /
            - - autoScalingGroup
              - !Ref myASGroup
          TargetTrackingConfigurations:
            - PredefinedScalingMetricSpecification:
                PredefinedScalingMetricType: "ASGAverageCPUUtilization"
              TargetValue: !Ref ASGTargetUtilization
              EstimatedInstanceWarmup: !Ref ASGEstimatedInstanceWarmup
          PredefinedLoadMetricSpecification:
            PredefinedLoadMetricType: "ASGTotalCPUUtilization"
          PredictiveScalingMode: "ForecastOnly"
          PredictiveScalingMaxCapacityBehavior: "SetMaxCapacityAboveForecastCapacity"
          PredictiveScalingMaxCapacityBuffer: 25
          ScheduledActionBufferTime: 600
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Auto Scaling

ApplicationSource

All content copied from https://docs.aws.amazon.com/.
