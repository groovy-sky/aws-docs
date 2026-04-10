This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::DetectorModel Payload

Information needed to configure the payload.

By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload
contains all attribute-value pairs that have the information about the detector model instance
and the event triggered the action. To configure the action payload, you can use
`contentExpression`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentExpression" : String,
  "Type" : String
}

```

### YAML

```yaml

  ContentExpression: String
  Type: String

```

## Properties

`ContentExpression`

The content of the payload. You can use a string expression that includes quoted strings
( `'<string>'`), variables ( `$variable.<variable-name>`),
input values ( `$input.<input-name>.<path-to-datum>`), string
concatenations, and quoted strings that contain `${}` as the content. The
recommended maximum size of a content expression is 1 KB.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The value of the payload type can be either `STRING` or
`JSON`.

_Required_: Yes

_Type_: String

_Allowed values_: `STRING | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnInput

ResetTimer

All content copied from https://docs.aws.amazon.com/.
