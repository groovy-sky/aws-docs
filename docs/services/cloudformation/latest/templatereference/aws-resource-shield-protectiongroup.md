This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Shield::ProtectionGroup

Creates a grouping of protected resources so they can be handled as a collective. This resource grouping improves the accuracy of detection and reduces false positives.

To configure this resource through CloudFormation, you must be subscribed to AWS Shield Advanced. You can subscribe
through the [Shield Advanced console](https://console.aws.amazon.com/wafv2/shieldv2) and through
the APIs. For more information, see
[Subscribe to AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/enable-ddos-prem.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Shield::ProtectionGroup",
  "Properties" : {
      "Aggregation" : String,
      "Members" : [ String, ... ],
      "Pattern" : String,
      "ProtectionGroupId" : String,
      "ResourceType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Shield::ProtectionGroup
Properties:
  Aggregation: String
  Members:
    - String
  Pattern: String
  ProtectionGroupId: String
  ResourceType: String
  Tags:
    - Tag

```

## Properties

`Aggregation`

Defines how AWS Shield combines resource data for the group in order to detect, mitigate, and report events.

- `Sum` \- Use the total traffic across the group. This is a good choice for most cases. Examples include Elastic IP addresses for EC2 instances that scale manually or automatically.

- `Mean` \- Use the average of the traffic across the group. This is a good choice for resources that share traffic uniformly. Examples include accelerators and load balancers.

- `Max` \- Use the highest traffic from each resource. This is useful for resources that don't share traffic and for resources that share that traffic in a non-uniform way. Examples include Amazon CloudFront distributions and origin resources for CloudFront distributions.

_Required_: Yes

_Type_: String

_Allowed values_: `SUM | MEAN | MAX`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Members`

The ARNs (Amazon Resource Names) of the resources to include in the protection group. You must set this when you set `Pattern` to `ARBITRARY` and you must not set it for any other `Pattern` setting.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2048 | 10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pattern`

The criteria to use to choose the protected resources for inclusion in the group. You can include all resources that have protections, provide a list of resource ARNs (Amazon Resource Names), or include all resources of a specified resource type.

_Required_: Yes

_Type_: String

_Allowed values_: `ALL | ARBITRARY | BY_RESOURCE_TYPE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtectionGroupId`

The name of the protection group. You use this to identify the protection group in lists and to manage the protection group, for example to update, delete, or describe it.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9\-]*`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceType`

The resource type to include in the protection group. All protected resources of this type are included in the protection group.
You must set this when you set `Pattern` to `BY_RESOURCE_TYPE` and you must not set it for any other `Pattern` setting.

_Required_: No

_Type_: String

_Allowed values_: `CLOUDFRONT_DISTRIBUTION | ROUTE_53_HOSTED_ZONE | ELASTIC_IP_ALLOCATION | CLASSIC_LOAD_BALANCER | APPLICATION_LOAD_BALANCER | GLOBAL_ACCELERATOR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key:value pairs associated with an AWS resource. The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-shield-protectiongroup-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN (Amazon Resource Name) of the protection group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ProtectionGroupArn`

The ARN (Amazon Resource Name) of the new protection group.

## Examples

- [Create a protection group for all protected resources](#aws-resource-shield-protectiongroup--examples--Create_a_protection_group_for_all_protected_resources)

- [Create a protection group for protected Elastic IP address resources](#aws-resource-shield-protectiongroup--examples--Create_a_protection_group_for_protected_Elastic_IP_address_resources)

### Create a protection group for all protected resources

The following shows an example protection group configuration for all protected resources.

#### YAML

```yaml

Resources:
  ProtectionGroup:
    Type: AWS::Shield::ProtectionGroup
    Properties:
      ProtectionGroupId: 'ProtectionGroupForAllResources'
      Aggregation: SUM
      Pattern: ALL
```

#### JSON

```json

{
    "Resources": {
        "ProtectionGroup": {
            "Type": "AWS::Shield::ProtectionGroup",
            "Properties": {
                "ProtectionGroupId": "ProtectionGroupForAllResources",
                "Aggregation": "SUM",
                "Pattern": "ALL"
            }
        }
    }
}
```

### Create a protection group for protected Elastic IP address resources

The following shows an example protection group configuration for all Elastic IP address resources that have AWS Shield Advanced protection.

#### YAML

```yaml

Resources:
  ProtectionGroup:
    Type: AWS::Shield::ProtectionGroup
    Properties:
      ProtectionGroupId: 'ProtectionGroupForAllEIPResources'
      Aggregation: SUM
      Pattern: BY_RESOURCE_TYPE
      ResourceType: ELASTIC_IP_ALLOCATION
```

#### JSON

```json

{
    "Resources": {
        "ProtectionGroup": {
            "Type": "AWS::Shield::ProtectionGroup",
            "Properties": {
                "ProtectionGroupId": "ProtectionGroupForAllEIPResources",
                "Aggregation": "SUM",
                "Pattern": "BY_RESOURCE_TYPE",
                "ResourceType": "ELASTIC_IP_ALLOCATION"
            }
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
