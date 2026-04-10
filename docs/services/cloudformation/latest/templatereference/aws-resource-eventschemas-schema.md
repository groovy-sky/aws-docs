This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EventSchemas::Schema

Use the `AWS::EventSchemas::Schema` resource to specify an event
schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EventSchemas::Schema",
  "Properties" : {
      "Content" : String,
      "Description" : String,
      "RegistryName" : String,
      "SchemaName" : String,
      "Tags" : [ TagsEntry, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::EventSchemas::Schema
Properties:
  Content: String
  Description: String
  RegistryName: String
  SchemaName: String
  Tags:
    - TagsEntry
  Type: String

```

## Properties

`Content`

The source of the schema definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the schema.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegistryName`

The name of the schema registry.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SchemaName`

The name of the schema.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags associated with the schema.

_Required_: No

_Type_: Array of [TagsEntry](aws-properties-eventschemas-schema-tagsentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of schema.

Valid types include `OpenApi3` and `JSONSchemaDraft4`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic
function, `Ref` returns the ARN of the schema. For example:

`{ "Ref": "MySchema" }`

Returns a value similar to the following:

`arn:aws:schemas:us-east-1:012345678901:schema/MyRegistry/MySchema`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`LastModified`

The date and time that schema was modified.

`SchemaArn`

The ARN of the schema.

`SchemaName`

The name of the schema.

`SchemaVersion`

The version number of the schema.

`VersionCreatedDate`

The date the schema version was created.

## Examples

#### YAML

```yaml

Resources:
  ExecutionStatusChangeSchema:
    Type: AWS::EventSchemas::Schema
    Properties:
      RegistryName: 'aws.events'
      SchemaName: ExecutionStatusChange
      Description: 'event emitted when the status of a state machine execution change'
      Type: OpenApi3
      Content: >
        {
          "openapi": "3.0.0",
          "info": {
            "version": "1.0.0",
            "title": "StepFunctionsExecutionStatusChange"
          },
          "paths":{},
          "components": {
            "schemas": {
              "StepFunctionsExecutionStatusChange": {
                "type": "object",
                "required": [ "output", "input", "executionArn", "name", "stateMachineArn", "startDate", "stopDate", "status" ],
                "properties": {
                  "output": {"type": "string","nullable": true},
                  "input": {"type": "string"},
                  "executionArn": {"type": "string"},
                  "name": {"type": "string"},
                  "stateMachineArn": {"type": "string"},
                  "startDate": {"type": "integer","format": "int64"},
                  "stopDate": {"type": "integer","format": "int64","nullable": true},
                  "status": {"type": "string","enum": [ "FAILED", "RUNNING", "SUCCEEDED", "ABORTED" ]}
                }
              }
            }
          }
        }

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EventSchemas::RegistryPolicy

TagsEntry

All content copied from https://docs.aws.amazon.com/.
