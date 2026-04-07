This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Network

The `AWS::MediaLive::Network` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::Network",
  "Properties" : {
      "IpPools" : [ IpPool, ... ],
      "Name" : String,
      "Routes" : [ Route, ... ],
      "Tags" : [ Tags, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::Network
Properties:
  IpPools:
    - IpPool
  Name: String
  Routes:
    - Route
  Tags:
    - Tags

```

## Properties

`IpPools`

Property description not available.

_Required_: Yes

_Type_: Array of [IpPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-network-ippool.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Routes`

Property description not available.

_Required_: No

_Type_: Array of [Route](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-network-route.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: [Array](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-network-tags.html) of [Tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-network-tags.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN of the network.

`AssociatedClusterIds`

The IDs of the clusters associated with this network.

`Id`

The unique ID of the network.

`State`

The current state of the network.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MultiplexVideoSettings

IpPool
