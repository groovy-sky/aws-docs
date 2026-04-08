# Scope

The wrapper that contains the AWS accounts that are in
scope for the assessment.

###### Note

You no longer need to specify which AWS services are in scope when you
create or update an assessment. Audit Manager infers the services in scope by
examining your assessment controls and their data sources, and then mapping this
information to the relevant AWS services.

If an underlying data source changes for your assessment, we automatically update the
services scope as needed to reflect the correct AWS services. This
ensures that your assessment collects accurate and comprehensive evidence about all of
the relevant services in your AWS environment.

## Contents

**awsAccounts**

The AWS accounts that are included in the scope of the assessment.

Type: Array of [AWSAccount](api-awsaccount.md) objects

Array Members: Minimum number of 1 item. Maximum number of 200 items.

Required: No

**awsServices**

_This member has been deprecated._

The AWS services that are included in the scope of the assessment.

###### Important

This API parameter is no longer supported. If you use this parameter to specify one
or more AWS services, Audit Manager ignores this input. Instead, the
value for `awsServices` will show as empty.

Type: Array of [AWSService](api-awsservice.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/scope.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/scope.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/scope.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Role

ServiceMetadata
