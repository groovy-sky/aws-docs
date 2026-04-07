This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::DedicatedIpPool

A request to create a new dedicated IP pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PinpointEmail::DedicatedIpPool",
  "Properties" : {
      "PoolName" : String,
      "Tags" : [ Tags, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::PinpointEmail::DedicatedIpPool
Properties:
  PoolName: String
  Tags:
    - Tags

```

## Properties

`PoolName`

The name of the dedicated IP pool.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An object that defines the tags (keys and values) that you want to associate with the
dedicated IP pool.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-dedicatedippool-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpointemail-dedicatedippool-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myDedicatedIpPool" }`

For the Amazon Pinpoint dedicated IP pool `myDedicatedIpPool`, Ref returns the
name of the IP pool.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SnsDestination

Tags
