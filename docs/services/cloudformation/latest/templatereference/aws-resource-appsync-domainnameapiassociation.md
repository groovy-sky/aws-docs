This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::DomainNameApiAssociation

The `AWS::AppSync::DomainNameApiAssociation` resource represents the
mapping of your custom domain name to the assigned API URL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::DomainNameApiAssociation",
  "Properties" : {
      "ApiId" : String,
      "DomainName" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::DomainNameApiAssociation
Properties:
  ApiId: String
  DomainName: String

```

## Properties

`ApiId`

The API ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The domain name.

_Required_: Yes

_Type_: String

_Pattern_: `^(\*[a-z\d-]*\.)?([a-z\d-]+\.)+[a-z\d-]+$`

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::AppSync::FunctionConfiguration
