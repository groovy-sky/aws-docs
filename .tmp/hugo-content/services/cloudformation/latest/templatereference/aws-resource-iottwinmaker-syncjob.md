This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::SyncJob

The SyncJob.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTTwinMaker::SyncJob",
  "Properties" : {
      "SyncRole" : String,
      "SyncSource" : String,
      "Tags" : {Key: Value, ...},
      "WorkspaceId" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoTTwinMaker::SyncJob
Properties:
  SyncRole: String
  SyncSource: String
  Tags:
    Key: Value
  WorkspaceId: String

```

## Properties

`SyncRole`

The SyncJob IAM role. This IAM role is used by the sync job to read from the syncSource,
and create, update or delete the corresponding resources.

_Required_: Yes

_Type_: String

_Pattern_: `arn:((aws)|(aws-cn)|(aws-us-gov)):iam::[0-9]{12}:role/.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SyncSource`

The sync source.

###### Note

Currently the only supported syncSoucre is `SITEWISE`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata you can use to manage the SyncJob.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkspaceId`

The ID of the workspace that contains the sync job.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z_0-9][a-zA-Z_\-0-9]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the workspace Id and the entity Id.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The SyncJob ARN.

`CreationDateTime`

The creation date and time of the SyncJob.

`State`

The SyncJob's state.

`UpdateDateTime`

The update date and time.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTTwinMaker::Scene

AWS::IoTTwinMaker::Workspace

All content copied from https://docs.aws.amazon.com/.
