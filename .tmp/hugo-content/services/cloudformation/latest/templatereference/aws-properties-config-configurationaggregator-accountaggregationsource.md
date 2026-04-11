This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigurationAggregator AccountAggregationSource

A collection of accounts and regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountIds" : [ String, ... ],
  "AllAwsRegions" : Boolean,
  "AwsRegions" : [ String, ... ]
}

```

### YAML

```yaml

  AccountIds:
    - String
  AllAwsRegions: Boolean
  AwsRegions:
    - String

```

## Properties

`AccountIds`

The 12-digit account ID of the account being aggregated.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllAwsRegions`

If true, aggregate existing AWS Config regions and future
regions.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsRegions`

The source regions being aggregated.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Config::ConfigurationAggregator

OrganizationAggregationSource

All content copied from https://docs.aws.amazon.com/.
