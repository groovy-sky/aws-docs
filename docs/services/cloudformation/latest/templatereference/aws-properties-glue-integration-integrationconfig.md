This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Integration IntegrationConfig

Properties associated with the integration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContinuousSync" : Boolean,
  "RefreshInterval" : String,
  "SourceProperties" : {Key: Value, ...}
}

```

### YAML

```yaml

  ContinuousSync: Boolean
  RefreshInterval: String
  SourceProperties:
    Key: Value

```

## Properties

`ContinuousSync`

Enables continuous synchronization for on-demand data extractions from SaaS applications to AWS data services like Amazon Redshift
and Amazon S3.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RefreshInterval`

Specifies the frequency at which CDC (Change Data Capture) pulls or incremental loads should occur. This parameter provides flexibility to align
the refresh rate with your specific data update patterns, system load considerations, and performance optimization goals. Time increment can be set from
15 minutes to 8640 minutes (six days).

_Required_: No

_Type_: String

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceProperties`

A collection of key-value pairs that specify additional properties for the integration source. These properties provide configuration options that
can be used to customize the behavior of the ODB source during data integration operations.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\s\S]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Glue::Integration

Tag
