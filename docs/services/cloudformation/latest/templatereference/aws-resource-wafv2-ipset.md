This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::IPSet

###### Note

This is the latest version of **AWS WAF**, named AWS WAFV2, released in November, 2019. For
information, including how to migrate your AWS WAF resources from the
prior release, see the [AWS WAF developer guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).

Use an AWS::WAFv2::IPSet to identify web requests that originate from specific IP
addresses or ranges of IP addresses. For example, if you're receiving a lot of requests
from a ranges of IP addresses, you can configure AWS WAF to block them
using an IP set that lists those IP addresses.

You use an IP set by providing its Amazon Resource Name (ARN) to the rule statement
`IPSetReferenceStatement`, when you add a rule to a rule group or web ACL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFv2::IPSet",
  "Properties" : {
      "Addresses" : [ String, ... ],
      "Description" : String,
      "IPAddressVersion" : String,
      "Name" : String,
      "Scope" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAFv2::IPSet
Properties:
  Addresses:
    - String
  Description: String
  IPAddressVersion: String
  Name: String
  Scope: String
  Tags:
    - Tag

```

## Properties

`Addresses`

Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses that you want AWS WAF to inspect for in incoming requests. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.

Example address strings:

- For requests that originated from the IP address 192.0.2.44, specify `192.0.2.44/32`.

- For requests that originated from IP addresses from 192.0.2.0 to 192.0.2.255, specify
`192.0.2.0/24`.

- For requests that originated from the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128`.

- For requests that originated from IP addresses 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64`.

For more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

Example JSON `Addresses` specifications:

- Empty array: `"Addresses": []`

- Array with one address: `"Addresses": ["192.0.2.44/32"]`

- Array with three addresses: `"Addresses": ["192.0.2.44/32", "192.0.2.0/24", "192.0.0.0/16"]`

- INVALID specification: `"Addresses": [""]` INVALID

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the IP set that helps with identification.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9=:#@/\-,.][a-zA-Z0-9+=:#@/\-,.\s]+[a-zA-Z0-9+=:#@/\-,.]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IPAddressVersion`

The version of the IP addresses, either `IPV4` or `IPV6`.

_Required_: Yes

_Type_: String

_Allowed values_: `IPV4 | IPV6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the IP set. You cannot change the name of an `IPSet` after you create it.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,128}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scope`

Specifies whether this is for an Amazon CloudFront distribution or for a regional
application. A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool, an AWS App Runner service, or an AWS Verified Access instance. Valid Values are `CLOUDFRONT` and
`REGIONAL`.

###### Note

For `CLOUDFRONT`, you must create your WAFv2 resources in the US East (N.
Virginia) Region, `us-east-1`.

_Required_: Yes

_Type_: String

_Allowed values_: `CLOUDFRONT | REGIONAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key:value pairs associated with an AWS resource. The key:value pair can
be anything you define. Typically, the tag key represents a category (such as
"environment") and the tag value represents a specific value within that category (such as
"test," "development," or "production"). You can add up to 50 tags to each AWS resource.

###### Note

To modify tags on existing resources, use the AWS WAF APIs or
command line interface. With AWS CloudFormation, you can only add tags to AWS WAF resources during resource creation.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-ipset-tag.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The `Ref` for the resource, containing the resource name, physical ID, and
scope, formatted as follows: `name|id|scope`.

For example:
`my-webacl-name|1234a1a-a1b1-12a1-abcd-a123b123456|REGIONAL`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the IP set.

`Id`

The ID of the IP set.

## Examples

### Create an IP set

The following shows an example IP set specification.

#### YAML

```yaml

 SampleIPSet:
      Type: 'AWS::WAFv2::IPSet'
      Properties:
        Description: SampleIPSet
        Name: SampleIPSet
        Scope: REGIONAL
        IPAddressVersion: IPV4
        Addresses:
          - 1.2.1.1/32
```

#### JSON

```json

 "SampleIPSet": {
      "Type": "AWS::WAFv2::IPSet",
      "Properties": {
        "Description": "SampleIPSet",
        "Name": "SampleIPSSet",
        "Scope": "REGIONAL",
        "IPAddressVersion": "IPV4",
        "Addresses": [
          "1.2.1.1/32"
        ]
      }
    }
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS WAF V2

Tag
