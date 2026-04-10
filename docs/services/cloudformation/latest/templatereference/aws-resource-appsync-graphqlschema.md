This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::GraphQLSchema

The `AWS::AppSync::GraphQLSchema` resource is used for your AWS AppSync GraphQL schema that controls the data model for your API. Schema
files are text written in Schema Definition Language (SDL) format. For more information
about schema authoring, see [Designing a GraphQL\
API](../../../appsync/latest/devguide/designing-a-graphql-api.md) in the _AWS AppSync Developer_
_Guide_.

###### Note

When you submit an update, AWS CloudFormation updates resources based on
differences between what you submit and the stack's current template. To cause this
resource to be updated you must change a property value for this resource in the
CloudFormation template. Changing the Amazon S3 file
content without changing a property value will not result in an update
operation.

See [Update Behaviors of Stack Resources](../userguide/using-cfn-updating-stacks-update-behaviors.md) in the _AWS CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::GraphQLSchema",
  "Properties" : {
      "ApiId" : String,
      "Definition" : String,
      "DefinitionS3Location" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::GraphQLSchema
Properties:
  ApiId: String
  Definition: String
  DefinitionS3Location: String

```

## Properties

`ApiId`

The AWS AppSync GraphQL API identifier to which you want to apply this
schema.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Definition`

The text representation of a GraphQL schema in SDL format.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionS3Location`

The location of a GraphQL schema file in an Amazon S3 bucket. Use this if
you want to provision with the schema living in Amazon S3 rather than
embedding it in your CloudFormation template.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::AppSync::GraphQLSchema` resource
to the intrinsic `Ref` function, the function returns the GraphQL API ID with
the literal String GraphQLSchema attached to it.

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`Id`

The ID value.

## Examples

### GraphQL Schema Creation Example

The following example creates a GraphQL Schema and associates it with an
existing GraphQL API by passing the GraphQL API ID as a parameter.

#### YAML

```yaml

Parameters:
  graphQlApiId:
    Type: String
  graphQlSchemaS3DescriptionLocation:
    Type: String
Resources:
  Schema:
    Type: AWS::AppSync::GraphQLSchema
    Properties:
      ApiId:
	Ref: graphQlApiId
      DefinitionS3Location:
        Ref: graphQlSchemaS3DescriptionLocation

```

#### JSON

```json

{
  "Parameters": {
    "graphQlApiId": {
      "Type": "String"
    },
    "graphQlSchemaS3DescriptionLocation": {
      "Type": "String"
    }
  },
  "Resources": {
    "Schema": {
      "Type": "AWS::AppSync::GraphQLSchema",
      "Properties": {
        "ApiId": {
          "Ref": "graphQlApiId"
        },
        "DefinitionS3Location": {
          "Ref": "graphQlSchemaS3DescriptionLocation"
        }
      }
    }
  }
}

```

## See also

- [StartSchemaCreation](../../../../reference/appsync/latest/apireference/api-startschemacreation.md) operation in the _AWS AppSync API Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserPoolConfig

AWS::AppSync::Resolver

All content copied from https://docs.aws.amazon.com/.
