This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::Schema

Creates an Amazon Personalize schema from the specified schema string. The schema you create
must be in Avro JSON format.

Amazon Personalize recognizes three schema variants. Each schema is associated with a dataset
type and has a set of required field and keywords. If you are creating a schema for a dataset in a Domain dataset group, you
provide the domain of the Domain dataset group.
You specify a schema when you call [CreateDataset](../../../personalize/latest/dg/api-createdataset.md).

For more information on schemas, see
[Datasets and schemas](../../../personalize/latest/dg/how-it-works-dataset-schema.md).

###### Related APIs

- [ListSchemas](../../../personalize/latest/dg/api-listschemas.md)

- [DescribeSchema](../../../personalize/latest/dg/api-describeschema.md)

- [DeleteSchema](../../../personalize/latest/dg/api-deleteschema.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Personalize::Schema",
  "Properties" : {
      "Domain" : String,
      "Name" : String,
      "Schema" : String
    }
}

```

### YAML

```yaml

Type: AWS::Personalize::Schema
Properties:
  Domain: String
  Name: String
  Schema: String

```

## Properties

`Domain`

The domain of a schema that you created for a dataset in a Domain dataset group.

_Required_: No

_Type_: String

_Allowed values_: `ECOMMERCE | VIDEO_ON_DEMAND`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the schema.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schema`

The schema.

_Required_: Yes

_Type_: String

_Maximum_: `10000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the resource.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`SchemaArn`

The Amazon Resource Name (ARN) of the schema.

## Examples

### Creating a schema

The following example creates an Amazon Personalize schema for an Interactions dataset.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "MySchema": {
            "Type": "AWS::Personalize::Schema",
            "Properties": {
                "Name": "my-schema-name",
                "Schema": "{\"type\": \"record\",\"name\": \"Interactions\", \"namespace\": \"com.amazonaws.personalize.schema\", \"fields\": [ { \"name\": \"USER_ID\", \"type\": \"string\" }, { \"name\": \"ITEM_ID\", \"type\": \"string\" }, { \"name\": \"TIMESTAMP\", \"type\": \"long\"}], \"version\": \"1.0\"}"
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MySchema:
    Type: AWS::Personalize::Schema
    Properties:
      Name: "my-schema-name"
      Schema: >-
        {"type": "record","name": "Interactions", "namespace":
        "com.amazonaws.personalize.schema", "fields": [ { "name": "USER_ID",
        "type": "string" }, { "name": "ITEM_ID", "type": "string" }, { "name":
        "TIMESTAMP", "type": "long"}], "version": "1.0"}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Personalize::DatasetGroup

AWS::Personalize::Solution

All content copied from https://docs.aws.amazon.com/.
