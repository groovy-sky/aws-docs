This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::UsageProfile ConfigurationObject

Specifies the values that an admin sets for each job or session parameter configured in a AWS Glue usage profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowedValues" : [ String, ... ],
  "DefaultValue" : String,
  "MaxValue" : String,
  "MinValue" : String
}

```

### YAML

```yaml

  AllowedValues:
    - String
  DefaultValue: String
  MaxValue: String
  MinValue: String

```

## Properties

`AllowedValues`

A list of allowed values for the parameter.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

A default value for the parameter.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxValue`

A maximum allowed value for the parameter.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinValue`

A minimum allowed value for the parameter.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::UsageProfile

ProfileConfiguration

All content copied from https://docs.aws.amazon.com/.
