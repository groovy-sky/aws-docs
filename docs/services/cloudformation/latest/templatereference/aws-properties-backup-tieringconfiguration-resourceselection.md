This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::TieringConfiguration ResourceSelection

This contains metadata about resource selection for tiering configurations.

You can specify up to 5 different resource selections per tiering configuration.
Data moved to lower-cost tier remains there until deletion (one-way transition).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Resources" : [ String, ... ],
  "ResourceType" : String,
  "TieringDownSettingsInDays" : Integer
}

```

### YAML

```yaml

  Resources:
    - String
  ResourceType: String
  TieringDownSettingsInDays: Integer

```

## Properties

`Resources`

An array of strings that either contains ARNs of the associated resources
or contains a wildcard `*` to specify all resources.
You can specify up to 100 specific resources per tiering configuration.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The type of AWS resource; for example, `S3` for Amazon S3.
For tiering configurations, this is currently limited to `S3`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\-\_\.]{1,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TieringDownSettingsInDays`

The number of days after creation within a backup vault that an object can transition to the low cost warm storage tier.
Must be a positive integer between 60 and 36500 days.

_Required_: Yes

_Type_: Integer

_Minimum_: `60`

_Maximum_: `36500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Backup::TieringConfiguration

Next
