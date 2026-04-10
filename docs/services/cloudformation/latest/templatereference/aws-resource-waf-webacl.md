This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAF::WebACL

###### Note

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

Contains the `Rules` that identify the requests that you want to allow, block, or count. In a `WebACL`, you also specify a
default action ( `ALLOW` or `BLOCK`), and the action for each `Rule` that you add to a
`WebACL`, for example, block requests from specified IP addresses or block requests from specified referrers.
You also associate the `WebACL` with a Amazon CloudFront distribution to identify the requests that you want AWS WAF to filter.
If you add more than one `Rule` to a `WebACL`, a request needs to match only one of the specifications
to be allowed, blocked, or counted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAF::WebACL",
  "Properties" : {
      "DefaultAction" : WafAction,
      "MetricName" : String,
      "Name" : String,
      "Rules" : [ ActivatedRule, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::WAF::WebACL
Properties:
  DefaultAction:
    WafAction
  MetricName: String
  Name: String
  Rules:
    - ActivatedRule

```

## Properties

`DefaultAction`

The action to perform if none of the `Rules` contained in the `WebACL` match. The action is specified by the
`WafAction` object.

_Required_: Yes

_Type_: [WafAction](aws-properties-waf-webacl-wafaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metrics for this `WebACL`. The name can contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128 and minimum length one. It can't contain
whitespace or metric names reserved for AWS WAF, including "All" and "Default\_Action." You can't change `MetricName` after you create the `WebACL`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A friendly name or description of the `WebACL`. You can't change the name of a `WebACL` after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Rules`

An array that contains the action for each `Rule` in a `WebACL`, the priority of the `Rule`,
and the ID of the `Rule`.

_Required_: No

_Type_: Array of [ActivatedRule](aws-properties-waf-webacl-activatedrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name, such as 1234a1a-a1b1-12a1-abcd-a123b123456.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Create a Web ACL](#aws-resource-waf-webacl--examples--Create_a_Web_ACL)

- [Associate a Web ACL with a Amazon CloudFront Distribution](#aws-resource-waf-webacl--examples--Associate_a_Web_ACL_with_a_Distribution)

### Create a Web ACL

The following example defines a web ACL that allows, by default, any web request.
However, if the request matches any rule, AWS WAF blocks the request. AWS WAF evaluates
each rule in priority order, starting with the lowest value.

#### JSON

```json

"MyWebACL": {
  "Type": "AWS::WAF::WebACL",
  "Properties": {
    "Name": "WebACL to with three rules",
    "DefaultAction": {
      "Type": "ALLOW"
    },
    "MetricName" : "MyWebACL",
    "Rules": [
      {
        "Action" : {
          "Type" : "BLOCK"
        },
        "Priority" : 1,
        "RuleId" : { "Ref" : "MyRule" }
      },
      {
        "Action" : {
          "Type" : "BLOCK"
        },
        "Priority" : 2,
        "RuleId" : { "Ref" : "BadReferersRule" }
      },
      {
        "Action" : {
          "Type" : "BLOCK"
        },
        "Priority" : 3,
        "RuleId" : { "Ref" : "SqlInjRule" }
      }
    ]
  }
}
```

#### YAML

```yaml

MyWebACL:
  Type: "AWS::WAF::WebACL"
  Properties:
    Name: "WebACL to with three rules"
    DefaultAction:
      Type: "ALLOW"
    MetricName: "MyWebACL"
    Rules:
      -
        Action:
          Type: "BLOCK"
        Priority: 1
        RuleId:
          Ref: "MyRule"
      -
        Action:
          Type: "BLOCK"
        Priority: 2
        RuleId:
          Ref: "BadReferersRule"
      -
        Action:
          Type: "BLOCK"
        Priority: 3
        RuleId:
          Ref: "SqlInjRule"
```

### Associate a Web ACL with a Amazon CloudFront Distribution

The follow example associates the `MyWebACL` web ACL with a Amazon CloudFront distribution.
The web ACL restricts which requests can access content served by Amazon CloudFront.

#### JSON

```json

"myDistribution": {
  "Type": "AWS::CloudFront::Distribution",
  "Properties": {
    "DistributionConfig": {
      "WebACLId": { "Ref" : "MyWebACL" },
      "Origins": [
        {
          "DomainName": "test.example.com",
          "Id": "myCustomOrigin",
          "CustomOriginConfig": {
            "HTTPPort": "80",
            "HTTPSPort": "443",
            "OriginProtocolPolicy": "http-only"
          }
        }
      ],
      "Enabled": "true",
      "Comment": "TestDistribution",
      "DefaultRootObject": "index.html",
      "DefaultCacheBehavior": {
        "TargetOriginId": "myCustomOrigin",
        "SmoothStreaming" : "false",
        "ForwardedValues": {
          "QueryString": "false",
          "Cookies" : { "Forward" : "all" }
        },
        "ViewerProtocolPolicy": "allow-all"
      },
      "CustomErrorResponses" : [
        {
          "ErrorCode" : "404",
          "ResponsePagePath" : "/error-pages/404.html",
          "ResponseCode" : "200",
          "ErrorCachingMinTTL" : "30"
        }
      ],
      "PriceClass" : "PriceClass_200",
      "Restrictions" : {
        "GeoRestriction" : {
          "RestrictionType" : "whitelist",
          "Locations" : [ "AQ", "CV" ]
        }
      },
      "ViewerCertificate" : { "CloudFrontDefaultCertificate" : "true" }
    }
  }
}
```

#### YAML

```yaml

myDistribution:
  Type: "AWS::CloudFront::Distribution"
  Properties:
    DistributionConfig:
      WebACLId:
        Ref: "MyWebACL"
      Origins:
        -
          DomainName: "test.example.com"
          Id: "myCustomOrigin"
          CustomOriginConfig:
            HTTPPort: "80"
            HTTPSPort: "443"
            OriginProtocolPolicy: "http-only"
      Enabled: "true"
      Comment: "TestDistribution"
      DefaultRootObject: "index.html"
      DefaultCacheBehavior:
        TargetOriginId: "myCustomOrigin"
        SmoothStreaming: "false"
        ForwardedValues:
          QueryString: "false"
          Cookies:
            Forward: "all"
        ViewerProtocolPolicy: "allow-all"
      CustomErrorResponses:
        -
          ErrorCode: "404"
          ResponsePagePath: "/error-pages/404.html"
          ResponseCode: "200"
          ErrorCachingMinTTL: "30"
      PriceClass: "PriceClass_200"
      Restrictions:
        GeoRestriction:
          RestrictionType: "whitelist"
          Locations:
            - "AQ"
            - "CV"
      ViewerCertificate:
        CloudFrontDefaultCertificate: "true"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SqlInjectionMatchTuple

ActivatedRule

All content copied from https://docs.aws.amazon.com/.
