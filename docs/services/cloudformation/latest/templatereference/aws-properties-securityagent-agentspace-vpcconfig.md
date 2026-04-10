This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace VpcConfig

The VPC configuration for a pentest, specifying the VPC, security groups, and subnets to use during testing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupArns" : [ String, ... ],
  "SubnetArns" : [ String, ... ],
  "VpcArn" : String
}

```

### YAML

```yaml

  SecurityGroupArns:
    - String
  SubnetArns:
    - String
  VpcArn: String

```

## Properties

`SecurityGroupArns`

The Amazon Resource Names (ARNs) of the security groups for the VPC configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetArns`

The Amazon Resource Names (ARNs) of the subnets for the VPC configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcArn`

The Amazon Resource Name (ARN) of the VPC.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::SecurityAgent::Application

All content copied from https://docs.aws.amazon.com/.
