This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::DataProtectionSettings InlineRedactionConfiguration

The configuration for in-session inline redaction.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GlobalConfidenceLevel" : Number,
  "GlobalEnforcedUrls" : [ String, ... ],
  "GlobalExemptUrls" : [ String, ... ],
  "InlineRedactionPatterns" : [ InlineRedactionPattern, ... ]
}

```

### YAML

```yaml

  GlobalConfidenceLevel: Number
  GlobalEnforcedUrls:
    - String
  GlobalExemptUrls:
    - String
  InlineRedactionPatterns:
    - InlineRedactionPattern

```

## Properties

`GlobalConfidenceLevel`

The global confidence level for the inline redaction configuration. This indicates the
certainty of data type matches in the redaction process. Confidence level 3 means high
confidence, and requires a formatted text pattern match in order for content to be
redacted. Confidence level 2 means medium confidence, and redaction considers both
formatted and unformatted text, and adds keyword associate to the logic. Confidence level 1
means low confidence, and redaction is enforced for both formatted pattern + unformatted
pattern without keyword. This is applied to patterns that do not have a pattern-level
confidence level. Defaults to confidence level 2.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalEnforcedUrls`

The global enforced URL configuration for the inline redaction configuration. This is applied to patterns that do not have a pattern-level enforced URL list.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalExemptUrls`

The global exempt URL configuration for the inline redaction configuration. This is applied to patterns that do not have a pattern-level exempt URL list.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InlineRedactionPatterns`

The inline redaction patterns to be enabled for the inline redaction configuration.

_Required_: Yes

_Type_: Array of [InlineRedactionPattern](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-dataprotectionsettings-inlineredactionpattern.html)

_Minimum_: `0`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomPattern

InlineRedactionPattern
