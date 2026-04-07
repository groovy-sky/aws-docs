This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::UserHierarchyGroup

Specifies a new user hierarchy group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::UserHierarchyGroup",
  "Properties" : {
      "InstanceArn" : String,
      "Name" : String,
      "ParentGroupArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Connect::UserHierarchyGroup
Properties:
  InstanceArn: String
  Name: String
  ParentGroupArn: String
  Tags:
    - Tag

```

## Properties

`InstanceArn`

The Amazon Resource Name (ARN) of the user hierarchy group.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the user hierarchy group.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParentGroupArn`

The Amazon Resource Name (ARN) of the parent group.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent-group/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-userhierarchygroup-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the user hierarchy group. For example:

`{ "Ref": "myUserHierarchyGroup" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`UserHierarchyGroupArn`

The Amazon Resource Name (ARN) of the user hierarchy group.

## Examples

### Specify a user hierarchy group

The following example specifies a user hierarchy group for an Amazon Connect instance. This example specifies a user hierachy group for the
instance on top of existing HierarchyStructure

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a user hierarchy group for an Amazon Connect instance
Resources:
    HierarchyGroup:
    Type: 'AWS::Connect::UserHierarchyGroup'
    Properties:
      Name: 'exampleUserHierarchyGroup'
      InstanceArn: 'arn:aws:connect:region-name:aws-account-id:instance/instance-arn'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VoiceEnhancementConfig

Tag
