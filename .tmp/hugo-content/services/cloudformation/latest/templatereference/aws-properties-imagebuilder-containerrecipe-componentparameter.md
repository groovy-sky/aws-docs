This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ContainerRecipe ComponentParameter

Contains a key/value pair that sets the named component parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : [ String, ... ]
}

```

### YAML

```yaml

  Name: String
  Value:
    - String

```

## Properties

`Name`

The name of the component parameter to set.

_Required_: Yes

_Type_: String

_Pattern_: `[^\x00]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

Sets the value for the named component parameter.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentConfiguration

EbsInstanceBlockDeviceSpecification

All content copied from https://docs.aws.amazon.com/.
