This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application LogPatternSet

The `AWS::ApplicationInsights::Application LogPatternSet` property type specifies the log pattern set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogPatterns" : [ LogPattern, ... ],
  "PatternSetName" : String
}

```

### YAML

```yaml

  LogPatterns:
    - LogPattern
  PatternSetName: String

```

## Properties

`LogPatterns`

A list of objects that define the log patterns that belong to `LogPatternSet`.

_Required_: Yes

_Type_: Array of [LogPattern](aws-properties-applicationinsights-application-logpattern.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternSetName`

The name of the log pattern. A log pattern name can contain up to 30 characters, and
it cannot be empty. The characters can be Unicode letters, digits, or one of the
following symbols: period, dash, underscore.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9.-_]*`

_Minimum_: `1`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogPattern

NetWeaverPrometheusExporter

All content copied from https://docs.aws.amazon.com/.
