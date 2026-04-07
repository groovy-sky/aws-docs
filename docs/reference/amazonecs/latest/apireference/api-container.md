# Container

A Docker container that's part of a task.

## Contents

**containerArn**

The Amazon Resource Name (ARN) of the container.

Type: String

Required: No

**cpu**

The number of CPU units set for the container. The value is `0` if no value
was specified in the container definition when the task definition was
registered.

Type: String

Required: No

**exitCode**

The exit code returned from the container.

Type: Integer

Required: No

**gpuIds**

The IDs of each GPU assigned to the container.

Type: Array of strings

Required: No

**healthStatus**

The health status of the container. If health checks aren't configured for this
container in its task definition, then it reports the health status as
`UNKNOWN`.

Type: String

Valid Values: `HEALTHY | UNHEALTHY | UNKNOWN`

Required: No

**image**

The image used for the container.

Type: String

Required: No

**imageDigest**

The container image manifest digest.

Type: String

Required: No

**lastStatus**

The last known status of the container.

Type: String

Required: No

**managedAgents**

The details of any Amazon ECS managed agents associated with the container.

Type: Array of [ManagedAgent](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ManagedAgent.html) objects

Required: No

**memory**

The hard limit (in MiB) of memory set for the container.

Type: String

Required: No

**memoryReservation**

The soft limit (in MiB) of memory set for the container.

Type: String

Required: No

**name**

The name of the container.

Type: String

Required: No

**networkBindings**

The network bindings associated with the container.

Type: Array of [NetworkBinding](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_NetworkBinding.html) objects

Required: No

**networkInterfaces**

The network interfaces associated with the container.

Type: Array of [NetworkInterface](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_NetworkInterface.html) objects

Required: No

**reason**

A short (1024 max characters) human-readable string to provide additional details
about a running or stopped container.

Type: String

Required: No

**runtimeId**

The ID of the Docker container.

Type: String

Required: No

**taskArn**

The ARN of the task.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ecs-2014-11-13/Container)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ecs-2014-11-13/Container)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ecs-2014-11-13/Container)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClusterSetting

ContainerDefinition
