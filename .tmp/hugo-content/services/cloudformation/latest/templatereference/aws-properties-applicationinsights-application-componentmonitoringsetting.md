This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application ComponentMonitoringSetting

The `AWS::ApplicationInsights::Application ComponentMonitoringSetting` property type defines the monitoring setting of the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentARN" : String,
  "ComponentConfigurationMode" : String,
  "ComponentName" : String,
  "CustomComponentConfiguration" : ComponentConfiguration,
  "DefaultOverwriteComponentConfiguration" : ComponentConfiguration,
  "Tier" : String
}

```

### YAML

```yaml

  ComponentARN: String
  ComponentConfigurationMode: String
  ComponentName: String
  CustomComponentConfiguration:
    ComponentConfiguration
  DefaultOverwriteComponentConfiguration:
    ComponentConfiguration
  Tier: String

```

## Properties

`ComponentARN`

The ARN of the component. Either the component ARN or the component name is required.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-[\w]+)*:[\w\d-]+:([\w\d-]*)?:[\w\d_-]*([:/].+)*$`

_Minimum_: `20`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentConfigurationMode`

Component monitoring can be configured in one of the following three modes:

- `DEFAULT`: The component will be configured with the recommended default monitoring settings of the selected `Tier`.

- `CUSTOM`: The component will be configured with the customized monitoring settings
that are specified in `CustomComponentConfiguration`. If used,
`CustomComponentConfiguration` must be provided.

- `DEFAULT_WITH_OVERWRITE`: The component will be configured with the recommended
default monitoring settings of the selected `Tier`, and merged with
customized overwrite settings that are specified in
`DefaultOverwriteComponentConfiguration`. If used,
`DefaultOverwriteComponentConfiguration` must be provided.

_Required_: Yes

_Type_: String

_Allowed values_: `DEFAULT | DEFAULT_WITH_OVERWRITE | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComponentName`

The name of the component. Either the component ARN or the component name is required.

_Required_: No

_Type_: String

_Pattern_: `^[\d\w\-_.+]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomComponentConfiguration`

Customized monitoring settings. Required if CUSTOM mode is configured in `ComponentConfigurationMode`.

_Required_: No

_Type_: [ComponentConfiguration](aws-properties-applicationinsights-application-componentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultOverwriteComponentConfiguration`

Customized overwrite monitoring settings. Required if CUSTOM mode is configured in `ComponentConfigurationMode`.

_Required_: No

_Type_: [ComponentConfiguration](aws-properties-applicationinsights-application-componentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

The tier of the application component. Supported tiers include
`DOT_NET_CORE`, `DOT_NET_WORKER`, `DOT_NET_WEB`,
`SQL_SERVER`, `SQL_SERVER_ALWAYSON_AVAILABILITY_GROUP`,
`SQL_SERVER_FAILOVER_CLUSTER_INSTANCE`, `MYSQL`,
`POSTGRESQL`, `JAVA_JMX`, `ORACLE`,
`SAP_HANA_MULTI_NODE`, `SAP_HANA_SINGLE_NODE`,
`SAP_HANA_HIGH_AVAILABILITY`, `SHAREPOINT`.
`ACTIVE_DIRECTORY`, and `DEFAULT`.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z][A-Z_]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ComponentConfiguration

ConfigurationDetails

All content copied from https://docs.aws.amazon.com/.
