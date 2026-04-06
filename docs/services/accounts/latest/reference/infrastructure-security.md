# Infrastructure security in AWS Account Management

As managed services, AWS services running in your AWS account are protected by the
AWS global network security. For information about AWS security services and how AWS
protects infrastructure, see [AWS Cloud\
Security](https://aws.amazon.com/security). To design your AWS environment using the best practices for
infrastructure security, see [Infrastructure Protection](../../../wellarchitected/latest/security-pillar/infrastructure-protection.md) in _Security Pillar AWS_
_Well‐Architected Framework_.

You use AWS published API calls to access account settings through the network. Clients
must support the following:

- Transport Layer Security (TLS). We require TLS 1.2 and recommend TLS 1.3.

- Cipher suites with perfect forward secrecy (PFS) such as DHE (Ephemeral
Diffie-Hellman) or ECDHE (Elliptic Curve Ephemeral Diffie-Hellman). Most modern
systems such as Java 7 and later support these modes.

Additionally, requests must be signed by using an access key ID and a secret access key
that is associated with an IAM principal. Or you can use the [AWS Security Token Service](../../../../reference/sts/latest/apireference/welcome.md) (AWS STS) to generate
temporary security credentials to sign requests.

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Resilience

Monitor your account
