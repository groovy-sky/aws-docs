This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::DomainNameV2

The `AWS::ApiGateway::DomainNameV2` resource specifies a custom domain name for your private APIs
in API Gateway. You can use a private custom domain name to provide a URL for your private API that's more
intuitive and easier to recall.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGateway::DomainNameV2",
  "Properties" : {
      "CertificateArn" : String,
      "DomainName" : String,
      "EndpointAccessMode" : String,
      "EndpointConfiguration" : EndpointConfiguration,
      "Policy" : Json,
      "RoutingMode" : String,
      "SecurityPolicy" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ApiGateway::DomainNameV2
Properties:
  CertificateArn: String
  DomainName: String
  EndpointAccessMode: String
  EndpointConfiguration:
    EndpointConfiguration
  Policy: Json
  RoutingMode: String
  SecurityPolicy: String
  Tags:
    - Tag

```

## Properties

`CertificateArn`

The reference to an AWS-managed certificate that will be used by the private endpoint for this domain name. AWS Certificate Manager is the only supported source.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

Represents a custom domain name as a user-friendly host name of an API (RestApi).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointAccessMode`

The endpoint access mode for your DomainName.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndpointConfiguration`

The endpoint configuration to indicate the types of endpoints an API (RestApi) or its custom domain name (DomainName) has and the IP address types that can invoke it.

_Required_: No

_Type_: [EndpointConfiguration](aws-properties-apigateway-domainnamev2-endpointconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

A stringified JSON policy document that applies to the `execute-api` service for this DomainName regardless of the caller and Method
configuration. You can use `Fn::ToJsonString` to enter your `policy`. For more information,
see [Fn::ToJsonString](../userguide/intrinsic-function-reference-tojsonstring.md).

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoutingMode`

The routing mode for this domain name. The routing mode determines how API Gateway sends traffic from your custom domain name to your private APIs.

_Required_: No

_Type_: String

_Allowed values_: `BASE_PATH_MAPPING_ONLY | ROUTING_RULE_THEN_BASE_PATH_MAPPING | ROUTING_RULE_ONLY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityPolicy`

The Transport Layer Security (TLS) version + cipher suite for this DomainName.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The collection of tags. Each tag element is associated with a given resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-apigateway-domainnamev2-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the domain name ARN.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DomainNameArn`

The ARN of the domain name.

`DomainNameId`

The domain name ID.

## Examples

- [Private custom domain name example](#aws-resource-apigateway-domainnamev2--examples--Private_custom_domain_name_example)

- [Private custom domain name example with routing mode](#aws-resource-apigateway-domainnamev2--examples--Private_custom_domain_name_example_with_routing_mode)

### Private custom domain name example

The following example creates a `DomainNameV2` resource named `MyDomainName`.

#### JSON

```json

{
    "MyDomainName": {
        "Type": "AWS::ApiGateway::DomainNameV2",
        "Properties": {
            "DomainName": "private.example.com",
            "CertificateArn": "arn:aws:acm:us-west-2:123456789:certificate/abcd-000-1234-0000-000000abcd",
            "EndpointConfiguration": {
                "Types": [
                    "PRIVATE"
                ]
            },
            "SecurityPolicy": "TLS_1_2",
            "Policy": "{\n\"Version\": \"2012-10-17\",\n\"Statement\": [\n{\n\"Effect\": \"Allow\",\n\"Principal\": \"*\",\n\"Action\": \"execute-api:Invoke\",\n\"Resource\": [\n\"execute-api:/*\"\n]\n},\n {\n\"Effect\": \"Deny\",\n\"Principal\": \"*\",\n\"Action\": \"execute-api:Invoke\",\n\"Resource\": [\n\"execute-api:/*\"\n],\n\"Condition\" : {\n\"StringNotEquals\": {\n\"aws:SourceVpce\": \"vpce-abcd1234efg\"\n}\n}\n}\n]\n}"
        }
    }
}
```

#### YAML

```yaml

MyDomainName:
  Type: AWS::ApiGateway::DomainNameV2
  Properties:
    DomainName: private.example.com
    CertificateArn: arn:aws:acm:us-west-2:123456789:certificate/abcd-000-1234-0000-000000abcd
    EndpointConfiguration:
      Types:
        - PRIVATE
    SecurityPolicy: TLS_1_2
    Policy:
        Statement:
            - Action: 'execute-api:Invoke'
              Effect: Allow
              Principal: '*'
              Resource: 'execute-api:/*'
            - Action: 'execute-api:Invoke'
              Condition:
                StringNotEquals:
                  'aws:SourceVpce': !Ref EndpointID
              Effect: Deny
              Principal: '*'
              Resource: 'execute-api:/*'
        Version: 2012-10-17
```

### Private custom domain name example with routing mode

The following example creates a `DomainNameV2` resource named `MyDomainName` with a RoutingMode of `ROUTING_RULE_ONLY`.

#### JSON

```json

{
  "MyDomainName": {
    "Type": "AWS::ApiGateway::DomainNameV2",
    "Properties": {
      "DomainName": "private.example.com",
      "CertificateArn": "arn:aws:acm:us-west-2:123456789:certificate/abcd-000-1234-0000-000000abcd",
      "EndpointConfiguration": {
        "Types": [
          "PRIVATE"
        ]
      },
      "SecurityPolicy": "TLS_1_2",
      "Policy": "{\n\"Version\": \"2012-10-17\",\n\"Statement\": [\n{\n\"Effect\": \"Allow\",\n\"Principal\": \"*\",\n\"Action\": \"execute-api:Invoke\",\n\"Resource\": [\n\"execute-api:/*\"\n]\n},\n {\n\"Effect\": \"Deny\",\n\"Principal\": \"*\",\n\"Action\": \"execute-api:Invoke\",\n\"Resource\": [\n\"execute-api:/*\"\n],\n\"Condition\" : {\n\"StringNotEquals\": {\n\"aws:SourceVpce\": \"vpce-abcd1234efg\"\n}\n}\n}\n]\n}",
      "RoutingMode": "ROUTING_RULE_ONLY"
    }
  }
}
```

#### YAML

```yaml

MyDomainName:
    Type: AWS::ApiGateway::DomainNameV2
    Properties:
      DomainName: private.example.com
      CertificateArn: arn:aws:acm:us-west-2:123456789:certificate/abcd-000-1234-0000-000000abcd
      EndpointConfiguration:
        Types:
          - PRIVATE
      SecurityPolicy: "TLS_1_2"
      Policy:
        Statement:
            - Action: 'execute-api:Invoke'
              Effect: Allow
              Principal: '*'
              Resource: 'execute-api:/*'
            - Action: 'execute-api:Invoke'
              Condition:
                StringNotEquals:
                  'aws:SourceVpce': !Ref EndpointID
              Effect: Deny
              Principal: '*'
              Resource: 'execute-api:/*'
        Version: 2012-10-17
      RoutingMode: ROUTING_RULE_ONLY
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EndpointConfiguration

All content copied from https://docs.aws.amazon.com/.
