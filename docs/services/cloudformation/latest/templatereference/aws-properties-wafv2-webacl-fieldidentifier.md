This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL FieldIdentifier

The identifier of a field in the web request payload that contains customer data.

This data type is used to specify fields in the `RequestInspection` and
`RequestInspectionACFP` configurations, which are used in the managed rule
group configurations `AWSManagedRulesATPRuleSet` and
`AWSManagedRulesACFPRuleSet`, respectively.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Identifier" : String
}

```

### YAML

```yaml

  Identifier: String

```

## Properties

`Identifier`

The name of the field.

When the `PayloadType` in the request inspection is `JSON`, this
identifier must be in JSON pointer syntax. For example `/form/username`. For
information about the JSON Pointer syntax, see the Internet Engineering Task Force (IETF)
documentation [JavaScript Object Notation\
(JSON) Pointer](https://tools.ietf.org/html/rfc6901).

When the `PayloadType` is `FORM_ENCODED`, use the HTML form names.
For example, `username`.

For more information, see the descriptions for each field type in the request inspection
properties.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExcludedRule

FieldToMatch

All content copied from https://docs.aws.amazon.com/.
