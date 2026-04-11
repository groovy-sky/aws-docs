This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Scene

Use the `AWS::IoTTwinMaker::Scene` resource to declare a scene.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTTwinMaker::Scene",
  "Properties" : {
      "Capabilities" : [ String, ... ],
      "ContentLocation" : String,
      "Description" : String,
      "SceneId" : String,
      "SceneMetadata" : {Key: Value, ...},
      "Tags" : {Key: Value, ...},
      "WorkspaceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoTTwinMaker::Scene
Properties:
  Capabilities:
    - String
  ContentLocation: String
  Description: String
  SceneId: String
  SceneMetadata:
    Key: Value
  Tags:
    Key: Value
  WorkspaceId: String

```

## Properties

`Capabilities`

A list of capabilities that the scene uses to render.

_Required_: No

_Type_: Array of String

_Minimum_: `0 | 0`

_Maximum_: `256 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContentLocation`

The relative path that specifies the location of the content definition file.

_Required_: Yes

_Type_: String

_Pattern_: `[sS]3://[A-Za-z0-9._/-]+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of this scene.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SceneId`

The ID of the scene.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z_0-9][a-zA-Z_\-0-9]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SceneMetadata`

The scene metadata.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z_\-0-9]+`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The ComponentType tags.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkspaceId`

The ID of the workspace.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z_0-9][a-zA-Z_\-0-9]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the workspace ID and the scene ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The scene ARN.

`CreationDateTime`

The date and time when the scene was created.

`UpdateDateTime`

The scene the update time.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Status

AWS::IoTTwinMaker::SyncJob

All content copied from https://docs.aws.amazon.com/.
