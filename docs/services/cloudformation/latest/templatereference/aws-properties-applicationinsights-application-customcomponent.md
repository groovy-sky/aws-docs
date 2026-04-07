This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application CustomComponent

The `AWS::ApplicationInsights::Application CustomComponent` property type describes a custom component by grouping similar standalone instances to monitor.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentName" : String,
  "ResourceList" : [ String, ... ]
}

```

### YAML

```yaml

  ComponentName: String
  ResourceList:
    - String

```

## Properties

`ComponentName`

The name of the component.

_Required_: Yes

_Type_: String

_Pattern_: `^[\d\w\-_.+]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceList`

The list of resource ARNs that belong to the component.

_Required_: Yes

_Type_: Array of String

_Maximum_: `300`

_Minimum_: `20 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConfigurationDetails

HAClusterPrometheusExporter
