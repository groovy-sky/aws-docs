This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::WebACLAssociation

###### Note

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

The AWS::WAFRegional::WebACLAssociation resource associates an AWS WAF Regional web access control group (ACL) with a resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::WAFRegional::WebACLAssociation",
  "Properties" : {
      "ResourceArn" : String,
      "WebACLId" : String
    }
}

```

### YAML

```yaml

Type: AWS::WAFRegional::WebACLAssociation
Properties:
  ResourceArn: String
  WebACLId: String

```

## Properties

`ResourceArn`

The Amazon Resource Name (ARN) of the resource to protect with the web ACL.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WebACLId`

A unique identifier (ID) for the web ACL.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Fn::GetAtt

## Examples

### Associate an Application Load Balancer resource with a web ACL

The following example associates an Application Load Balancer resource with a web ACL.

#### JSON

```json

"MyWebACLAssociation": {
  "Type": "AWS::WAFRegional::WebACLAssociation",
  "Properties": {
    "ResourceArn": { "Ref": "MyLoadBalancer" },
    "WebACLId": { "Ref": "MyWebACL" }
  }
}
```

#### YAML

```yaml

MyWebACLAssociation:
  Type: "AWS::WAFRegional::WebACLAssociation"
  Properties:
    ResourceArn:
      Ref: MyLoadBalancer
    WebACLId:
      Ref: MyWebACL
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

AWS::WAFRegional::XssMatchSet

All content copied from https://docs.aws.amazon.com/.
