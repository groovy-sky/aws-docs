# Standard mode for burstable performance instances

A burstable performance instance configured as `standard` is suited to
workloads with an average CPU utilization that is consistently below the baseline CPU
utilization of the instance. To burst above the baseline, the instance spends credits that
it has accrued in its CPU credit balance. If the instance is running low on accrued
credits, CPU utilization is gradually lowered to the baseline level, so that the instance
does not experience a sharp performance drop-off when its accrued CPU credit balance is
depleted. For more information, see [Key concepts for burstable performance instances](burstable-credits-baseline-concepts.md).

###### Contents

- [Standard mode concepts for burstable instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-standard-mode-concepts.html)

  - [How standard burstable performance instances work](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-standard-mode-concepts.html#how-burstable-performance-instances-standard-works)

  - [Launch credits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-standard-mode-concepts.html#launch-credits)

  - [Launch credit limits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-standard-mode-concepts.html#launch-credit-limits)

  - [Differences between launch credits and earned credits](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances-standard-mode-concepts.html#burstable-performance-instances-diff-launch-earned-credits)
- [Standard mode examples for burstable instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html)

  - [Example 1: Explain credit use with T3 Standard](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#t3_standard_example)

  - [Example 2: Explain credit use with T2 Standard](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#t2-standard-example)

    - [Period 1: 1 – 24 hours](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#period-1)

    - [Period 2: 25 – 36 hours](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#period-2)

    - [Period 3: 37 – 61 hours](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#period-3)

    - [Period 4: 62 – 72 hours](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#period-4)

    - [Period 5: 73 – 75 hours](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#period-5)

    - [Period 6: 76 – 90 hours](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#period-6)

    - [Period 7: 91 – 96 hours](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/standard-mode-examples.html#period-7)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Examples

Concepts
