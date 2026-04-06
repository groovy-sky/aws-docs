# Quotas

Amazon Route 53 API requests and entities are subject to the following quotas (formerly referred to as "limits").

###### Topics

- [Using Service Quotas to view and manage quotas](#limits-service-quotas)

- [Quotas on entities](#limits-api-entities)

- [Maximums on API requests](#limits-api-requests)

## Using Service Quotas to view and manage quotas

You can use the Service Quotas service to view quotas and to request quota increases for many
AWS services. For more information, see the [Service Quotas User Guide](https://docs.aws.amazon.com/servicequotas/latest/userguide). (You can currently use
Service Quotas to view and manage domains, Route 53 ,and Route 53 VPC Resolver quotas.)

###### Note

To view quotas and request higher quotas for Route 53, you must change the Region to US East (N. Virginia). To view quotas and
request higher quotas for VPC Resolver, change to the applicable Region.

## Quotas on entities

Amazon Route 53 entities are subject to the following quotas.

For information on getting current quotas (formerly referred to as "limits"), see the
following Route 53 actions:

- [GetAccountLimit](https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html)
– Gets quotas on health checks, hosted zones, reusable delegation sets,
traffic flow policies, and traffic flow policy records

- [GetHostedZoneLimit](https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetHostedZoneLimit.html) – Gets quotas on records in a hosted zone
and on Amazon VPCs that you can associate with a private hosted zone

- [GetReusableDelegationSetLimit](https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetReusableDelegationSetLimit.html) – Gets the quota on the number
of hosted zones that you can associate with a reusable delegation set

###### Topics

- [Quotas on domains](#limits-api-entities-domains)

- [Quotas on hosted zones](#limits-api-entities-hosted-zones)

- [Quotas on records](#limits-api-entities-records)

- [Quotas on Route 53 VPC Resolver](#limits-api-entities-resolver)

- [Quotas on health checks](#limits-api-entities-health-checks)

- [Quotas on query log configurations](#limits-api-entities-query-log-configs)

- [Quotas on traffic flow policies and policy records](#limits-api-entities-traffic-flow)

- [Quotas on reusable delegation sets](#limits-api-entities-reusable-delegation-sets)

- [Quotas on Route 53 Profiles](#limits-api-entities-route53-profiles)

### Quotas on domains

EntityQuota

Domains

20\* per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

**\*** The limit is 20 for new customers as of March
2021.

If you have an existing account and your default limit is 50 now, it will remain at
50.

### Quotas on hosted zones

EntityQuota

Hosted zones

Initial quota of 500 per AWS account, but you can request a higher quota as needed.

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Hosted zones that can use the same reusable delegation set

100

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Amazon VPCs that you can associate with a private hosted zone per hosted zone

300

If you want more than 300 associations, we recommend you use Route 53 Profiles. For more
information, see
[What are Amazon Route 53 Profiles?](profiles.md).

Private hosted zones that you can associate a VPC with

No quota **\***

Authorizations that you can create so you can associate VPCs
that were created by one account with a hosted zone that was
created by another account

1000

The number of key signing keys (KSK) that you can create per hosted zone

2

**\*** You can associate a VPC with any or all of the
private hosted zones that you control through your AWS accounts. For example,
suppose you have three AWS accounts and all three have the default quota of 500
hosted zones. If you create 500 private hosted zones for all three accounts, you can
associate a VPC with all 1,500 private hosted zones.

### Quotas on records

EntityQuota

Records

10,000 per hosted zone

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

For a quota greater than 10,000 records in a hosted zone, an additional charge applies.For more information, see
[Amazon Route 53 Pricing](https://aws.amazon.com/route53/pricing).

Records in a record set

400 per record set

Geolocation, latency, multivalue answer, weighted, and IP-based records

100 records that have the same name and type

Geoproximity records

30 records that have the same name and type

CIDR collections

5 per AWS account.

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

CIDR blocks

1000 per CIDR collection.

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

### Quotas on Route 53 VPC Resolver

This section includes all the Route 53 VPC Resolver quotas

#### Quotas on Route 53 VPC Resolver

Use the following procedure to increase quotas for Route 53 VPC Resolver.

###### To increase Resolver quotas

1. Open the Service Quotas console at [https://console.aws.amazon.com/servicequotas/home/services/route53resolver/quotas](https://console.aws.amazon.com/servicequotas/home/services/route53resolver/quotas).

2. Go to the region where you want to increase the limit.

3. Select the Route 53 VPC Resolver **Quota name** you want to increase.

4. Select **Request quota increase**, enter the quota value, and then select **Request**.

#### Quotas on Route 53 VPC Resolver endpoints

EntityQuota

Endpoints per AWS Region

4 per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home/services/route53resolver/quotas).

IP addresses per endpoint

6

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home/services/route53resolver/quotas).

IP addresses per rule

6

Rules per AWS Region

1000 per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home/services/route53resolver/quotas).

Associations between rules and VPCs per AWS Region

2000 per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home/services/route53resolver/quotas).

UDP Queries per second per IP address in an endpoint

10,000\*

**\*** Each IP address in an endpoint can process up to 10,000
UDP DNS queries per second (QPS). The number of DNS QPS varies by the type of
query, size of the response, health of the target name servers, query response
times, round trip latency, and the protocol in use. For example, queries to a
target name server that is slow to respond can significantly reduce the capacity
of the network interface. Additionally, to ensure high availability, Route 53
Resolver generates redundant outbound queries for each DNS request that it
receives. As a result, the QPS for each outbound network interface will not
match the QPS sent to Route 53 VPC Resolver. Use CloudWatch metrics to measure how
many queries are being sent to each network interface. For more information, see
[Metrics for Route 53 VPC Resolver IP addresses](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/monitoring-resolver-with-cloudwatch.html#cloudwatch-metrics-resolver-ip-address). If your maximum
query rate exceeds 50% of the capacity for any network interface in the
endpoint, you can add more network interfaces to increase the endpoint
capacity.

Connections made through applications like Network Load Balancer and AWS Lambda (for a full
list see [Automatically tracked connections](../../../ec2/latest/userguide/security-group-connection-tracking.md#automatic-tracking) ) are automatically tracked, even
if the security group configuration does not otherwise require tracking.

If the connection tracking is enforced either by using restrictive security group rules or queries are routed through Network Load Balancer,
the overall maximum queries per second per IP address for an inbound endpoint can be as low as 1500.

#### Quotas on Route 53 VPC Resolver query logs

EntityQuota

Query log configurations per AWS Region

20

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Query log configuration VPC associations per AWS
Region

100 **\***

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Query log configuration VPC associations per account per
AWS Region (including query long confgurations shared using RAM) for the account that the
configuration was shared to.

100

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

**\*** This is a Regional limit that applies cumulatively
across all VPC Resolver query log configurations within a single Region. Creating
additional query log configurations in the same Region does not provide
additional VPC association capacity.

#### Quotas on Resolver DNS Firewall

EntityQuota

Number of rule groups associated to a VPC for a single account per AWS
Region

5

Number of DNS Firewall domains in a single Amazon S3 file for a single account per AWS Region

250,000

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Number of DNS Firewall rule groups for a single account per AWS Region

1,000

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Number of rules within a rule group for a single account per AWS Region

100

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Number of domain lists for a single account per AWS Region

1000

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

The maximum number of domains that you can specify across
all of the domain lists for a single account per AWS
Region

100,000

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

#### Quotas on Resolver on Outpost

EntityQuota

Resolver on Outpost instance limit

6 (with a minimum of 4 required)

**Resolver on Outpost instance types and the number of DNS**
**queries per second each instance type can accommodate:**

Instance typeQueries per second

c5.large

Up to 7,000

c5.xlarge

Up to 12,000

c5.2xlarge

Up to 24,000

c5.4xlarge

Up to 56,000

c5d.large

Up to 7,000

c5d.xlarge

Up to 12,000

c5d.2xlarge

Up to 24,000

c5d.4xlarge

Up to 56,000

m5.large

Up to 7,000

m5.xlarge

Up to 12,000

m5.2xlarge

Up to 24,000

m5.4xlarge

Up to 56,000

m5d.large

Up to 7,000

m5d.xlarge

Up to 12,000

m5d.2xlarge

Up to 24,000

m5d.4xlarge

Up to 56,000

r5.large

Up to 7,000

r5.xlarge

Up to 12,000

r5.2xlarge

Up to 24,000

r5.4xlarge

Up to 56,000

r5d.large

Up to 7,000

r5d.xlarge

Up to 12,000

r5d.2xlarge

Up to 24,000

r5d.4xlarge

Up to 56,000

**Resolver on Outpost endpoint instance types and the number of**
**DNS queries per second each instance type can accommodate:**

Instance typeQueries per second

c5.large

Up to 5,000

c5.xlarge

Up to 10,000

c5.2xlarge

Up to 18,000

c5.4xlarge

Up to 30,000

c5d.large

Up to 5,000

c5d.xlarge

Up to 10,000

c5d.2xlarge

Up to 18,000

c5d.4xlarge

Up to 30,000

m5.large

Up to 5,000

m5.xlarge

Up to 10,000

m5.2xlarge

Up to 18,000

m5.4xlarge

Up to 30,000

m5d.large

Up to 5,000

m5d.xlarge

Up to 10,000

m5d.2xlarge

Up to 18,000

m5d.4xlarge

Up to 30,000

r5.large

Up to 5,000

r5.xlarge

Up to 10,000

r5.2xlarge

Up to 18,000

r5.4xlarge

Up to 30,000

r5d.large

Up to 5,000

r5d.xlarge

Up to 10,000

r5d.2xlarge

Up to 18,000

r5d.4xlarge

Up to 30,000

### Quotas on health checks

EntityQuota

Health checks

200 active health checks per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Child health checks that a calculated health check can
monitor

255

Maximum total length of headers in the response to a health
check request

16,384 bytes (16K)

### Quotas on query log configurations

EntityQuota

Query log configurations

1 per hosted zone

### Quotas on traffic flow policies and policy records

EntityQuota

Traffic policies

For more information about Route 53 traffic flow, see [Using Traffic Flow to route DNS traffic](traffic-flow.md).

50 per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Traffic policy versions

1,000 per traffic policy

Traffic policy records (referred to as "policy instances" in
the Route 53 API, AWS SDKs, AWS Command Line Interface, and AWS Tools for Windows PowerShell)

5 per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

### Quotas on reusable delegation sets

EntityQuota

Reusable delegation sets

100 per AWS account

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

### Quotas on Route 53 Profiles

EntityQuota

Number of Route 53 Profiles per AWS account in a Region

5

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Number of VPCs that can be associated to a Profile

1000

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Number of DNS Firewall rule groups per Profile

5

Number of Resolver rules per Profile

1000

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Number of private hosted zones per a Profile

5000

[Request a higher quota](https://console.aws.amazon.com/servicequotas/home?region=us-east-1).

Number of VPC Resolver query logging configurations per
Profile

2

## Maximums on API requests

Amazon Route 53 API requests are subject to the following maximums.

###### Topics

- [Number of elements and characters in ChangeResourceRecordSets requests](#limits-api-requests-changeresourcerecordsets)

- [Frequency of Amazon Route 53 API requests](#limits-api-requests-route-53)

- [Frequency of Route 53 VPC Resolver API requests](#limits-api-requests-route-53-resolver)

### Number of elements and characters in `ChangeResourceRecordSets` requests

**`ResourceRecord` elements**

A request cannot contain more than 1,000 `ResourceRecord` elements (including alias
records). When the value of the `Action` element is
`UPSERT`, each `ResourceRecord` element is
counted twice.

**Maximum number of characters**

The sum of the number of characters (including spaces) in all `Value` elements in a request
cannot exceed 32,000 characters. When the value of the `Action` element is `UPSERT`, each character
in a `Value` element is counted twice.

### Frequency of Amazon Route 53 API requests

**All Amazon Route 53 API requests**

For the [Amazon Route 53 APIs](https://docs.aws.amazon.com/Route53/latest/APIReference/API_Operations_Amazon_Route_53.html) five requests per second per AWS account. If you submit more than five requests per second,
Amazon Route 53 returns an HTTP 400 error ( `Bad request`). The response header also includes a `Code` element
with a value of `Throttling` and a `Message` element with a value of `Rate exceeded`.

###### Note

If your application exceeds this limit, we recommend that you implement exponential backoff for retries.
For more information, see
[Error Retries and Exponential Backoff in AWS](https://docs.aws.amazon.com/general/latest/gr/api-retries.html)
in the _Amazon Web Services General Reference_.

**`ChangeResourceRecordSets` requests**

If Route 53 can't process a request before the next request arrives, it will reject subsequent requests
for the same hosted zone and return an HTTP 400 error ( `Bad request`). The response header also includes a
`Code` element with a value of `PriorRequestNotComplete` and a `Message` element
with a value of `The request was rejected because Route 53 was still processing a prior request.`

**`CreateHealthCheck` requests**

You can submit one `CreateHealthCheck` request every 2 seconds per AWS account.

### Frequency of Route 53 VPC Resolver API requests

**All requests**

Five requests per second per AWS account per Region. If you submit more than five requests per second in a Region,
VPC Resolver returns an HTTP 400 error ( `Bad request`). The response header also includes a `Code` element
with a value of `Throttling` and a `Message` element with a value of `Rate exceeded`.

###### Note

If your application exceeds this limit, we recommend that you implement exponential backoff for retries.
For more information, see
[Error Retries and Exponential Backoff in AWS](https://docs.aws.amazon.com/general/latest/gr/api-retries.html)
in the _Amazon Web Services General Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Best practices for Amazon Route 53 health checks

Related information
