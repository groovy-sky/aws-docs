This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet SemanticModelConfiguration

Configuration for the semantic model that defines how prepared data is structured for analysis and reporting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TableMap" : {Key: Value, ...}
}

```

### YAML

```yaml

  TableMap:
    Key: Value

```

## Properties

`TableMap`

A map of semantic tables that define the analytical structure.

_Required_: No

_Type_: Object of [SemanticTable](aws-properties-quicksight-dataset-semantictable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SaaSTable

SemanticTable

All content copied from https://docs.aws.amazon.com/.
