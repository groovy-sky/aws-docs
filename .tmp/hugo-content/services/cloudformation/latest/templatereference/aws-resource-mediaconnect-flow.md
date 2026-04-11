This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow

The `AWS::MediaConnect::Flow` resource defines a connection between one or more video
sources and one or more outputs. For each flow, you specify the transport protocol to
use, encryption information, and details for any outputs or entitlements that you want.
AWS Elemental MediaConnect returns an ingest endpoint where you can
send your live video as a single unicast stream. The service replicates and distributes
the video to every output that you specify, whether inside or outside the AWS Cloud. You can also set up entitlements on a flow to allow other
AWS accounts to access your content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaConnect::Flow",
  "Properties" : {
      "AvailabilityZone" : String,
      "EncodingConfig" : EncodingConfig,
      "FlowSize" : String,
      "Maintenance" : Maintenance,
      "MediaStreams" : [ MediaStream, ... ],
      "Name" : String,
      "NdiConfig" : NdiConfig,
      "Source" : Source,
      "SourceFailoverConfig" : FailoverConfig,
      "SourceMonitoringConfig" : SourceMonitoringConfig,
      "Tags" : [ Tag, ... ],
      "VpcInterfaces" : [ VpcInterface, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaConnect::Flow
Properties:
  AvailabilityZone: String
  EncodingConfig:
    EncodingConfig
  FlowSize: String
  Maintenance:
    Maintenance
  MediaStreams:
    - MediaStream
  Name: String
  NdiConfig:
    NdiConfig
  Source:
    Source
  SourceFailoverConfig:
    FailoverConfig
  SourceMonitoringConfig:
    SourceMonitoringConfig
  Tags:
    - Tag
  VpcInterfaces:
    - VpcInterface

```

## Properties

`AvailabilityZone`

The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS Region.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncodingConfig`

The encoding configuration to apply to the NDI® source when transcoding it to a transport stream for downstream distribution.

_Required_: No

_Type_: [EncodingConfig](aws-properties-mediaconnect-flow-encodingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlowSize`

Determines the processing capacity and feature set of the flow.

_Required_: No

_Type_: String

_Allowed values_: `MEDIUM | LARGE | LARGE_4X`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Maintenance`

The maintenance settings you want to use for the flow.

_Required_: No

_Type_: [Maintenance](aws-properties-mediaconnect-flow-maintenance.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaStreams`

The media streams that are associated with the flow. After you associate a media stream with a source, you can also associate it with outputs on the flow.

_Required_: No

_Type_: Array of [MediaStream](aws-properties-mediaconnect-flow-mediastream.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the flow.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NdiConfig`

Specifies the configuration settings for a flow's NDI source or output. Required when the flow includes an NDI source or output.

_Required_: No

_Type_: [NdiConfig](aws-properties-mediaconnect-flow-ndiconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

The settings for the source that you want to use for the new flow.

_Required_: Yes

_Type_: [Source](aws-properties-mediaconnect-flow-source.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceFailoverConfig`

The settings for source failover.

_Required_: No

_Type_: [FailoverConfig](aws-properties-mediaconnect-flow-failoverconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceMonitoringConfig`

The settings for source monitoring.

_Required_: No

_Type_: [SourceMonitoringConfig](aws-properties-mediaconnect-flow-sourcemonitoringconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediaconnect-flow-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcInterfaces`

The VPC Interfaces for this flow.

_Required_: No

_Type_: Array of [VpcInterface](aws-properties-mediaconnect-flow-vpcinterface.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the flow ARN. For example:

`{ "Ref":
            "arn:aws:mediaconnect:us-east-1:111122223333:flow:1-23aBC45dEF67hiJ8-12AbC34DE5fG:BasketballGame"
            }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`EgressIp`

The IP address from which video will be sent to output destinations.

`FlowArn`

The Amazon Resource Name (ARN) of the flow.

`FlowAvailabilityZone`

The Availability Zone that the flow was created in. These options are limited to the
Availability Zones within the current AWS Region.

`FlowNdiMachineName`

This read-only value represents the automatically-generated NDI machine name that MediaConnect generated for this flow. These NDI machine names are only generated when you don't specify your own custom name.

`Source.IngestIp`

The IP address that the flow listens on for incoming content.

`Source.SourceArn`

The ARN of the source.

`Source.SourceIngestPort`

The port that the flow listens on for incoming content. If the protocol of the source is Zixi, the port must
be set to 2088.

## Examples

- [Enabling thumbnails](#aws-resource-mediaconnect-flow--examples--Enabling_thumbnails)

- [Disabling thumbnails](#aws-resource-mediaconnect-flow--examples--Disabling_thumbnails)

- [Setting a thumbnail state to None](#aws-resource-mediaconnect-flow--examples--Setting_a_thumbnail_state_to_None)

### Enabling thumbnails

This example demonstrates a flow that sets the thumbnail state to
`ENABLED`. You can verify the thumbnail state using the
`DescribeFlow` API operation.

#### JSON

```json

{
  "Parameters": {
    "Name": {
      "Type": "String"
    }
  },
  "Resources": {
    "TestFlow": {
      "Type": "AWS::MediaConnect::Flow",
      "Properties": {
        "Name": {
          "Ref": "Name"
        },
        "Source": {
          "Name": "testSource",
          "Protocol": "rtp",
          "IngestPort": 1234,
          "Description": "CFN test Source",
          "WhitelistCidr": "0.0.0.0/0"
        },
        "Maintenance": {
          "MaintenanceDay": "Tuesday",
          "MaintenanceStartHour": "02:00"
        },
        "SourceMonitoringConfig": {
          "ThumbnailState": "ENABLED"
        }
      }
    }
  },
  "Outputs": {
    "FlowArn": {
      "Value": {
        "Ref": "TestFlow"
      }
    }
  }
}
```

#### YAML

```yaml

Parameters:
  Name:
    Type: String

Resources:
  TestFlow:
    Type: AWS::MediaConnect::Flow
    Properties:
      Name: !Ref Name
      Source:
        Name: testSource
        Protocol: rtp
        IngestPort: 1234
        Description: CFN test Source
        WhitelistCidr: 0.0.0.0/0
      Maintenance:
        MaintenanceDay: Tuesday
        MaintenanceStartHour: 02:00
      SourceMonitoringConfig:
        ThumbnailState: ENABLED

Outputs:
  FlowArn:
    Value: !Ref TestFlow
```

### Disabling thumbnails

This example demonstrates a flow that sets the thumbnail state to
`DISABLED`. You can verify the thumbnail state using the
`DescribeFlow` API operation.

#### JSON

```json

{
  "Parameters": {
    "Name": {
      "Type": "String"
    }
  },
  "Resources": {
    "TestFlow": {
      "Type": "AWS::MediaConnect::Flow",
      "Properties": {
        "Name": {
          "Ref": "Name"
        },
        "Source": {
          "Name": "testSource",
          "Protocol": "rtp",
          "IngestPort": 1234,
          "Description": "CFN test Source",
          "WhitelistCidr": "0.0.0.0/0"
        },
        "Maintenance": {
          "MaintenanceDay": "Tuesday",
          "MaintenanceStartHour": "02:00"
        },
        "SourceMonitoringConfig": {
          "ThumbnailState": "DISABLED"
        }
      }
    }
  },
  "Outputs": {
    "FlowArn": {
      "Value": {
        "Ref": "TestFlow"
      }
    }
  }
}
```

#### YAML

```yaml

Parameters:
  Name:
    Type: String

Resources:
  TestFlow:
    Type: AWS::MediaConnect::Flow
    Properties:
      Name: !Ref Name
      Source:
        Name: testSource
        Protocol: rtp
        IngestPort: 1234
        Description: CFN test Source
        WhitelistCidr: 0.0.0.0/0
      Maintenance:
        MaintenanceDay: Tuesday
        MaintenanceStartHour: 02:00
      SourceMonitoringConfig:
        ThumbnailState: DISABLED

Outputs:
  FlowArn:
    Value: !Ref TestFlow
```

### Setting a thumbnail state to `None`

If you don't specify a thumbnail state, it affects new and existing flows
differently. Here's how it works in each scenario.

**Creating a flow**

When you create a new flow without specifying a thumbnail state, no state has been
set yet. After you create your flow in this way, the `DescribeFlow` API
operation shows the thumbnail state as `NONE`.

**Updating an existing flow**

When you update an existing flow and remove its thumbnail state (previously set to
either `ENABLED` or `DISABLED`), you're altering a previously
set state. In this case, the state doesn't become `NONE`. Instead, it
defaults to `DISABLED`. This happens because after you've specified a
state for an existing flow, it can't revert to `NONE`. The
`NONE` state only applies when no state has been set before. After you
update your flow in this way, the `DescribeFlow` API operation shows the
thumbnail state as `DISABLED`, not `NONE`.

#### JSON

```json

{
  "Parameters": {
    "Name": {
      "Type": "String"
    }
  },
  "Resources": {
    "TestFlow": {
      "Type": "AWS::MediaConnect::Flow",
      "Properties": {
        "Name": {
          "Ref": "Name"
        },
        "Source": {
          "Name": "testSource",
          "Protocol": "rtp",
          "IngestPort": 1234,
          "Description": "CFN test Source",
          "WhitelistCidr": "0.0.0.0/0"
        },
        "Maintenance": {
          "MaintenanceDay": "Tuesday",
          "MaintenanceStartHour": "02:00"
        }
      }
    }
  },
  "Outputs": {
    "FlowArn": {
      "Value": {
        "Ref": "TestFlow"
      }
    }
  }
}
```

#### YAML

```yaml

Parameters:
  Name:
    Type: String

Resources:
  TestFlow:
    Type: AWS::MediaConnect::Flow
    Properties:
      Name: !Ref Name
      Source:
        Name: testSource
        Protocol: rtp
        IngestPort: 1234
        Description: CFN test Source
        WhitelistCidr: 0.0.0.0/0
      Maintenance:
        MaintenanceDay: Tuesday
        MaintenanceStartHour: 02:00

Outputs:
  FlowArn:
    Value: !Ref TestFlow
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcInterfaceAttachment

AudioMonitoringSetting

All content copied from https://docs.aws.amazon.com/.
