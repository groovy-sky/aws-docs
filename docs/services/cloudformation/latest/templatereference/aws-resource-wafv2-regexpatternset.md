This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RegexPatternSet

###### Note

This is the latest version of **AWS WAF**, named AWS WAFV2, released in November, 2019. For
information, including how to migrate your AWS WAF resources from the
prior release, see the [AWS WAF developer guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).

Use an AWS::WAFv2::RegexPatternSet to have AWS WAF inspect a web
request component for a specific set of regular expression patterns.

You use a regex pattern set by providing its Amazon Resource Name (ARN) to the rule
statement `RegexPatternSetReferenceStatement`, when you add a rule to a rule
group or web ACL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFv2::RegexPatternSet",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "RegularExpressionList" : [ String, ... ],
      "Scope" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAFv2::RegexPatternSet
Properties:
  Description: String
  Name: String
  RegularExpressionList:
    - String
  Scope: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the set that helps with identification.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9=:#@/\-,.][a-zA-Z0-9+=:#@/\-,.\s]+[a-zA-Z0-9+=:#@/\-,.]{1,256}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the set. You cannot change the name after you create the set.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_-]{1,128}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RegularExpressionList`

The regular expression patterns in the set.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Scope`

Specifies whether this is for an Amazon CloudFront distribution or for a regional
application. For an AWS Amplify application, use `CLOUDFRONT`.
A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway
REST API, an AWS AppSync GraphQL API, an Amazon Cognito user pool,
an AWS App Runner service, or an AWS Verified Access instance. Valid Values are
`CLOUDFRONT` and `REGIONAL`.

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

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-regexpatternset-tag.html)

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

The Amazon Resource Name (ARN) of the regex pattern set.

`Id`

The ID of the regex pattern set.

## Examples

### Create a regex pattern set

The following shows an example regex pattern set specification.

#### YAML

```yaml

 ExampleRegexPatternSet:
    Type: AWS::WAFv2::RegexPatternSet
    Properties:
      Name: ExampleRegexPatternSet
      Scope: REGIONAL
      Description: This is an example RegexPatternSet
      RegularExpressionList:
        - ^foobar$
        - ^example$
```

#### JSON

```json

 "ExampleRegexPatternSet": {
      "Type": "AWS::WAFv2::RegexPatternSet",
      "Properties": {
        "Name": "ExampleRegexPatternSet1",
        "Scope": "REGIONAL",
        "Description": "This is an example RegexPatternSet",
        "RegularExpressionList": [
          "^foobar$",
          "^example$"
        ]
      }
    }
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SingleHeader

Tag
