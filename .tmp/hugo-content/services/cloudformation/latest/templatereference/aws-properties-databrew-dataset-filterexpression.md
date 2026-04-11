This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset FilterExpression

Represents a structure for defining parameter conditions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "ValuesMap" : [ FilterValue, ... ]
}

```

### YAML

```yaml

  Expression: String
  ValuesMap:
    - FilterValue

```

## Properties

`Expression`

The expression which includes condition names followed by substitution variables,
possibly grouped and combined with other conditions. For example, "(starts\_with :prefix1
or starts\_with :prefix2) and (ends\_with :suffix1 or ends\_with :suffix2)". Substitution
variables should start with ':' symbol.

_Required_: Yes

_Type_: String

_Pattern_: `^[><0-9A-Za-z_.,:)(!= ]+$`

_Minimum_: `4`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValuesMap`

The map of substitution variable names to their values used in this filter
expression.

_Required_: Yes

_Type_: Array of [FilterValue](aws-properties-databrew-dataset-filtervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilesLimit

FilterValue

All content copied from https://docs.aws.amazon.com/.
