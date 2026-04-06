# Creating, updating, and deleting health checks

###### Important

If you're updating or deleting health checks that are associated with records,
review the tasks in [Updating or deleting health checks when DNS failover is configured](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-updating-deleting-tasks.html) before you
proceed.

This section covers the following topics related to managing Route 53 health checks:

1. **Creating and updating health checks:**

- Learn how to create and update health checks using the Route 53 console.

- Understand the values you need to specify when creating or updating health checks, such as endpoint monitoring, protocol, IP address, domain name, and advanced configuration options.

2. **Values displayed when creating a health check:**

- Discover the values that the Route 53 console displays based on your input when creating a health check, such as the full URL or IP address and port.

3. **Updating health checks for CloudWatch alarm changes:**

- Find out how to update a health check when you change the settings of the associated CloudWatch alarm.

4. **Deleting health checks:**

- Follow the procedure to delete health checks by using the Route 53 console.

5. **Updating or deleting health checks when DNS failover is configured:**

- Learn the recommended tasks to perform when updating or deleting health checks associated with DNS records to ensure proper routing and failover configuration.

6. **Configuring router and firewall rules:**

- Understand how to configure your router and firewall rules to allow inbound traffic from Route 53 health checkers, ensuring successful health checks.

By following the information provided in this section, you can effectively create, update,
and delete Route 53 health checks, manage their configuration, and ensure proper integration with DNS failover and routing policies.

###### Topics

- [Creating and updating health checks](health-checks-creating.md)

- [Values that you specify when you create or update health checks](health-checks-creating-values.md)

- [Values that Amazon Route 53 displays when you create a health check](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-values-displayed.html)

- [Updating health checks when you change CloudWatch alarm settings (health checks that monitor a CloudWatch alarm only)](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-updating-cloudwatch-alarm-settings.html)

- [Disabling or enabling health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-disable.html)

- [Inverting health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-invert.html)

- [Deleting health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-deleting.html)

- [Updating or deleting health checks when DNS failover is configured](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-updating-deleting-tasks.html)

- [Configuring router and firewall rules for Amazon Route 53 health checks](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-router-firewall-rules.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

How Route 53
determines whether a health check is healthy

Creating and updating health checks
