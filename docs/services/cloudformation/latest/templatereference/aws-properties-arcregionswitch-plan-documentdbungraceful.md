This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan DocumentDbUngraceful

Configuration for handling failures when performing operations on DocumentDB global clusters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ungraceful" : String
}

```

### YAML

```yaml

  Ungraceful: String

```

## Properties

`Ungraceful`

The settings for ungraceful execution.

_Required_: No

_Type_: String

_Allowed values_: `failover`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentDbConfiguration

Ec2AsgCapacityIncreaseConfiguration

All content copied from https://docs.aws.amazon.com/.
