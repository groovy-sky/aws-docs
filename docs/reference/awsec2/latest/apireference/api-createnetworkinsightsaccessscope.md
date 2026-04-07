# CreateNetworkInsightsAccessScope

Creates a Network Access Scope.

AWS Network Access Analyzer enables cloud networking and cloud operations teams
to verify that their networks on AWS conform to their network security and governance
objectives. For more information, see the [AWS Network Access Analyzer Guide](https://docs.aws.amazon.com/vpc/latest/network-access-analyzer).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information,
see [How to ensure idempotency](https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html).

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ExcludePath.N**

The paths to exclude.

Type: Array of [AccessScopePathRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AccessScopePathRequest.html) objects

Required: No

**MatchPath.N**

The paths to match.

Type: Array of [AccessScopePathRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AccessScopePathRequest.html) objects

Required: No

**TagSpecification.N**

The tags to apply.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**networkInsightsAccessScope**

The Network Access Scope.

Type: [NetworkInsightsAccessScope](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInsightsAccessScope.html) object

**networkInsightsAccessScopeContent**

The Network Access Scope content.

Type: [NetworkInsightsAccessScopeContent](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkInsightsAccessScopeContent.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateNetworkInsightsAccessScope)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateNetworkAclEntry

CreateNetworkInsightsPath
