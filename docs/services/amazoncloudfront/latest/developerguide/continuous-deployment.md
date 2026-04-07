# Use CloudFront continuous deployment to safely test CDN configuration changes

With Amazon CloudFront _continuous deployment_ you can safely
deploy changes to your CDN configuration by testing first with a subset of production
traffic. You can use a _staging distribution_ and a
_continuous deployment policy_ to send some traffic from
real (production) viewers to the new CDN configuration and validate that it works as
expected. You can monitor the performance of the new configuration in real time, and promote
the new configuration to serve all traffic via the _primary_
_distribution_ when you're ready.

The following diagram shows the benefit of using CloudFront continuous deployment. Without it,
you would have to test CDN configuration changes with simulated traffic. With continuous
deployment you can test the changes with a subset of production traffic, then promote the
changes to the primary distribution when you're ready.

![Graphic of CloudFront continuous deployment that sends production traffic to a staging distribution.](https://docs.aws.amazon.com/images/AmazonCloudFront/latest/DeveloperGuide/images/cloudfront-continuous-deployment.png)

Learn more about working with continuous deployment in the following topics.

###### Topics

- [CloudFront continuous deployment workflow](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/continuous-deployment-workflow.html)

- [Work with a staging distribution and continuous deployment policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/working-with-staging-distribution-continuous-deployment-policy.html)

- [Monitor a staging distribution](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/monitoring-staging-distribution.html)

- [Learn how continuous deployment works](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/understanding-continuous-deployment.html)

- [Quotas and other considerations for continuous deployment](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/continuous-deployment-quotas-considerations.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Enable IPv6

CloudFront continuous deployment workflow
