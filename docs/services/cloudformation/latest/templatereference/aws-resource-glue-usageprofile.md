This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::UsageProfile

Creates an AWS Glue usage profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::UsageProfile",
  "Properties" : {
      "Configuration" : ProfileConfiguration,
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Glue::UsageProfile
Properties:
  Configuration:
    ProfileConfiguration
  Description: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Configuration`

Property description not available.

_Required_: No

_Type_: [ProfileConfiguration](aws-properties-glue-usageprofile-profileconfiguration.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the usage profile.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9\-\:\_]{1,64}`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the usage profile.

_Required_: Yes

_Type_: String

_Minimum_: `5`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-glue-usageprofile-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedOn`

The date and time when the usage profile was created.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Predicate

ConfigurationObject

All content copied from https://docs.aws.amazon.com/.
