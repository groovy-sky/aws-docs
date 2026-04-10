This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Macie::AllowList Criteria

Specifies the criteria for an allow list, which is a list that defines specific text
or a text pattern to ignore when inspecting data sources for sensitive data. The
criteria can be:

- The location and name of an Amazon Simple Storage Service (Amazon S3)
object that lists specific predefined text to ignore ( `S3WordsList`), or

- A regular expression ( `Regex`) that defines a text pattern to
ignore.

The criteria must specify either an S3 object or a regular expression. It can't
specify both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Regex" : String,
  "S3WordsList" : S3WordsList
}

```

### YAML

```yaml

  Regex: String
  S3WordsList:
    S3WordsList

```

## Properties

`Regex`

The regular expression ( _regex_) that defines the text pattern to
ignore. The expression can contain 1-512 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3WordsList`

The location and name of an Amazon S3 object that lists specific text to
ignore.

_Required_: No

_Type_: [S3WordsList](aws-properties-macie-allowlist-s3wordslist.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Macie::AllowList

S3WordsList

All content copied from https://docs.aws.amazon.com/.
