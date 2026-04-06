# Using Amazon Route 53 as the DNS service for subdomains without migrating the parent domain

Amazon Route 53 provides flexibility in managing DNS for subdomains, allowing you to
leverage its features without the need to migrate the entire parent domain.

You can either create a new subdomain or migrate an existing one to Route 53, while
keeping the parent domain hosted with another DNS service provider.

**Creating a new subdomain with Route 53:**

1. Create a hosted zone for the new subdomain.

2. Add the desired DNS records (e.g., A, CNAME, MX) for the subdomain to the
    hosted zone.

3. Obtain the Route 53 name servers assigned to the hosted zone.

4. Update the DNS configuration of the parent domain by adding NS (Name Server)
    records for the subdomain, pointing to the Route 53 name servers.

**Migrating an existing subdomain to Route 53:**

1. Create a hosted zone for the subdomain.

2. Obtain the current DNS configuration for the subdomain from your existing DNS
    service provider.

3. Add the corresponding DNS records to the hosted zone.

4. Obtain the Route 53 name servers assigned to the hosted zone.

5. Update the DNS configuration of the parent domain by adding NS records for the
    subdomain, pointing to the Route 53 name servers.

By following these steps, you can leverage Route 53's advanced features, such as health
checks, routing policies, and traffic flow management, for your subdomains while
maintaining the parent domain's DNS configuration with your existing provider.

###### Topics

- [Creating a subdomain that uses Amazon Route 53 as the DNS service without migrating the parent domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/CreatingNewSubdomain.html)

- [Migrating DNS service for a subdomain to Amazon Route 53 without migrating the parent domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/MigratingSubdomain.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tutorials

Creating a subdomain that uses Amazon Route 53 as the DNS service without migrating the parent domain
