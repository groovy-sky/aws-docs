This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::ResponderGateway AutoScalingGroupsConfiguration

Describes the configuration of an auto scaling group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingGroupNameList" : [ String, ... ],
  "RoleArn" : String
}

```

### YAML

```yaml

  AutoScalingGroupNameList:
    - String
  RoleArn: String

```

## Properties

`AutoScalingGroupNameList`

The names of the auto scaling group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`RoleArn`

The role ARN of the auto scaling group.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RTBFabric::ResponderGateway

EksEndpointsConfiguration

All content copied from https://docs.aws.amazon.com/.
