This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::CustomEntityType

Creates a custom pattern that is used to detect sensitive data across the columns and rows of your structured data.

Each custom pattern you create specifies a regular expression and an optional list of context words. If no context words are passed only a regular expression is checked.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::CustomEntityType",
  "Properties" : {
      "ContextWords" : [ String, ... ],
      "Name" : String,
      "RegexString" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Glue::CustomEntityType
Properties:
  ContextWords:
    - String
  Name: String
  RegexString:
    String
  Tags:
    - Tag

```

## Properties

`ContextWords`

A list of context words. If none of these context words are found within the vicinity of the regular expression the data will not be detected as sensitive data.

If no context words are passed only a regular expression is checked.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the custom pattern that allows it to be retrieved or deleted later. This name must be unique per AWS account.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegexString`

A regular expression string that is used for detecting sensitive data in a custom pattern.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

AWS tags that contain a key value pair and may be searched by console, command line, or API.

_Required_: No

_Type_: Array of [`Tag`](aws-properties-resource-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Targets

AWS::Glue::Database

All content copied from https://docs.aws.amazon.com/.
