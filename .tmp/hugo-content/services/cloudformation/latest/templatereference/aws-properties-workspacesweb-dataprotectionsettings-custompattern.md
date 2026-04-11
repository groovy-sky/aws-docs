This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::DataProtectionSettings CustomPattern

The pattern configuration for redacting custom data types in session.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KeywordRegex" : String,
  "PatternDescription" : String,
  "PatternName" : String,
  "PatternRegex" : String
}

```

### YAML

```yaml

  KeywordRegex: String
  PatternDescription: String
  PatternName: String
  PatternRegex: String

```

## Properties

`KeywordRegex`

The keyword regex for the customer pattern. After there is a match to the pattern regex, the keyword regex is used to search within the proximity of the match. If there is a keyword match, then the match is confirmed. If no keyword regex is provided, the pattern regex match will automatically be confirmed. The format must follow JavaScript regex format. The pattern must be enclosed between slashes, and can have flags behind the second slash. For example, “/ab+c/gi”

_Required_: No

_Type_: String

_Pattern_: `^\/((?:[^\n])+)\/([gimsuyvd]{0,8})$`

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternDescription`

The pattern description for the customer pattern.

_Required_: No

_Type_: String

_Pattern_: `^[ _\-\d\w]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternName`

The pattern name for the custom pattern.

_Required_: Yes

_Type_: String

_Pattern_: `^[_\-\d\w]+$`

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternRegex`

The pattern regex for the customer pattern. The format must follow JavaScript regex format. The pattern must be enclosed between slashes, and can have flags behind the second slash. For example: “/ab+c/gi”.

_Required_: Yes

_Type_: String

_Pattern_: `^\/((?:[^\n])+)\/([gimsuyvd]{0,8})$`

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::WorkSpacesWeb::DataProtectionSettings

InlineRedactionConfiguration

All content copied from https://docs.aws.amazon.com/.
