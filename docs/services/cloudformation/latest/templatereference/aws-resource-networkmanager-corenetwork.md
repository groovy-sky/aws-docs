This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::CoreNetwork

Describes a core network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::CoreNetwork",
  "Properties" : {
      "Description" : String,
      "GlobalNetworkId" : String,
      "PolicyDocument" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::CoreNetwork
Properties:
  Description: String
  GlobalNetworkId: String
  PolicyDocument: Json
  Tags:
    - Tag

```

## Properties

`Description`

The description of a core network.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalNetworkId`

The ID of the global network that your core network is a part of.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyDocument`

Describes a core network policy. For more information, see [Core network policies](https://docs.aws.amazon.com/network-manager/latest/cloudwan/cloudwan-policy-change-sets.html).

If you update the policy document, CloudFormation will apply the core network change set generated from the updated policy document, and then set it as the LIVE policy.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of key-value tags associated with a core network.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-networkmanager-corenetwork-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the core network. For example, `{ "Ref: "core-network-060ea2740fe60fd38" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CoreNetworkArn`

The ARN of the core network.

`CoreNetworkId`

The ID of the core network.

`CreatedAt`

The timestamp when the core network was created.

`Edges`

The edges.

`NetworkFunctionGroups`

The network function groups associated with a core network.

`OwnerAccount`

The owner of the core network.

`Segments`

The segments.

`State`

The current state of the core network. These states are: `CREATING` \| `UPDATING` \| `AVAILABLE` \| `DELETING`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CoreNetworkEdge
