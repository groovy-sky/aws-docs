This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentConfig TimeBasedLinear

A configuration that shifts traffic from one version of a Lambda function
or ECS task set to another in equal increments, with an equal number of minutes between
each increment. The original and target Lambda function versions or ECS task
sets are specified in the deployment's AppSpec file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LinearInterval" : Integer,
  "LinearPercentage" : Integer
}

```

### YAML

```yaml

  LinearInterval: Integer
  LinearPercentage: Integer

```

## Properties

`LinearInterval`

The number of minutes between each incremental traffic shift of a
`TimeBasedLinear` deployment.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LinearPercentage`

The percentage of traffic that is shifted at the start of each increment of a
`TimeBasedLinear` deployment.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeBasedCanary

TrafficRoutingConfig

All content copied from https://docs.aws.amazon.com/.
