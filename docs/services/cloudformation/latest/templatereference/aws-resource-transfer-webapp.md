This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::WebApp

Creates a web app based on specified parameters, and returns the ID for the new web
app. You can configure the web app to be publicly accessible or hosted within a VPC.

For more information about using VPC endpoints with AWS Transfer Family, see [Create a Transfer Family web app in a VPC](https://docs.aws.amazon.com/transfer/latest/userguide/create-webapp-in-vpc.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::WebApp",
  "Properties" : {
      "AccessEndpoint" : String,
      "EndpointDetails" : EndpointDetails,
      "IdentityProviderDetails" : IdentityProviderDetails,
      "Tags" : [ Tag, ... ],
      "WebAppCustomization" : WebAppCustomization,
      "WebAppEndpointPolicy" : String,
      "WebAppUnits" : WebAppUnits
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::WebApp
Properties:
  AccessEndpoint: String
  EndpointDetails:
    EndpointDetails
  IdentityProviderDetails:
    IdentityProviderDetails
  Tags:
    - Tag
  WebAppCustomization:
    WebAppCustomization
  WebAppEndpointPolicy: String
  WebAppUnits:
    WebAppUnits

```

## Properties

`AccessEndpoint`

The `AccessEndpoint` is the URL that you provide to your users for them to
interact with the Transfer Family web app. You can specify a custom URL or use the default
value.

Before you enter a custom URL for this parameter, follow the steps described in [Update\
your access endpoint with a custom URL](https://docs.aws.amazon.com/transfer/latest/userguide/webapp-customize.html).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointDetails`

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
the use of VPC\_ENDPOINT](https://docs.aws.amazon.com/transfer/latest/userguide/create-server-in-vpc.html#deprecate-vpc-endpoint).

It is recommended that you use `VPC` as the `EndpointType`. With
this endpoint type, you have the option to directly associate up to three Elastic IPv4
addresses (BYO IP included) with your server's endpoint and use VPC security groups to
restrict traffic by the client's public IP address. This is not possible with
`EndpointType` set to `VPC_ENDPOINT`.

_Required_: No

_Type_: [EndpointDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-webapp-endpointdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProviderDetails`

You can provide a structure that contains the details for the identity provider to use
with your web app.

For more details about this parameter, see [Configure your identity\
provider for Transfer Family web apps](https://docs.aws.amazon.com/transfer/latest/userguide/webapp-identity-center.html).

_Required_: Yes

_Type_: [IdentityProviderDetails](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-webapp-identityproviderdetails.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs that can be used to group and search for web apps. Tags are metadata
attached to web apps for any purpose.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-webapp-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebAppCustomization`

A structure that contains the customization fields for the web app. You can provide a
title, logo, and icon to customize the appearance of your web app.

_Required_: No

_Type_: [WebAppCustomization](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-webapp-webappcustomization.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WebAppEndpointPolicy`

Setting for the type of endpoint policy for the web app. The default value is
`STANDARD`.

If your web app was created in an AWS GovCloud (US) Region, the value of this
parameter can be `FIPS`, which indicates the web app endpoint is
FIPS-compliant.

_Required_: No

_Type_: String

_Allowed values_: `STANDARD | FIPS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WebAppUnits`

A union that contains the value for number of concurrent connections or the user
sessions on your web app.

_Required_: No

_Type_: [WebAppUnits](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-transfer-webapp-webappunits.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the web app ARN, such as
`arn:aws:transfer:us-east-2:123456789012:webapp/webapp-01234567890abcdef`
.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the web app.

`IdentityProviderDetails.ApplicationArn`

Property description not available.

`VpcEndpointId`

The identifier of the VPC endpoint created for the web app.

`WebAppId`

The unique identifier for the web app.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

EndpointDetails
