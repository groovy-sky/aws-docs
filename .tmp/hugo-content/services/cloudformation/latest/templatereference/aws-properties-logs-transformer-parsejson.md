This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer ParseJSON

This processor parses log events that are in JSON format. It can extract JSON
key-value pairs and place them under a destination that you specify.

Additionally, because you must have at least one parse-type processor in a
transformer, you can use `ParseJSON` as that processor for JSON-format logs,
so that you can also apply other processors, such as mutate processors, to these
logs.

For more information about this processor including examples, see [parseJSON](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-parseJSON) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String,
  "Source" : String
}

```

### YAML

```yaml

  Destination: String
  Source: String

```

## Properties

`Destination`

The location to put the parsed key value pair into. If you omit this parameter, it is
placed under the root node.

_Required_: No

_Type_: String

_Pattern_: `^.*[a-zA-Z0-9]+.*$`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

Path to the field in the log event that will be parsed. Use dot notation to access
child fields. For example, `store.book`

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParseCloudfront

ParseKeyValue

All content copied from https://docs.aws.amazon.com/.
