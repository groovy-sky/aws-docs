This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterOutput

The `AWS::MediaConnect::RouterOutput` resource defines a connection point in the MediaConnect router
that can send content to your destination endpoint. You can configure a router output with either a Regional routing scope or a global routing scope.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::RouterOutput",
  "Properties" : {
      "AvailabilityZone" : String,
      "Configuration" : RouterOutputConfiguration,
      "MaintenanceConfiguration" : MaintenanceConfiguration,
      "MaximumBitrate" : Integer,
      "Name" : String,
      "RegionName" : String,
      "RoutingScope" : String,
      "Tags" : [ Tag, ... ],
      "Tier" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::RouterOutput
Properties:
  AvailabilityZone: String
  Configuration:
    RouterOutputConfiguration
  MaintenanceConfiguration:
    MaintenanceConfiguration
  MaximumBitrate: Integer
  Name: String
  RegionName: String
  RoutingScope: String
  Tags:
    - Tag
  Tier: String

```

## Properties

`AvailabilityZone`

The Availability Zone of the router output.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

The configuration settings for a router output.

_Required_: Yes

_Type_: [RouterOutputConfiguration](aws-properties-mediaconnect-routeroutput-routeroutputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceConfiguration`

The maintenance configuration settings applied to this router output.

_Required_: No

_Type_: [MaintenanceConfiguration](aws-properties-mediaconnect-routeroutput-maintenanceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBitrate`

The maximum bitrate for the router output.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the router output.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionName`

The AWS Region where the router output is located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoutingScope`

Indicates whether the router output is configured for Regional or global routing.

_Required_: Yes

_Type_: String

_Allowed values_: `REGIONAL | GLOBAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs that can be used to tag and organize this router output.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediaconnect-routeroutput-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

The tier level of the router output.

_Required_: Yes

_Type_: String

_Allowed values_: `OUTPUT_100 | OUTPUT_50 | OUTPUT_20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the router output ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-west-2:111122223333:routerOutput:56eb95d755a1"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the router output.

`CreatedAt`

The timestamp when the router output was created.

`Id`

The unique identifier of the router output.

`IpAddress`

The IP address of the router output.

`MaintenanceType`

The type of maintenance configuration applied to this router output.

`OutputType`

The type of the router output.

`RoutedState`

The current state of the association between the router output and its input.

`State`

The overall state of the router output.

`UpdatedAt`

The timestamp when the router output was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcRouterNetworkInterfaceConfiguration

FlowTransitEncryption

All content copied from https://docs.aws.amazon.com/.
