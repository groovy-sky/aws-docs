This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL TextTransformation

Text transformations eliminate some of the unusual formatting that attackers use in web
requests in an effort to bypass detection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Priority" : Integer,
  "Type" : String
}

```

### YAML

```yaml

  Priority: Integer
  Type: String

```

## Properties

`Priority`

Sets the relative processing order for multiple transformations.
AWS WAF processes all transformations, from lowest priority to highest,
before inspecting the transformed content. The priorities don't need to be consecutive, but
they must all be different.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

For detailed descriptions of each of the transformation types, see [Text transformations](../../../waf/latest/developerguide/waf-rule-statement-transformation.md)
in the _AWS WAF Developer Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `NONE | COMPRESS_WHITE_SPACE | HTML_ENTITY_DECODE | LOWERCASE | CMD_LINE | URL_DECODE | BASE64_DECODE | HEX_DECODE | MD5 | REPLACE_COMMENTS | ESCAPE_SEQ_DECODE | SQL_HEX_DECODE | CSS_DECODE | JS_DECODE | NORMALIZE_PATH | NORMALIZE_PATH_WIN | REMOVE_NULLS | REPLACE_NULLS | BASE64_DECODE_EXT | URL_DECODE_UNI | UTF8_TO_UNICODE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

UriFragment

All content copied from https://docs.aws.amazon.com/.
