This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Partnership X12Delimiters

In X12 EDI messages, delimiters are used to mark the end of segments or elements, and
are defined in the interchange control header. The delimiters are part of the message's
syntax and divide up its different elements.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentSeparator" : String,
  "DataElementSeparator" : String,
  "SegmentTerminator" : String
}

```

### YAML

```yaml

  ComponentSeparator: String
  DataElementSeparator: String
  SegmentTerminator: String

```

## Properties

`ComponentSeparator`

The component, or sub-element, separator. The default value is `:`
(colon).

_Required_: No

_Type_: String

_Pattern_: ``^[!&'()*+,\-./:;?=%@\[\]_{}|<>~^`"]$``

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataElementSeparator`

The data element separator. The default value is `*` (asterisk).

_Required_: No

_Type_: String

_Pattern_: ``^[!&'()*+,\-./:;?=%@\[\]_{}|<>~^`"]$``

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SegmentTerminator`

The segment terminator. The default value is `~` (tilde).

_Required_: No

_Type_: String

_Pattern_: ``^[!&'()*+,\-./:;?=%@\[\]_{}|<>~^`"]$``

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

X12ControlNumbers

X12Envelope

All content copied from https://docs.aws.amazon.com/.
