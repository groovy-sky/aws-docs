This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule DynamoDBAction

Describes an action to write to a DynamoDB table.

The `tableName`, `hashKeyField`, and `rangeKeyField`
values must match the values used when you created the table.

The `hashKeyValue` and `rangeKeyvalue` fields use a
substitution template syntax. These templates provide data at runtime. The syntax is as
follows: ${ _sql-expression_}.

You can specify any valid expression in a WHERE or SELECT clause, including JSON
properties, comparisons, calculations, and functions. For example, the following field uses
the third level of the topic:

`"hashKeyValue": "${topic(3)}"`

The following field uses the timestamp:

`"rangeKeyValue": "${timestamp()}"`

For more information, see [DynamoDBv2 Action](../../../iot/latest/developerguide/iot-rule-actions.md) in the _AWS IoT Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HashKeyField" : String,
  "HashKeyType" : String,
  "HashKeyValue" : String,
  "PayloadField" : String,
  "RangeKeyField" : String,
  "RangeKeyType" : String,
  "RangeKeyValue" : String,
  "RoleArn" : String,
  "TableName" : String
}

```

### YAML

```yaml

  HashKeyField: String
  HashKeyType: String
  HashKeyValue: String
  PayloadField: String
  RangeKeyField: String
  RangeKeyType: String
  RangeKeyValue: String
  RoleArn: String
  TableName: String

```

## Properties

`HashKeyField`

The hash key name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HashKeyType`

The hash key type. Valid values are "STRING" or "NUMBER"

_Required_: No

_Type_: String

_Allowed values_: `STRING | NUMBER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HashKeyValue`

The hash key value.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadField`

The action payload. This name can be customized.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeKeyField`

The range key name.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeKeyType`

The range key type. Valid values are "STRING" or "NUMBER"

_Required_: No

_Type_: String

_Allowed values_: `STRING | NUMBER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeKeyValue`

The range key value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants access to the DynamoDB table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the DynamoDB table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudwatchMetricAction

DynamoDBv2Action

All content copied from https://docs.aws.amazon.com/.
