This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet TargetGroupsConfig

Describes the target groups to attach to a Spot Fleet. Spot Fleet registers the
running Spot Instances with these target groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetGroups" : [ TargetGroup, ... ]
}

```

### YAML

```yaml

  TargetGroups:
    - TargetGroup

```

## Properties

`TargetGroups`

One or more target groups.

_Required_: Yes

_Type_: Array of [TargetGroup](aws-properties-ec2-spotfleet-targetgroup.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetGroup

TotalLocalStorageGBRequest

All content copied from https://docs.aws.amazon.com/.
