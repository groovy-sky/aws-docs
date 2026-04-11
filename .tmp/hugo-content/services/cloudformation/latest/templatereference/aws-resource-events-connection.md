This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection

Creates a connection. A connection defines the authorization type and credentials to use
for authorization with an API destination HTTP endpoint.

For more information, see [Connections for endpoint targets](../../../eventbridge/latest/userguide/eb-target-connection.md) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Events::Connection",
  "Properties" : {
      "AuthorizationType" : String,
      "AuthParameters" : AuthParameters,
      "Description" : String,
      "InvocationConnectivityParameters" : InvocationConnectivityParameters,
      "KmsKeyIdentifier" : String,
      "Name" : String
    }
}

```

### YAML

```yaml

Type: AWS::Events::Connection
Properties:
  AuthorizationType: String
  AuthParameters:
    AuthParameters
  Description: String
  InvocationConnectivityParameters:
    InvocationConnectivityParameters
  KmsKeyIdentifier: String
  Name: String

```

## Properties

`AuthorizationType`

The type of authorization to use for the connection.

###### Note

OAUTH tokens are refreshed when a 401 or 407 response is returned.

_Required_: No

_Type_: String

_Allowed values_: `API_KEY | BASIC | OAUTH_CLIENT_CREDENTIALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthParameters`

The
authorization parameters to use to authorize with the endpoint.

You must include only authorization parameters for the `AuthorizationType` you specify.

_Required_: No

_Type_: [AuthParameters](aws-properties-events-connection-authparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the connection to create.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationConnectivityParameters`

For connections to private APIs, the parameters to use for invoking the API.

For more information, see [Connecting to private APIs](../../../eventbridge/latest/userguide/connection-private.md) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: [InvocationConnectivityParameters](aws-properties-events-connection-invocationconnectivityparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyIdentifier`

The identifier of the AWS KMS
customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this connection. The identifier can be the key
Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

If you do not specify a customer managed key identifier, EventBridge uses an
AWS owned key to encrypt the connection.

For more information, see [Identify and view keys](../../../kms/latest/developerguide/viewing-keys.md) in the _AWS Key Management Service_
_Developer Guide_.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-/:]*$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the connection to create.

_Required_: No

_Type_: String

_Pattern_: `[\.\-_A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the connection that was created by the
request.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the connection that was created by the request.

`ArnForPolicy`

Returns the Amazon Resource Name (ARN) of a connection in resource format, so it can be used in the `Resource` element of IAM permission policy statements.
For more information, see [Resource types defined by Amazon EventBridge](../../../service-authorization/latest/reference/list-amazoneventbridge.md#amazoneventbridge-resources-for-iam-policies) in the _Service Authorization Reference_.

For example, the following resource defines an IAM policy that grants permission to update a specific connection.

```

Resources:
  ExamplePolicy:
    Type: AWS::IAM::Policy
    Properties:
      PolicyName: ExamplePolicy
      PolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Action:
              - events:UpdateConnection
            Resource:
              - !GetAtt myConnection.ArnForPolicy
```

`AuthParameters.ConnectivityParameters.ResourceParameters.ResourceAssociationArn`

For connections to private APIs, the Amazon Resource Name (ARN) of the resource association EventBridge created between the connection and the private API's resource configuration.

For more information, see [Managing service network resource associations for connections](../../../eventbridge/latest/userguide/connection-private.md#connection-private-snra) in the _Amazon EventBridge User Guide_.

`InvocationConnectivityParameters.ResourceParameters.ResourceAssociationArn`

For connections to private APIs, the Amazon Resource Name (ARN) of the resource association EventBridge created between the connection and the private API's resource configuration.

For more information, see [Managing service network resource associations for connections](../../../eventbridge/latest/userguide/connection-private.md#connection-private-snra) in the _Amazon EventBridge User Guide_.

`SecretArn`

The ARN for the secret created for the connection.

## Examples

- [Create a connection with ApiKey authorization parameters](#aws-resource-events-connection--examples--Create_a_connection_with_ApiKey_authorization_parameters)

- [Create a connection with OAuth authorization parameters](#aws-resource-events-connection--examples--Create_a_connection_with_OAuth_authorization_parameters)

### Create a connection with ApiKey authorization parameters

The following example creates a connection named pagerduty-connection using ApiKey authorization and stores a secret from Secrets Manager.

#### JSON

```json

{
  "Resources": {
    "Connection": {
      "Type": "AWS::Events::Connection",
      "Properties": {
        "Name": "pagerduty-connection",
        "AuthorizationType": "API_KEY",
        "AuthParameters": {
          "ApiKeyAuthParameters": {
            "ApiKeyName": "Authorization",
            "ApiKeyValue": "{{resolve:secretsmanager:arn:aws:secretsmanager:us-west-2:123456789012:secret:pagerdutyApiToken-S9SoDa}}"
          },
          "InvocationHttpParameters": {
            "BodyParameters": [
              {
                "Key": "routing_key",
                "Value": "my-pagerduty-integration-key",
                "IsValueSecret": true
              }
            ]
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  Connection:
    Type: AWS::Events::Connection
    Properties:
      Name: pagerduty-connection
      AuthorizationType: API_KEY
      AuthParameters:
        ApiKeyAuthParameters:
          ApiKeyName: Authorization
          ApiKeyValue: '{{resolve:secretsmanager:arn:aws:secretsmanager:us-west-2:123456789012:secret:pagerdutyApiToken-S9SoDa}}'
        InvocationHttpParameters:
          BodyParameters:
            - Key: routing_key
              Value: my-pagerduty-integration-key
              IsValueSecret: true
```

### Create a connection with OAuth authorization parameters

The following example creates a connection named auth0-connection using OAuth authorization and stores a secret from Secrets Manager.

#### JSON

```json

{
  "Resources": {
    "Auth0Connection": {
      "Type": "AWS::Events::Connection",
      "Properties": {
        "Name": "auth0-connection",
        "AuthorizationType": "OAUTH_CLIENT_CREDENTIALS",
        "AuthParameters": {
          "OAuthParameters": {
            "AuthorizationEndpoint": "https://yourUserName.us.auth0.com/oauth/token",
            "ClientParameters": {
              "ClientID": "{{resolve:secretsmanager:arn:aws:secretsmanager:us-west-2:123456789012:secret:auth0ClientId}}",
              "ClientSecret": "{{resolve:secretsmanager:arn:aws:secretsmanager:us-west-2:123456789012:secret:auth0ClientSecret}}"
            },
            "HttpMethod": "POST",
            "OAuthHttpParameters": {
              "BodyParameters": [
                {
                  "Key": "audience",
                  "Value": "my-auth0-identifier",
                  "IsValueSecret": true
                }
              ]
            }
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  Auth0Connection:
    Type: AWS::Events::Connection
    Properties:
      Name: auth0-connection
      AuthorizationType: OAUTH_CLIENT_CREDENTIALS
      AuthParameters:
        OAuthParameters:
          AuthorizationEndpoint: https://yourUserName.us.auth0.com/oauth/token
          ClientParameters:
            ClientID: '{{resolve:secretsmanager:arn:aws:secretsmanager:us-west-2:123456789012:secret:auth0ClientId}}'
            ClientSecret: '{{resolve:secretsmanager:arn:aws:secretsmanager:us-west-2:123456789012:secret:auth0ClientSecret}}'
          HttpMethod: POST
          OAuthHttpParameters:
            BodyParameters:
              - Key: audience
                Value: my-auth0-identifier
                IsValueSecret: true
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Events::Archive

ApiKeyAuthParameters

All content copied from https://docs.aws.amazon.com/.
