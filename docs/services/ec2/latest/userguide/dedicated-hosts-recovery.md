# Amazon EC2 Dedicated Host recovery

Dedicated Host auto recovery restarts your instances on to a new replacement host when certain
problematic conditions are detected on your Dedicated Host. Host recovery reduces the need for
manual intervention and lowers the operational burden if there is an unexpected Dedicated Host
failure concerning system power or network connectivity events. Other Dedicated Host issues will
require manual intervention to recover from.

###### Contents

- [How host recovery works](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery-basics.html)

- [Supported instance types](#dedicated-hosts-recovery-instances)

- [Pricing](#dedicated-hosts-recovery-pricing)

- [Manage host recovery](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery-enable.html)

- [View host recovery setting](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery-view.html)

- [Manually recovery unsupported instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery-unsupported.html)

## Supported instance types

Host recovery is supported for the following instance families:

- **General purpose:** A1 \| M3 \| M4 \| M5 \| M5n \| M5zn \| M6a \| M6g \| M6i \| T3 \| Mac1 \| Mac2 \| Mac2-m1ultra \| Mac2-m2 \| Mac2-m2pro

- **Compute optimized:** C3 \| C4 \| C5 \| C5n \| C6a \| C6g \| C6i

- **Memory optimized:** R3 \| R4 \| R5 \| R5b \| R5n \| R6g \| R6i \| U-6tb1 \| U-9tb1 \| U-12tb1 \| U-18tb1 \| U-24tb1 \| X1 \| X1e \| X2iezn

- **Accelerated computing:** Inf1 \| G3 \| G5g \| P2 \| P3

To recover instances that are not supported, see [Manually recover instances that are not supported by Amazon EC2 Dedicated Host recovery](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/dedicated-hosts-recovery-unsupported.html).

###### Note

Dedicated Host auto recovery of supported metal instance types
will take longer to detect and recover from than non-metal instance
types.

## Pricing

There are no additional charges for using host recovery, but the usual Dedicated Host
charges apply. For more information, see [Amazon EC2 Dedicated Hosts\
Pricing](https://aws.amazon.com/ec2/dedicated-hosts/pricing).

As soon as host recovery is initiated, you are no longer billed for the impaired
Dedicated Host. Billing for the replacement Dedicated Host begins only after it enters the
`available` state.

If the impaired Dedicated Host was billed using the On-Demand rate, the replacement Dedicated Host is
also billed using the On-Demand rate. If the impaired Dedicated Host had an active Dedicated Host Reservation, it
is transferred to the replacement Dedicated Host.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Allocate Dedicated Host on Outpost

How host recovery works
