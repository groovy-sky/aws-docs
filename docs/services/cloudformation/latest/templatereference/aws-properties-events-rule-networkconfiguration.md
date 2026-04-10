This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule NetworkConfiguration

This structure specifies the network configuration for an ECS task.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsVpcConfiguration" : AwsVpcConfiguration
}

```

### YAML

```yaml

  AwsVpcConfiguration:
    AwsVpcConfiguration

```

## Properties

`AwsVpcConfiguration`

Use this structure to specify the VPC subnets and security groups for the task, and
whether a public IP address is to be used. This structure is relevant only for ECS tasks that
use the `awsvpc` network mode.

_Required_: No

_Type_: [AwsVpcConfiguration](aws-properties-events-rule-awsvpcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [NetworkConfiguration](../../../../reference/eventbridge/latest/apireference/api-networkconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisParameters

PlacementConstraint

All content copied from https://docs.aws.amazon.com/.
