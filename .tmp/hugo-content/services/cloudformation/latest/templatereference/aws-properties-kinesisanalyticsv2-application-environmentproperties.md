This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application EnvironmentProperties

Describes execution properties for a Managed Service for Apache Flink application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PropertyGroups" : [ PropertyGroup, ... ]
}

```

### YAML

```yaml

  PropertyGroups:
    - PropertyGroup

```

## Properties

`PropertyGroups`

Describes the execution property groups.

_Required_: No

_Type_: Array of [PropertyGroup](aws-properties-kinesisanalyticsv2-application-propertygroup.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [EnvironmentProperties](../../../managed-flink/latest/apiv2/api-environmentproperties.md) in the _Amazon Kinesis Data_
_Analytics API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeployAsApplicationConfiguration

FlinkApplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
