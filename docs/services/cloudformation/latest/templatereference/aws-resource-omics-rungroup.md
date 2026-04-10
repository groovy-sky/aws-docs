This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::RunGroup

Creates a run group to limit the compute resources for the runs that are added to the group.
Returns an ARN, ID, and tags for the run group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::RunGroup",
  "Properties" : {
      "MaxCpus" : Number,
      "MaxDuration" : Number,
      "MaxGpus" : Number,
      "MaxRuns" : Number,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Omics::RunGroup
Properties:
  MaxCpus: Number
  MaxDuration: Number
  MaxGpus: Number
  MaxRuns: Number
  Name: String
  Tags:
    Key: Value

```

## Properties

`MaxCpus`

The group's maximum CPU count setting.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxDuration`

The group's maximum duration setting in minutes.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxGpus`

The maximum GPUs that can be used by a run group.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxRuns`

The group's maximum concurrent run setting.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The group's name.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags for the group.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the details of this resource. For example:

`{ "Ref": "RunGroup.CreationTime" }` `Ref` returns the timestamp for a run group.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The run group's ARN.

`CreationTime`

When the run group was created.

`Id`

The run group's ID.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SseConfig

AWS::Omics::SequenceStore

All content copied from https://docs.aws.amazon.com/.
