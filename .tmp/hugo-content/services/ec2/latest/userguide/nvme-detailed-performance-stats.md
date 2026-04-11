# Detailed performance statistics for Amazon EC2 instance store volumes

Amazon EC2 provides real-time, high-resolution performance statistics for NVMe instance store volumes
attached to Nitro-based Amazon EC2 instances. These statistics are presented as aggregated counters that
are retained for the duration of the instance's lifetime. The statistics provide details about the
cumulative number of operations, bytes sent and received, time spent on read and write I/O operations,
and histograms for read and write I/O operations. While these statistics maintain consistency with
[Amazon EBS \
detailed performance statistics](../../../ebs/latest/userguide/nvme-detailed-performance-stats.md), they also include detailed latency histograms broken down
by I/O size, which can provide more granular insights into your storage performance patterns. This
enhanced visibility enables you to precisely identify which specific I/O sizes are experiencing
latency issues, allowing you to optimize application performance and troubleshoot issues more
effectively.

You can collect these statistics at a granularity of up to 1 second intervals. If
requests are made more frequently than 1 second intervals, the NVMe driver might queue
the requests, along with other admin commands, to be processed at a later time.

###### Considerations

- The statistics are supported only for NVMe instance store volumes attached to
Nitro-based instances.

- The counters are not persistent across instance stops and restarts.

- The statistics are available at no additional cost.

## Statistics

The NVMe block device vends the following statistics:

Statistic nameFull nameTypeDescription`total_read_ops`Total read operationsCounterThe total number of completed read operations.`total_write_ops`Total write operationsCounterThe total number of completed write operations.`total_read_bytes`Total read bytesCounterThe total number of read bytes transferred.`total_write_bytes`Total write bytesCounterThe total number of write bytes transferred.`total_read_time`Total read timeCounterThe total time spent, in microseconds, by all completed read operations.`total_write_time`Total write timeCounterThe total time spent, in microseconds, by all completed write operations.`instance_store_volume_performance_exceeded_iops`Total time demand exceeded volume's maximum IOPSCounterThe total time, in microseconds, that IOPS requests exceeded the volume's maximum IOPS. Any
value above `0` indicates that your workload demanded more IOPS than the volume could
deliver. Ideally, the incremental count of this metric, between two snapshot times, should be
minimal.`instance_store_volume_performance_exceeded_tp`Total time demand exceeded volume's maximum throughputCounterThe total time, in microseconds, that throughput requests exceeded the volume's maximum
throughput. Any value above `0` indicates that your workload demanded more throughput
than the volume could deliver. Ideally, the incremental count of this metric, between two snapshot
times, should be minimal.`volume_queue_length`Volume queue lengthPoint in timeThe number of read and write operations waiting to be completed.`read_io_latency_histogram`Read I/O histogramHistogram \*The number of read operations completed within each latency bin, in microseconds. `write_io_latency_histogram`Write I/O histogramHistogram \*The number of write operations completed within each latency bin, in microseconds.

###### Note

\\* Histogram statistics represent only I/O operations that have completed successfully. Stalled or
impaired I/O operations are not included, but will be evident in the `volume_queue_length`
statistics, which is presented as a point-in-time statistic.

## Accessing the statistics

The statistics must be accessed directly from the instance to which the instance store volumes is
attached. You can access the statistics using one of the following methods.

Amazon CloudWatch

You can configure the Amazon CloudWatch agent to collect the statistics from your instance and make them
available as custom metrics in CloudWatch. You can then use the metrics in CloudWatch to analyze I/O patterns,
track performance trends, create custom dashboards, and set up automated alarms based on performance
thresholds.

For more information about configuring the CloudWatch agent, see [Collect Amazon EC2 instance store volume metrics](../../../amazoncloudwatch/latest/monitoring/container-insights-metrics-instance-store-collect.md).

nvme-cli tool

###### To access the statistics

1. Connect to the instance to which the volume is attached.

2. Amazon Linux 2023 AMIs released after September 15, 2025 include the latest version of the
    `nvme-cli` tool. If you are using an older Amazon Linux AMI, update the
    `nvme-cli` tool.

```nohighlight

sudo yum install nvme-cli
```

3. Run the following command and specify the device name for the volume.

```nohighlight

sudo nvme amzn stats /dev/nvme0n1
```

The statistics also provide detailed latency histograms broken down by I/O size. To view
statistics broken down by I/O size, include the `--details` option. For example:

```nohighlight

sudo nvme amzn stats --details /dev/nvme0n1
```

You can get more information about how to use the tool by specifying the `--help`
option. For example:

```nohighlight

sudo nvme amzn stats --help
```

nvme\_amzn.exe tool

###### To access the statistics

1. Connect to the instance to which the volume is attached.

2. Make sure that you're using AWSNVMe driver version `1.7.0` or later.
    For more information about updating the AWSNVMe driver, see [AWS NVMe drivers](aws-nvme-drivers.md).

3. Get the disk number for the volume. For more information, see [Map NVMe disks on Amazon EC2 Windows \
    instance to volumes](windows-list-disks-nvme.md).

4. Run the following command as Administrator and specify the disk number for the
    volume.

```nohighlight

.\nvme_amzn.exe stats disk_number
```

The statistics also provide detailed latency histograms broken down by I/O size. To
view statistics broken down by I/O size, include the `--details` option. For
example:

```nohighlight

.\nvme_amzn.exe stats --details disk_number
```

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Initialize instance store volumes

Root volumes
