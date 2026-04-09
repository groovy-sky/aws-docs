# UserPoolConfig

Describes an Amazon Cognito user pool configuration.

## Contents

**awsRegion**

The AWS Region in which the user pool was created.

Type: String

Required: Yes

**defaultAction**

The action that you want your GraphQL API to take when a request that uses Amazon Cognito user pool authentication doesn't match the Amazon Cognito user pool
configuration.

Type: String

Valid Values: `ALLOW | DENY`

Required: Yes

**userPoolId**

The user pool ID.

Type: String

Required: Yes

**appIdClientRegex**

A regular expression for validating the incoming Amazon Cognito user pool app client
ID. If this value isn't set, no filtering is applied.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appsync-2017-07-25/userpoolconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appsync-2017-07-25/userpoolconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appsync-2017-07-25/userpoolconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Type

Common Parameters

All content copied from https://docs.aws.amazon.com/.
