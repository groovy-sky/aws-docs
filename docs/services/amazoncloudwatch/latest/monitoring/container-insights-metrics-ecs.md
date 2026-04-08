# Amazon ECS Container Insights metrics

Container Insights metrics provides additional network, storage, and ephemeral storage metrics. These metrics provide more information than standard Amazon ECS metrics. Container Insights integrates with CloudWatch Logs. You can correlate metric changes with log entries for easier troubleshooting.
Container Insights also shows cluster, service, and daemon-wide statistics by averaging data across all tasks. This provides a higher-level view of your service and daemon health, assisting in both environment monitoring and capacity planning.

**Use cases**

- **Problem identification and troubleshooting** – Track failed deployments by analyzing task state transition patterns, enabling rapid identification of failure points. Diagnose configuration issues through comprehensive examination of task startup sequences and initialization behaviors

- **Cluster and Service-Level health assessment** – Shows average task performance across the cluster. This approach moderates outliers to deliver a more stable view of cluster and service health.
Use these insights for general service monitoring where extreme values could be misleading

- **Service availability issues** – Detect deployment failures by monitoring running task count metrics. Correlate service event logs with performance metrics to understand infrastructure impacts.
Track task restart patterns to identify unstable services or infrastructure issues

- **Capacity planning for average load** – It helps determine resource requirements based on typical task behavior patterns, provides consistent metrics that support effective long-term planning, and reduces the impact of short-lived spikes on capacity decisions

- **Provides additional metrics** – Collects additional network, storage, and ephemeral storage metrics not available in vended metrics

For more information on Amazon ECS metrics, see [Amazon ECS service utilization metrics use cases](../../../amazonecs/latest/developerguide/service-utilization-metrics-explanation.md) and for information on container insights with enhanced observability
[Amazon ECS Container Insights with enhanced observability metrics](container-insights-enhanced-observability-metrics-ecs.md).

###### Note

Amazon ECS Managed Daemon metrics use the same `ECS/ContainerInsights` namespace and
the same `ServiceName` dimension as service metrics. For daemon metrics, the
`ServiceName` dimension value uses the format
`daemon:daemon-name`. For example, a daemon named
`my-daemon` has a `ServiceName` dimension value of
`daemon:my-daemon`. All metrics in the table below that include the
`ServiceName` dimension also apply to Managed Daemons.

The following table lists the metrics and dimensions that Container Insights collects
for Amazon ECS. These metrics are in the `ECS/ContainerInsights` namespace. For more
information, see [Metrics](cloudwatch-concepts.md#Metric).

If you do not see any Container Insights metrics in your console, be sure that you have
completed the setup of Container Insights. Metrics do not appear before Container Insights
has been set up completely. For more information, see [Setting up Container Insights](deploy-container-insights.md).

The following metrics are available when you complete the steps in [Setting up Container Insights on Amazon ECS](deploy-container-insights-ecs-cluster.md).

Metric nameDimensionsDescription

`ContainerInstanceCount`

`ClusterName`

The number of EC2 instances running the Amazon ECS agent that are registered with a
cluster.

This metric is collected only for container instances that are running Amazon ECS
tasks in the cluster. It is not collected for empty container instances that do
not have any Amazon ECS tasks.

Unit: Count

`CpuUtilized`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The CPU units used by tasks in the resource that is specified by the dimension
set that you're using.

Also applies to Managed Daemons.

Unit: None

`CpuReserved`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The CPU units reserved by tasks in the resource that is specified by the
dimension set that you're using.

This metric is collected based on the CPU reservation defined in the task
definition, for example, at the task or all containers level. If this is not
specified in the task definition, then the instance CPU reservation is
used.

Also applies to Managed Daemons.

Unit: None

`DeploymentCount`

`ServiceName`, `ClusterName`

The number of deployments in an Amazon ECS service.

Unit: Count

`DesiredTaskCount`

`ServiceName`, `ClusterName`

The desired number of tasks for an Amazon ECS service.

Unit: Count

`EBSFilesystemSize`

`VolumeName`, `TaskDefinitionFamily`,
`ClusterName`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

The total amount, in gigabytes (GB), of Amazon EBS filesystem storage that is
allocated to the resources specified by the dimensions you're using.

This metric is only available for tasks that run on Amazon ECS infrastructure
running on Fargate using platform version `1.4.0` or Amazon EC2 instances
using container agent version `1.79.0` or later.

Also applies to Managed Daemons.

Unit: Gigabytes (GB)

`EBSFilesystemUtilized`

`VolumeName`, `TaskDefinitionFamily`,
`ClusterName`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

The total amount, in gigabytes (GB), of Amazon EBS filesystem storage that is being
used by the resources specified by the dimensions that you're using.

This metric is only available for tasks that run on Amazon ECS infrastructure
running on Fargate using platform version `1.4.0` or Amazon EC2 instances
using container agent version `1.79.0` or later.

For tasks run on Fargate, Fargate reserves space on the disk that is only
used by Fargate. There is no cost associated with the space Fargate uses, but
you will see this additional storage using tools like `df`.

Also applies to Managed Daemons.

Unit: Gigabytes (GB)

`EphemeralStorageReserved` [1](#ci-metrics-ecs-storage-fargate-note)

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The number of bytes reserved from ephemeral storage in the resource that is
specified by the dimensions that you're using. Ephemeral storage is used for the
container root filesystem and any bind mount host volumes defined in the container
image and task definition. The amount of ephemeral storage can’t be changed in a
running task.

This metric is only available for tasks that run on Fargate Linux platform
version 1.4.0 or later.

Also applies to Managed Daemons.

Unit: Gigabytes (GB)

`EphemeralStorageUtilized` [1](#ci-metrics-ecs-storage-fargate-note)

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The number of bytes used from ephemeral storage in the resource that is
specified by the dimensions that you're using. Ephemeral storage is used for the
container root filesystem and any bind mount host volumes defined in the container
image and task definition. The amount of ephemeral storage can’t be changed in a
running task.

This metric is only available for tasks that run on Fargate Linux platform
version 1.4.0 or later.

Also applies to Managed Daemons.

Unit: Gigabytes (GB)

`InstanceOSFilesystemUtilization`

`CapacityProviderName`, `ClusterName`, `ContainerInstanceId`, `EC2InstanceId`

`ClusterName`

The percentage of total disk space that is used for OS volume.

`InstanceDataFilesystemUtilization`

`CapacityProviderName`, `ClusterName`, `ContainerInstanceId`, `EC2InstanceId`

`ClusterName`

The percentage of total disk space that is used for data volume.

`MemoryUtilized`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The memory being used by tasks in the resource that is specified by the
dimension set that you're using.

###### Note

If you're using the Java ZGC garbage collector for your application, this
metric might be inaccurate.

Although `MemoryUtilized` and `MemoryReserved` are
marked as "Megabytes", the actual units are in MiB (Mebibytes).

Also applies to Managed Daemons.

Unit: Megabytes

`MemoryReserved`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The memory that is reserved by tasks in the resource that is specified by the
dimension set that you're using. This metric is collected based on the memory
reservation defined in the task definition, for example, at the task or all
containers level. If this is not specified in the task definition, then the
instance memory reservation is used.

Also applies to Managed Daemons.

Unit: Megabytes

###### Note

Although `MemoryUtilized` and `MemoryReserved` are
marked as "Megabytes", the actual units are in MiB (Mebibytes).

`NetworkRxBytes`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The number of bytes received by the resource that is specified by the
dimensions that you're using. This metric is obtained from the Docker
runtime.

This metric is available only for containers in tasks using the
`awsvpc` or `bridge` network modes.

Also applies to Managed Daemons.

Unit: Bytes/Second

`NetworkTxBytes`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The number of bytes transmitted by the resource that is specified by the
dimensions that you're using. This metric is obtained from the Docker
runtime.

This metric is available only for containers in tasks using the
`awsvpc` or `bridge` network modes.

Also applies to Managed Daemons.

Unit: Bytes/Second

`PendingTaskCount`

`ServiceName`, `ClusterName`

The number of tasks currently in the `PENDING` state.

Unit: Count

`RunningTaskCount`

`ServiceName`, `ClusterName`

The number of tasks currently in the `RUNNING` state.

Unit: Count

`RestartCount`

`ClusterName`

`ClusterName`, `ServiceName`

`ClusterName`, `TaskDefinitionFamily`

The number of times a container in an Amazon ECS task has been restarted.

This metric is collected only for containers that have a restart policy
enabled.

Also applies to Managed Daemons.

Unit: Count

`ServiceCount`

`ClusterName`

The number of services in the cluster.

Unit: Count

`StorageReadBytes`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The number of bytes read from storage on the instance in the resource that is
specified by the dimensions that you're using. This does not include read bytes
for your storage devices. This metric is obtained from the Docker runtime.

Also applies to Managed Daemons.

Unit: Bytes

`StorageWriteBytes`

`TaskDefinitionFamily`, `ClusterName`

`ServiceName`, `ClusterName`

`ClusterName`

The number of bytes written to storage in the resource that is specified by
the dimensions that you're using. This metric is obtained from the Docker
runtime.

Also applies to Managed Daemons.

Unit: Bytes

`TaskCount`

`ClusterName`

The number of tasks running in the cluster.

Unit: Count

`TaskSetCount`

`ServiceName`, `ClusterName`

The number of task sets in the service.

Unit: Count

###### Note

The
`EphemeralStorageReserved` and `EphemeralStorageUtilized` metrics
are only available for tasks that run on Fargate Linux platform version 1.4.0 or
later.

Fargate reserves space on disk. It is only used by Fargate. You aren't billed for it.
It isn't shown in these metrics. However, you can see this additional storage in other
tools such as `df`.

The following metrics are available when you complete the steps in [Deploying the CloudWatch agent to collect EC2 instance-level metrics on Amazon ECS](deploy-container-insights-ecs-instancelevel.md)

Metric nameDimensionsDescription

`instance_cpu_limit`

`ClusterName`

The maximum number of CPU units that can be assigned to a single EC2 Instance
in the cluster.

Unit: None

`instance_cpu_reserved_capacity`

`ClusterName`

`InstanceId`, `ContainerInstanceId`,
`ClusterName`

The percentage of CPU currently being reserved on a single EC2 instance in the
cluster.

Unit: Percent

`instance_cpu_usage_total`

`ClusterName`

The number of CPU units being used on a Single EC2 instance in the
cluster.

Unit: None

`instance_cpu_utilization`

`ClusterName`

`InstanceId`, `ContainerInstanceId`,
`ClusterName`

The total percentage of CPU units being used on a single EC2 instance in the
cluster.

Unit: Percent

`instance_filesystem_utilization`

`ClusterName`

`InstanceId`, `ContainerInstanceId`,
`ClusterName`

The total percentage of file system capacity being used on a single EC2
instance in the cluster.

Unit: Percent

`instance_memory_limit`

`ClusterName`

The maximum amount of memory, in bytes, that can be assigned to a single EC2
Instance in this cluster.

Unit: Bytes

`instance_memory_reserved_capacity`

`ClusterName`

`InstanceId`, `ContainerInstanceId`,
`ClusterName`

The percentage of Memory currently being reserved on a single EC2 Instance in
the cluster.

Unit: Percent

`instance_memory_utilization`

`ClusterName`

`InstanceId`, `ContainerInstanceId`,
`ClusterName`

The total percentage of memory being used on a single EC2 Instance in the
cluster.

###### Note

If you're using the Java ZGC garbage collector for your application, this
metric might be inaccurate.

Unit: Percent

`instance_memory_working_set`

`ClusterName`

The amount of memory, in bytes, being used on a single EC2 Instance in the
cluster.

###### Note

If you're using the Java ZGC garbage collector for your application, this
metric might be inaccurate.

Unit: Bytes

`instance_network_total_bytes`

`ClusterName`

The total number of bytes per second transmitted and received over the network
on a single EC2 Instance in the cluster.

Unit: Bytes/second

`instance_number_of_running_tasks`

`ClusterName`

The number of running tasks on a single EC2 Instance in the cluster.

Unit: Count

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon ECS Container Insights with enhanced observability metrics

Amazon EKS and Kubernetes Container Insights with enhanced observability metrics

All content copied from https://docs.aws.amazon.com/.
