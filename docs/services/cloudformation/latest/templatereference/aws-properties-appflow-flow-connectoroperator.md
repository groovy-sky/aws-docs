This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow ConnectorOperator

The operation to be performed on the provided source fields.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amplitude" : String,
  "CustomConnector" : String,
  "Datadog" : String,
  "Dynatrace" : String,
  "GoogleAnalytics" : String,
  "InforNexus" : String,
  "Marketo" : String,
  "Pardot" : String,
  "S3" : String,
  "Salesforce" : String,
  "SAPOData" : String,
  "ServiceNow" : String,
  "Singular" : String,
  "Slack" : String,
  "Trendmicro" : String,
  "Veeva" : String,
  "Zendesk" : String
}

```

### YAML

```yaml

  Amplitude: String
  CustomConnector: String
  Datadog: String
  Dynatrace: String
  GoogleAnalytics: String
  InforNexus: String
  Marketo: String
  Pardot: String
  S3: String
  Salesforce: String
  SAPOData: String
  ServiceNow: String
  Singular: String
  Slack: String
  Trendmicro: String
  Veeva: String
  Zendesk: String

```

## Properties

`Amplitude`

The operation to be performed on the provided Amplitude source fields.

_Required_: No

_Type_: String

_Allowed values_: `BETWEEN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomConnector`

Operators supported by the custom connector.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | LESS_THAN | GREATER_THAN | CONTAINS | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Datadog`

The operation to be performed on the provided Datadog source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | BETWEEN | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dynatrace`

The operation to be performed on the provided Dynatrace source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | BETWEEN | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GoogleAnalytics`

The operation to be performed on the provided Google Analytics source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | BETWEEN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InforNexus`

The operation to be performed on the provided Infor Nexus source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | BETWEEN | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Marketo`

The operation to be performed on the provided Marketo source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | LESS_THAN | GREATER_THAN | BETWEEN | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pardot`

The operation to be performed on the provided Salesforce Pardot source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | EQUAL_TO | NO_OP | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

The operation to be performed on the provided Amazon S3 source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | LESS_THAN | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Salesforce`

The operation to be performed on the provided Salesforce source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | LESS_THAN | CONTAINS | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SAPOData`

The operation to be performed on the provided SAPOData source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | LESS_THAN | CONTAINS | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNow`

The operation to be performed on the provided ServiceNow source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | LESS_THAN | CONTAINS | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Singular`

The operation to be performed on the provided Singular source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Slack`

The operation to be performed on the provided Slack source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | BETWEEN | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Trendmicro`

The operation to be performed on the provided Trend Micro source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Veeva`

The operation to be performed on the provided Veeva source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | LESS_THAN | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Zendesk`

The operation to be performed on the provided Zendesk source fields.

_Required_: No

_Type_: String

_Allowed values_: `PROJECTION | GREATER_THAN | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConnectorOperator](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ConnectorOperator.html) in
the _Amazon AppFlow API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AmplitudeSourceProperties

CustomConnectorDestinationProperties
