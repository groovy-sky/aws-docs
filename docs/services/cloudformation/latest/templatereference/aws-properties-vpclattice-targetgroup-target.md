This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::TargetGroup Target

Describes a target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "Port" : Integer
}

```

### YAML

```yaml

  Id: String
  Port: Integer

```

## Properties

`Id`

The ID of the target. If the target group type is `INSTANCE`, this is an instance
ID. If the target group type is `IP`, this is an IP address. If the target group type
is `LAMBDA`, this is the ARN of a Lambda function. If the target group type is
`ALB`, this is the ARN of an Application Load Balancer.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port on which the target is listening. For HTTP, the default is 80. For HTTPS, the
default is 443.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

TargetGroupConfig
