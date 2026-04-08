# Parameter

A value such as an Amazon Resource Name (ARN) or an Amazon Simple Notification Service topic entered
in an extension when invoked. Parameter values are specified in an extension association.
For more information about extensions, see [Extending\
workflows](../../../../services/appconfig/latest/userguide/working-with-appconfig-extensions.md) in the _AWS AppConfig User Guide_.

## Contents

**Description**

Information about the parameter.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**Dynamic**

Indicates whether this parameter's value can be supplied at the extension's action point
instead of during extension association. Dynamic parameters can't be marked
`Required`.

Type: Boolean

Required: No

**Required**

A parameter value must be specified in the extension association.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/parameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/parameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/parameter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Monitor

Validator
