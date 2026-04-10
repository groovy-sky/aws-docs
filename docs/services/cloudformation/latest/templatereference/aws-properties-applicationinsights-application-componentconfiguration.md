This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application ComponentConfiguration

The `AWS::ApplicationInsights::Application ComponentConfiguration` property type defines the configuration settings of the component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigurationDetails" : ConfigurationDetails,
  "SubComponentTypeConfigurations" : [ SubComponentTypeConfiguration, ... ]
}

```

### YAML

```yaml

  ConfigurationDetails:
    ConfigurationDetails
  SubComponentTypeConfigurations:
    - SubComponentTypeConfiguration

```

## Properties

`ConfigurationDetails`

The configuration settings.

_Required_: No

_Type_: [ConfigurationDetails](aws-properties-applicationinsights-application-configurationdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubComponentTypeConfigurations`

Sub-component configurations of the component.

_Required_: No

_Type_: Array of [SubComponentTypeConfiguration](aws-properties-applicationinsights-application-subcomponenttypeconfiguration.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AlarmMetric

ComponentMonitoringSetting

All content copied from https://docs.aws.amazon.com/.
