This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RUM::AppMonitor JavaScriptSourceMaps

A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Uri" : String,
  "Status" : String
}

```

### YAML

```yaml

  S3Uri: String
  Status: String

```

## Properties

`S3Uri`

The S3Uri of the bucket or folder that stores the source map files. It is required if status is ENABLED.

_Required_: No

_Type_: String

_Pattern_: `^s3://[a-z0-9][-.a-z0-9]{1,62}(?:/[-!_*'().a-z0-9A-Z]+(?:/[-!_*'().a-z0-9A-Z]+)*)?/?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Specifies whether JavaScript error stack traces should be unminified for this app monitor. The default is for JavaScript error stack trace unminification to be `DISABLED`.

_Required_: Yes

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeobfuscationConfiguration

MetricDefinition

All content copied from https://docs.aws.amazon.com/.
