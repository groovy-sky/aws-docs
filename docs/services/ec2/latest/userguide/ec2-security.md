# Security in Amazon EC2

Cloud security at AWS is the highest priority. As an AWS customer, you benefit from a
data center and network architecture that are built to meet the requirements of the most
security-sensitive organizations.

Security is a shared responsibility between AWS and you. The [shared responsibility\
model](https://aws.amazon.com/compliance/shared-responsibility-model) describes this as security of the cloud and security in the cloud:

- **Security of the cloud** – AWS is responsible for
protecting the infrastructure that runs AWS services in the AWS Cloud. AWS also
provides you with services that you can use securely. Third-party auditors regularly test
and verify the effectiveness of our security as part of the [AWS Compliance Programs](https://aws.amazon.com/compliance/programs). To learn about the compliance programs that
apply to Amazon EC2, see [AWS Services in Scope by Compliance Program](https://aws.amazon.com/compliance/services-in-scope).

- **Security in the cloud** – Your responsibility
includes the following areas:

- Controlling network access to your instances, for example, through configuring
your VPC and security groups. For more information, see [Controlling network traffic](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/infrastructure-security.html#control-network-traffic).

- Managing the credentials used to connect to your instances.

- Managing the guest operating system and software deployed to the guest operating
system, including updates and security patches. For more information, see [Update management for Amazon EC2 instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/update-management.html).

- Configuring the IAM roles that are attached to the instance and the permissions
associated with those roles. For more information, see [IAM roles for Amazon EC2](iam-roles-for-amazon-ec2.md).

This documentation helps you understand how to apply the shared responsibility model when
using Amazon EC2. It shows you how to configure Amazon EC2 to meet your security and compliance objectives.
You also learn how to use other AWS services that help you to monitor and secure your Amazon EC2
resources.

###### Contents

- [Data protection](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/data-protection.html)

- [Infrastructure security](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/infrastructure-security.html)

- [Resilience](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/disaster-recovery-resiliency.html)

- [Compliance validation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/compliance-validation.html)

- [Identity and access management](security-iam.md)

- [Update management](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/update-management.html)

- [Best practices for Windows instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-windows-security-best-practices.html)

- [Key pairs](ec2-key-pairs.md)

- [Security groups](ec2-security-groups.md)

- [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html)

- [EC2 instance attestation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm-attestation.html)

- [Credential Guard for Windows instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/credential-guard.html)

- [AWS PrivateLink](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/interface-vpc-endpoints.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Secondary Networks

Data protection
