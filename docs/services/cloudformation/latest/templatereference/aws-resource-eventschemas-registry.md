This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EventSchemas::Registry

Use the `AWS::EventSchemas::Registry` to specify a schema registry. Schema
registries are containers for Schemas. Registries collect and organize schemas so that
your schemas are in logical groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EventSchemas::Registry",
  "Properties" : {
      "Description" : String,
      "RegistryName" : String,
      "Tags" : [ TagsEntry, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EventSchemas::Registry
Properties:
  Description: String
  RegistryName: String
  Tags:
    - TagsEntry

```

## Properties

`Description`

A description of the registry to be created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegistryName`

The name of the schema registry.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags to associate with the registry.

_Required_: No

_Type_: Array of [TagsEntry](aws-properties-eventschemas-registry-tagsentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic
function, `Ref` returns the ARN of the schema. For example:

`{ "Ref": "MyRegistry" }`

Returns a value similar to the following:

`arn:aws:schemas:us-east-1:012345678901:registry/MyRegistry`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RegistryArn`

The ARN of the registry.

`RegistryName`

The name of the registry.

## Examples

### Create a Schema Registry for Events Emitted by AWS Step Functions

#### YAML

```yaml

Resources:
  StatesSchemasRegistry:
    Type: AWS::EventSchemas::Registry
    Properties:
      RegistryName: 'aws.states'
      Description: 'Contains the schemas of events emitted by AWS Step Functions'

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagsEntry

TagsEntry

All content copied from https://docs.aws.amazon.com/.
