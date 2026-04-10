This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::GlobalNetwork

Creates a new, empty global network.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::GlobalNetwork",
  "Properties" : {
      "CreatedAt" : String,
      "Description" : String,
      "State" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::GlobalNetwork
Properties:
  CreatedAt: String
  Description: String
  State: String
  Tags:
    - Tag

```

## Properties

`CreatedAt`

The date and time that the global network was created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the global network.

Constraints: Maximum length of 256 characters.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state of the global network.

_Required_: No

_Type_: String

_Allowed values_: `PENDING | AVAILABLE | DELETING | UPDATING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the global network.

_Required_: No

_Type_: Array of [Tag](aws-properties-networkmanager-globalnetwork-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the global network. For example: `global-network-01231231231231231`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the global network. For example,
`arn:aws:networkmanager::123456789012:global-network/global-network-01231231231231231`.

`Id`

The ID of the global network. For example,
`global-network-01231231231231231`.

## Examples

### Global Network

The following example creates a global network.

#### JSON

```json

{
    "Type": "AWS::NetworkManager::GlobalNetwork",
    "Properties": {
        "Description": "Global network for USA sites",
        "Tags": [
            {
                "Key": "Name",
                "Value": "global-network-usa"
            }
        ]
    }
}
```

#### YAML

```yaml

Type: 'AWS::NetworkManager::GlobalNetwork'
Properties:
  Description: Global network for USA sites
  Tags:
    - Key: Name
      Value: global-network-usa
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
