This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSO::Instance

Creates an instance of IAM Identity Center for a standalone AWS account that is not
managed by AWS Organizations or a member AWS account in an organization. You can create only
one instance per account and across all AWS Regions.

The CreateInstance request is rejected if the following apply:

- The instance is created within the organization management account.

- An instance already exists in the same account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SSO::Instance",
  "Properties" : {
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SSO::Instance
Properties:
  Name: String
  Tags:
    - Tag

```

## Properties

`Name`

The name of the Identity Center instance.

_Required_: No

_Type_: String

_Pattern_: `^[\w+=,.@-]+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies tags to be attached to the instance of IAM Identity Center.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-sso-instance-tag.html)

_Maximum_: `75`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a generated ID, combined by all fields with the delimiter
`|`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`IdentityStoreId`

The identifier of the identity store that is connected to the Identity Center
instance.

`InstanceArn`

The ARN of the Identity Center instance under which the operation will be executed.
For more information about ARNs, see [Amazon Resource\
Names (ARNs) and AWS Service Namespaces](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in the _AWS General Reference_.

`OwnerAccountId`

The AWS account ID number of the owner of the Identity Center instance.

`Status`

The current status of this Identity Center instance.

## Examples

### Creating a new instance of IAM Identity Center

The following example creates an instance of IAM Identity Center for a specific
AWS account.

#### JSON

```json

"Instance": {
    "Type": "AWS::SSO::Instance",
    "Properties": {
        "Name": "InstanceExample",
        "Tags": {
            "InstanceTagKey1": "InstanceTagValue1"
        }
    }
}
```

#### YAML

```yaml

 Instance:
    Type: AWS::SSO::Instance
    Properties:
        Name: InstanceExample
        Tags:
            InstanceTagKey1: 'InstanceTagValue1'

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::SSO::Assignment

Tag
