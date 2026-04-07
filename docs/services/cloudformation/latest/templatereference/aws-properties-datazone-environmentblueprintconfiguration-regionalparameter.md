This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::EnvironmentBlueprintConfiguration RegionalParameter

The regional parameters in the environment blueprint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : {Key: Value, ...},
  "Region" : String
}

```

### YAML

```yaml

  Parameters:
    Key: Value
  Region: String

```

## Properties

`Parameters`

A string to string map containing parameters for the region.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The region specified in the environment parameter.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-?(iso|gov)?-{1}[a-z]*-{1}[0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProvisioningConfiguration

AWS::DataZone::EnvironmentProfile
