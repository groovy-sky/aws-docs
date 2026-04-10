This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::WebApp EndpointDetails

The virtual private cloud (VPC) endpoint settings that are configured for your file
transfer protocol-enabled server. With a VPC endpoint, you can restrict access to your
server and resources only within your VPC. To control incoming internet traffic, invoke
the `UpdateServer` API and attach an Elastic IP address to your server's
endpoint.

###### Note

After May 19, 2021, you won't be able to create a server using
`EndpointType=VPC_ENDPOINT` in your AWS account if your account hasn't already
done so before May 19, 2021. If you have already created servers with
`EndpointType=VPC_ENDPOINT` in your AWS account on or before May 19, 2021,
you will not be affected. After this date, use
`EndpointType` = `VPC`.

For more information, see [Discontinuing\
the use of VPC\_ENDPOINT](../../../transfer/latest/userguide/create-server-in-vpc.md#deprecate-vpc-endpoint).

It is recommended that you use `VPC` as the `EndpointType`. With
this endpoint type, you have the option to directly associate up to three Elastic IPv4
addresses (BYO IP included) with your server's endpoint and use VPC security groups to
restrict traffic by the client's public IP address. This is not possible with
`EndpointType` set to `VPC_ENDPOINT`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Vpc" : Vpc
}

```

### YAML

```yaml

  Vpc:
    Vpc

```

## Properties

`Vpc`

Property description not available.

_Required_: No

_Type_: [Vpc](aws-properties-transfer-webapp-vpc.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Transfer::WebApp

IdentityProviderDetails

All content copied from https://docs.aws.amazon.com/.
