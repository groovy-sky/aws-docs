---
title: "AWS-managed prefix lists"
---

# AWS-managed prefix lists

AWS-managed prefix lists are sets of IP address ranges for AWS services. These prefix lists are maintained by Amazon Web Services and provide a way to reference the IP addresses used by various AWS offerings. This can be particularly useful when configuring security groups or other network-level controls within a VPC.

The prefix lists cover a wide range of AWS services, including S3 and DynamoDB, and many others. By using the managed prefix lists, you can ensure that your network configurations are up-to-date and properly account for the IP addresses used by the AWS services you depend on. This can help simplify networking tasks and reduce the administrative overhead of manually maintaining lists of IP addresses.

In addition to the practical benefits, using the managed prefix lists also aligns with AWS security best practices. By relying on the authoritative IP address information provided by AWS, you can minimize the risk of misconfiguration or unexpected connectivity issues. This can be especially important for mission-critical applications or workloads with strict compliance requirements.

###### Contents

- [Available AWS-managed prefix lists](#available-aws-managed-prefix-lists)

- [AWS-managed prefix list weight](#aws-managed-prefix-list-weights)

- [Use an AWS-managed prefix list](#use-aws-managed-prefix-list)

## Available AWS-managed prefix lists

The following services provide AWS-managed prefix lists.

AWS servicePrefix list nameWeight[Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/locationsofedgeservers.md#managed-prefix-list)

com.amazonaws.global.cloudfront.origin-facing (IPv4)

com.amazonaws.global.ipv6.cloudfront.origin-facing (IPv6)

55Amazon DynamoDBcom.amazonaws. `region`.dynamodb1[Amazon EC2 Instance Connect](../../../ec2/latest/userguide/ec2-instance-connect-prerequisites.md#ec2-instance-connect-setup-security-group)com.amazonaws. `region`.ec2-instance-connect2com.amazonaws. `region`.ipv6.ec2-instance-connect2AWS Ground Stationcom.amazonaws.global.groundstation5[Amazon Route 53](../../../route53/latest/developerguide/dns-failover-router-firewall-rules.md)com.amazonaws. `region`.ipv6.route53-healthchecks25com.amazonaws. `region`.route53-healthchecks25Amazon S3com.amazonaws. `region`.s31Amazon S3 Express One Zone com.amazonaws. `region`.s3express6[Amazon VPC Lattice](../../../vpc-lattice/latest/ug/security-groups.md#managed-prefix-list)com.amazonaws. `region`.vpc-lattice10com.amazonaws. `region`.ipv6.vpc-lattice10

###### To view the AWS-managed prefix lists using the console

1. Open the Amazon VPC console at
    [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).

2. In the navigation pane, choose **Managed Prefix Lists**.

3. In the search field, add the **Owner ID: AWS** filter.

###### To view the AWS-managed prefix lists using the AWS CLI

Use the [describe-managed-prefix-lists](../../../cli/latest/reference/ec2/describe-managed-prefix-lists.md)
command as follows.

```nohighlight

aws ec2 describe-managed-prefix-lists --filters Name=owner-id,Values=AWS
```

## AWS-managed prefix list weight

The weight of an AWS-managed prefix list refers to the number of entries that
it takes up in a resource.

For example, the weight of a Amazon CloudFront managed prefix list is 55. Here's how the this
affects your Amazon VPC quotas:

- Security groups – The [default quota](amazon-vpc-limits.md#vpc-limits-security-groups)
is 60 rules, leaving room for only 5 additional rules in a security group. You can [request a quota increase](https://console.aws.amazon.com/servicequotas/home/services/vpc/quotas/L-0EA8095F) for this quota.

- Route tables – The [default quota](amazon-vpc-limits.md#vpc-limits-route-tables)
is 50 routes, so you must [request a quota increase](https://console.aws.amazon.com/servicequotas/home/services/vpc/quotas/L-93826ACB)
before you can add the prefix list to a route table.

## Use an AWS-managed prefix list

AWS-managed prefix lists are created and maintained by AWS and can be used by anyone
with an AWS account. You cannot create, modify, share, or delete an AWS-managed prefix
list.

As with customer-managed prefix lists, you can use AWS-managed prefix lists with
AWS resources such as security groups and route tables. For more information, see
[Optimize AWS infrastructure management with prefix lists](managed-prefix-lists-referencing.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Work with shared prefix lists

Optimize AWS infrastructure management with prefix lists

All content copied from https://docs.aws.amazon.com/.
