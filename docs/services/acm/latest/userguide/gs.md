# Getting started with AWS Certificate Manager certificates

ACM manages public, private, and imported certificates. Certificates are used to
establish secure communications across the internet or within an internal network. You can
request a publicly trusted certificate directly from ACM (an "ACM certificate"), import
a publicly trusted certificate issued by a third party. Self-signed certificates are also
supported. To provision your organization's internal PKI, you can issue ACM certificates
signed by a private certificate authority (CA) created and managed by [AWS Private CA](../../../privateca/latest/userguide/pcawelcome.md). The CA may either reside in your
account or be shared with you by a different account.

###### Note

Public ACM certificates can be installed on Amazon EC2 instances that are connected to a
[Nitro Enclave](acm-services.md#acm-nitro-enclave). You can also [export a public certificate](export-public-certificate.md)
to use on any Amazon EC2 instance. For information about
setting up a standalone web server on an Amazon EC2 instance not connected to a Nitro Enclave, see [Tutorial: Install a LAMP web server\
on Amazon Linux 2](../../../ec2/latest/userguide/ec2-lamp-amazon-linux-2.md) or [Tutorial:\
Install a LAMP web server with the Amazon Linux AMI](../../../ec2/latest/userguide/install-lamp.md).

###### Note

Because certificates signed by a private CA are not trusted by default,
administrators must install them in client trust stores.

To begin issuing certificates, sign into the AWS Management Console and open the ACM
console at [https://console.aws.amazon.com/acm/home](https://console.aws.amazon.com/acm/home).
If the introductory page appears, choose **Get Started**. Otherwise, choose
**Certificate Manager** or **Private CAs** in the left
navigation pane.

###### Topics

- [Set up to use AWS Certificate Manager](setup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

What is the right AWS certificate service for my needs?

Set up

All content copied from https://docs.aws.amazon.com/.
