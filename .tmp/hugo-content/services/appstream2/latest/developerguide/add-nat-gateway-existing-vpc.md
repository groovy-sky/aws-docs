# Add a NAT Gateway to an Existing VPC

If you have already configured a VPC, complete the following steps to add a NAT gateway to your VPC. If you need to create a new VPC, see [Create and Configure a New VPC](create-configure-new-vpc-with-private-public-subnets-nat.md).

###### To add a NAT gateway to an existing VPC

1. To create your NAT gateway, complete the steps in [Creating a NAT Gateway](../../../vpc/latest/userguide/vpc-nat-gateway.md#nat-gateway-creating) in the
    _Amazon VPC User Guide_.

2. Verify that your VPC has at least one private subnet. We recommend that you specify two private subnets from different
    Availability Zones for high availability and fault tolerance. For information about how to create a second private subnet, see
    [Step 3: Add a Second Private Subnet](create-configure-new-vpc-with-private-public-subnets-nat.md#vpc-with-private-and-public-subnets-add-private-subnet-nat).

3. Update the route table
    associated with one or more of your private subnets to point
    internet-bound traffic to the NAT gateway. This enables the streaming
    instances in your private subnets to communicate with the internet. To
    do so, complete the steps in [Configure route tables](../../../vpc/latest/userguide/vpc-route-tables.md).

**Next Steps**

To enable your fleet instances, app block builders, and image builders to
access the internet, complete the steps in [Enable Internet Access for Your Fleet, Image Builder, or App Block Builder in Amazon WorkSpaces Applications](managing-network-manual-enable-internet-access.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Create and Configure a New VPC

Internet Access for Your Fleet, Image Builder, or App Block Builder

All content copied from https://docs.aws.amazon.com/.
