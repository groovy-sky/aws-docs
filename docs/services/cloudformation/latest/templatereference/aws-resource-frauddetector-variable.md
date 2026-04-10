This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FraudDetector::Variable

Manages a variable.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FraudDetector::Variable",
  "Properties" : {
      "DataSource" : String,
      "DataType" : String,
      "DefaultValue" : String,
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "VariableType" : String
    }
}

```

### YAML

```yaml

Type: AWS::FraudDetector::Variable
Properties:
  DataSource: String
  DataType: String
  DefaultValue: String
  Description: String
  Name: String
  Tags:
    - Tag
  VariableType: String

```

## Properties

`DataSource`

The data source of the variable.

Valid values: `EVENT | EXTERNAL_MODEL_SCORE`

When defining a variable within a detector, you can only use the `EVENT` value for DataSource when the _Inline_ property is set to true.
If the _Inline_ property is set false, you can use either `EVENT` or `MODEL_SCORE` for DataSource.

_Required_: Yes

_Type_: String

_Allowed values_: `EVENT | EXTERNAL_MODEL_SCORE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The data type of the variable.

Valid data types: `STRING | INTEGER | BOOLEAN | FLOAT`

_Required_: Yes

_Type_: String

_Allowed values_: `STRING | INTEGER | FLOAT | BOOLEAN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value of the variable.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the variable.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the variable.

Pattern: `^[0-9a-z_-]+$`

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z_][a-z0-9_]{0,99}?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-frauddetector-variable-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariableType`

The type of the variable. For more information see [Variable types](../../../frauddetector/latest/ug/create-a-variable.md#variable-types).

Valid Values: `AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 | BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE | BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS | FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID | PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 | SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME | SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT `

_Required_: No

_Type_: String

_Allowed values_: `AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 | BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE | BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS | FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID | PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 | SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME | SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the primary identifier for the resource, which is the ARN.

Example: `{ "Ref": "arn:aws:frauddetector:us-west-2:123123123123:outcome/outcome_name"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the variable.

`CreatedTime`

Timestamp of when variable was created.

`LastUpdatedTime`

Timestamp of when variable was last updated.

## See also

- [CreateVariable](../../../frauddetector/latest/api/api-createvariable.md) in the _Amazon Fraud Detector API Reference_

- [Create a variable](../../../frauddetector/latest/ug/create-a-variable.md) in the _Amazon Fraud Detector User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
