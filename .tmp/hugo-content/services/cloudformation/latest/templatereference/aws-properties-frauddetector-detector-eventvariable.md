This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FraudDetector::Detector EventVariable

The event type variable for the detector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "CreatedTime" : String,
  "DataSource" : String,
  "DataType" : String,
  "DefaultValue" : String,
  "Description" : String,
  "Inline" : Boolean,
  "LastUpdatedTime" : String,
  "Name" : String,
  "Tags" : [ Tag, ... ],
  "VariableType" : String
}

```

### YAML

```yaml

  Arn: String
  CreatedTime: String
  DataSource: String
  DataType: String
  DefaultValue: String
  Description: String
  Inline: Boolean
  LastUpdatedTime: String
  Name: String
  Tags:
    - Tag
  VariableType: String

```

## Properties

`Arn`

The event variable ARN.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedTime`

Timestamp for when the event variable was created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSource`

The data source of the event variable.

Valid values: `EVENT | EXTERNAL_MODEL_SCORE`

When defining a variable within a detector, you can only use the `EVENT` value for DataSource when the _Inline_ property is set to true.
If the _Inline_ property is set false, you can use either `EVENT` or `MODEL_SCORE` for DataSource.

_Required_: No

_Type_: String

_Allowed values_: `EVENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataType`

The data type of the event variable.

Valid values: `STRING | INTEGER | BOOLEAN | FLOAT`

_Required_: No

_Type_: String

_Allowed values_: `STRING | INTEGER | FLOAT | BOOLEAN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultValue`

The default value of the event variable. This is required if you are providing the details of your variables instead of the ARN.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the event variable.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Inline`

Indicates whether the resource is defined within this CloudFormation template and impacts the create, update, and delete behavior of the stack. If the value is `true`,
CloudFormation will create/update/delete the resource when creating/updating/deleting the stack. If the value is `false`, CloudFormation will validate that the object exists and
then use it within the resource without making changes to the object.

For example, when creating `AWS::FraudDetector::Detector` you must define at least two variables. You can set `Inline=true` for these variables and CloudFormation will
create/update/delete the variables as part of stack operations. However, if you set `Inline=false`, CloudFormation will associate the variables to your detector but not execute any
changes to the variables.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastUpdatedTime`

Timestamp for when the event variable was last updated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the event variable.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-frauddetector-detector-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VariableType`

The type of event variable. For more information, see [Variable types](../../../frauddetector/latest/ug/create-a-variable.md#variable-types).

_Required_: No

_Type_: String

_Allowed values_: `AUTH_CODE | AVS | BILLING_ADDRESS_L1 | BILLING_ADDRESS_L2 | BILLING_CITY | BILLING_COUNTRY | BILLING_NAME | BILLING_PHONE | BILLING_STATE | BILLING_ZIP | CARD_BIN | CATEGORICAL | CURRENCY_CODE | EMAIL_ADDRESS | FINGERPRINT | FRAUD_LABEL | FREE_FORM_TEXT | IP_ADDRESS | NUMERIC | ORDER_ID | PAYMENT_TYPE | PHONE_NUMBER | PRICE | PRODUCT_CATEGORY | SHIPPING_ADDRESS_L1 | SHIPPING_ADDRESS_L2 | SHIPPING_CITY | SHIPPING_COUNTRY | SHIPPING_NAME | SHIPPING_PHONE | SHIPPING_STATE | SHIPPING_ZIP | USERAGENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventType

Label

All content copied from https://docs.aws.amazon.com/.
