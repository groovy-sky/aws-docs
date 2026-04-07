This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentConfig TrafficRoutingConfig

The configuration that specifies how traffic is shifted from one version of a Lambda function to another version during an AWS Lambda deployment,
or from one Amazon ECS task set to another during an Amazon ECS
deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TimeBasedCanary" : TimeBasedCanary,
  "TimeBasedLinear" : TimeBasedLinear,
  "Type" : String
}

```

### YAML

```yaml

  TimeBasedCanary:
    TimeBasedCanary
  TimeBasedLinear:
    TimeBasedLinear
  Type: String

```

## Properties

`TimeBasedCanary`

A configuration that shifts traffic from one version of a Lambda function
or ECS task set to another in two increments. The original and target Lambda
function versions or ECS task sets are specified in the deployment's AppSpec
file.

_Required_: No

_Type_: [TimeBasedCanary](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentconfig-timebasedcanary.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeBasedLinear`

A configuration that shifts traffic from one version of a Lambda function
or Amazon ECS task set to another in equal increments, with an equal number of
minutes between each increment. The original and target Lambda function
versions or Amazon ECS task sets are specified in the deployment's AppSpec
file.

_Required_: No

_Type_: [TimeBasedLinear](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codedeploy-deploymentconfig-timebasedlinear.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of traffic shifting ( `TimeBasedCanary` or
`TimeBasedLinear`) used by a deployment configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `TimeBasedCanary | TimeBasedLinear | TimeBasedFlexible | AllAtOnce`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TimeBasedLinear

ZonalConfig
