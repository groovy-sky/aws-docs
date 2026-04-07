This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VerifiedAccessEndpoint

An AWS Verified Access endpoint specifies the application that AWS Verified Access provides access to. It must be
attached to an AWS Verified Access group. An AWS Verified Access endpoint must also have an attached access policy
before you attached it to a group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VerifiedAccessEndpoint",
  "Properties" : {
      "ApplicationDomain" : String,
      "AttachmentType" : String,
      "CidrOptions" : CidrOptions,
      "Description" : String,
      "DomainCertificateArn" : String,
      "EndpointDomainPrefix" : String,
      "EndpointType" : String,
      "LoadBalancerOptions" : LoadBalancerOptions,
      "NetworkInterfaceOptions" : NetworkInterfaceOptions,
      "PolicyDocument" : String,
      "PolicyEnabled" : Boolean,
      "RdsOptions" : RdsOptions,
      "SecurityGroupIds" : [ String, ... ],
      "SseSpecification" : SseSpecification,
      "Tags" : [ Tag, ... ],
      "VerifiedAccessGroupId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VerifiedAccessEndpoint
Properties:
  ApplicationDomain: String
  AttachmentType: String
  CidrOptions:
    CidrOptions
  Description: String
  DomainCertificateArn: String
  EndpointDomainPrefix: String
  EndpointType: String
  LoadBalancerOptions:
    LoadBalancerOptions
  NetworkInterfaceOptions:
    NetworkInterfaceOptions
  PolicyDocument: String
  PolicyEnabled: Boolean
  RdsOptions:
    RdsOptions
  SecurityGroupIds:
    - String
  SseSpecification:
    SseSpecification
  Tags:
    - Tag
  VerifiedAccessGroupId: String

```

## Properties

`ApplicationDomain`

The DNS name for users to reach your application.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AttachmentType`

The type of attachment used to provide connectivity between the AWS Verified Access endpoint and the
application.

_Required_: Yes

_Type_: String

_Allowed values_: `vpc`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CidrOptions`

The options for a CIDR endpoint.

_Required_: No

_Type_: [CidrOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-cidroptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the AWS Verified Access endpoint.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainCertificateArn`

The ARN of a public TLS/SSL certificate imported into or created with ACM.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointDomainPrefix`

A custom identifier that is prepended to the DNS name that is generated for the
endpoint.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointType`

The type of AWS Verified Access endpoint. Incoming application requests will be sent to an IP
address, load balancer or a network interface depending on the endpoint type
specified.

_Required_: Yes

_Type_: String

_Allowed values_: `load-balancer | network-interface | rds | cidr`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoadBalancerOptions`

The load balancer details if creating the AWS Verified Access endpoint as
`load-balancer` type.

_Required_: No

_Type_: [LoadBalancerOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-loadbalanceroptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceOptions`

The options for network-interface type endpoint.

_Required_: No

_Type_: [NetworkInterfaceOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-networkinterfaceoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyDocument`

The Verified Access policy document.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyEnabled`

The status of the Verified Access policy.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RdsOptions`

The options for an RDS endpoint.

_Required_: No

_Type_: [RdsOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-rdsoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The IDs of the security groups for the endpoint.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SseSpecification`

The options for additional server side encryption.

_Required_: No

_Type_: [SseSpecification](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-ssespecification.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-verifiedaccessendpoint-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VerifiedAccessGroupId`

The ID of the AWS Verified Access group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the Verified Access endpoint.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationTime`

The creation time.

`DeviceValidationDomain`

Use this to construct the redirect URI to add to your OIDC provider's allow list.

`EndpointDomain`

The DNS name generated for the endpoint.

`LastUpdatedTime`

The last updated time.

`Status`

The endpoint status.

`VerifiedAccessEndpointId`

The ID of the Verified Access endpoint.

`VerifiedAccessInstanceId`

The instance identifier.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CidrOptions
