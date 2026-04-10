This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link

Creates a new link between gateways.

Establishes a connection that allows gateways to communicate and exchange bid requests
and responses.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RTBFabric::Link",
  "Properties" : {
      "GatewayId" : String,
      "HttpResponderAllowed" : Boolean,
      "LinkAttributes" : LinkAttributes,
      "LinkLogSettings" : LinkLogSettings,
      "ModuleConfigurationList" : [ ModuleConfiguration, ... ],
      "PeerGatewayId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RTBFabric::Link
Properties:
  GatewayId: String
  HttpResponderAllowed: Boolean
  LinkAttributes:
    LinkAttributes
  LinkLogSettings:
    LinkLogSettings
  ModuleConfigurationList:
    - ModuleConfiguration
  PeerGatewayId: String
  Tags:
    - Tag

```

## Properties

`GatewayId`

The unique identifier of the gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^rtb-gw-[a-z0-9-]{1,25}$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`HttpResponderAllowed`

Boolean to specify if an HTTP responder is allowed.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LinkAttributes`

Attributes of the link.

_Required_: No

_Type_: [LinkAttributes](aws-properties-rtbfabric-link-linkattributes.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LinkLogSettings`

Settings for the application logs.

_Required_: Yes

_Type_: [LinkLogSettings](aws-properties-rtbfabric-link-linklogsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModuleConfigurationList`

Property description not available.

_Required_: No

_Type_: Array of [ModuleConfiguration](aws-properties-rtbfabric-link-moduleconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PeerGatewayId`

The unique identifier of the peer gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^rtb-gw-[a-z0-9-]{1,25}$`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Tags`

A map of the key-value pairs of the tag or tags to assign to the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-rtbfabric-link-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Property description not available.

`CreatedTimestamp`

Property description not available.

`LinkDirection`

Property description not available.

`LinkId`

The unique identifier of the link.

`LinkStatus`

Property description not available.

`UpdatedTimestamp`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Action

All content copied from https://docs.aws.amazon.com/.
