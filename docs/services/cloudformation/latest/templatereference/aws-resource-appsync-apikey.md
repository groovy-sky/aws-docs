This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ApiKey

The `AWS::AppSync::ApiKey` resource creates a unique key that you can
distribute to clients who are executing GraphQL operations with AWS AppSync that require an API key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppSync::ApiKey",
  "Properties" : {
      "ApiId" : String,
      "Description" : String,
      "Expires" : Number
    }
}

```

### YAML

```yaml

Type: AWS::AppSync::ApiKey
Properties:
  ApiId: String
  Description: String
  Expires: Number

```

## Properties

`ApiId`

Unique AWS AppSync GraphQL API ID for this API key.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

Unique description of your API key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expires`

The time after which the API key expires. The date is represented as seconds since the epoch, rounded down
to the nearest hour.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of an `AWS::AppSync::ApiKey` resource to the
intrinsic `Ref` function, the function returns the ARN of the API key, such
as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/apikey/apikeya1bzhi`.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The
following are the available attributes and sample return values.

For more information about using `Fn::GetAtt`, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt).

`ApiKey`

The API key.

`ApiKeyId`

The API key ID.

`Arn`

The Amazon Resource Name (ARN) of the API key, such as
`arn:aws:appsync:us-east-1:123456789012:apis/graphqlapiid/apikey/apikeya1bzhi`.

## Examples

### API Key Creation Example

The following example creates an API key and associates it with an existing
GraphQL API by passing the GraphQL API ID as a parameter.

#### YAML

```yaml

Parameters:
  graphQlApiId:
    Type: String
  apiKeyDescription:
    Type: String
  apiKeyExpires:
    Type: Number
Resources:
  ApiKey:
    Type: AWS::AppSync::ApiKey
    Properties:
      ApiId:
	Ref: graphQlApiId
      Description:
        Ref: apiKeyDescription
      Expires:
        Ref: apiKeyExpires

```

#### JSON

```json

{
  "Parameters": {
    "graphQlApiId": {
      "Type": "String"
    },
    "apiKeyDescription": {
      "Type": "String"
    },
    "apiKeyExpires": {
      "Type": "Number"
    }
  },
  "Resources": {
    "ApiKey": {
      "Type": "AWS::AppSync::ApiKey",
      "Properties": {
        "ApiId": {
           "Ref": "graphQlApiId"
        },
        "Description": {
          "Ref": "apiKeyDescription"
        },
        "Expires": {
          "Ref": "apiKeyExpires"
        }
      }
    }
  }
}

```

## See also

- [CreateApiKey](https://docs.aws.amazon.com/appsync/latest/APIReference/API_CreateApiKey.html) operation in the _AWS AppSync API Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppSync::ApiCache

AWS::AppSync::ChannelNamespace
