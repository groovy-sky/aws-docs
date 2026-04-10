This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index DocumentMetadataConfiguration

Specifies the properties, such as relevance tuning and searchability, of an index
field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Relevance" : Relevance,
  "Search" : Search,
  "Type" : String
}

```

### YAML

```yaml

  Name: String
  Relevance:
    Relevance
  Search:
    Search
  Type: String

```

## Properties

`Name`

The name of the index field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Relevance`

Provides tuning parameters to determine how the field affects the search
results.

_Required_: No

_Type_: [Relevance](aws-properties-kendra-index-relevance.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Search`

Provides information about how the field is used during a search.

_Required_: No

_Type_: [Search](aws-properties-kendra-index-search.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The data type of the index field.

_Required_: Yes

_Type_: String

_Allowed values_: `STRING_VALUE | STRING_LIST_VALUE | LONG_VALUE | DATE_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityUnitsConfiguration

JsonTokenTypeConfiguration

All content copied from https://docs.aws.amazon.com/.
