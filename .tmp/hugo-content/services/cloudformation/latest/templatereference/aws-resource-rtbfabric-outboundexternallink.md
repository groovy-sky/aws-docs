This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::OutboundExternalLink

The `AWS::RTBFabric::OutboundExternalLink` resource Property description not available. for RTBFabric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RTBFabric::OutboundExternalLink",
  "Properties" : {
      "GatewayId" : String,
      "LinkAttributes" : LinkAttributes,
      "LinkLogSettings" : LinkLogSettings,
      "PublicEndpoint" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RTBFabric::OutboundExternalLink
Properties:
  GatewayId: String
  LinkAttributes:
    LinkAttributes
  LinkLogSettings:
    LinkLogSettings
  PublicEndpoint: String
  Tags:
    - Tag

```

## Properties

`GatewayId`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^rtb-gw-[a-z0-9-]{1,25}$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LinkAttributes`

Property description not available.

_Required_: No

_Type_: [LinkAttributes](aws-properties-rtbfabric-outboundexternallink-linkattributes.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LinkLogSettings`

Property description not available.

_Required_: Yes

_Type_: [LinkLogSettings](aws-properties-rtbfabric-outboundexternallink-linklogsettings.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PublicEndpoint`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^(https|http)://.+$`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-rtbfabric-outboundexternallink-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

Property description not available.

`CreatedTimestamp`

Property description not available.

`LinkId`

Property description not available.

`LinkStatus`

Property description not available.

`UpdatedTimestamp`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ApplicationLogs

All content copied from https://docs.aws.amazon.com/.
