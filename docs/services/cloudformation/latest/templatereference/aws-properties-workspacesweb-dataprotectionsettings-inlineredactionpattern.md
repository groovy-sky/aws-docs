This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::DataProtectionSettings InlineRedactionPattern

The set of patterns that determine the data types redacted in session.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BuiltInPatternId" : String,
  "ConfidenceLevel" : Number,
  "CustomPattern" : CustomPattern,
  "EnforcedUrls" : [ String, ... ],
  "ExemptUrls" : [ String, ... ],
  "RedactionPlaceHolder" : RedactionPlaceHolder
}

```

### YAML

```yaml

  BuiltInPatternId: String
  ConfidenceLevel: Number
  CustomPattern:
    CustomPattern
  EnforcedUrls:
    - String
  ExemptUrls:
    - String
  RedactionPlaceHolder:
    RedactionPlaceHolder

```

## Properties

`BuiltInPatternId`

The built-in pattern from the list of preconfigured patterns. Either a customPattern or builtInPatternId is required. To view the entire list of data types and their corresponding built-in pattern IDs, see [Base inline redaction](../../../workspaces-web/latest/adminguide/base-inline-redaction.md).

_Required_: No

_Type_: String

_Pattern_: `^[_\-\d\w]+$`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfidenceLevel`

The confidence level for inline redaction pattern. This indicates the certainty of data
type matches in the redaction process. Confidence level 3 means high confidence, and
requires a formatted text pattern match in order for content to be redacted. Confidence
level 2 means medium confidence, and redaction considers both formatted and unformatted
text, and adds keyword associate to the logic. Confidence level 1 means low confidence, and
redaction is enforced for both formatted pattern + unformatted pattern without keyword.
This overrides the global confidence level.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomPattern`

The configuration for a custom pattern. Either a customPattern or builtInPatternId is required.

_Required_: No

_Type_: [CustomPattern](aws-properties-workspacesweb-dataprotectionsettings-custompattern.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnforcedUrls`

The enforced URL configuration for the inline redaction pattern. This will override the global enforced URL configuration.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExemptUrls`

The exempt URL configuration for the inline redaction pattern. This will override the global exempt URL configuration for the inline redaction pattern.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedactionPlaceHolder`

The redaction placeholder that will replace the redacted text in session for the inline redaction pattern.

_Required_: Yes

_Type_: [RedactionPlaceHolder](aws-properties-workspacesweb-dataprotectionsettings-redactionplaceholder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InlineRedactionConfiguration

RedactionPlaceHolder

All content copied from https://docs.aws.amazon.com/.
