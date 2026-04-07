This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::ClientCertificate

The `AWS::ApiGateway::ClientCertificate` resource creates a client certificate that API Gateway uses to configure client-side SSL authentication for sending requests to the integration endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::ClientCertificate",
  "Properties" : {
      "Description" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::ClientCertificate
Properties:
  Description: String
  Tags:
    - Tag

```

## Properties

`Description`

The description of the client certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigateway-clientcertificate-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the client certificate name, such as `abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ClientCertificateId`

The ID for the client certificate. For example: `abc123`.

## Examples

### Create client certificate

The following example creates a client certificate that you can use with an API Gateway deployment and stage.

#### JSON

```json

{
    "TestClientCertificate": {
        "Type": "AWS::ApiGateway::ClientCertificate",
        "Properties": {
            "Description": "A test client certificate"
        }
    }
}
```

#### YAML

```yaml

TestClientCertificate:
  Type: 'AWS::ApiGateway::ClientCertificate'
  Properties:
    Description: A test client certificate

```

## See also

- [clientcertificate:generate](https://docs.aws.amazon.com/apigateway/latest/api/API_GenerateClientCertificate.html) in the _Amazon API Gateway REST API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGateway::BasePathMappingV2

Tag
