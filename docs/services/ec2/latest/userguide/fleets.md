# EC2 Fleet and Spot Fleet

EC2 Fleet and Spot Fleet are designed to be a useful way to launch a fleet of tens, hundreds, or
thousands of Amazon EC2 instances in a single operation. Each instance in a fleet is either
configured by a [launch template](ec2-launch-templates.md) or a set of
launch parameters that you configure manually at launch.

###### Topics

- [Features and benefits](#ec2-fleet-features-and-benefits)

- [Which is the best fleet method to use?](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/which-fleet-method-to-use.html)

- [Configuration options for your EC2 Fleet or Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-configuration-strategies.html)

- [Work with EC2 Fleet](manage-ec2-fleet.md)

- [Work with Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/work-with-spot-fleets.html)

- [Monitor your EC2 Fleet or Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/fleet-monitor.html)

- [Tutorials for EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/fleet-tutorials.html)

- [Example CLI configurations for EC2 Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-examples.html)

- [Example CLI configurations Spot Fleet](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-fleet-examples.html)

- [Quotas for EC2 Fleet and Spot Fleet](fleet-quotas.md)

## Features and benefits

Fleets provide the following features and benefits, enabling you to maximize cost
savings and optimize availability and performance when running applications on multiple
EC2 instances.

**Multiple instance types**

A fleet can launch multiple instance types, ensuring it isn't dependent on
the availability of any single instance type. This increases the overall
availability of instances in your fleet.

**Distributing instances across Availability Zones**

A fleet can launch into multiple Availability Zones, enabling you to
reduce costs and improve availability. If your fleet includes Spot Instances, the
fleet automatically selects Availability Zones based on your preferences
regarding price and interruptions.

**Multiple purchasing options**

A fleet can launch multiple purchase options (Spot and On-Demand Instances), allowing
you to optimize costs through Spot Instance usage. You can also take advantage of
Reserved Instance and Savings Plans discounts by using them in conjunction with On-Demand Instances in
the fleet.

**Automated replacement of Spot Instances**

If your fleet includes Spot Instances, it can automatically request replacement
Spot capacity if your Spot Instances are interrupted. Through [Capacity Rebalancing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-capacity-rebalance.html), a
fleet can also monitor and proactively replace your Spot Instances that are at an
elevated risk of interruption.

**Reserve On-Demand capacity**

A fleet can use an [On-Demand Capacity Reservation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-fleet-on-demand-capacity-reservations.html) to
reserve On-Demand capacity. A fleet can also include [Capacity Blocks for ML](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-capacity-blocks.html), allowing you to reserve
GPU instances on a future date to support short duration machine
learning (ML) workloads.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Option 3: Manually connect

Which fleet method to
use?
