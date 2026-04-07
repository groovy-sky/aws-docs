This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Authorizer

Specifies an authorizer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::Authorizer",
  "Properties" : {
      "AuthorizerFunctionArn" : String,
      "AuthorizerName" : String,
      "EnableCachingForHttp" : Boolean,
      "SigningDisabled" : Boolean,
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "TokenKeyName" : String,
      "TokenSigningPublicKeys" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::IoT::Authorizer
Properties:
  AuthorizerFunctionArn: String
  AuthorizerName: String
  EnableCachingForHttp: Boolean
  SigningDisabled: Boolean
  Status: String
  Tags:
    - Tag
  TokenKeyName: String
  TokenSigningPublicKeys:
    Key: Value

```

## Properties

`AuthorizerFunctionArn`

The authorizer's Lambda function ARN.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizerName`

The authorizer name.

_Required_: No

_Type_: String

_Pattern_: `[\w=,@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableCachingForHttp`

When `true`, the result from the authorizer's Lambda function is cached for
clients that use persistent HTTP connections. The results are cached for the time specified
by the Lambda function in `refreshAfterInSeconds`. This value doesn't affect
authorization of clients that use MQTT connections.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SigningDisabled`

Specifies whether AWS IoT validates the token signature in an authorization request.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

The status of the authorizer.

Valid values: `ACTIVE` \| `INACTIVE`

_Required_: No

_Type_: String

_Allowed values_: `ACTIVE | INACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the custom authorizer.

###### Note

For URI Request parameters use format: ...key1=value1&key2=value2...

For the CLI command-line parameter use format: &&tags
"key1=value1&key2=value2..."

For the cli-input-json file use format: "tags":
"key1=value1&key2=value2..."

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-authorizer-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenKeyName`

The key used to extract the token from the HTTP headers.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenSigningPublicKeys`

The public keys used to validate the token signature returned by your custom
authentication service.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the authorizer name. For example:

`{ "Ref": "MyAuthorizer" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the authorizer.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeviceCertExpirationAuditCheckConfiguration

Tag
