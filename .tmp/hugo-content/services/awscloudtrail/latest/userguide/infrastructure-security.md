# Infrastructure security in AWS CloudTrail

As a managed service, AWS CloudTrail is protected by AWS global network security. For
information about AWS security services and how AWS protects infrastructure, see [AWS Cloud Security](https://aws.amazon.com/security). To design your AWS
environment using the best practices for infrastructure security, see [Infrastructure\
Protection](../../../wellarchitected/latest/security-pillar/infrastructure-protection.md) in _Security Pillar AWS Well‐Architected_
_Framework_.

You use AWS published API calls to access CloudTrail through the network. Clients must
support the following:

- Transport Layer Security (TLS). We require TLS 1.2 and recommend TLS 1.3.

- Cipher suites with perfect forward secrecy (PFS) such as DHE (Ephemeral
Diffie-Hellman) or ECDHE (Elliptic Curve Ephemeral Diffie-Hellman). Most modern systems
such as Java 7 and later support these modes.

The following security best practices also address infrastructure security in CloudTrail:

- [Consider Amazon VPC endpoints for\
trail access.](cloudtrail-and-interface-vpc.md)

- Consider Amazon VPC endpoints for Amazon S3 bucket access. For more information, see [Controlling access from VPC endpoints with bucket policies](../../../s3/latest/userguide/example-bucket-policies-vpc-endpoint.md).

- Identify and audit all Amazon S3 buckets that contain CloudTrail log files. Consider using
tags to help identify both your CloudTrail trails and the Amazon S3 buckets that contain CloudTrail
log files. You can then use resource groups for your CloudTrail resources. For more
information, see [AWS Resource Groups](../../../arg/latest/userguide/resource-groups.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Resilience

Cross-service confused deputy prevention

All content copied from https://docs.aws.amazon.com/.
