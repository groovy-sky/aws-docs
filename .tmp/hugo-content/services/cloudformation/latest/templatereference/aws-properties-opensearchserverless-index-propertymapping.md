This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchServerless::Index PropertyMapping

Property mappings for the OpenSearch Serverless index.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimension" : Integer,
  "Index" : Boolean,
  "Method" : Method,
  "Properties" : {Key: Value, ...},
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Dimension: Integer
  Index: Boolean
  Method:
    Method
  Properties:
    Key: Value
  Type: String
  Value: String

```

## Properties

`Dimension`

Dimension size for vector fields, defines the number of dimensions in the vector.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Index`

Whether a field should be indexed.

_Required_: No

_Type_: [Boolean](aws-properties-opensearchserverless-index-index.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Method`

Configuration for k-NN search method.

_Required_: No

_Type_: [Method](aws-properties-opensearchserverless-index-method.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

Defines the fields within the mapping, including their types and configurations.

_Required_: No

_Type_: Object of [PropertyMapping](aws-properties-opensearchserverless-index-propertymapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The field data type. Must be a valid OpenSearch field type.

_Required_: Yes

_Type_: String

_Allowed values_: `text | knn_vector`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Default value for the field when not specified in a document.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Parameters

AWS::OpenSearchServerless::LifecyclePolicy

All content copied from https://docs.aws.amazon.com/.
