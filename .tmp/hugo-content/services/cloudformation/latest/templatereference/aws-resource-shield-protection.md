This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Shield::Protection

Enables AWS Shield Advanced for a specific AWS resource. The resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone, AWS Global Accelerator standard accelerator, Elastic IP Address, Application Load Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and Network Load Balancers by association with protected Amazon EC2 Elastic IP addresses.

**Configure a single `AWS::Shield::Protection`**

Use this protection to protect a single resource at a time.

To configure this Shield Advanced protection through CloudFormation, you must be subscribed to Shield Advanced. You can subscribe
through the [Shield Advanced console](https://console.aws.amazon.com/wafv2/shieldv2) and through
the APIs. For more information, see
[Subscribe to AWS Shield Advanced](../../../waf/latest/developerguide/enable-ddos-prem.md).

See example templates for Shield Advanced in CloudFormation at [aws-samples/aws-shield-advanced-examples](https://github.com/aws-samples/aws-shield-advanced-examples).

**Configure Shield Advanced using AWS CloudFormation and AWS Firewall Manager**

You might be able to use Firewall Manager with AWS CloudFormation to configure Shield Advanced across multiple accounts and protected resources. To do this, your accounts must be part of an organization in AWS Organizations. You can use Firewall Manager to configure Shield Advanced protections for any resource types except for Amazon Route 53 or AWS Global Accelerator.

For an example of this, see the one-click configuration guidance published by the AWS technical community at
[One-click deployment of Shield Advanced](https://youtu.be/LCA3FwMk_QE).

**Configure multiple protections through the Shield Advanced console**

You can add protection to multiple resources at once through the [Shield Advanced console](https://console.aws.amazon.com/wafv2/shieldv2).
For more information see
[Getting Started with AWS Shield Advanced](../../../waf/latest/developerguide/getting-started-ddos.md)
and [Managing resource protections in AWS Shield Advanced](../../../waf/latest/developerguide/ddos-manage-protected-resources.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Shield::Protection",
  "Properties" : {
      "ApplicationLayerAutomaticResponseConfiguration" : ApplicationLayerAutomaticResponseConfiguration,
      "HealthCheckArns" : [ String, ... ],
      "Name" : String,
      "ResourceArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Shield::Protection
Properties:
  ApplicationLayerAutomaticResponseConfiguration:
    ApplicationLayerAutomaticResponseConfiguration
  HealthCheckArns:
    - String
  Name: String
  ResourceArn: String
  Tags:
    - Tag

```

## Properties

`ApplicationLayerAutomaticResponseConfiguration`

The automatic application layer DDoS mitigation settings for the protection.
This configuration determines whether Shield Advanced automatically
manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.

If you use CloudFormation to manage the web ACLs that you use with Shield Advanced automatic mitigation, see the additional guidance
about web ACL management in the `AWS::WAFv2::WebACL` resource description.

_Required_: No

_Type_: [ApplicationLayerAutomaticResponseConfiguration](aws-properties-shield-protection-applicationlayerautomaticresponseconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HealthCheckArns`

The ARN (Amazon Resource Name) of the health check to associate with the protection. Health-based detection provides improved responsiveness and accuracy in attack detection and mitigation.

You can use this option with any resource type except for Route 53 hosted zones.

For more information, see [Configuring health-based detection using health checks](../../../waf/latest/developerguide/ddos-advanced-health-checks.md) in the _AWS Shield Advanced Developer Guide_.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2048 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the protection. For example, `My CloudFront distributions`.

###### Note

If you change the name of an existing protection, Shield Advanced deletes the
protection and replaces it with a new one. While this is happening, the protection isn't available on the AWS resource.

_Required_: Yes

_Type_: String

_Pattern_: `[ a-zA-Z0-9_\.\-]*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceArn`

The ARN (Amazon Resource Name) of the AWS resource that is protected.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-shield-protection-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN (Amazon Resource Name) of the protection.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ProtectionArn`

The ARN (Amazon Resource Name) of the new protection.

`ProtectionId`

The ID of the new protection.

## Examples

- [Create a network layer protection](#aws-resource-shield-protection--examples--Create_a_network_layer_protection)

- [Create an application layer protection](#aws-resource-shield-protection--examples--Create_an_application_layer_protection)

### Create a network layer protection

The following shows an example protection configuration for an Amazon EC2 Elastic IP address.

#### YAML

```yaml

Resources:
  EIP:
    Type: AWS::EC2::EIP
  Protection:
    Type: AWS::Shield::Protection
    Properties:
      Name: 'MyEIPProtection'
            ResourceArn: !Sub 'arn:${AWS::Partition}:ec2:${AWS::Region}:${AWS::AccountId}:eip-allocation/${EIP.AllocationId}'
```

#### JSON

```json

{
    "Resources": {
        "EIP": {
            "Type": "AWS::EC2::EIP"
        },
        "Protection": {
            "Type": "AWS::Shield::Protection",
            "Properties": {
                "Name": "MyEIPProtection",
                "ResourceArn": {
                    "Fn::Sub": "arn:${AWS::Partition}:ec2:${AWS::Region}:${AWS::AccountId}:eip-allocation/${EIP.AllocationId}"
                }
            }
        }
    }
}
```

### Create an application layer protection

The following shows an example layer 7 protection configuration for an application load balancer. The protection includes a health check and has application layer automatic response enabled. The load balancer must be associated with an AWS WAF web ACL that has a rate-based rule defined in it.

#### YAML

```yaml

Resources:
  # Create L7 Protection
  Protection:
    Type: AWS::Shield::Protection
    DependsOn:
    - WebACLAssociation
    Properties:
      Name: 'MyL7Protection'
      ResourceArn: !Ref ALB
      HealthCheckArns:
        - !Sub 'arn:${AWS::Partition}:route53:::healthcheck/${HealthCheck}'
      ApplicationLayerAutomaticResponseConfiguration:
        Status: ENABLED
        Action:
          Block: { }
```

#### JSON

```json

{
    "Resources": {
        "Protection": {
            "Type": "AWS::Shield::Protection",
            "DependsOn": [
                "WebACLAssociation"
            ],
            "Properties": {
                "Name": "MyL7Protection",
                "ResourceArn": {
                    "Ref": "ALB"
                },
                "HealthCheckArns": [
                    {
                        "Fn::Sub": "arn:${AWS::Partition}:route53:::healthcheck/${HealthCheck}"
                    }
                ],
                "ApplicationLayerAutomaticResponseConfiguration": {
                    "Status": "ENABLED",
                    "Action": {
                        "Block": {}
                    }
                }
            }
        }
    }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EmergencyContact

Action

All content copied from https://docs.aws.amazon.com/.
