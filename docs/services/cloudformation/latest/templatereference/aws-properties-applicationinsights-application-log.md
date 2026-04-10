This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application Log

The `AWS::ApplicationInsights::Application Log` property type specifies a log to monitor for the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Encoding" : String,
  "LogGroupName" : String,
  "LogPath" : String,
  "LogType" : String,
  "PatternSet" : String
}

```

### YAML

```yaml

  Encoding: String
  LogGroupName: String
  LogPath: String
  LogType: String
  PatternSet: String

```

## Properties

`Encoding`

The type of encoding of the logs to be monitored. The specified encoding should be included in the list of CloudWatch agent supported encodings. If not provided, CloudWatch Application Insights uses the default encoding type for the log type:

- `APPLICATION/DEFAULT`: utf-8 encoding

- `SQL_SERVER`: utf-16 encoding

- `IIS`: ascii encoding

_Required_: No

_Type_: String

_Allowed values_: `utf-8 | utf-16 | ascii`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupName`

The CloudWatch log group name to be associated with the monitored log.

_Required_: No

_Type_: String

_Pattern_: `[\.\-_/#A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogPath`

The path of the logs to be monitored. The log path must be an absolute Windows or Linux system file path. For more information, see [CloudWatch Agent Configuration File: Logs Section](../../../amazoncloudwatch/latest/monitoring/cloudwatch-agent-configuration-file-details.md#CloudWatch-Agent-Configuration-File-Logssection).

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z]:\\[\\\S|*\S]?.*|/[^"']*)$`

_Minimum_: `1`

_Maximum_: `260`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogType`

The log type decides the log patterns against which Application Insights analyzes the
log. The log type is selected from the following: `SQL_SERVER`,
`MYSQL`, `MYSQL_SLOW_QUERY`, `POSTGRESQL`,
`ORACLE_ALERT`, `ORACLE_LISTENER`, `IIS`,
`APPLICATION`, `WINDOWS_EVENTS`,
`WINDOWS_EVENTS_ACTIVE_DIRECTORY`, `WINDOWS_EVENTS_DNS `,
`WINDOWS_EVENTS_IIS `, `WINDOWS_EVENTS_SHAREPOINT`,
`SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP`,
`SQL_SERVER_FAILOVER_CLUSTER_INSTANCE`, `STEP_FUNCTION`,
`API_GATEWAY_ACCESS`, `API_GATEWAY_EXECUTION`,
`SAP_HANA_LOGS`, `SAP_HANA_TRACE`,
`SAP_HANA_HIGH_AVAILABILITY`, and `DEFAULT`.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z][A-Z_]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternSet`

The log pattern set.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9.-_]*`

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JMXPrometheusExporter

LogPattern

All content copied from https://docs.aws.amazon.com/.
