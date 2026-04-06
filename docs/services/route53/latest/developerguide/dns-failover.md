# Creating Amazon Route 53 health checks

Amazon Route 53 health checks monitor the health and performance of your web applications, web
servers, and other resources. Each health check that you create can monitor one of the
following:

- The health of a specified resource, such as a web server.

- The status of other health checks.

- The status of an Amazon CloudWatch alarm.

- Additionally, with Amazon Application Recovery Controller (ARC), you can set up routing control health checks with DNS
failover records to manage traffic failover for your application. To learn more, see
[Amazon Application Recovery Controller (ARC) Developer\
Guide](https://docs.aws.amazon.com/r53recovery/latest/dg/what-is-route-53-recovery.html).

For an overview of the types of health checks, see [Types of Amazon Route 53 health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-types.html). For information
about creating health checks, see [Creating and updating health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating.html).

After you create a health check, you can get the status of the health check, get
notifications when the status changes, and configure DNS failover:

**Getting health check status and notifications**

You can view the current and recent status of your health checks on the Route 53
console. You can also work with health checks programmatically through one of
the AWS SDKs, the AWS Command Line Interface, AWS Tools for Windows PowerShell, or the Route 53 API.

If you want to receive a notification when the status of a health check
changes, you can configure an Amazon CloudWatch alarm for each health check.

For information about viewing health check status and receiving notifications,
see [Monitoring health check status and getting notifications](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-monitor-view-status.html).

**Configuring DNS failover**

If you have multiple resources that perform the same function, you can
configure DNS failover so that Route 53 will route your traffic from an unhealthy
resource to a healthy resource. For example, if you have two web servers and one
web server becomes unhealthy, Route 53 can route traffic to the other web server.
For more information, see [Configuring DNS failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-configuring.html).

###### Topics

- [Types of Amazon Route 53 health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-types.html)

- [How Amazon Route 53 determines whether a health check is healthy](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html)

- [Creating, updating, and deleting health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html)

- [Configuring DNS failover](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-configuring.html)

- [Naming and tagging health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-tagging.html)

- [Using health checks with Amazon Route 53 API versions earlier than 2012-12-12](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-using-old-apis.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Other AWS resources

Types of health checks
