This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job StatisticOverride

Override of a particular evaluation for a profile job.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Parameters" : {Key: Value, ...},
  "Statistic" : String
}

```

### YAML

```yaml

  Parameters:
    Key: Value
  Statistic: String

```

## Properties

`Parameters`

A map that includes overrides of an evaluation’s parameters.

_Required_: Yes

_Type_: Object of String

_Pattern_: `^[A-Za-z0-9]{1,128}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The name of an evaluation

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z\_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3TableOutputOptions

StatisticsConfiguration

All content copied from https://docs.aws.amazon.com/.
