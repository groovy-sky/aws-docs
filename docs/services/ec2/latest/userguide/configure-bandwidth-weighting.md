# EC2 instance bandwidth weighting configuration

Some instance types support configurable bandwidth weighting, where you can select
baseline bandwidth weighting that favors either network processing or EBS operations. Default
settings for baseline bandwidth are determined by your instance type. You can configure the
bandwidth weighting during launch, or modify your instance settings with the following
weighting preferences:

- **default** – This option uses the standard
bandwidth configuration for your instance type.

- **vpc-1** – This option increases the baseline
bandwidth available for networking and decreases the baseline bandwidth for EBS operations.

- **ebs-1** – This option increases the baseline
bandwidth available for EBS operations, and decreases the baseline bandwidth for
networking.

## Bandwidth weighting considerations

The following are some considerations that might affect your bandwidth weighting strategy.

- Setting bandwidth weighting preferences only affects bandwidth specifications. The
network packets per second (PPS) and EBS input/output operations per second (IOPS)
specifications don't change.

- The combined bandwidth specification
between networking and EBS does not change. When you select a bandwidth weighting
configuration, the baseline bandwidth available for the selected option increases, and
the baseline bandwidth for the remaining option is reduced by the same absolute amount.
For all instances except Flex instances, the available burst bandwidth remains the same for your selected option, and is reduced for the remaining option. For Flex instances up to 4xlarge, burst bandwidth remains unchanged. For Flex instances 8xlarge and larger, burst bandwidth increases by the same amount as the baseline bandwidth.

- It's important to understand how changes in bandwidth allocation can affect I/O
performance for EBS. For EC2 instances that have `vpc-1` configuration
(increased networking bandwidth), you might experience lower IOPS for EBS volumes
if you reach the EBS bandwidth limit before you've reached the IOPS limit. This is
more noticeable with larger I/O sizes.

For example, on an instance type that normally supports 240,000 IOPS with 16 KiB I/O size,
if you select `vpc-1` weighting, that might reduce the achievable IOPS due to the
adjusted EBS baseline bandwidth limit.

When planning your workload, consider your I/O size and patterns. Smaller I/O sizes are less
likely to be affected by bandwidth limitations, while larger I/O sizes or sequential workloads
might see more impact from bandwidth changes. Always test your specific workload to ensure
optimal performance with your chosen configuration.

- The networking multi-flow bandwidth specification for traffic that goes through an internet
gateway or a local gateway is adjusted to 50% of the baseline bandwidth of the configured
option or 5 Gbps, where applicable. For more information, see [Amazon EC2 instance network bandwidth](ec2-instance-network-bandwidth.md).

The following example is based on an instance type that has a default baseline bandwidth of
40 Gbps, and a default border bandwidth of 20 Gbps. If you choose `vpc-1`
bandwidth weighting for this instance, the weighted baseline bandwidth changes to 50 Gbps,
and the border bandwidth changes to 25 Gbps.

- This feature is available in all commercial regions, aligned with EC2 instance availablilty and
support.

- This feature adds no additional cost to your EC2 instance.

## Supported instance types for bandwidth weighting

Instance types in the following instance families support configurable bandwidth weighting.

- **General purpose:** M8a, M8g, M8gd, M8i, M8id, M8i-flex

- **Compute optimized:** C8a, C8g, C8gd, C8i, C8id, C8i-flex

- **Memory optimized:** R8a, R8g, R8gd, R8i, R8id, R8i-flex, X8g, X8aedz, X8i

## Check current bandwidth settings

To see the current bandwidth settings for your instance, select one of the tabs for
instructions.

Console

###### To get the bandwidth setting for an instance

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. In the navigation pane, choose **Instances**.

3. Select the instance that you want to check from the list, and navigate
    to the **Networking** tab. Your current setting is shown in
    the **Configured bandwidth** field. Amazon EC2 uses default
    settings for your instance type if the bandwidth is not set to a specific
    value.

AWS CLI

###### To get the bandwidth setting for an instance

Use the [describe-instances](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-instances.html) command.

```nohighlight

aws ec2 describe-instances \
    --instance-ids i-1234567890abcdef0 \
    --query Reservations[].Instances[].NetworkPerformanceOptions.BandwidthWeighting \
    --output text
```

The following is example output.

```nohighlight

default
```

This example lists all of your instances that have the
bandwidth weighting preference set to `vpc-1`, for higher
networking bandwidth.

```nohighlight

aws ec2 describe-instances \
    --filters "Name=network-performance-options.bandwidth-weighting,Values=vpc-1" \
    --query Reservations[].Instances[].InstanceId \
    --output text
```

PowerShell

###### To get the bandwidth setting for an instance

Use the [Get-EC2Instance](https://docs.aws.amazon.com/powershell/latest/reference/items/Get-EC2Instance.html)
cmdlet.

```powershell

(Get-EC2Instance `
    -InstanceId i-1234567890abcdef0).Instances.NetworkPerformanceOptions.BandwidthWeighting.Value
```

The following is example output.

```nohighlight

default
```

This example lists all of your instances that have the
bandwidth weighting preference set to `vpc-1`, for higher
networking bandwidth.

```powershell

(Get-EC2Instance `
    -Filter @{Name="network-performance-options.bandwidth-weighting";Values="vpc-1"}).Instances.InstanceId
```

## Configure bandwidth weighting for your instance

You can configure bandwidth weighting either at launch or by modifying existing instances
from the EC2 console, API/SDKs or CLI.

### Configure bandwidth weighting when you launch an instance

To configure bandwidth settings when you launch an instance, select one of the tabs for
instructions.

You can also specify bandwidth weighting in a launch template. To create a
launch template, see [Create an Amazon EC2 launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-launch-template.html). The parameter to set is in the same
location as it is for launching an instance directly from the console. Expand
the **Advanced details** section, and set the **Instance**
**bandwidth configuration**.

To launch an instance with your launch template, see
[Launch EC2 instances using a launch template](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launch-instances-from-launch-template.html).

Console

###### To launch an instance with configurable bandwidth weighting

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. In the navigation pane, choose **Instances**.

3. Choose **Launch instances**. This opens the
    **Launch an instance** dialog. There are several
    additional ways that you can get to the launch dialog, depending
    on your preference. For example, you can launch an instance directly
    from an AMI or from the Amazon EC2 dashboard itself.

4. The Amazon Machine Image (AMI) that you launch from must be based
    on `Arm` architecture. Many **Quick Start**
    images support both `x86` and `Arm` architectures,
    After you choose the operating system for your instance, select the
    `Arm` option from the **Architecture** list.

5. The instance type must be one of the [Supported instance types](#config-bw-support) for
    this feature.

6. When you expand the **Advanced details** section,
    you can scroll down to find the **Instance bandwidth configuration**
    settings. Select the bandwidth configuration option for your instance.

7. Configure all of the other settings for your instance as you
    normally would, and choose **Launch instance**.

AWS CLI

###### To launch an instance with configurable bandwidth weighting

Use the [run-instances](https://docs.aws.amazon.com/cli/latest/reference/ec2/run-instances.html)
command with the following option to launch instances that are configured for
higher network bandwidth weighting.

```nohighlight

--network-performance-options BandwidthWeighting=vpc-1
```

Use the [run-instances](https://docs.aws.amazon.com/cli/latest/reference/ec2/run-instances.html)
command with the following option to launch instances that are configured for
higher EBS bandwidth weighting.

```nohighlight

--network-performance-options BandwidthWeighting=ebs-1
```

PowerShell

###### To launch an instance with configurable bandwidth weighting

Use the [New-EC2Instance](https://docs.aws.amazon.com/powershell/latest/reference/items/New-EC2Instance.html)
cmdlet with the following parameter to launch instances that are configured
for higher network bandwidth weighting.

```powershell

-NetworkPerformanceOptions_BandwidthWeighting vpc-1
```

Use the [New-EC2Instance](https://docs.aws.amazon.com/powershell/latest/reference/items/New-EC2Instance.html)
cmdlet with the following parameter to launch instances that are configured
for higher EBS bandwidth weighting.

```powershell

-NetworkPerformanceOptions_BandwidthWeighting ebs-1
```

### Update bandwidth weighting for an existing instance

To update bandwidth weighting for an existing instance, your instance must be in the
`Stopped` state.

Console

###### To update bandwidth weighting

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. In the navigation pane, choose **Instances**.

3. Select the instance that you want to update from the list.

4. Before you change the bandwidth configuration, your instance must be
    in a `Stopped` state. If your instance is running, select
    **Stop instance** from the **Instance**
**state** menu.

5. Choose **Manage bandwidth** from the
    **Actions > Networking** menu. This opens the
    **Manage bandwidth** dialog.

###### Note

If your instance type doesn't support configuration for bandwidth
weighting, that menu item is disabled.

6. Select the option to update your instance, and choose **Change**
    to save your settings.

AWS CLI

###### To update bandwidth weighting

Use the [modify-instance-network-performance-options](https://docs.aws.amazon.com/cli/latest/reference/ec2/modify-instance-network-performance-options.html) command to configure
higher network bandwidth weighting for the specified instance.

```nohighlight

aws ec2 modify-instance-network-performance-options \
    --instance-id i-1234567890abcdef0 \
    --bandwidth-weighting=vpc-1
```

The following example configures higher EBS bandwidth weighting
for the specified instance.

```nohighlight

aws ec2 modify-instance-network-performance-options \
    --instance-id i-1234567890abcdef0 \
    --bandwidth-weighting=ebs-1
```

PowerShell

###### To update bandwidth weighting

Use the [Edit-EC2InstanceNetworkPerformanceOption](https://docs.aws.amazon.com/powershell/latest/reference/items/Edit-EC2InstanceNetworkPerformanceOption.html)
cmdlet to configure higher network bandwidth weighting for the specified instance.

```powershell

Edit-EC2InstanceNetworkPerformanceOption `
    -InstanceId i-1234567890abcdef0 `
    -BandwidthWeighting vpc-1
```

The following example configures higher EBS bandwidth weighting
for the specified instance.

```powershell

Edit-EC2InstanceNetworkPerformanceOption `
    -InstanceId i-1234567890abcdef0 `
    -BandwidthWeighting ebs-1
```

## Networking performance

The following table shows the networking performance, in Gbps, that can be achieved with
the `default`, `vpc-1`, and `ebs-1` configurations.

Instance type

**`default`**

(Baseline / Burst)

**`vpc-1`**

(Baseline / Burst)

**`ebs-1`**

(Baseline / Burst)

c8a.medium0.52 / 12.50.65 / 12.50.438 / 10.0c8a.large0.937 / 12.51.172 / 12.50.774 / 10.0c8a.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0c8a.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5c8a.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5c8a.8xlarge15.018.7512.5c8a.12xlarge22.528.12518.75c8a.16xlarge30.037.525.0c8a.24xlarge40.050.032.5c8a.48xlarge75.093.7560.0c8a.metal-24xl40.050.032.5c8a.metal-48xl75.093.7560.0c8g.medium0.52 / 12.50.65 / 12.50.441 / 10.0c8g.large0.937 / 12.51.171 / 12.50.779 / 10.0c8g.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0c8g.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5c8g.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5c8g.8xlarge15.018.7512.5c8g.12xlarge22.528.12518.75c8g.16xlarge30.037.525.0c8g.24xlarge40.050.032.5c8g.48xlarge50.062.540.0c8g.metal-24xl40.050.032.5c8g.metal-48xl50.062.540.0c8gd.medium0.52 / 12.50.65 / 12.50.441 / 10.0c8gd.large0.937 / 12.51.171 / 12.50.779 / 10.0c8gd.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0c8gd.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5c8gd.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5c8gd.8xlarge15.018.7512.5c8gd.12xlarge22.528.12518.75c8gd.16xlarge30.037.525.0c8gd.24xlarge40.050.032.5c8gd.48xlarge50.062.540.0c8gd.metal-24xl40.050.032.5c8gd.metal-48xl50.062.540.0c8i.large0.937 / 12.51.172 / 12.50.774 / 10.0c8i.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0c8i.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5c8i.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5c8i.8xlarge15.018.7512.5c8i.12xlarge22.528.12518.75c8i.16xlarge30.037.525.0c8i.24xlarge40.050.032.5c8i.32xlarge50.062.540.0c8i.48xlarge75.093.7560.0c8i.96xlarge100.0125.080.0c8i.metal-48xl75.093.7560.0c8i.metal-96xl100.0125.080.0c8id.large0.937 / 12.51.172 / 12.50.774 / 10.0c8id.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0c8id.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5c8id.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5c8id.8xlarge15.018.7512.5c8id.12xlarge22.528.12518.75c8id.16xlarge30.037.525.0c8id.24xlarge40.050.032.5c8id.32xlarge50.062.540.0c8id.48xlarge75.093.7560.0c8id.96xlarge100.0125.080.0c8id.metal-48xl75.093.7560.0c8id.metal-96xl100.0125.080.0c8i-flex.large0.468 / 12.50.585 / 12.50.389 / 10.0c8i-flex.xlarge0.937 / 12.51.172 / 12.50.779 / 10.0c8i-flex.2xlarge1.875 / 15.02.344 / 15.01.562 / 12.5c8i-flex.4xlarge3.75 / 15.04.688 / 15.03.125 / 12.5c8i-flex.8xlarge7.5 / 15.09.375 / 18.756.25 / 12.5c8i-flex.12xlarge11.25 / 22.514.063 / 28.1259.375 / 18.75c8i-flex.16xlarge15.0 / 30.018.75 / 37.512.5 / 25.0m8a.medium0.52 / 12.50.65 / 12.50.438 / 10.0m8a.large0.937 / 12.51.172 / 12.50.774 / 10.0m8a.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0m8a.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5m8a.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5m8a.8xlarge15.018.7512.5m8a.12xlarge22.528.12518.75m8a.16xlarge30.037.525.0m8a.24xlarge40.050.032.5m8a.48xlarge75.093.7560.0m8a.metal-24xl40.050.032.5m8a.metal-48xl75.093.7560.0m8g.medium0.52 / 12.50.65 / 12.50.441 / 10.0m8g.large0.937 / 12.51.171 / 12.50.779 / 10.0m8g.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0m8g.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5m8g.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5m8g.8xlarge15.018.7512.5m8g.12xlarge22.528.12518.75m8g.16xlarge30.037.525.0m8g.24xlarge40.050.032.5m8g.48xlarge50.062.540.0m8g.metal-24xl40.050.032.5m8g.metal-48xl50.062.540.0m8gd.medium0.52 / 12.50.65 / 12.50.441 / 10.0m8gd.large0.937 / 12.51.171 / 12.50.779 / 10.0m8gd.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0m8gd.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5m8gd.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5m8gd.8xlarge15.018.7512.5m8gd.12xlarge22.528.12518.75m8gd.16xlarge30.037.525.0m8gd.24xlarge40.050.032.5m8gd.48xlarge50.062.540.0m8gd.metal-24xl40.050.032.5m8gd.metal-48xl50.062.540.0m8i.large0.937 / 12.51.172 / 12.50.774 / 10.0m8i.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0m8i.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5m8i.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5m8i.8xlarge15.018.7512.5m8i.12xlarge22.528.12518.75m8i.16xlarge30.037.525.0m8i.24xlarge40.050.032.5m8i.32xlarge50.062.540.0m8i.48xlarge75.093.7560.0m8i.96xlarge100.0125.080.0m8i.metal-48xl75.093.7560.0m8i.metal-96xl100.0125.080.0m8id.large0.937 / 12.51.172 / 12.50.774 / 10.0m8id.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0m8id.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5m8id.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5m8id.8xlarge15.018.7512.5m8id.12xlarge22.528.12518.75m8id.16xlarge30.037.525.0m8id.24xlarge40.050.032.5m8id.32xlarge50.062.540.0m8id.48xlarge75.093.7560.0m8id.96xlarge100.0125.080.0m8id.metal-48xl75.093.7560.0m8id.metal-96xl100.0125.080.0m8i-flex.large0.468 / 12.50.585 / 12.50.389 / 10.0m8i-flex.xlarge0.937 / 12.51.172 / 12.50.779 / 10.0m8i-flex.2xlarge1.875 / 15.02.344 / 15.01.562 / 12.5m8i-flex.4xlarge3.75 / 15.04.688 / 15.03.125 / 12.5m8i-flex.8xlarge7.5 / 15.09.375 / 18.756.25 / 12.5m8i-flex.12xlarge11.25 / 22.514.063 / 28.1259.375 / 18.75m8i-flex.16xlarge15.0 / 30.018.75 / 37.512.5 / 25.0r8a.medium0.52 / 12.50.65 / 12.50.438 / 10.0r8a.large0.937 / 12.51.172 / 12.50.774 / 10.0r8a.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0r8a.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5r8a.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5r8a.8xlarge15.018.7512.5r8a.12xlarge22.528.12518.75r8a.16xlarge30.037.525.0r8a.24xlarge40.050.032.5r8a.48xlarge75.093.7560.0r8a.metal-24xl40.050.032.5r8a.metal-48xl75.093.7560.0r8g.medium0.52 / 12.50.65 / 12.50.441 / 10.0r8g.large0.937 / 12.51.171 / 12.50.779 / 10.0r8g.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0r8g.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5r8g.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5r8g.8xlarge15.018.7512.5r8g.12xlarge22.528.12518.75r8g.16xlarge30.037.525.0r8g.24xlarge40.050.032.5r8g.48xlarge50.062.540.0r8g.metal-24xl40.050.032.5r8g.metal-48xl50.062.540.0r8gd.medium0.52 / 12.50.65 / 12.50.441 / 10.0r8gd.large0.937 / 12.51.171 / 12.50.779 / 10.0r8gd.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0r8gd.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5r8gd.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5r8gd.8xlarge15.018.7512.5r8gd.12xlarge22.528.12518.75r8gd.16xlarge30.037.525.0r8gd.24xlarge40.050.032.5r8gd.48xlarge50.062.540.0r8gd.metal-24xl40.050.032.5r8gd.metal-48xl50.062.540.0r8i.large0.937 / 12.51.172 / 12.50.774 / 10.0r8i.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0r8i.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5r8i.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5r8i.8xlarge15.018.7512.5r8i.12xlarge22.528.12518.75r8i.16xlarge30.037.525.0r8i.24xlarge40.050.032.5r8i.32xlarge50.062.540.0r8i.48xlarge75.093.7560.0r8i.96xlarge100.0125.080.0r8i.metal-48xl75.093.7560.0r8i.metal-96xl100.0125.080.0r8id.large0.937 / 12.51.172 / 12.50.774 / 10.0r8id.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0r8id.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5r8id.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5r8id.8xlarge15.018.7512.5r8id.12xlarge22.528.12518.75r8id.16xlarge30.037.525.0r8id.24xlarge40.050.032.5r8id.32xlarge50.062.540.0r8id.48xlarge75.093.7560.0r8id.96xlarge100.0125.080.0r8id.metal-48xl75.093.7560.0r8id.metal-96xl100.0125.080.0r8i-flex.large0.468 / 12.50.585 / 12.50.389 / 10.0r8i-flex.xlarge0.937 / 12.51.172 / 12.50.779 / 10.0r8i-flex.2xlarge1.875 / 15.02.344 / 15.01.562 / 12.5r8i-flex.4xlarge3.75 / 15.04.688 / 15.03.125 / 12.5r8i-flex.8xlarge7.5 / 15.09.375 / 18.756.25 / 12.5r8i-flex.12xlarge11.25 / 22.514.063 / 28.1259.375 / 18.75r8i-flex.16xlarge15.0 / 30.018.75 / 37.512.5 / 25.0x8g.medium0.52 / 12.50.65 / 12.50.441 / 10.0x8g.large0.937 / 12.51.171 / 12.50.779 / 10.0x8g.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0x8g.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5x8g.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5x8g.8xlarge15.018.7512.5x8g.12xlarge22.528.12518.75x8g.16xlarge30.037.525.0x8g.24xlarge40.050.032.5x8g.48xlarge50.062.540.0x8g.metal-24xl40.050.032.5x8g.metal-48xl50.062.540.0x8aedz.large1.562 / 18.751.953 / 18.751.249 / 15.0x8aedz.xlarge3.125 / 18.753.907 / 18.752.5 / 15.0x8aedz.3xlarge9.375 / 18.7511.719 / 18.757.5 / 15.0x8aedz.6xlarge18.7523.43815.0x8aedz.12xlarge37.546.87530.0x8aedz.24xlarge75.093.7560.0x8aedz.metal-12xl37.546.87530.0x8aedz.metal-24xl75.093.7560.0x8i.large0.937 / 12.51.172 / 12.50.774 / 10.0x8i.xlarge1.875 / 12.52.344 / 12.51.562 / 10.0x8i.2xlarge3.75 / 15.04.688 / 15.03.125 / 12.5x8i.4xlarge7.5 / 15.09.375 / 15.06.25 / 12.5x8i.8xlarge15.018.7512.5x8i.12xlarge22.528.12518.75x8i.16xlarge30.037.525.0x8i.24xlarge40.050.032.5x8i.32xlarge50.062.540.0x8i.48xlarge75.093.7560.0x8i.64xlarge80.0100.062.5x8i.96xlarge100.0125.080.0x8i.metal-48xl75.093.7560.0x8i.metal-96xl100.0125.080.0

## Amazon EBS performance

The following table shows the Amazon EBS performance, in Gbps, that can be achieved with
the `default`, `vpc-1`, and `ebs-1` configurations.

Instance type

**`default`**

(Baseline / Burst)

**`vpc-1`**

(Baseline / Burst)

**`ebs-1`**

(Baseline / Burst)

c8a.medium0.325 / 10.00.195 / 6.250.407 / 10.0c8a.large0.65 / 10.00.415 / 6.250.813 / 10.0c8a.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0c8a.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0c8a.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0c8a.8xlarge10.06.2512.5c8a.12xlarge15.09.37518.75c8a.16xlarge20.012.525.0c8a.24xlarge30.020.037.5c8a.48xlarge60.041.2575.0c8a.metal-24xl30.020.037.5c8a.metal-48xl60.041.2575.0c8g.medium0.315 / 10.00.185 / 6.250.394 / 10.0c8g.large0.63 / 10.00.396 / 6.250.788 / 10.0c8g.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0c8g.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0c8g.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0c8g.8xlarge10.06.2512.5c8g.12xlarge15.09.37518.75c8g.16xlarge20.012.525.0c8g.24xlarge30.020.037.5c8g.48xlarge40.027.550.0c8g.metal-24xl30.020.037.5c8g.metal-48xl40.027.550.0c8gd.medium0.315 / 10.00.185 / 6.250.394 / 10.0c8gd.large0.63 / 10.00.396 / 6.250.788 / 10.0c8gd.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0c8gd.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0c8gd.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0c8gd.8xlarge10.06.2512.5c8gd.12xlarge15.09.37518.75c8gd.16xlarge20.012.525.0c8gd.24xlarge30.020.037.5c8gd.48xlarge40.027.550.0c8gd.metal-24xl30.020.037.5c8gd.metal-48xl40.027.550.0c8i.large0.65 / 10.00.415 / 6.250.813 / 10.0c8i.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0c8i.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0c8i.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0c8i.8xlarge10.06.2512.5c8i.12xlarge15.09.37518.75c8i.16xlarge20.012.525.0c8i.24xlarge30.020.037.5c8i.32xlarge40.027.550.0c8i.48xlarge60.041.2575.0c8i.96xlarge80.055.0100.0c8i.metal-48xl60.041.2575.0c8i.metal-96xl80.055.0100.0c8id.large0.65 / 10.00.415 / 6.250.813 / 10.0c8id.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0c8id.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0c8id.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0c8id.8xlarge10.06.2512.5c8id.12xlarge15.09.37518.75c8id.16xlarge20.012.525.0c8id.24xlarge30.020.037.5c8id.32xlarge40.027.550.0c8id.48xlarge60.041.2575.0c8id.96xlarge80.055.0100.0c8id.metal-48xl60.041.2575.0c8id.metal-96xl80.055.0100.0c8i-flex.large0.315 / 10.00.198 / 6.250.394 / 10.0c8i-flex.xlarge0.63 / 10.00.395 / 6.250.788 / 10.0c8i-flex.2xlarge1.25 / 10.00.781 / 6.251.563 / 10.0c8i-flex.4xlarge2.5 / 10.01.562 / 6.253.125 / 10.0c8i-flex.8xlarge5.0 / 10.03.125 / 6.256.25 / 12.5c8i-flex.12xlarge7.5 / 15.04.687 / 9.3759.375 / 18.75c8i-flex.16xlarge10.0 / 20.06.25 / 12.512.5 / 25.0m8a.medium0.325 / 10.00.195 / 6.250.407 / 10.0m8a.large0.65 / 10.00.415 / 6.250.813 / 10.0m8a.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0m8a.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0m8a.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0m8a.8xlarge10.06.2512.5m8a.12xlarge15.09.37518.75m8a.16xlarge20.012.525.0m8a.24xlarge30.020.037.5m8a.48xlarge60.041.2575.0m8a.metal-24xl30.020.037.5m8a.metal-48xl60.041.2575.0m8g.medium0.315 / 10.00.185 / 6.250.394 / 10.0m8g.large0.63 / 10.00.396 / 6.250.788 / 10.0m8g.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0m8g.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0m8g.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0m8g.8xlarge10.06.2512.5m8g.12xlarge15.09.37518.75m8g.16xlarge20.012.525.0m8g.24xlarge30.020.037.5m8g.48xlarge40.027.550.0m8g.metal-24xl30.020.037.5m8g.metal-48xl40.027.550.0m8gd.medium0.315 / 10.00.185 / 6.250.394 / 10.0m8gd.large0.63 / 10.00.396 / 6.250.788 / 10.0m8gd.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0m8gd.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0m8gd.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0m8gd.8xlarge10.06.2512.5m8gd.12xlarge15.09.37518.75m8gd.16xlarge20.012.525.0m8gd.24xlarge30.020.037.5m8gd.48xlarge40.027.550.0m8gd.metal-24xl30.020.037.5m8gd.metal-48xl40.027.550.0m8i.large0.65 / 10.00.415 / 6.250.813 / 10.0m8i.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0m8i.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0m8i.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0m8i.8xlarge10.06.2512.5m8i.12xlarge15.09.37518.75m8i.16xlarge20.012.525.0m8i.24xlarge30.020.037.5m8i.32xlarge40.027.550.0m8i.48xlarge60.041.2575.0m8i.96xlarge80.055.0100.0m8i.metal-48xl60.041.2575.0m8i.metal-96xl80.055.0100.0m8id.large0.65 / 10.00.415 / 6.250.813 / 10.0m8id.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0m8id.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0m8id.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0m8id.8xlarge10.06.2512.5m8id.12xlarge15.09.37518.75m8id.16xlarge20.012.525.0m8id.24xlarge30.020.037.5m8id.32xlarge40.027.550.0m8id.48xlarge60.041.2575.0m8id.96xlarge80.055.0100.0m8id.metal-48xl60.041.2575.0m8id.metal-96xl80.055.0100.0m8i-flex.large0.315 / 10.00.198 / 6.250.394 / 10.0m8i-flex.xlarge0.63 / 10.00.395 / 6.250.788 / 10.0m8i-flex.2xlarge1.25 / 10.00.781 / 6.251.563 / 10.0m8i-flex.4xlarge2.5 / 10.01.562 / 6.253.125 / 10.0m8i-flex.8xlarge5.0 / 10.03.125 / 6.256.25 / 12.5m8i-flex.12xlarge7.5 / 15.04.687 / 9.3759.375 / 18.75m8i-flex.16xlarge10.0 / 20.06.25 / 12.512.5 / 25.0r8a.medium0.325 / 10.00.195 / 6.250.407 / 10.0r8a.large0.65 / 10.00.415 / 6.250.813 / 10.0r8a.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0r8a.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0r8a.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0r8a.8xlarge10.06.2512.5r8a.12xlarge15.09.37518.75r8a.16xlarge20.012.525.0r8a.24xlarge30.020.037.5r8a.48xlarge60.041.2575.0r8a.metal-24xl30.020.037.5r8a.metal-48xl60.041.2575.0r8g.medium0.315 / 10.00.185 / 6.250.394 / 10.0r8g.large0.63 / 10.00.396 / 6.250.788 / 10.0r8g.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0r8g.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0r8g.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0r8g.8xlarge10.06.2512.5r8g.12xlarge15.09.37518.75r8g.16xlarge20.012.525.0r8g.24xlarge30.020.037.5r8g.48xlarge40.027.550.0r8g.metal-24xl30.020.037.5r8g.metal-48xl40.027.550.0r8gd.medium0.315 / 10.00.185 / 6.250.394 / 10.0r8gd.large0.63 / 10.00.396 / 6.250.788 / 10.0r8gd.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0r8gd.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0r8gd.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0r8gd.8xlarge10.06.2512.5r8gd.12xlarge15.09.37518.75r8gd.16xlarge20.012.525.0r8gd.24xlarge30.020.037.5r8gd.48xlarge40.027.550.0r8gd.metal-24xl30.020.037.5r8gd.metal-48xl40.027.550.0r8i.large0.65 / 10.00.415 / 6.250.813 / 10.0r8i.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0r8i.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0r8i.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0r8i.8xlarge10.06.2512.5r8i.12xlarge15.09.37518.75r8i.16xlarge20.012.525.0r8i.24xlarge30.020.037.5r8i.32xlarge40.027.550.0r8i.48xlarge60.041.2575.0r8i.96xlarge80.055.0100.0r8i.metal-48xl60.041.2575.0r8i.metal-96xl80.055.0100.0r8id.large0.65 / 10.00.415 / 6.250.813 / 10.0r8id.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0r8id.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0r8id.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0r8id.8xlarge10.06.2512.5r8id.12xlarge15.09.37518.75r8id.16xlarge20.012.525.0r8id.24xlarge30.020.037.5r8id.32xlarge40.027.550.0r8id.48xlarge60.041.2575.0r8id.96xlarge80.055.0100.0r8id.metal-48xl60.041.2575.0r8id.metal-96xl80.055.0100.0r8i-flex.large0.315 / 10.00.198 / 6.250.394 / 10.0r8i-flex.xlarge0.63 / 10.00.395 / 6.250.788 / 10.0r8i-flex.2xlarge1.25 / 10.00.781 / 6.251.563 / 10.0r8i-flex.4xlarge2.5 / 10.01.562 / 6.253.125 / 10.0r8i-flex.8xlarge5.0 / 10.03.125 / 6.256.25 / 12.5r8i-flex.12xlarge7.5 / 15.04.687 / 9.3759.375 / 18.75r8i-flex.16xlarge10.0 / 20.06.25 / 12.512.5 / 25.0x8g.medium0.315 / 10.00.185 / 6.250.394 / 10.0x8g.large0.63 / 10.00.396 / 6.250.788 / 10.0x8g.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0x8g.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0x8g.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0x8g.8xlarge10.06.2512.5x8g.12xlarge15.09.37518.75x8g.16xlarge20.012.525.0x8g.24xlarge30.020.037.5x8g.48xlarge40.027.550.0x8g.metal-24xl30.020.037.5x8g.metal-48xl40.027.550.0x8aedz.large1.25 / 15.00.859 / 10.3121.563 / 15.0x8aedz.xlarge2.5 / 15.01.718 / 10.3123.125 / 15.0x8aedz.3xlarge7.5 / 15.05.156 / 10.3129.375 / 15.0x8aedz.6xlarge15.010.31218.75x8aedz.12xlarge30.020.62537.5x8aedz.24xlarge60.041.2575.0x8aedz.metal-12xl30.020.62537.5x8aedz.metal-24xl60.041.2575.0x8i.large0.65 / 10.00.415 / 6.250.813 / 10.0x8i.xlarge1.25 / 10.00.781 / 6.251.563 / 10.0x8i.2xlarge2.5 / 10.01.562 / 6.253.125 / 10.0x8i.4xlarge5.0 / 10.03.125 / 6.256.25 / 10.0x8i.8xlarge10.06.2512.5x8i.12xlarge15.09.37518.75x8i.16xlarge20.012.525.0x8i.24xlarge30.020.037.5x8i.32xlarge40.027.550.0x8i.48xlarge60.041.2575.0x8i.64xlarge70.050.087.5x8i.96xlarge80.055.0100.0x8i.metal-48xl60.041.2575.0x8i.metal-96xl80.055.0100.0

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Network bandwidth

Enhanced networking
