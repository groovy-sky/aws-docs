This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::Configuration

Create a new configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::Configuration",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "RunConfigurations" : RunConfigurations,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Omics::Configuration
Properties:
  Description: String
  Name: String
  RunConfigurations:
    RunConfigurations
  Tags:
    Key: Value

```

## Properties

`Description`

Description for the configuration.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

User-friendly name for the configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9\-\._]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RunConfigurations`

Run-specific configuration settings.

_Required_: Yes

_Type_: [RunConfigurations](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-omics-configuration-runconfigurations.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Unique resource identifier for the configuration.

`CreationTime`

Configuration creation timestamp.

`Status`

Current configuration status.

`Uuid`

Unique identifier for the configuration.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TsvStoreOptions

RunConfigurations
