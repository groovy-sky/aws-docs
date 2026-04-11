This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::Input Attribute

The attributes from the JSON payload that are made available by the input. Inputs are
derived from messages sent to the AWS IoT Events system using `BatchPutMessage`. Each such
message contains a JSON payload. Those attributes (and their paired values) specified here are
available for use in the `condition` expressions used by detectors.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JsonPath" : String
}

```

### YAML

```yaml

  JsonPath: String

```

## Properties

`JsonPath`

An expression that specifies an attribute-value pair in a JSON structure. Use this to
specify an attribute from the JSON payload that is made available by the input. Inputs are
derived from messages sent to AWS IoT Events ( `BatchPutMessage`). Each such message contains
a JSON payload. The attribute (and its paired value) specified here are available for use in
the `condition` expressions used by detectors.

Syntax: `<field-name>.<field-name>...`

_Required_: Yes

_Type_: String

_Pattern_: ``^((`[a-zA-Z0-9_\- ]+`)|([a-zA-Z0-9_\-]+))(\.((`[a-zA-Z0-9_\- ]+`)|([a-zA-Z0-9_\-]+)))*$``

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTEvents::Input

InputDefinition

All content copied from https://docs.aws.amazon.com/.
