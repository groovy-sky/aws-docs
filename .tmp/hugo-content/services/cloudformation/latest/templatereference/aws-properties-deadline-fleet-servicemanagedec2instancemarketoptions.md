This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet ServiceManagedEc2InstanceMarketOptions

The details of the Amazon EC2 instance market options for a service managed fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

The Amazon EC2 instance type.

_Required_: Yes

_Type_: String

_Allowed values_: `on-demand | spot | wait-and-save`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceManagedEc2InstanceCapabilities

Tag

All content copied from https://docs.aws.amazon.com/.
