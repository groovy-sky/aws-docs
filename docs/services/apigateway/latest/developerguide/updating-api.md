# Updates to REST APIs that require redeployment

Maintaining an API amounts to viewing, updating and deleting the existing API setups. You
can maintain an API using the API Gateway console, AWS CLI, CloudFormation, an SDK or the API Gateway REST API. Updating
an API involves modifying certain resource properties or configuration settings of the API.
Resource updates require redeploying the API, where configuration updates do not.

The following table describes API resources that require redeployment of your API when you update them.

ResourceNotes[ApiKey](../api/api-apikey.md)For applicable properties and supported operations, see [apikey:update](../api/api-updateapikey.md). The update requires redeploying the API.[Authorizer](../api/api-authorizer.md)For applicable properties and supported operations, see [authorizer:update](../api/api-updateauthorizer.md). The update requires redeploying the
API.[disableExecuteApiEndpoint](../api/api-updaterestapi.md#apigw-UpdateRestApi-response-disableExecuteApiEndpoint)The update requires modifying any stage on the API such as redeploying the API to a stage.[DocumentationPart](../api/api-documentationpart.md)For applicable properties and supported operations, see [documentationpart:update](../api/api-updatedocumentationpart.md). The update requires redeploying the
API.[DocumentationVersion](../api/api-documentationversion.md)For applicable properties and supported operations, see [documentationversion:update](../api/api-updatedocumentationversion.md). The update requires redeploying
the API.[GatewayResponse](../api/api-gatewayresponse.md)For applicable properties and supported operations, see [gatewayresponse:update](../api/api-updategatewayresponse.md#remarks). The update requires redeploying the
API.[Integration](../api/api-integration.md)

For applicable properties and supported operations, see [integration:update](../api/api-updateintegration.md). The update requires redeploying the
API.

[IntegrationResponse](../api/api-integrationresponse.md)For applicable properties and supported operations, see [integrationresponse:update](../api/api-updateintegrationresponse.md). The update requires redeploying the
API.[Method](../api/api-method.md)For applicable properties and supported operations, see [method:update](../api/api-updatemethod.md). The update requires redeploying the API.[MethodResponse](../api/api-methodresponse.md)For applicable properties and supported operations, see [methodresponse:update](../api/api-updatemethodresponse.md). The update requires redeploying the
API.[Model](../api/api-model.md)For applicable properties and supported operations, see [model:update](../api/api-updatemodel.md). The update requires redeploying the API.[RequestValidator](../api/api-requestvalidator.md)For applicable properties and supported operations, see [requestvalidator:update](../api/api-updaterequestvalidator.md). The update requires redeploying the
API.[Resource](../api/api-resource.md)For applicable properties and supported operations, see [resource:update](../api/api-updateresource.md). The update requires redeploying the
API.[RestApi](../api/api-updaterestapi.md)For applicable properties and supported operations, see [restapi:update](../api/api-updaterestapi.md). The update requires redeploying the
API. This includes modifying resource policies.[VpcLink](../api/api-vpclink.md)For applicable properties and supported operations, see [vpclink:update](../api/api-updatevpclink.md). The update requires redeploying the
API.

The following table describes API configurations that don't require redeployment of your API when you update them.

ConfigurationNotes[Account](../api/api-getaccount.md)

For applicable properties and supported operations, see [account:update](../api/api-updateaccount.md). The update does not require redeploying the
API.

[Deployment](../api/api-deployment.md)For applicable properties and supported operations, see [deployment:update](../api/api-updatedeployment.md). [DomainName](../api/api-domainname.md)For applicable properties and supported operations, see [domainname:update](../api/api-updatedomainname.md). The update does not require redeploying the
API.[BasePathMapping](../api/api-basepathmapping.md)

For applicable properties and supported operations, see [basepathmapping:update](../api/api-updatebasepathmapping.md). The update does not require
redeploying the API.

[IP address type](../api/api-createrestapi.md)

The update does not require redeploying the
API.

[Stage](../api/api-stage.md)

For applicable properties and supported operations, see [stage:update](../api/api-updatestage.md). The update does not require redeploying the
API.

[Usage](../api/api-getusage.md)

For applicable properties and supported operations, see [usage:update](../api/api-updateusage.md). The update does not require redeploying the
API.

[UsagePlan](../api/api-usageplan.md)For applicable properties and supported operations, see [usageplan:update](../api/api-updateusageplan.md). The update does not require redeploying the
API.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Turn off a canary release

Custom domain names

All content copied from https://docs.aws.amazon.com/.
