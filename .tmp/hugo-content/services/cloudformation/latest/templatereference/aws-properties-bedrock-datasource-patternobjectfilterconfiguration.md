This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource PatternObjectFilterConfiguration

The configuration of filtering certain objects or content types of the data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filters" : [ PatternObjectFilter, ... ]
}

```

### YAML

```yaml

  Filters:
    - PatternObjectFilter

```

## Properties

`Filters`

The configuration of specific filters applied to your data source content. You can
filter out or include certain content.

_Required_: Yes

_Type_: Array of [PatternObjectFilter](aws-properties-bedrock-datasource-patternobjectfilter.md)

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PatternObjectFilter

S3DataSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
