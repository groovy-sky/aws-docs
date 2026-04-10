This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::ProjectProfile Region

The AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RegionName" : String
}

```

### YAML

```yaml

  RegionName: String

```

## Properties

`RegionName`

The AWS Region name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z]{2}-?(iso|gov)?-{1}[a-z]*-{1}[0-9]$`

_Minimum_: `4`

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnvironmentConfigurationParametersDetails

ResourceTagParameter

All content copied from https://docs.aws.amazon.com/.
