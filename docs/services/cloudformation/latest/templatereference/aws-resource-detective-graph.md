This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Detective::Graph

The `AWS::Detective::Graph` resource is an Amazon Detective resource type
that creates a Detective behavior graph. The requesting account becomes the
administrator account for the behavior graph.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Detective::Graph",
  "Properties" : {
      "AutoEnableMembers" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Detective::Graph
Properties:
  AutoEnableMembers: Boolean
  Tags:
    - Tag

```

## Properties

`AutoEnableMembers`

Indicates whether to automatically enable new organization accounts as member accounts in the organization behavior graph.

By default, this property is set to `false`. If you want to change the value of this property, you must be the
Detective administrator for the organization. For more information on setting a Detective administrator account,
see [AWS::Detective::OrganizationAdmin](../userguide/aws-resource-detective-organizationadmin.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tag values to assign to the new behavior graph.

_Required_: No

_Type_: Array of [Tag](aws-properties-detective-graph-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the new behavior graph.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the new behavior graph.

## Examples

### Creating a new Detective behavior graph

This example shows how to declare a new `AWS:Detective:Graph` resource to create a new Detective behavior graph

#### JSON

```json

"NewGraph": {
    "Type": "AWS::Detective::Graph"
    "Properties": {
        "Tags": [
            {
                "Key": "Tag1Name",
                "Value": "Tag1Value"
            },
            {
                "Key": "Tag2Name".
                "Value": "Tag2alue"
            }
        ]
    }
}
```

#### YAML

```yaml

NewGraph:
    Type: AWS::Detective::Graph
    Properties:
        Tags:
            - Key: Tag1Name
              Value: Tag1Value
            - Key: Tag2Name
              Value: Tag2Value
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Detective

Tag

All content copied from https://docs.aws.amazon.com/.
