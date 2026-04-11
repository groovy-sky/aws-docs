This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentConfig TimeBasedCanary

A configuration that shifts traffic from one version of a Lambda function
or Amazon ECS task set to another in two increments. The original and target
Lambda function versions or ECS task sets are specified in the
deployment's AppSpec file.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CanaryInterval" : Integer,
  "CanaryPercentage" : Integer
}

```

### YAML

```yaml

  CanaryInterval: Integer
  CanaryPercentage: Integer

```

## Properties

`CanaryInterval`

The number of minutes between the first and second traffic shifts of a
`TimeBasedCanary` deployment.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CanaryPercentage`

The percentage of traffic to shift in the first increment of a
`TimeBasedCanary` deployment.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MinimumHealthyHostsPerZone

TimeBasedLinear

All content copied from https://docs.aws.amazon.com/.
