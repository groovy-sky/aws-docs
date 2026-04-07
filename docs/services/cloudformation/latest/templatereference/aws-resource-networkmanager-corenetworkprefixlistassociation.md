This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::CoreNetworkPrefixListAssociation

Creates an association between a core network and a prefix list for routing control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::NetworkManager::CoreNetworkPrefixListAssociation",
  "Properties" : {
      "CoreNetworkId" : String,
      "PrefixListAlias" : String,
      "PrefixListArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::NetworkManager::CoreNetworkPrefixListAssociation
Properties:
  CoreNetworkId: String
  PrefixListAlias: String
  PrefixListArn: String

```

## Properties

`CoreNetworkId`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrefixListAlias`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrefixListArn`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-]+:ec2:[a-z]+-[a-z]+-[0-9]:([0-9]{12}):prefix-list/pl-[a-z0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::NetworkManager::CustomerGatewayAssociation
