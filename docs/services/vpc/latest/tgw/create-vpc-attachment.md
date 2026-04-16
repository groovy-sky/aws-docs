---
title: "Create a VPC attachment in AWS Transit Gateway"
---

# Create a VPC attachment in AWS Transit Gateway

###### To create a VPC attachment using the console

01. Open the Amazon VPC console at
     [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).

02. On the navigation pane, choose **Transit Gateway**
    **Attachments**.

03. Choose **Create transit gateway attachment**.

04. For **Name tag**, optionally enter a name for the transit gateway attachment.

05. For **Transit gateway ID**, choose the transit gateway for the attachment. You can
     choose a transit gateway that you own or a transit gateway that was shared with you.

06. For **Attachment type**, choose **VPC**.

07. Choose whether to enable **DNS Support**, **IPv6**
    **Support** and **Appliance mode**
    **support**.

    If appliance mode is chosen, traffic flow between a source and destination
     uses the same Availability Zone for the VPC attachment for the lifetime of
     that flow.

08. Choose whether to enable **Security Group Referencing support**. Enable this feature to reference
     a security group across VPCs attached to a transit gateway. For more information about security group referencing, see [Security group referencing](tgw-vpc-attachments.md#vpc-attachment-security).

09. Choose whether to enable **IPv6**
    **Support**.

10. For **VPC ID**, choose the VPC to attach to the transit gateway.

    This VPC must have at least one subnet associated with it.

11. For **Subnet IDs**, select one subnet for each Availability Zone
     to be used by the transit gateway to route traffic. You must select at least one subnet. You can
     select only one subnet per Availability Zone.

12. Choose **Create transit gateway attachment**.

###### To create a VPC attachment using the AWS CLI

Use the [create-transit-gateway-vpc-attachment](../../../cli/latest/reference/ec2/create-transit-gateway-vpc-attachment.md) command.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VPC attachments

Modify a VPC attachment

All content copied from https://docs.aws.amazon.com/.
