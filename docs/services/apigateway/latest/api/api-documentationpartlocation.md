# DocumentationPartLocation

Specifies the target API entity to which the documentation applies.

## Contents

**type**

The type of API entity to which the documentation content applies. Valid values are `API`, `AUTHORIZER`, `MODEL`, `RESOURCE`, `METHOD`, `PATH_PARAMETER`, `QUERY_PARAMETER`, `REQUEST_HEADER`, `REQUEST_BODY`, `RESPONSE`, `RESPONSE_HEADER`, and `RESPONSE_BODY`. Content inheritance does not apply to any entity of the `API`, `AUTHORIZER`, `METHOD`, `MODEL`, `REQUEST_BODY`, or `RESOURCE` type.

Type: String

Valid Values: `API | AUTHORIZER | MODEL | RESOURCE | METHOD | PATH_PARAMETER | QUERY_PARAMETER | REQUEST_HEADER | REQUEST_BODY | RESPONSE | RESPONSE_HEADER | RESPONSE_BODY`

Required: Yes

**method**

The HTTP verb of a method. It is a valid field for the API entity types of `METHOD`, `PATH_PARAMETER`, `QUERY_PARAMETER`, `REQUEST_HEADER`, `REQUEST_BODY`, `RESPONSE`, `RESPONSE_HEADER`, and `RESPONSE_BODY`. The default value is `*` for any method. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other `location` attributes, the child entity's `method` attribute must match that of the parent entity exactly.

Type: String

Required: No

**name**

The name of the targeted API entity. It is a valid and required field for the API entity types of `AUTHORIZER`, `MODEL`, `PATH_PARAMETER`, `QUERY_PARAMETER`, `REQUEST_HEADER`, `REQUEST_BODY` and `RESPONSE_HEADER`. It is an invalid field for any other entity type.

Type: String

Required: No

**path**

The URL path of the target. It is a valid field for the API entity types of `RESOURCE`, `METHOD`, `PATH_PARAMETER`, `QUERY_PARAMETER`, `REQUEST_HEADER`, `REQUEST_BODY`, `RESPONSE`, `RESPONSE_HEADER`, and `RESPONSE_BODY`. The default value is `/` for the root resource. When an applicable child entity inherits the content of another entity of the same type with more general specifications of the other `location` attributes, the child entity's `path` attribute must match that of the parent entity as a prefix.

Type: String

Required: No

**statusCode**

The HTTP status code of a response. It is a valid field for the API entity types of `RESPONSE`, `RESPONSE_HEADER`, and `RESPONSE_BODY`. The default value is `*` for any status code. When an applicable child entity inherits the content of an entity of the same type with more general specifications of the other `location` attributes, the child entity's `statusCode` attribute must match that of the parent entity exactly.

Type: String

Pattern: `^([1-5]\d\d|\*|\s*)$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apigateway-2015-07-09/documentationpartlocation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apigateway-2015-07-09/documentationpartlocation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apigateway-2015-07-09/documentationpartlocation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentationPart

DocumentationVersion

All content copied from https://docs.aws.amazon.com/.
