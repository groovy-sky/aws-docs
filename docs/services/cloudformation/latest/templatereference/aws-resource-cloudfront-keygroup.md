This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::KeyGroup

A key group.

A key group contains a list of public keys that you can use with [CloudFront signed URLs and signed cookies](../../../amazoncloudfront/latest/developerguide/privatecontent.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::KeyGroup",
  "Properties" : {
      "KeyGroupConfig" : KeyGroupConfig
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::KeyGroup
Properties:
  KeyGroupConfig:
    KeyGroupConfig

```

## Properties

`KeyGroupConfig`

The key group configuration.

_Required_: Yes

_Type_: [KeyGroupConfig](aws-properties-cloudfront-keygroup-keygroupconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the key group. For example:
`e9fcd3cf-f3f4-4b61-bd85-9ba9e091b309`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The identifier for the key group.

`LastModifiedTime`

The date and time when the key group was last modified.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyValueStoreAssociation

KeyGroupConfig

All content copied from https://docs.aws.amazon.com/.
