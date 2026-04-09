# Best Practices for Enabling Application Settings Persistence

To enable application settings persistence without providing internet access to
your instances, use a VPC endpoint. This endpoint must be in the VPC to which your
WorkSpaces Applications instances are connected. You must attach a custom policy to enable WorkSpaces Applications
access to the endpoint. For information about how to create the custom policy, see
the _Home Folders and VPC Endpoints_ section in [Networking and Access for Amazon WorkSpaces Applications](managing-network.md). For more
information about private Amazon S3 endpoints, see [VPC Endpoints](../../../vpc/latest/userguide/vpc-endpoints.md) and
[Endpoints for Amazon S3](../../../vpc/latest/userguide/vpc-endpoints-s3.md) in the _Amazon VPC User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Prerequisites for Enabling Application Settings
Persistence

How to Enable Application Settings Persistence

All content copied from https://docs.aws.amazon.com/.
