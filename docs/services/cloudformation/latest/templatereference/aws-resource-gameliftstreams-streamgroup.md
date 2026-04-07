This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLiftStreams::StreamGroup

Stream groups manage how Amazon GameLift Streams allocates resources and handles concurrent streams, allowing you to effectively manage capacity and
costs. Within a stream group, you specify an application to stream, streaming locations and their capacity, and the stream class you want
to use when streaming applications to your end-users. A stream class defines the hardware configuration of the compute resources that
Amazon GameLift Streams will use when streaming, such as the CPU, GPU, and memory.

Stream capacity represents the number of concurrent streams that can be active at a time. You set stream
capacity per location, per stream group. The following capacity settings are available:

- **Always-on capacity**: This setting, if non-zero, indicates minimum streaming capacity
which is allocated to you and is never released back to the service. You pay for this
base level of capacity at all times, whether used or idle.

- **Maximum capacity**: This indicates the maximum capacity that the service can allocate
for you. Newly created streams may take a few minutes to start. Capacity is released back
to the service when idle. You pay for capacity that is allocated to you until it is released.

- **Target-idle capacity**: This indicates idle capacity which the service pre-allocates and
holds for you in anticipation of future activity. This helps to insulate your users from
capacity-allocation delays. You pay for capacity which is held in this intentional idle
state.

Values for capacity must be whole number multiples of the tenancy value of the stream group's stream class.

###### Note

Application association is not currently supported in CloudFormation. To link additional applications to a stream group, use the
Amazon GameLift Streams console or the AWS CLI.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::GameLiftStreams::StreamGroup",
  "Properties" : {
      "DefaultApplication" : DefaultApplication,
      "Description" : String,
      "LocationConfigurations" : [ LocationConfiguration, ... ],
      "StreamClass" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::GameLiftStreams::StreamGroup
Properties:
  DefaultApplication:
    DefaultApplication
  Description: String
  LocationConfigurations:
    - LocationConfiguration
  StreamClass: String
  Tags:
    Key: Value

```

## Properties

`DefaultApplication`

Object that identifies the Amazon GameLift Streams application to stream with this stream group.

_Required_: No

_Type_: [DefaultApplication](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gameliftstreams-streamgroup-defaultapplication.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A descriptive label for the stream group.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_.!+@/][a-zA-Z0-9-_.!+@/ ]*$`

_Minimum_: `1`

_Maximum_: `80`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationConfigurations`

A set of one or more locations and the streaming capacity for each location. One of the locations MUST be your primary location, which is
the AWS Region where you are specifying this resource.

_Required_: Yes

_Type_: Array of [LocationConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-gameliftstreams-streamgroup-locationconfiguration.html)

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamClass`

The target stream quality for sessions that are hosted in this stream group. Set a stream class that is appropriate to the type of
content that you're streaming. Stream class determines the type of computing resources Amazon GameLift Streams uses and impacts the cost of streaming. The
following options are available:

A stream class can be one of the following:

- **`gen6n_pro_win2022` (NVIDIA, pro)** Supports applications with extremely high 3D scene complexity which require maximum resources. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.6, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 16 vCPUs, 64 GB RAM, 24 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen6n_pro` (NVIDIA, pro)** Supports applications with extremely high 3D scene complexity which require maximum resources.
Uses dedicated NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 16 vCPUs, 64 GB RAM, 24 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen6n_ultra_win2022` (NVIDIA, ultra)** Supports applications with high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.6, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen6n_ultra` (NVIDIA, ultra)** Supports applications with high 3D scene complexity.
Uses dedicated NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen6n_high` (NVIDIA, high)** Supports applications with moderate to high 3D scene complexity.
Uses NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 4 vCPUs, 16 GB RAM, 12 GB VRAM

- Tenancy: Supports up to 2 concurrent stream sessions

- **`gen6n_medium` (NVIDIA, medium)** Supports applications with moderate 3D scene complexity.
Uses NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 2 vCPUs, 8 GB RAM, 6 GB VRAM

- Tenancy: Supports up to 4 concurrent stream sessions

- **`gen6n_small` (NVIDIA, small)** Supports applications with lightweight 3D scene complexity and low CPU usage.
Uses NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 1 vCPUs, 4 GB RAM, 2 GB VRAM

- Tenancy: Supports up to 12 concurrent stream sessions

- **`gen6n_medium_win2022` (NVIDIA, medium)** Supports applications with low 3D scene complexity.
Uses NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 8 vCPUs, 32 GB RAM, 6 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen6n_small_win2022` (NVIDIA, small)** Supports applications with low 3D scene complexity.
Uses NVIDIA L4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 2 vCPUs, 8 GB RAM, 3 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen5n_win2022` (NVIDIA, ultra)** Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.6, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA A10G Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen5n_high` (NVIDIA, high)** Supports applications with moderate to high 3D scene complexity.
Uses NVIDIA A10G Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 4 vCPUs, 16 GB RAM, 12 GB VRAM

- Tenancy: Supports up to 2 concurrent stream sessions

- **`gen5n_ultra` (NVIDIA, ultra)** Supports applications with extremely high 3D scene complexity.
Uses dedicated NVIDIA A10G Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 8 vCPUs, 32 GB RAM, 24 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen4n_win2022` (NVIDIA, ultra)** Supports applications with extremely high 3D scene complexity. Runs applications on Microsoft Windows Server 2022 Base and supports DirectX 12. Compatible with Unreal Engine versions up through 5.6, 32 and 64-bit applications, and anti-cheat technology. Uses NVIDIA T4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM

- Tenancy: Supports 1 concurrent stream session

- **`gen4n_high` (NVIDIA, high)** Supports applications with moderate to high 3D scene complexity.
Uses NVIDIA T4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 4 vCPUs, 16 GB RAM, 8 GB VRAM

- Tenancy: Supports up to 2 concurrent stream sessions

- **`gen4n_ultra` (NVIDIA, ultra)** Supports applications with high 3D scene complexity.
Uses dedicated NVIDIA T4 Tensor Core GPU.

- Reference resolution: 1080p

- Reference frame rate: 60 fps

- Workload specifications: 8 vCPUs, 32 GB RAM, 16 GB VRAM

- Tenancy: Supports 1 concurrent stream session

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of labels to assign to the new stream group resource. Tags are developer-defined key-value pairs. Tagging AWS resources is
useful for resource management, access management and cost allocation. See [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the _AWS General Reference_.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns an [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html)
that uniquely identifies the stream group resource across all AWS Regions. For example:

`arn:aws:gameliftstreams:us-west-2:123456789012:streamgroup/sg-1AB2C3De4`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

An [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that uniquely
identifies the stream group resource. For example: `arn:aws:gameliftstreams:us-west-2:123456789012:streamgroup/sg-1AB2C3De4`.

`Id`

An ID that uniquely identifies the stream group resource. For example: `sg-1AB2C3De4`.

## See also

- [Manage streaming with an Amazon GameLift\
Streams stream group](https://docs.aws.amazon.com/gameliftstreams/latest/developerguide/stream-groups.html) in the _Amazon GameLift Streams Developer Guide_

- [CreateStreamGroup](https://docs.aws.amazon.com/gameliftstreams/latest/apireference/API_CreateStreamGroup.html) in the
_Amazon GameLift Streams API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RuntimeEnvironment

DefaultApplication
