# Working with records

After you create a hosted zone for your domain, such as example.com, you create records to tell the Domain Name System (DNS)
how you want traffic to be routed for that domain.

For example, you might create records that cause DNS to do the following:

- Route internet traffic for example.com to the IP address of a host in your data center.

- Route email for that domain (ichiro@example.com) to a mail server (mail.example.com).

- Route traffic for a subdomain called operations.tokyo.example.com to the IP address of a different host.

Each record includes the name of a domain or a subdomain, a record type (for example, a record
with a type of MX routes email), and other information applicable to the record type (for MX records, the host name of one or more
mail servers and a priority for each server). For information about the different record types, see
[Supported DNS record types](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html).

The name of each record in a hosted zone must end with the name of the hosted zone. For example,
the example.com hosted zone can contain records for www.example.com and accounting.tokyo.example.com subdomains,
but cannot contain records for a www.example.ca subdomain.

###### Note

To create records for complex routing configurations, you can also use the Traffic Flow visual editor and
save the configuration as a traffic policy. You can then associate the traffic policy with one or more domain names
(such as example.com) or subdomain names (such as www.example.com), in the same hosted zone or in multiple hosted zones.
In addition, you can roll back the updates if the new configuration isn't performing as you expected it to. For more information,
see [Using Traffic Flow to route DNS traffic](traffic-flow.md).

Amazon Route 53 doesn't charge for the records that you add to a hosted zone. For information about the maximum number of
records that you can create in a hosted zone, see [Quotas](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html).

###### Topics

- [Choosing a routing policy](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html)

- [Choosing between alias and non-alias records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html)

- [Supported DNS record types](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html)

- [Creating records by using the Amazon Route 53 console](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-creating.html)

- [Resource record set permissions](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-permissions.html)

- [Values that you specify when you create or edit Amazon Route 53 records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values.html)

- [Creating records by importing a zone file](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-creating-import.html)

- [Editing records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-editing.html)

- [Deleting records](resource-record-sets-deleting.md)

- [Listing records](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-listing.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Migrating a hosted zone to a different AWS account

Choosing a routing policy
