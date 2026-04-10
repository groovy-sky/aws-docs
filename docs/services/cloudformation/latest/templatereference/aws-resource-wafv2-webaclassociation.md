This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACLAssociation

###### Note

This is the latest version of **AWS WAF**, named AWS WAFV2, released in November, 2019. For
information, including how to migrate your AWS WAF resources from the
prior release, see the [AWS WAF developer guide](../../../waf/latest/developerguide/waf-chapter.md).

Use a web ACL association to define an association between a web ACL and a regional
application resource, to protect the resource. A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, an AWS AppSync
GraphQL API, an Amazon Cognito user pool, an AWS App Runner service,
an AWS Amplify application, or an AWS Verified Access instance.

For Amazon CloudFront, don't use this resource. Instead, use your CloudFront distribution configuration. To associate a web ACL with a distribution,
provide the Amazon Resource Name (ARN) of the [AWS::WAFv2::WebACL](aws-resource-wafv2-webacl.md) to your CloudFront distribution configuration. To disassociate a web ACL, provide an empty ARN.
For information, see [AWS::CloudFront::Distribution](../userguide/aws-resource-cloudfront-distribution.md).

**Required permissions for customer-managed IAM**
**policies**

This call requires permissions that are specific to the protected resource type. For
details, see [Permissions for AssociateWebACL](../../../waf/latest/developerguide/security-iam-service-with-iam.md#security_iam_action-AssociateWebACL) in the _AWS WAF Developer Guide_.

**Temporary inconsistencies during updates**

When you create or change a web ACL or other AWS WAF resources, the
changes take a small amount of time to propagate to all areas where the resources are
stored. The propagation time can be from a few seconds to a number of minutes.

The following are examples of the temporary inconsistencies that you might notice during
change propagation:

- After you create a web ACL, if you try to associate it with a resource, you might
get an exception indicating that the web ACL is unavailable.

- After you add a rule group to a web ACL, the new rule group rules might be in
effect in one area where the web ACL is used and not in another.

- After you change a rule action setting, you might see the old action in some
places and the new action in others.

- After you add an IP address to an IP set that is in use in a blocking rule, the
new address might be blocked in one area while still allowed in another.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFv2::WebACLAssociation",
  "Properties" : {
      "ResourceArn" : String,
      "WebACLArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::WAFv2::WebACLAssociation
Properties:
  ResourceArn: String
  WebACLArn: String

```

## Properties

`ResourceArn`

The Amazon Resource Name (ARN) of the resource to associate with the web ACL.

The ARN must be in one of the following formats:

- For an Application Load Balancer:
`arn:partition:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id`

- For an Amazon API Gateway REST API:
`arn:partition:apigateway:region::/restapis/api-id/stages/stage-name`

- For an AWS AppSync GraphQL API:
`arn:partition:appsync:region:account-id:apis/GraphQLApiId`

- For an Amazon Cognito user pool:
`arn:partition:cognito-idp:region:account-id:userpool/user-pool-id`

- For an AWS App Runner service:
`arn:partition:apprunner:region:account-id:service/apprunner-service-name/apprunner-service-id`

- For an AWS Verified Access instance:
`arn:partition:ec2:region:account-id:verified-access-instance/instance-id`

- For an AWS Amplify instance:
`arn:partition:amplify:region:account-id:apps/app-id`

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WebACLArn`

The Amazon Resource Name (ARN) of the web ACL that you want to associate with the
resource.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

The `Ref` for the resource, containing the resource name, physical ID, and
scope, formatted as follows: `name|id|scope`.

For example:
`my-webacl-name|1234a1a-a1b1-12a1-abcd-a123b123456|REGIONAL`.

## Examples

### Create a web ACL association

The following shows an example web ACL association specification.

#### YAML

```yaml

  SampleWebACLAssociation:
    Type: 'AWS::WAFv2::WebACLAssociation'
    Properties:
      WebACLArn: ExampleARNForWebACL
      ResourceArn: ExampleARNForRegionalResource
```

#### JSON

```json

  "SampleWebACLAssociation": {
    "Type": "AWS::WAFv2::WebACLAssociation",
    "Properties": {
      "WebACLArn": "WebACLArn",
      "ResourceArn": "APIGatewayOrALBArn"
    }
  }
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

XssMatchStatement

Next

All content copied from https://docs.aws.amazon.com/.
