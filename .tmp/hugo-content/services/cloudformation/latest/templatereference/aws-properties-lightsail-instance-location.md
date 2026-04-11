This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::Instance Location

`Location` is a property of the [AWS::Lightsail::Instance](../userguide/aws-resource-lightsail-instance.md) resource. It describes the location for an
instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZone" : String,
  "RegionName" : String
}

```

### YAML

```yaml

  AvailabilityZone: String
  RegionName: String

```

## Properties

`AvailabilityZone`

The Availability Zone for the instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionName`

The name of the AWS Region for the instance.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Hardware

MonthlyTransfer

All content copied from https://docs.aws.amazon.com/.
