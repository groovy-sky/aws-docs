# OperatorResponse

Describes whether the resource is managed by a service provider and, if so, describes
the service provider that manages it.

## Contents

**managed**

If `true`, the resource is managed by a service provider.

Type: Boolean

Required: No

**principal**

If `managed` is `true`, then the principal is returned. The
principal is the service provider that manages the resource.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/operatorresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/operatorresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/operatorresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

OperatorRequest

OutpostLag
