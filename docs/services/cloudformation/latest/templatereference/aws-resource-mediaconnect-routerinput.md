This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::RouterInput

The `AWS::MediaConnect::RouterInput` resource defines a connection point in the MediaConnect router
that can receive content from your source endpoint. You can configure a router input with either a Regional routing scope or a global routing scope.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::RouterInput",
  "Properties" : {
      "AvailabilityZone" : String,
      "Configuration" : RouterInputConfiguration,
      "MaintenanceConfiguration" : MaintenanceConfiguration,
      "MaximumBitrate" : Integer,
      "Name" : String,
      "RegionName" : String,
      "RoutingScope" : String,
      "Tags" : [ Tag, ... ],
      "Tier" : String,
      "TransitEncryption" : RouterInputTransitEncryption
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::RouterInput
Properties:
  AvailabilityZone: String
  Configuration:
    RouterInputConfiguration
  MaintenanceConfiguration:
    MaintenanceConfiguration
  MaximumBitrate: Integer
  Name: String
  RegionName: String
  RoutingScope: String
  Tags:
    - Tag
  Tier: String
  TransitEncryption:
    RouterInputTransitEncryption

```

## Properties

`AvailabilityZone`

The Availability Zone of the router input.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

The configuration settings for a router input.

_Required_: Yes

_Type_: [RouterInputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-routerinput-routerinputconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaintenanceConfiguration`

The maintenance configuration settings applied to this router input.

_Required_: No

_Type_: [MaintenanceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-routerinput-maintenanceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBitrate`

The maximum bitrate for the router input.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the router input.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionName`

The AWS Region where the router input is located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoutingScope`

Indicates whether the router input is configured for Regional or global routing.

_Required_: Yes

_Type_: String

_Allowed values_: `REGIONAL | GLOBAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs that can be used to tag and organize this router input.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-routerinput-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

The tier level of the router input.

_Required_: Yes

_Type_: String

_Allowed values_: `INPUT_100 | INPUT_50 | INPUT_20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitEncryption`

Encryption information.

_Required_: No

_Type_: [RouterInputTransitEncryption](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediaconnect-routerinput-routerinputtransitencryption.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the router input ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-west-2:111122223333:routerInput:86f52099b545"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the router input.

`CreatedAt`

The timestamp when the router input was created.

`Id`

The unique identifier of the router input.

`InputType`

The type of the router input.

`IpAddress`

The IP address of the router input.

`MaintenanceType`

The type of maintenance configuration applied to this router input.

`RoutedOutputs`

The number of router outputs associated with the router input.

`State`

The current state of the router input.

`UpdatedAt`

The timestamp when the router input was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GatewayNetwork

FailoverRouterInputConfiguration
