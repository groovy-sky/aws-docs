This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::UserHierarchyStructure

Contains information about a hierarchy structure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::UserHierarchyStructure",
  "Properties" : {
      "InstanceArn" : String,
      "UserHierarchyStructure" : UserHierarchyStructure
    }
}

```

### YAML

```yaml

Type: AWS::Connect::UserHierarchyStructure
Properties:
  InstanceArn: String
  UserHierarchyStructure:
    UserHierarchyStructure

```

## Properties

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserHierarchyStructure`

Contains information about a hierarchy structure.

_Required_: No

_Type_: [UserHierarchyStructure](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-userhierarchystructure-userhierarchystructure.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the user hierarchy structure. For example:

`{ "Ref": "myhierarchystructure" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`UserHierarchyStructureArn`

The identifier for the user hierarchy structure.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

LevelFive
