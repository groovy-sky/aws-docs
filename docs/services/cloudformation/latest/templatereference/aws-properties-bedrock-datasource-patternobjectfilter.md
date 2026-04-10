This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource PatternObjectFilter

The specific filters applied to your data source content. You can filter out or
include certain content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExclusionFilters" : [ String, ... ],
  "InclusionFilters" : [ String, ... ],
  "ObjectType" : String
}

```

### YAML

```yaml

  ExclusionFilters:
    - String
  InclusionFilters:
    - String
  ObjectType: String

```

## Properties

`ExclusionFilters`

A list of one or more exclusion regular expression patterns to exclude certain
object types that adhere to the pattern. If you specify an inclusion and exclusion
filter/pattern and both match a document, the exclusion filter takes precedence
and the document isn’t crawled.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000 | 25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InclusionFilters`

A list of one or more inclusion regular expression patterns to include certain
object types that adhere to the pattern. If you specify an inclusion and exclusion
filter/pattern and both match a document, the exclusion filter takes precedence
and the document isn’t crawled.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000 | 25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectType`

The supported object type or content type of the data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParsingPrompt

PatternObjectFilterConfiguration

All content copied from https://docs.aws.amazon.com/.
