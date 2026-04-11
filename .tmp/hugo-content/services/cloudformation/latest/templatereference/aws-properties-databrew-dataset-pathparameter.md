This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset PathParameter

Represents a single entry in the path parameters of a dataset. Each
`PathParameter` consists of a name and a parameter definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatasetParameter" : DatasetParameter,
  "PathParameterName" : String
}

```

### YAML

```yaml

  DatasetParameter:
    DatasetParameter
  PathParameterName: String

```

## Properties

`DatasetParameter`

The path parameter definition.

_Required_: Yes

_Type_: [DatasetParameter](aws-properties-databrew-dataset-datasetparameter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PathParameterName`

The name of the path parameter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PathOptions

S3Location

All content copied from https://docs.aws.amazon.com/.
