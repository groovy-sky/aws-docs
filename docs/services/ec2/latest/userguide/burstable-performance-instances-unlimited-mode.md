# Unlimited mode for burstable performance instances

A burstable performance instance configured as `unlimited` can sustain high
CPU utilization for any period of time whenever required. The hourly instance price
automatically covers all CPU usage spikes if the average CPU utilization of the instance is
at or below the baseline over a rolling 24-hour period or the instance lifetime, whichever
is shorter.

For the vast majority of general-purpose workloads, instances configured as
`unlimited` provide ample performance without any additional charges. If the
instance runs at higher CPU utilization for a prolonged period, it can do so for a flat
additional rate per vCPU-hour. For information about pricing, see [Amazon EC2 pricing](https://aws.amazon.com/ec2/pricing)
and [T2/T3/T4 Unlimited Mode Pricing](https://aws.amazon.com/ec2/pricing/on-demand).

If your created your AWS account before July 15, 2025 and you use a
`t2.micro` or `t3.micro` instance under the [AWS Free Tier](https://aws.amazon.com/free) offer and use it in
`unlimited` mode, charges might apply if your average utilization over a
rolling 24-hour period exceeds the [baseline\
utilization](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-credits-baseline-concepts.html#baseline_performance) of the instance.

T4g, T3a, and T3 instances launch as `unlimited` by default (unless you [change the default](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-how-to.html#burstable-performance-instance-set-default-credit-specification-for-account)). If the average CPU usage over a 24-hour period exceeds the
baseline, you incur charges for surplus credits. If you launch
Spot Instances as `unlimited` and plan to use them immediately and for a short
duration, with no idle time for accruing CPU credits, you incur charges for surplus
credits. We recommend that you launch your Spot Instances in [standard](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-standard-mode.html) mode to avoid
paying higher costs. For more information, see [Surplus credits can incur charges](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-surplus-credits) and [Launch burstable performance instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-spot-instances-work.html#burstable-spot-instances).

###### Note

T3 instances launched on a Dedicated Host launch as `standard` by default;
`unlimited` mode is not supported for T3 instances on a Dedicated Host.

###### Contents

- [Unlimited mode concepts for burstable instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html)

  - [How Unlimited burstable performance instances work](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#how-burstable-performance-instances-unlimited-works)

  - [When to use unlimited mode versus fixed CPU](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#when-to-use-unlimited-mode)

  - [Surplus credits can incur charges](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-surplus-credits)

  - [How much does unlimited burstable performance cost?](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#how-much-does-unlimited-burstable-performance-cost)

  - [No launch credits for T2 Unlimited instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-no-launch-credits)

  - [Enable unlimited mode](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-enabling)

  - [What happens to credits when switching between Unlimited and Standard](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-switching-and-credits)

  - [Monitor credit usage](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-unlimited-mode-concepts.html#unlimited-mode-monitoring-credit-usage)
- [Unlimited mode examples for burstable instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/unlimited-mode-examples.html)

  - [Example 1: Explain credit use with T3 Unlimited](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/unlimited-mode-examples.html#t3_unlimited_example)

  - [Example 2: Explain credit use with T2 Unlimited](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/unlimited-mode-examples.html#t2_unlimited_example)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Key concepts

Concepts
