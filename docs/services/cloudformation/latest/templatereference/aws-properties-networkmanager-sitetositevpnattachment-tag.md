This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkManager::SiteToSiteVpnAttachment Tag

Describes a tag.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The tag key.

Constraints: Maximum length of 128 characters.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `10000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value.

Constraints: Maximum length of 256 characters.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `10000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProposedSegmentChange

AWS::NetworkManager::TransitGatewayPeering
