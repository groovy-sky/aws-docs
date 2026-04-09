# Connect to third-party services

###### Topics

- [OpenAPI Connector vs. API Connector](#add-connector-third-party-openapi-vs-api)

- [Connect to third-party services and APIs (generic)](connectors-api.md)

- [Connect to services with OpenAPI](connectors-openapi.md)

- [Connect to Salesforce](connectors-salesforce.md)

## OpenAPI Connector vs. API Connector

To send API requests to third-party services from App Studio applications, you must create and configure a connector that the application uses to authenticate with
the service and configure the API calls. App Studio provides both the `API Connector` and `OpenAPI Connector` connector types to accomplish this, which
are described as follows:

- **API Connector:** Used to configure authentication and request information for any type of REST API.

- **OpenAPI Connector:** Used to configure authentication and request information for APIs that have adopted the OpenAPI
Specification (OAS). APIs that adhere to the OAS provide several benefits, including standardization, security, governance, and documentation.

App Studio recommends using the `OpenAPI Connector` for any APIs that adhere to the OAS, and provide an OpenAPI Specification File. For more
information about OpenAPI, see [What is OpenAPI?](https://swagger.io/docs/specification/v3_0/about) in the Swagger documentation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use encrypted data sources with CMKs

Connect to third-party services and APIs (generic)

All content copied from https://docs.aws.amazon.com/.
