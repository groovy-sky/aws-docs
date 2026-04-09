# Amazon WorkSpaces Applications Service Quotas

WorkSpaces Applications provides different resources that you can use. WorkSpaces Applications resources include stacks, fleets, images, and image builders. When you create your Amazon Web Services account, we set default quotas (also referred to as limits) on the number of resources that you can create, and on the number of users who can use the WorkSpaces Applications service.

To request a quota increase, you can use the Service Quotas console at [https://console.aws.amazon.com/servicequotas/](https://console.aws.amazon.com/servicequotas). For more information, see [Requesting a quota increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the _Service Quotas User Guide_.

The following table lists the default quotas for each WorkSpaces Applications resource and for users in the
WorkSpaces Applications user pool. The actual quotas for your account could be higher or lower, depending on
when you created your account.

###### Note

Graphics Pro instances are no longer available after 10/31/2025 due to End of Life of
hardware supporting Graphics Pro instance types. WorkSpaces Applications does not accept limit increase
requests for Graphics Pro instances.

Graphics Design instances will no longer be available from AWS after 12/31/2025 due
to End of Life of hardware supporting Graphics Design instance types. WorkSpaces Applications does not
accept limit increase requests for Graphics Design instances.

NameDefaultAdjustableStacks10[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-A8B8B901)Fleets10[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-3FEADC0C)Compute-optimized fleet instances \*

- stream.compute.large: 5

- stream.compute.xlarge: 2

- stream.compute.2xlarge: 0

- stream.compute.4xlarge: 0

- stream.compute.8xlarge: 0

YesGraphics fleet instances \*

- stream.graphics-design.large: 3

- stream.graphics-design.xlarge: 3

- stream.graphics-design.2xlarge: 3

- stream.graphics-design.4xlarge: 0

- stream.graphics.g4dn.xlarge: 0

- stream.graphics.g4dn.2xlarge: 0

- stream.graphics.g4dn.4xlarge: 0

- stream.graphics.g4dn.8xlarge: 0

- stream.graphics.g4dn.12xlarge: 0

- stream.graphics.g4dn.16xlarge: 0

- stream.graphics.g5.xlarge: 0

- stream.graphics.g5.2xlarge: 0

- stream.graphics.g5.4xlarge: 0

- stream.graphics.g5.8xlarge: 0

- stream.graphics.g5.12xlarge: 0

- stream.graphics.g5.16xlarge: 0

- stream.graphics.g5.24xlarge: 0

- stream.graphics.g6.xlarge: 0

- stream.graphics.g6.2xlarge: 0

- stream.graphics.g6.4xlarge: 0

- stream.graphics.g6.8xlarge: 0

- stream.graphics.g6.16xlarge: 0

- stream.graphics.g6.12xlarge: 0

- stream.graphics.g6.24xlarge: 0

- stream.graphics.gr6.4xlarge: 0

- stream.graphics.gr6.8xlarge: 0

- stream.graphics.g6f.large: 0

- stream.graphics.g6f.xlarge: 0

- stream.graphics.g6f.2xlarge: 0

- stream.graphics.g6f.4xlarge: 0

- stream.graphics.gr6f.4xlarge: 0

YesMemory-optimized fleet instances \*

- stream.memory.large: 5

- stream.memory.xlarge: 2

- stream.memory.2xlarge: 0

- stream.memory.4xlarge: 0

- stream.memory.8xlarge: 0

- stream.memory.z1d.large: 5

- stream.memory.z1d.xlarge: 2

- stream.memory.z1d.2xlarge: 0

- stream.memory.z1d.3xlarge: 0

- stream.memory.z1d.6xlarge: 0

- stream.memory.z1d.12xlarge: 0

YesStandard fleet instances \*

- stream.standard.small: 50

- stream.standard.medium: 50

- stream.standard.large: 50

- stream.standard.xlarge: 10

- stream.standard.2xlarge: 10

YesGeneral purpose fleet instances\*

- t3/t3a: small/medium/large/xlarge/2xlarge: 50/50/10/10/10

- m6i/m7i: large/xlarge/2xlarge+: 5/2/0

- m7i-flex: large/xlarge/2xlarge+: 5/2/0

- m6a/m7a: large/xlarge/2xlarge+: 5/2/0

YesCompute optimize fleet instances\*

- c6i/c7i: large/xlarge/2xlarge+: 5/2/0

- c6a/c7a: large/xlarge/2xlarge+: 5/2/0

- c7i-flex: large/xlarge/2xlarge+: 5/2/0

YesMemory optimize fleet instances\*

- r6i/r7i: large/xlarge/2xlarge+: 5/2/0

- r6a/r7a: large/xlarge/2xlarge+: 5/2/0

YesAccelerated fleet instances\*

- G4dn/G5/G6/G6f/Gr6/Gr6f/G6e: All sizes: 0

YesImage builders (total)10[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-DE32F884)Images10[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-E1182CCA)Compute-optimized image builder instances

- stream.compute.large: 3

- stream.compute.xlarge: 3

- stream.compute.2xlarge: 0

- stream.compute.4xlarge: 0

- stream.compute.8xlarge: 0

YesGraphics image builder instances

- stream.graphics-design.large: 1

- stream.graphics-design.xlarge: 1

- stream.graphics-design.2xlarge: 1

- stream.graphics-design.4xlarge: 0

- stream.graphics.g4dn.xlarge: 0

- stream.graphics.g4dn.2xlarge: 0

- stream.graphics.g4dn.4xlarge: 0

- stream.graphics.g4dn.8xlarge: 0

- stream.graphics.g4dn.12xlarge: 0

- stream.graphics.g4dn.16xlarge: 0

- stream.graphics.g5.xlarge: 0

- stream.graphics.g5.2xlarge: 0

- stream.graphics.g5.4xlarge: 0

- stream.graphics.g5.8xlarge: 0

- stream.graphics.g5.12xlarge: 0

- stream.graphics.g5.16xlarge: 0

- stream.graphics.g5.24xlarge: 0

- stream.graphics.g6.xlarge: 0

- stream.graphics.g6.2xlarge: 0

- stream.graphics.g6.4xlarge: 0

- stream.graphics.g6.8xlarge: 0

- stream.graphics.g6.16xlarge: 0

- stream.graphics.g6.12xlarge: 0

- stream.graphics.g6.24xlarge: 0

- stream.graphics.gr6.4xlarge: 0

- stream.graphics.gr6.8xlarge: 0

- stream.graphics.g6f.large: 0

- stream.graphics.g6f.xlarge: 0

- stream.graphics.g6f.2xlarge: 0

- stream.graphics.g6f.4xlarge: 0

- stream.graphics.gr6f.4xlarge: 0

YesMemory-optimized image builder instances

- stream.memory.large: 3

- stream.memory.xlarge: 3

- stream.memory.2xlarge: 0

- stream.memory.4xlarge: 0

- stream.memory.8xlarge: 0

- stream.memory.z1d.large: 3

- stream.memory.z1d.xlarge: 3

- stream.memory.z1d.2xlarge: 0

- stream.memory.z1d.3xlarge: 0

- stream.memory.z1d.6xlarge: 0

- stream.memory.z1d.12xlarge: 0

YesStandard image builder instances

- stream.standard.small: 5

- stream.standard.medium: 5

- stream.standard.large: 5

- stream.standard.xlarge: 3

- stream.standard.2xlarge: 3

YesGeneral purpose image builder instances

- t3/t3a: small/medium/large/xlarge/2xlarge: 5/5/5/3/3

- m6i/m7i: large/xlarge/2xlarge+: 3/3/0

- m7i-flex: large/xlarge/2xlarge+: 3/3/0

- m6a/m7a: large/xlarge/2xlarge+: 3/3/0

YesCompute optimize image builder instances

- c6i/c7i: large/xlarge/2xlarge+: 3/3/0

- c6a/c7a: large/xlarge/2xlarge+: 3/3/0

- c7i-flex: large/xlarge/2xlarge+: 3/3/0

YesMemory optimize image builder instances

- r6i/r7i: large/xlarge/2xlarge+: 3/3/0

- r6a/r7a: large/xlarge/2xlarge+: 3/3/0

YesAccelerated image builder instances

- G4dn/G5/G6/G6f/Gr6/Gr6f/G6e: All sizes: 0

YesNumber of AWS accounts an image can be shared with100[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-99A44980)Concurrent image copies per destination Region2[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-60244546)Concurrent image updates5[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-9B6418E0)Users in the user pool50[Yes](https://console.aws.amazon.com/servicequotas/home/services/appstream2/quotas/L-6A8C9986)Max concurrent sessions for Elastic fleets

Amazon Linux 2

- stream.standard.small: 10

- stream.standard.medium: 10

- stream.standard.large: 5

- stream.standard.xlarge: 2

- stream.standard.2xlarge: 2

Ubuntu Pro 24.04 LTS

- stream.standard.small: 10

- stream.standard.medium: 10

- stream.standard.large: 5

- stream.standard.xlarge: 2

- stream.standard.2xlarge: 2

Windows Server 2019

- stream.standard.small: 10

- stream.standard.medium: 10

- stream.standard.large: 5

- stream.standard.xlarge: 2

- stream.standard.2xlarge: 2

Yes App block builders (total)10Yes Max number of app block builders

- stream.standard.small: 1

- stream.standard.medium: 1

- stream.standard.large: 1

- stream.standard.xlarge: 1

- stream.standard.2xlarge: 1

Yes

\\* WorkSpaces Applications instance type and size quotas are per AWS account per AWS Region. If you
have multiple fleets in the same Region that use the same instance type and size, the total
number of instances in all fleets in that Region must be less than or equal to the
applicable quota. To determine which instance types are available in which Regions or
Availability Zones, see _Pricing by AWS Region – Always-On, On-Demand,_
_app block builders,_
_and image builder instances_ in [WorkSpaces Applications Pricing](https://aws.amazon.com/appstream2/pricing).

For fleets that have **Default Internet Access** enabled, the quota is 100 fleet instances. If your deployment must support more than 100 concurrent users, use the [NAT gateway configuration](managing-network-internet-nat-gateway.md) instead. For more information about enabling internet access for a fleet, see [Internet Access](internet-access.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting Notification Codes

Guidance for WorkSpaces Applications Users

All content copied from https://docs.aws.amazon.com/.
