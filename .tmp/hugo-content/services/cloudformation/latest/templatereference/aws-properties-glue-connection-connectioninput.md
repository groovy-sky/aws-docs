This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Connection ConnectionInput

A structure that is used to specify a connection to create or update.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AthenaProperties" : Json,
  "AuthenticationConfiguration" : AuthenticationConfigurationInput,
  "ConnectionProperties" : Json,
  "ConnectionType" : String,
  "Description" : String,
  "MatchCriteria" : [ String, ... ],
  "Name" : String,
  "PhysicalConnectionRequirements" : PhysicalConnectionRequirements,
  "PythonProperties" : Json,
  "SparkProperties" : Json,
  "ValidateCredentials" : Boolean,
  "ValidateForComputeEnvironments" : [ String, ... ]
}

```

### YAML

```yaml

  AthenaProperties: Json
  AuthenticationConfiguration:
    AuthenticationConfigurationInput
  ConnectionProperties: Json
  ConnectionType: String
  Description: String
  MatchCriteria:
    - String
  Name: String
  PhysicalConnectionRequirements:
    PhysicalConnectionRequirements
  PythonProperties: Json
  SparkProperties: Json
  ValidateCredentials: Boolean
  ValidateForComputeEnvironments:
    - String

```

## Properties

`AthenaProperties`

Connection properties specific to the Athena compute environment.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationConfiguration`

The authentication properties of the connection.

_Required_: No

_Type_: [AuthenticationConfigurationInput](aws-properties-glue-connection-authenticationconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionProperties`

These key-value pairs define parameters for the connection.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionType`

The type of the connection. Currently, these types are supported:

- `JDBC` \- Designates a connection to a database through Java Database Connectivity (JDBC).

`JDBC` Connections use the following ConnectionParameters.

- Required: All of ( `HOST`, `PORT`, `JDBC_ENGINE`) or `JDBC_CONNECTION_URL`.

- Required: All of ( `USERNAME`, `PASSWORD`) or `SECRET_ID`.

- Optional: `JDBC_ENFORCE_SSL`, `CUSTOM_JDBC_CERT`, `CUSTOM_JDBC_CERT_STRING`, `SKIP_CUSTOM_JDBC_CERT_VALIDATION`. These parameters are used to configure SSL with JDBC.

- `KAFKA` \- Designates a connection to an Apache Kafka streaming platform.

`KAFKA` Connections use the following ConnectionParameters.

- Required: `KAFKA_BOOTSTRAP_SERVERS`.

- Optional: `KAFKA_SSL_ENABLED`, `KAFKA_CUSTOM_CERT`, `KAFKA_SKIP_CUSTOM_CERT_VALIDATION`. These parameters are used to configure SSL with `KAFKA`.

- Optional: `KAFKA_CLIENT_KEYSTORE`, `KAFKA_CLIENT_KEYSTORE_PASSWORD`, `KAFKA_CLIENT_KEY_PASSWORD`, `ENCRYPTED_KAFKA_CLIENT_KEYSTORE_PASSWORD`, `ENCRYPTED_KAFKA_CLIENT_KEY_PASSWORD`. These parameters are used to configure TLS client configuration with SSL in `KAFKA`.

- Optional: `KAFKA_SASL_MECHANISM`. Can be specified as `SCRAM-SHA-512`, `GSSAPI`, or `AWS_MSK_IAM`.

- Optional: `KAFKA_SASL_SCRAM_USERNAME`, `KAFKA_SASL_SCRAM_PASSWORD`, `ENCRYPTED_KAFKA_SASL_SCRAM_PASSWORD`. These parameters are used to configure SASL/SCRAM-SHA-512 authentication with `KAFKA`.

- Optional: `KAFKA_SASL_GSSAPI_KEYTAB`, `KAFKA_SASL_GSSAPI_KRB5_CONF`, `KAFKA_SASL_GSSAPI_SERVICE`, `KAFKA_SASL_GSSAPI_PRINCIPAL`. These parameters are used to configure SASL/GSSAPI authentication with `KAFKA`.

- `MONGODB` \- Designates a connection to a MongoDB document database.

`MONGODB` Connections use the following ConnectionParameters.

- Required: `CONNECTION_URL`.

- Required: All of ( `USERNAME`, `PASSWORD`) or `SECRET_ID`.

- `VIEW_VALIDATION_REDSHIFT` \- Designates a connection used for view validation by Amazon Redshift.

- `VIEW_VALIDATION_ATHENA` \- Designates a connection used for view validation by Amazon Athena.

- `NETWORK` \- Designates a network connection to a data source within an Amazon Virtual Private Cloud environment (Amazon VPC).

`NETWORK` Connections do not require ConnectionParameters. Instead, provide a PhysicalConnectionRequirements.

- `MARKETPLACE` \- Uses configuration settings contained in a connector purchased from AWS Marketplace to read from and write to data stores that are not natively supported by AWS Glue.

`MARKETPLACE` Connections use the following ConnectionParameters.

- Required: `CONNECTOR_TYPE`, `CONNECTOR_URL`, `CONNECTOR_CLASS_NAME`, `CONNECTION_URL`.

- Required for `JDBC` `CONNECTOR_TYPE` connections: All of ( `USERNAME`, `PASSWORD`) or `SECRET_ID`.

- `CUSTOM` \- Uses configuration settings contained in a custom connector to read from and write to data stores that are not natively supported by AWS Glue.

For more information on the connection parameters needed for a particular connector, see the documentation for the connector in [Adding an AWS Glue connection](../../../glue/latest/dg/console-connections.md) in the AWS Glue User Guide.

`SFTP` is not supported.

For more information about how optional ConnectionProperties are used to configure features in AWS Glue, consult [AWS Glue connection properties](../../../glue/latest/dg/connection-defining.md).

For more information about how optional ConnectionProperties are used to configure features in AWS Glue Studio, consult [Using connectors and connections](../../../glue/latest/ug/connectors-chapter.md).

_Required_: Yes

_Type_: String

_Allowed values_: `JDBC | SFTP | MONGODB | KAFKA | NETWORK | MARKETPLACE | CUSTOM | SALESFORCE | VIEW_VALIDATION_REDSHIFT | VIEW_VALIDATION_ATHENA | GOOGLEADS | GOOGLESHEETS | GOOGLEANALYTICS4 | SERVICENOW | MARKETO | SAPODATA | ZENDESK | JIRACLOUD | NETSUITEERP | HUBSPOT | FACEBOOKADS | INSTAGRAMADS | ZOHOCRM | SALESFORCEPARDOT | SALESFORCEMARKETINGCLOUD | ADOBEANALYTICS | SLACK | LINKEDIN | MIXPANEL | ASANA | STRIPE | SMARTSHEET | DATADOG | WOOCOMMERCE | INTERCOM | SNAPCHATADS | PAYPAL | QUICKBOOKS | FACEBOOKPAGEINSIGHTS | FRESHDESK | TWILIO | DOCUSIGNMONITOR | FRESHSALES | ZOOM | GOOGLESEARCHCONSOLE | SALESFORCECOMMERCECLOUD | SAPCONCUR | DYNATRACE | MICROSOFTDYNAMIC365FINANCEANDOPS | MICROSOFTTEAMS | BLACKBAUDRAISEREDGENXT | MAILCHIMP | GITLAB | PENDO | PRODUCTBOARD | CIRCLECI | PIPEDIVE | SENDGRID | AZURECOSMOS | AZURESQL | BIGQUERY | BLACKBAUD | CLOUDERAHIVE | CLOUDERAIMPALA | CLOUDWATCH | CLOUDWATCHMETRICS | CMDB | DATALAKEGEN2 | DB2 | DB2AS400 | DOCUMENTDB | DOMO | DYNAMODB | GOOGLECLOUDSTORAGE | HBASE | KUSTOMER | MICROSOFTDYNAMICS365CRM | MONDAY | MYSQL | OKTA | OPENSEARCH | ORACLE | PIPEDRIVE | POSTGRESQL | SAPHANA | SQLSERVER | SYNAPSE | TERADATA | TERADATANOS | TIMESTREAM | TPCDS | VERTICA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the connection.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchCriteria`

A list of criteria that can be used in selecting this connection.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the connection.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhysicalConnectionRequirements`

The physical connection requirements, such as virtual private cloud (VPC) and `SecurityGroup`, that are needed to successfully make this connection.

_Required_: No

_Type_: [PhysicalConnectionRequirements](aws-properties-glue-connection-physicalconnectionrequirements.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PythonProperties`

Connection properties specific to the Python compute environment.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparkProperties`

Connection properties specific to the Spark compute environment.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidateCredentials`

A flag to validate the credentials during create connection. Default is true.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidateForComputeEnvironments`

The compute environments that the specified connection properties are validated against.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthenticationCredentials

OAuth2ClientApplication

All content copied from https://docs.aws.amazon.com/.
