This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolUser

The `AWS::Cognito::UserPoolUser` resource creates an Amazon Cognito user
pool user.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::UserPoolUser",
  "Properties" : {
      "ClientMetadata" : {Key: Value, ...},
      "DesiredDeliveryMediums" : [ String, ... ],
      "ForceAliasCreation" : Boolean,
      "MessageAction" : String,
      "UserAttributes" : [ AttributeType, ... ],
      "Username" : String,
      "UserPoolId" : String,
      "ValidationData" : [ AttributeType, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::UserPoolUser
Properties:
  ClientMetadata:
    Key: Value
  DesiredDeliveryMediums:
    - String
  ForceAliasCreation: Boolean
  MessageAction: String
  UserAttributes:
    - AttributeType
  Username: String
  UserPoolId: String
  ValidationData:
    - AttributeType

```

## Properties

`ClientMetadata`

A map of custom key-value pairs that you can provide as input for any custom workflows
that this action triggers. You create custom workflows by assigning AWS Lambda functions
to user pool triggers.

When Amazon Cognito invokes any of these functions, it passes a JSON payload, which the
function receives as input. This payload contains a `clientMetadata`
attribute that provides the data that you assigned to the ClientMetadata parameter in
your request. In your function code, you can process the `clientMetadata`
value to enhance your workflow for your specific needs.

To review the Lambda trigger types that Amazon Cognito invokes at runtime with API requests, see [Connecting API actions to Lambda triggers](https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-working-with-lambda-triggers.html#lambda-triggers-by-event) in the _Amazon Cognito Developer Guide_.

###### Note

When you use the `ClientMetadata` parameter, note that Amazon Cognito won't do the
following:

- Store the `ClientMetadata` value. This data is available only
to AWS Lambda triggers that are assigned to a user pool to support custom
workflows. If your user pool configuration doesn't include triggers, the
`ClientMetadata` parameter serves no purpose.

- Validate the `ClientMetadata` value.

- Encrypt the `ClientMetadata` value. Don't send sensitive
information in this parameter.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DesiredDeliveryMediums`

Specify `EMAIL` if email will be used to send the welcome message. Specify
`SMS` if the phone number will be used. The default value is
`SMS`. You can specify more than one value.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ForceAliasCreation`

This parameter is used only if the `phone_number_verified` or
`email_verified` attribute is set to `True`. Otherwise, it is
ignored.

If this parameter is set to `True` and the phone number or email address
specified in the `UserAttributes` parameter already exists as an alias with a
different user, this request migrates the alias from the previous user to the
newly-created user. The previous user will no longer be able to log in using that
alias.

If this parameter is set to `False`, the API throws an
`AliasExistsException` error if the alias already exists. The default
value is `False`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MessageAction`

Set to `RESEND` to resend the invitation message to a user that already
exists, and to reset the temporary-password duration with a new temporary password. Set
to `SUPPRESS` to suppress sending the message. You can specify only one
value.

_Required_: No

_Type_: String

_Allowed values_: `RESEND | SUPPRESS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserAttributes`

An array of name-value pairs that contain user attributes and attribute values to be
set for the user to be created. You can create a user without specifying any attributes
other than `Username`. However, any attributes that you specify as required
(when creating a user pool or in the **Attributes** tab of
the console) either you should supply (in your call to `AdminCreateUser`) or
the user should supply (when they sign up in response to your welcome message).

For custom attributes, you must prepend the `custom:` prefix to the
attribute name.

To send a message inviting the user to sign up, you must specify the user's email
address or phone number. You can do this in your call to AdminCreateUser or in the
**Users** tab of the Amazon Cognito console for
managing your user pools.

You must also provide an email address or phone number when you expect the user to do
passwordless sign-in with an email or SMS OTP. These attributes must be provided when
passwordless options are the only available, or when you don't submit a
`TemporaryPassword`.

In your call to `AdminCreateUser`, you can set the
`email_verified` attribute to `True`, and you can set the
`phone_number_verified` attribute to `True`.

- **email**: The email address of the user to whom
the message that contains the code and username will be sent. Required if the
`email_verified` attribute is set to `True`, or if
`"EMAIL"` is specified in the `DesiredDeliveryMediums`
parameter.

- **phone\_number**: The phone number of the user to
whom the message that contains the code and username will be sent. Required if
the `phone_number_verified` attribute is set to `True`, or
if `"SMS"` is specified in the `DesiredDeliveryMediums`
parameter.

_Required_: No

_Type_: Array of [AttributeType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpooluser-attributetype.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Username`

The value that you want to set as the username sign-in attribute. The following
conditions apply to the username parameter.

- The username can't be a duplicate of another username in the same user
pool.

- You can't change the value of a username after you create it.

- You can only provide a value if usernames are a valid sign-in attribute for
your user pool. If your user pool only supports phone numbers or email addresses
as sign-in attributes, Amazon Cognito automatically generates a username value. For more
information, see [Customizing sign-in attributes](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-attributes.html#user-pool-settings-aliases).

_Required_: No

_Type_: String

_Pattern_: `[\p{L}\p{M}\p{S}\p{N}\p{P}]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UserPoolId`

The ID of the user pool where you want to create a user.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationData`

Temporary user attributes that contribute to the outcomes of your pre sign-up Lambda
trigger. This set of key-value pairs are for custom validation of information that you
collect from your users but don't need to retain.

Your Lambda function can analyze this additional data and act on it. Your function
can automatically confirm and verify select users or perform external API operations
like logging user attributes and validation data to Amazon CloudWatch Logs.

For more information about the pre sign-up Lambda trigger, see [Pre sign-up Lambda trigger](https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-sign-up.html).

_Required_: No

_Type_: Array of [AttributeType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cognito-userpooluser-attributetype.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the user. For example:
`admin`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Creating a new user in a user pool

The following example creates a new user with the username testuser.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Provision a new user to an Amazon Cognito user pool\n",
    "Resources": {
        "TestUser": {
            "Properties": {
                "DesiredDeliveryMediums": [
                    "SMS"
                ],
                "MessageAction": "SUPPRESS",
                "UserAttributes": [
                    {
                        "Name": "name",
                        "Value": "John"
                    },
                    {
                        "Name": "phone_number",
                        "Value": "+12065551212"
                    },
                    {
                        "Name": "email",
                        "Value": "testuser@example.com"
                    }
                ],
                "Username": "testuser",
                "UserPoolId": "us-west-2_EXAMPLE"
            },
            "Type": "AWS::Cognito::UserPoolUser"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"

Description: |
  Provision a new user to an Amazon Cognito user pool

Resources:
  TestUser:
    Type: AWS::Cognito::UserPoolUser
    Properties:
      DesiredDeliveryMediums:
        - SMS
      MessageAction: SUPPRESS
      UserAttributes:
        - Name: name
          Value: John
        - Name: phone_number
          Value: "+12065551212"
        - Name: email
          Value: testuser@example.com
      Username: testuser
      UserPoolId: us-west-2_EXAMPLE
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Cognito::UserPoolUICustomizationAttachment

AttributeType
