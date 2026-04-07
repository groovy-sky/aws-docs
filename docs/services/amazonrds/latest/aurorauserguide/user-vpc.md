# Amazon VPC and Amazon Aurora

Amazon Virtual Private Cloud (Amazon VPC) makes it possible for you to launch AWS resources,
such as Aurora DB clusters, into a virtual private cloud (VPC).

When you use a VPC, you have control over your virtual networking environment. You can
choose your own IP address range, create subnets, and configure routing and access control
lists. There is no additional cost to run your DB cluster in a VPC.

Accounts have a default VPC. All new DB clusters are created in the
default VPC unless you specify otherwise.

###### Topics

- [Working with a DB cluster in a VPC](user-vpc-workingwithrdsinstanceinavpc.md)

- [Scenarios for accessing a DB cluster in a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.Scenarios.html)

- [Tutorial: Create a VPC for use with a DB cluster (IPv4 only)](chap-tutorials-webserverdb-createvpc.md)

- [Tutorial: Create a VPC for use with a DB cluster (dual-stack mode)](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_Tutorials.CreateVPCDualStack.html)

Following, you can find a discussion about VPC functionality relevant to Amazon Aurora DB
clusters. For more information about Amazon VPC, see [Amazon VPC Getting Started Guide](https://docs.aws.amazon.com/AmazonVPC/latest/GettingStartedGuide) and
[Amazon VPC User Guide](https://docs.aws.amazon.com/vpc/latest/userguide).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Service-linked roles

Working with a DB
cluster in a VPC
