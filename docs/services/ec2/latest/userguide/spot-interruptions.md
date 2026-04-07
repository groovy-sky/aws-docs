# Spot Instance interruptions

You can launch Spot Instances on spare EC2 capacity for steep discounts in exchange for returning
them when Amazon EC2 needs the capacity back. When Amazon EC2 reclaims a Spot Instance, we call this event
a _Spot Instance interruption_.

Demand for Spot Instances can vary significantly from moment to moment, and the availability of Spot Instances
can also vary significantly depending on how many unused EC2 instances are available. It
is always possible that your Spot Instance might be interrupted. The following are the possible
reasons that Amazon EC2 might interrupt your Spot Instances:

**Capacity**

Amazon EC2 can interrupt your Spot Instance when it needs it back. EC2 reclaims your
instance mainly to repurpose capacity, but it can also occur for other
reasons such as host maintenance or hardware decommission.

**Price**

The Spot price is higher than your maximum price.

You can specify the maximum price in your Spot request. However, if
you specify a maximum price, your instances will be interrupted more
frequently than if you do not specify it.

**Constraints**

If your Spot request includes a constraint such as a launch group or an Availability
Zone group, the Spot Instances are terminated as a group when the constraint can
no longer be met.

When Amazon EC2 interrupts a Spot Instance, it either terminates, stops, or hibernates the instance,
depending on the interruption behavior that you specified when you created the Spot request.

###### Contents

- [Interruption behavior](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/interruption-behavior.html)

- [Prepare for interruptions](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/prepare-for-interruptions.html)

- [Initiate an interruption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/initiate-a-spot-instance-interruption.html)

- [Spot Instance interruption notices](spot-instance-termination-notices.md)

- [Find interrupted Spot Instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/finding-an-interrupted-Spot-Instance.html)

- [Determine whether Amazon EC2 terminated a Spot Instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/BidEvictedEvent.html)

- [Billing](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/billing-for-interrupted-spot-instances.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Manage your Spot Instances

Interruption behavior
