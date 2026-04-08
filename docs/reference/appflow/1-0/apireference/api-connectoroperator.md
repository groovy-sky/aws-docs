# ConnectorOperator

The operation to be performed on the provided source fields.

## Contents

**Amplitude**

The operation to be performed on the provided Amplitude source fields.

Type: String

Valid Values: `BETWEEN`

Required: No

**CustomConnector**

Operators supported by the custom connector.

Type: String

Valid Values: `PROJECTION | LESS_THAN | GREATER_THAN | CONTAINS | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Datadog**

The operation to be performed on the provided Datadog source fields.

Type: String

Valid Values: `PROJECTION | BETWEEN | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Dynatrace**

The operation to be performed on the provided Dynatrace source fields.

Type: String

Valid Values: `PROJECTION | BETWEEN | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**GoogleAnalytics**

The operation to be performed on the provided Google Analytics source fields.

Type: String

Valid Values: `PROJECTION | BETWEEN`

Required: No

**InforNexus**

The operation to be performed on the provided Infor Nexus source fields.

Type: String

Valid Values: `PROJECTION | BETWEEN | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Marketo**

The operation to be performed on the provided Marketo source fields.

Type: String

Valid Values: `PROJECTION | LESS_THAN | GREATER_THAN | BETWEEN | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Pardot**

The operation to be performed on the provided Salesforce Pardot source fields.

Type: String

Valid Values: `PROJECTION | EQUAL_TO | NO_OP | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC`

Required: No

**S3**

The operation to be performed on the provided Amazon S3 source fields.

Type: String

Valid Values: `PROJECTION | LESS_THAN | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Salesforce**

The operation to be performed on the provided Salesforce source fields.

Type: String

Valid Values: `PROJECTION | LESS_THAN | CONTAINS | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**SAPOData**

The operation to be performed on the provided SAPOData source fields.

Type: String

Valid Values: `PROJECTION | LESS_THAN | CONTAINS | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**ServiceNow**

The operation to be performed on the provided ServiceNow source fields.

Type: String

Valid Values: `PROJECTION | CONTAINS | LESS_THAN | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Singular**

The operation to be performed on the provided Singular source fields.

Type: String

Valid Values: `PROJECTION | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Slack**

The operation to be performed on the provided Slack source fields.

Type: String

Valid Values: `PROJECTION | LESS_THAN | GREATER_THAN | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Trendmicro**

The operation to be performed on the provided Trend Micro source fields.

Type: String

Valid Values: `PROJECTION | EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Veeva**

The operation to be performed on the provided Veeva source fields.

Type: String

Valid Values: `PROJECTION | LESS_THAN | GREATER_THAN | CONTAINS | BETWEEN | LESS_THAN_OR_EQUAL_TO | GREATER_THAN_OR_EQUAL_TO | EQUAL_TO | NOT_EQUAL_TO | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

**Zendesk**

The operation to be performed on the provided Zendesk source fields.

Type: String

Valid Values: `PROJECTION | GREATER_THAN | ADDITION | MULTIPLICATION | DIVISION | SUBTRACTION | MASK_ALL | MASK_FIRST_N | MASK_LAST_N | VALIDATE_NON_NULL | VALIDATE_NON_ZERO | VALIDATE_NON_NEGATIVE | VALIDATE_NUMERIC | NO_OP`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/connectoroperator.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/connectoroperator.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/connectoroperator.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ConnectorOAuthRequest

ConnectorProfile
