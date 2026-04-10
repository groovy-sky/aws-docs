This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTEvents::AlarmModel DynamoDB

Defines an action to write to the Amazon DynamoDB table that you created. The standard action
payload contains all the information about the detector model instance and the event that
triggered the action. You can customize the [payload](../../../../reference/iotevents/latest/apireference/api-payload.md). One column of the
DynamoDB table receives all attribute-value pairs in the payload that you specify.

You must use expressions for all parameters in `DynamoDBAction`. The expressions
accept literals, operators, functions, references, and substitution templates.

###### Examples

- For literal values, the expressions must contain single quotes. For example, the value
for the `hashKeyType` parameter can be `'STRING'`.

- For references, you must specify either variables or input values. For example, the
value for the `hashKeyField` parameter can be
`$input.GreenhouseInput.name`.

- For a substitution template, you must use `${}`, and the template must be
in single quotes. A substitution template can also contain a combination of literals,
operators, functions, references, and substitution templates.

In the following example, the value for the `hashKeyValue` parameter uses a
substitution template.

`'${$input.GreenhouseInput.temperature * 6 / 5 + 32} in Fahrenheit'`

- For a string concatenation, you must use `+`. A string concatenation can
also contain a combination of literals, operators, functions, references, and substitution
templates.

In the following example, the value for the `tableName` parameter uses a
string concatenation.

`'GreenhouseTemperatureTable ' + $input.GreenhouseInput.date`

For more information,
see [Expressions](../../../iotevents/latest/developerguide/iotevents-expressions.md)
in the _AWS IoT Events Developer Guide_.

If the defined payload type is a string, `DynamoDBAction` writes non-JSON data to
the DynamoDB table as binary data. The DynamoDB console displays the data as Base64-encoded text.
The value for the `payloadField` parameter is
`<payload-field>_raw`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HashKeyField" : String,
  "HashKeyType" : String,
  "HashKeyValue" : String,
  "Operation" : String,
  "Payload" : Payload,
  "PayloadField" : String,
  "RangeKeyField" : String,
  "RangeKeyType" : String,
  "RangeKeyValue" : String,
  "TableName" : String
}

```

### YAML

```yaml

  HashKeyField: String
  HashKeyType: String
  HashKeyValue: String
  Operation: String
  Payload:
    Payload
  PayloadField: String
  RangeKeyField: String
  RangeKeyType: String
  RangeKeyValue: String
  TableName: String

```

## Properties

`HashKeyField`

The name of the hash key (also called the partition key). The `hashKeyField`
value must match the partition key of the target DynamoDB table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HashKeyType`

The data type for the hash key (also called the partition key). You can specify the
following values:

- `'STRING'` \- The hash key is a string.

- `'NUMBER'` \- The hash key is a number.

If you don't specify `hashKeyType`, the default value is
`'STRING'`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HashKeyValue`

The value of the hash key (also called the partition key).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operation`

The type of operation to perform. You can specify the following values:

- `'INSERT'` \- Insert data as a new item into the DynamoDB table. This item uses
the specified hash key as a partition key. If you specified a range key, the item uses the
range key as a sort key.

- `'UPDATE'` \- Update an existing item of the DynamoDB table with new data. This
item's partition key must match the specified hash key. If you specified a range key, the
range key must match the item's sort key.

- `'DELETE'` \- Delete an existing item of the DynamoDB table. This item's
partition key must match the specified hash key. If you specified a range key, the range
key must match the item's sort key.

If you don't specify this parameter, AWS IoT Events triggers the `'INSERT'`
operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Payload`

Information needed to configure the payload.

By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload
contains all attribute-value pairs that have the information about the detector model instance
and the event triggered the action. To configure the action payload, you can use
`contentExpression`.

_Required_: No

_Type_: [Payload](aws-properties-iotevents-alarmmodel-payload.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PayloadField`

The name of the DynamoDB column that receives the action payload.

If you don't specify this parameter, the name of the DynamoDB column is
`payload`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeKeyField`

The name of the range key (also called the sort key). The `rangeKeyField` value
must match the sort key of the target DynamoDB table.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeKeyType`

The data type for the range key (also called the sort key), You can specify the following
values:

- `'STRING'` \- The range key is a string.

- `'NUMBER'` \- The range key is number.

If you don't specify `rangeKeyField`, the default value is
`'STRING'`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RangeKeyValue`

The value of the range key (also called the sort key).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The name of the DynamoDB table. The `tableName` value must match the table name of
the target DynamoDB table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AssetPropertyVariant

DynamoDBv2

All content copied from https://docs.aws.amazon.com/.
