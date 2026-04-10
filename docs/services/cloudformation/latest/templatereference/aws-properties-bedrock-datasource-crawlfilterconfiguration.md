This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource CrawlFilterConfiguration

The configuration of filtering the data source content. For example,
configuring regular expression patterns to include or exclude certain content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PatternObjectFilter" : PatternObjectFilterConfiguration,
  "Type" : String
}

```

### YAML

```yaml

  PatternObjectFilter:
    PatternObjectFilterConfiguration
  Type: String

```

## Properties

`PatternObjectFilter`

The configuration of filtering certain objects or content types of the data source.

_Required_: No

_Type_: [PatternObjectFilterConfiguration](aws-properties-bedrock-datasource-patternobjectfilterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of filtering that you want to apply to certain objects or content of the
data source. For example, the `PATTERN` type is regular expression patterns
you can apply to filter your content.

_Required_: Yes

_Type_: String

_Allowed values_: `PATTERN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContextEnrichmentConfiguration

CustomTransformationConfiguration

All content copied from https://docs.aws.amazon.com/.
