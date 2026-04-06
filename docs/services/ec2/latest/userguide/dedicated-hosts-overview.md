# Amazon EC2 Dedicated Hosts

An Amazon EC2 Dedicated Host is a physical server that is fully dedicated for your use. You can optionally choose to share the instance capacity with other
AWS accounts. For more information, see [Cross-account Amazon EC2 Dedicated Host sharing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dh-sharing.html).

Dedicated Hosts provide visibility and control over instance placement and they support host
affinity. This means that you can launch and run instances on specific hosts, and you can
ensure that instances run only on specific hosts. For more information, see [Amazon EC2 Dedicated Host auto-placement and host affinity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-understanding.html).

Dedicated Hosts provide comprehensive Bring Your Own License (BYOL) support. They allow you to use
your existing per-socket, per-core, or per-VM software licenses, including Windows Server,
SQL Server, SUSE Linux Enterprise Server, Red Hat Enterprise Linux, or other software
licenses that are bound to VMs, sockets, or physical cores, subject to your license
terms.

If you require your instances to run on dedicated hardware, but you do not need visibility
or control over instance placement, and you do not need to use per-socket or per-core
software licenses, you can consider using Dedicated Instances instead. Dedicated Instances and Dedicated Hosts can both be used to
launch Amazon EC2 instances onto dedicated physical servers. There are no performance, security,
or physical differences between Dedicated Instances and instances on Dedicated Hosts. However, there are some key
differences between them. The following table highlights some of the key differences between
Dedicated Instances and Dedicated Hosts:

Dedicated HostDedicated Instance

Dedicated physical server

Physical server with instance capacity fully dedicated to
your use.

Physical server that's dedicated to a single customer
account.

Instance capacity sharing

Can share instance capacity with other accounts.

Not supported

Billing

Per-host billing

Per-instance billing

Visibility of sockets, cores, and host ID

Provides visibility of the number of sockets and physical
cores

No visibility

Host and instance affinity

Allows you to consistently deploy your instances to the same
physical server over time

Not supported

Targeted instance placement

Provides additional visibility and control over how instances
are placed on a physical server

Not supported

Automatic instance recovery

Supported. For more information, see [Amazon EC2 Dedicated Host recovery](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery.html).

Supported

Bring Your Own License (BYOL)

Supported

Partial support \*

Capacity Reservations

Not supported

Supported

\\* Microsoft SQL Server with License Mobility through Software Assurance, and Windows Virtual Desktop
Access (VDA) licenses can be used with Dedicated Instance.

For more information about Dedicated Instances, see [Amazon EC2 Dedicated Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-instance.html).

## Dedicated Hosts restrictions

Before you allocate Dedicated Hosts, take note of the following limitations and
restrictions:

- To run RHEL and SUSE Linux on Dedicated Hosts, you must bring your own AMIs.
You can't use the RHEL and SUSE Linux AMIs that are offered by AWS or that are
available on AWS Marketplace with Dedicated Hosts. For more information about how to create your own
AMI, see [Bring your own software licenses to Amazon EC2 Dedicated Hosts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-BYOL.html).

This restriction does not apply to hosts allocated for
high memory instances ( `u-6tb1.metal`, `u-9tb1.metal`,
`u-12tb1.metal`, `u-18tb1.metal`, and
`u-24tb1.metal`). RHEL and SUSE Linux AMIs that are offered by
AWS or that are available on AWS Marketplace can be used with these hosts.

- There is a limit on the number of running Dedicated Hosts per instance family per AWS
account per Region. Quotas apply to running instances only. If your instance is
pending, stopping, or stopped, it does not count towards your quota. To view the
quotas for your account, or to request a quota increase, use the [Service Quotas\
console](https://console.aws.amazon.com/servicequotas/home/services/ec2/quotas).

- Auto Scaling groups are supported when using a launch template that specifies a host
resource group. For more information, see [Create a launch template using advanced settings](https://docs.aws.amazon.com/autoscaling/ec2/userguide/advanced-settings-for-your-launch-template.html) in the
_Amazon EC2 Auto Scaling User Guide_.

- Amazon RDS instances are not supported.

- The AWS Free Usage tier is not available for Dedicated Hosts.

- Instance placement control refers to managing instance launches onto Dedicated Hosts.
You can't launch Dedicated Hosts into placement groups.

- If you allocate a host for a virtualized instance type, you can't modify the
instance type to a `.metal` instance type after the host is
allocated. For example, if you allocate a host for the `m5.large`
instance type, you can't modify the instance type to
`m5.metal`.

Similarly, if you allocate a host for a `.metal` instance type, you
can't modify the instance type to a virtualized instance type after the host is
allocated. For example, if you allocate a host for the `m5.metal`
instance type, you can't modify the instance type to
`m5.large`.

###### Contents

- [Pricing and billing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-billing.html)

- [Instance capacity configurations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-limits.html)

- [Burstable instances on Dedicated Hosts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-t3.html)

- [Bring your own licenses](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-BYOL.html)

- [Auto-placement and affinity](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-understanding.html)

- [Allocate a Dedicated Host](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-allocating.html)

- [Launch instances on a Dedicated Host](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launching-dedicated-hosts-instances.html)

- [Launch instances into a host resource group](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launching-hrg-instances.html)

- [Modify Dedicated Host auto-placement](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/modify-host-auto-placement.html)

- [Modify supported instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/modify-host-support.html)

- [Modify tenancy and affinity for an instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/moving-instances-dedicated-hosts.html)

- [Release Dedicated Host](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-releasing.html)

- [Migrate to Nitro-based Amazon EC2 Dedicated Hosts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dh-migrate.html)

- [Purchase a Dedicated Host Reservation](#purchasing-dedicated-host-reservations)

- [Cross-account sharing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dh-sharing.html)

- [Dedicated Hosts on Outposts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dh-outposts.html)

- [Host recovery](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery.html)

- [Host maintenance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-maintenance.html)

- [Monitor Dedicated Hosts](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-monitoring.html)

- [Track configuration changes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-aws-config.html)

## Purchase Dedicated Host Reservations for Dedicated Host billing discounts

Dedicated Host Reservations provide you with a discount of up to 70 percent compared to On-Demand
Dedicated Host pricing. You must have active Dedicated Hosts allocated in your account before you can
purchase Dedicated Host Reservations. For more information, see [Dedicated Host Reservations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-billing.html#dedicated-host-reservations).

Console

###### To purchase reservations

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. Choose **Dedicated Hosts**,
    **Dedicated Host Reservations**, **Purchase**
**Dedicated Host Reservation**.

3. On the **Find offerings** screen, do the
    following:
1. For **Instance family**, select the
       instance family of the Dedicated Host for which to purchase the
       Dedicated Host Reservation.

2. For **Payment option**, select and
       configure your preferred payment option.
4. Choose **Next**.

5. Select the Dedicated Hosts with which to associate the Dedicated Host Reservation, and then
    choose **Next**.

6. ( _Optional_) Assign tags to the
    Dedicated Host Reservation.

7. Review your order and choose
    **Purchase**.

AWS CLI

###### To purchase reservations

1. Use the [describe-host-reservation-offerings](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-host-reservation-offerings.html) command
    to list the available offerings that match your needs. The
    following example lists the offerings that support instances in
    the `m4` instance family and have a one-year
    term.

The term is specified in seconds. A one-year term includes
    31,536,000 seconds, and a three-year term includes
    94,608,000 seconds.

```nohighlight

aws ec2 describe-host-reservation-offerings \
       --filter Name=instance-family,Values=m4 \
       --max-duration 31536000
```

The command returns a list of offerings that match your
    criteria. Note the ID of the offering to purchase.

2. Use the [purchase-host-reservation](https://docs.aws.amazon.com/cli/latest/reference/ec2/purchase-host-reservation.html) command to purchase
    the offering and provide the `offeringId` noted in
    the previous step. The following example purchases the specified
    reservation and associates it with a specific Dedicated Host that is
    already allocated in the AWS account, and it applies a tag
    with a key of `purpose` and a value of
    `production`.

```nohighlight

aws ec2 purchase-host-reservation \
       --offering-id hro-03f707bf363b6b324 \
       --host-id-set h-013abcd2a00cbd123 \
       --tag-specifications 'ResourceType=host-reservation,Tags={Key=purpose,Value=production}'
```

PowerShell

###### To purchase reservations

1. Use the [Get-EC2HostReservationOffering](https://docs.aws.amazon.com/powershell/latest/reference/items/Get-EC2HostReservationOffering.html) cmdlet to
    list the available offerings that match your needs. The
    following examples list the offerings that support instances in
    the `m5` instance family and have a one-year
    term.

The term is specified in seconds. A one-year term includes
    31,536,000 seconds, and a three-year term includes
    94,608,000 seconds.

```powershell

$filter = @{Name="instance-family"; Values="m5"}
Get-EC2HostReservationOffering `
       -Filter $filter `
       -MaxDuration 31536000
```

The command returns a list of offerings that match your
    criteria. Note the ID of the offering to purchase.

2. Use the [New-EC2HostReservation](https://docs.aws.amazon.com/powershell/latest/reference/items/New-EC2HostReservation.html) cmdlet to
    purchase the offering and provide the offering ID
    noted in the previous step. The following example purchases the
    specified reservation and associates it with a specific Dedicated Host
    that is already allocated in the AWS account.

```powershell

New-EC2HostReservation `
       -OfferingId hro-03f707bf363b6b324 `
       -HostIdSet h-013abcd2a00cbd123
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Spot Instance quotas

Pricing and billing
