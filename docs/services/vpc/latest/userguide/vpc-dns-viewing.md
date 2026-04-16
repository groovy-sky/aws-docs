---
title: "View DNS hostnames for your EC2 instance"
---

# View DNS hostnames for your EC2 instance

You can view the DNS hostnames for a running instance or a network interface using the
Amazon EC2 console or the command line. Knowing these hostnames is important for
connecting to your resources.

The **Public DNS (IPv4)** and **Private DNS** fields are
available when the DNS options are enabled for the VPC that is associated with the
instance. For more information, see [DNS attributes for your VPC](amazondns-concepts.md#vpc-dns-support).

## Instance

###### To view DNS hostnames for an instance using the console

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. In the navigation pane, choose **Instances**.

3. Select your instance from the list.

4. In the details pane, the **Public DNS (IPv4)** and
    **Private DNS** fields display the DNS hostnames, if
    applicable.

###### To view DNS hostnames for an instance using the command line

- [describe-instances](../../../cli/latest/reference/ec2/describe-instances.md) (AWS CLI)

- [Get-EC2Instance](../../../powershell/latest/reference/items/get-ec2instance.md) (AWS Tools for Windows PowerShell)

## Network interface

###### To view the private DNS hostname for a network interface using the console

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. In the navigation pane, choose **Network Interfaces**.

3. Select the network interface from the list.

4. In the details pane, the **Private DNS (IPv4)** field displays the
    private DNS hostname.

###### To view DNS hostnames for a network interface using the command line

- [describe-network-interfaces](../../../cli/latest/reference/ec2/describe-network-interfaces.md) (AWS CLI)

- [Get-EC2NetworkInterface](../../../powershell/latest/reference/items/get-ec2networkinterface.md) (AWS Tools for Windows PowerShell)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Understanding Amazon DNS

View and update DNS attributes for your VPC

All content copied from https://docs.aws.amazon.com/.
