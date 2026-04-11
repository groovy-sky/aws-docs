This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Asg

Configuration for an Amazon EC2 Auto Scaling group used in a Region switch plan.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "CrossAccountRole" : String,
  "ExternalId" : String
}

```

### YAML

```yaml

  Arn: String
  CrossAccountRole: String
  ExternalId: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the EC2 Auto Scaling group.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z-]*:autoscaling:[a-z0-9-]+:\d{12}:autoScalingGroup:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}:autoScalingGroupName/[\S\s]{1,255}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CrossAccountRole`

The cross account role for the configuration.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[a-zA-Z0-9-]*:iam::[0-9]{12}:role/.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExternalId`

The external ID (secret key) for the configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArcRoutingControlConfiguration

AssociatedAlarm

All content copied from https://docs.aws.amazon.com/.
