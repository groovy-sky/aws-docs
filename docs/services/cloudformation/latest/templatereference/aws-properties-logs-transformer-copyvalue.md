This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Transformer CopyValue

This processor copies values within a log event. You can also use this processor to
add metadata to log events by copying the values of the following metadata keys into the
log events: `@logGroupName`, `@logGroupStream`,
`@accountId`, `@regionName`.

For more information about this processor including examples, see [copyValue](../../../amazoncloudwatch/latest/logs/cloudwatch-logs-transformation-processors.md#CloudWatch-Logs-Transformation-copyValue) in the _CloudWatch Logs User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Entries" : [ CopyValueEntry, ... ]
}

```

### YAML

```yaml

  Entries:
    - CopyValueEntry

```

## Properties

`Entries`

An array of `CopyValueEntry` objects, where each object contains the
information about one field value to copy.

_Required_: Yes

_Type_: Array of [CopyValueEntry](aws-properties-logs-transformer-copyvalueentry.md)

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AddKeys

CopyValueEntry

All content copied from https://docs.aws.amazon.com/.
