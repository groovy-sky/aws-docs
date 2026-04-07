This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Inspector::ResourceGroup

The `AWS::Inspector::ResourceGroup` resource is used to create Amazon
Inspector resource groups. A resource group defines a set of tags that, when queried,
identify the AWS resources that make up the assessment target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Inspector::ResourceGroup",
  "Properties" : {
      "ResourceGroupTags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Inspector::ResourceGroup
Properties:
  ResourceGroupTags:
    - Tag

```

## Properties

`ResourceGroupTags`

The tags (key and value pairs) that will be associated with the resource
group.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: Yes

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-inspector-resourcegroup-tag.html)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the new resource group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) that specifies the resource group that is created.

## Examples

### Declaring an Amazon Inspector Resource Group Resource

The following example shows how to declare an
`AWS::Inspector::ResourceGroup` resource to create an Amazon Inspector
resource group.

#### JSON

```json

{
    "Type": "AWS::Inspector::ResourceGroup",
    "Properties": {
        "ResourceGroupTags": [{
            "Key": "Name",
            "Value": "example"
        }]
    }
}
```

#### YAML

```yaml

Type: AWS::Inspector::ResourceGroup
Properties:
  ResourceGroupTags:
    - Key: Name
      Value: example
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
