This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion ComponentPlatform

Contains information about a platform that a component supports.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : {Key: Value, ...},
  "Name" : String
}

```

### YAML

```yaml

  Attributes:
    Key: Value
  Name: String

```

## Properties

`Attributes`

A dictionary of attributes for the platform. The AWS IoT Greengrass Core software defines
the `os` and `platform` by default. You can specify additional platform
attributes for a core device when you deploy the AWS IoT Greengrass nucleus component. For
more information, see the [AWS IoT Greengrass\
nucleus component](../../../greengrass/v2/developerguide/greengrass-nucleus-component.md) in the _AWS IoT Greengrass V2 Developer_
_Guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The friendly name of the platform. This name helps you identify the platform.

If you omit this parameter, AWS IoT Greengrass creates a friendly name from the
`os` and `architecture` of the platform.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentDependencyRequirement

LambdaContainerParams

All content copied from https://docs.aws.amazon.com/.
