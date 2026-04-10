This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link ModuleConfiguration

Describes the configuration of a module.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DependsOn" : [ String, ... ],
  "ModuleParameters" : ModuleParameters,
  "Name" : String,
  "Version" : String
}

```

### YAML

```yaml

  DependsOn:
    - String
  ModuleParameters:
    ModuleParameters
  Name: String
  Version: String

```

## Properties

`DependsOn`

The dependencies of the module.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModuleParameters`

Describes the parameters of a module.

_Required_: No

_Type_: [ModuleParameters](aws-properties-rtbfabric-link-moduleparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the module.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9 -]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the module.

_Required_: No

_Type_: String

_Pattern_: `^[a-z0-9]{1,25}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LinkLogSettings

ModuleParameters

All content copied from https://docs.aws.amazon.com/.
