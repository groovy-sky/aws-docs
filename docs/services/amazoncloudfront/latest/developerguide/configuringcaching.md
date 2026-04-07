# Caching and availability

You can use CloudFront to reduce the number of requests that your origin server must respond to
directly. With CloudFront caching, more objects are served from CloudFront edge locations, which are
closer to your users. This reduces the load on your origin server and reduces
latency.

The more requests that CloudFront can serve from edge caches, the fewer viewer requests that
CloudFront must forward to your origin to get the latest version or a unique version of an object.
To optimize CloudFront to make as few requests to your origin as possible, consider using a CloudFront
Origin Shield. For more information, see [Use Amazon CloudFront Origin Shield](origin-shield.md).

The proportion of requests that are served directly from the CloudFront cache compared to all
requests is called the _cache hit ratio_. You can view the percentage of
viewer requests that are hits, misses, and errors in the CloudFront console. For more information,
see [View CloudFront cache statistics reports](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cache-statistics.html).

A number of factors affect the cache hit ratio. You can adjust your CloudFront distribution
configuration to improve the cache hit ratio by following the guidance in [Increase the proportion of requests that are served directly from the CloudFront caches (cache hit ratio)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cache-hit-ratio.html).

To learn about adding and removing the content that you want CloudFront to serve, see [Add, remove, or replace content that CloudFront distributes](addremovereplaceobjects.md).

###### Topics

- [Increase the proportion of requests that are served directly from the CloudFront caches (cache hit ratio)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cache-hit-ratio.html)

- [Use Amazon CloudFront Origin Shield](origin-shield.md)

- [Optimize high availability with CloudFront origin failover](high-availability-origin-failover.md)

- [Manage how long content stays in the cache (expiration)](expiration.md)

- [Cache content based on query string parameters](querystringparameters.md)

- [Cache content based on cookies](cookies.md)

- [Cache content based on request headers](header-caching.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Working with shared resources in CloudFront

Improve your cache hit ratio
