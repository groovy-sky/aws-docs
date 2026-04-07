Menu

- [Aws](namespace-aws.md)
- [IoTSiteWise](namespace-aws-iotsitewise.md)

## IoTSiteWiseClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **AWS IoT SiteWise** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2019-12-02**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html)

  - [AssociateAssets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#associateassets)
  - [AssociateTimeSeriesToAssetProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#associatetimeseriestoassetproperty)
  - [BatchAssociateProjectAssets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#batchassociateprojectassets)
  - [BatchDisassociateProjectAssets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#batchdisassociateprojectassets)
  - [BatchGetAssetPropertyAggregates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#batchgetassetpropertyaggregates)
  - [BatchGetAssetPropertyValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#batchgetassetpropertyvalue)
  - [BatchGetAssetPropertyValueHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#batchgetassetpropertyvaluehistory)
  - [BatchPutAssetPropertyValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#batchputassetpropertyvalue)
  - [CreateAccessPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createaccesspolicy)
  - [CreateAsset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createasset)
  - [CreateAssetModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createassetmodel)
  - [CreateAssetModelCompositeModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createassetmodelcompositemodel)
  - [CreateBulkImportJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createbulkimportjob)
  - [CreateComputationModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createcomputationmodel)
  - [CreateDashboard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createdashboard)
  - [CreateDataset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createdataset)
  - [CreateGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#creategateway)
  - [CreatePortal](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createportal)
  - [CreateProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#createproject)
  - [DeleteAccessPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deleteaccesspolicy)
  - [DeleteAsset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deleteasset)
  - [DeleteAssetModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deleteassetmodel)
  - [DeleteAssetModelCompositeModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deleteassetmodelcompositemodel)
  - [DeleteAssetModelInterfaceRelationship](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deleteassetmodelinterfacerelationship)
  - [DeleteComputationModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deletecomputationmodel)
  - [DeleteDashboard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deletedashboard)
  - [DeleteDataset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deletedataset)
  - [DeleteGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deletegateway)
  - [DeletePortal](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deleteportal)
  - [DeleteProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deleteproject)
  - [DeleteTimeSeries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#deletetimeseries)
  - [DescribeAccessPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeaccesspolicy)
  - [DescribeAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeaction)
  - [DescribeAsset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeasset)
  - [DescribeAssetCompositeModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeassetcompositemodel)
  - [DescribeAssetModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeassetmodel)
  - [DescribeAssetModelCompositeModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeassetmodelcompositemodel)
  - [DescribeAssetModelInterfaceRelationship](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeassetmodelinterfacerelationship)
  - [DescribeAssetProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeassetproperty)
  - [DescribeBulkImportJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describebulkimportjob)
  - [DescribeComputationModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describecomputationmodel)
  - [DescribeComputationModelExecutionSummary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describecomputationmodelexecutionsummary)
  - [DescribeDashboard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describedashboard)
  - [DescribeDataset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describedataset)
  - [DescribeDefaultEncryptionConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describedefaultencryptionconfiguration)
  - [DescribeExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeexecution)
  - [DescribeGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describegateway)
  - [DescribeGatewayCapabilityConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describegatewaycapabilityconfiguration)
  - [DescribeLoggingOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeloggingoptions)
  - [DescribePortal](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeportal)
  - [DescribeProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describeproject)
  - [DescribeStorageConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describestorageconfiguration)
  - [DescribeTimeSeries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#describetimeseries)
  - [DisassociateAssets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#disassociateassets)
  - [DisassociateTimeSeriesFromAssetProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#disassociatetimeseriesfromassetproperty)
  - [ExecuteAction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#executeaction)
  - [ExecuteQuery](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#executequery)
  - [GetAssetPropertyAggregates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#getassetpropertyaggregates)
  - [GetAssetPropertyValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#getassetpropertyvalue)
  - [GetAssetPropertyValueHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#getassetpropertyvaluehistory)
  - [GetInterpolatedAssetPropertyValues](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#getinterpolatedassetpropertyvalues)
  - [InvokeAssistant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#invokeassistant)
  - [ListAccessPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listaccesspolicies)
  - [ListActions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listactions)
  - [ListAssetModelCompositeModels](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listassetmodelcompositemodels)
  - [ListAssetModelProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listassetmodelproperties)
  - [ListAssetModels](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listassetmodels)
  - [ListAssetProperties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listassetproperties)
  - [ListAssetRelationships](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listassetrelationships)
  - [ListAssets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listassets)
  - [ListAssociatedAssets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listassociatedassets)
  - [ListBulkImportJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listbulkimportjobs)
  - [ListCompositionRelationships](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listcompositionrelationships)
  - [ListComputationModelDataBindingUsages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listcomputationmodeldatabindingusages)
  - [ListComputationModelResolveToResources](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listcomputationmodelresolvetoresources)
  - [ListComputationModels](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listcomputationmodels)
  - [ListDashboards](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listdashboards)
  - [ListDatasets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listdatasets)
  - [ListExecutions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listexecutions)
  - [ListGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listgateways)
  - [ListInterfaceRelationships](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listinterfacerelationships)
  - [ListPortals](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listportals)
  - [ListProjectAssets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listprojectassets)
  - [ListProjects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listprojects)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listtagsforresource)
  - [ListTimeSeries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#listtimeseries)
  - [PutAssetModelInterfaceRelationship](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#putassetmodelinterfacerelationship)
  - [PutDefaultEncryptionConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#putdefaultencryptionconfiguration)
  - [PutLoggingOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#putloggingoptions)
  - [PutStorageConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#putstorageconfiguration)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#untagresource)
  - [UpdateAccessPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updateaccesspolicy)
  - [UpdateAsset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updateasset)
  - [UpdateAssetModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updateassetmodel)
  - [UpdateAssetModelCompositeModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updateassetmodelcompositemodel)
  - [UpdateAssetProperty](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updateassetproperty)
  - [UpdateComputationModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updatecomputationmodel)
  - [UpdateDashboard](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updatedashboard)
  - [UpdateDataset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updatedataset)
  - [UpdateGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updategateway)
  - [UpdateGatewayCapabilityConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updategatewaycapabilityconfiguration)
  - [UpdatePortal](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updateportal)
  - [UpdateProject](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-iotsitewise-2019-12-02.html#updateproject)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.IoTSiteWise.IoTSiteWiseClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.IoTSiteWise.IoTSiteWiseClient.html\#toc-methods)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-awsclient.md#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-awsclient.md#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-awsclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.IoTSiteWise.IoTSiteWiseClient.html\#methods)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-awsclient.md\#method___construct)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

- api\_provider: (callable) An optional PHP callable that accepts a
type, service, and version argument, and returns an array of
corresponding configuration data. The type value can be one of api,
waiter, or paginator.
- credentials:
(Aws\\Credentials\\CredentialsInterface\|array\|bool\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\Credentials\\CredentialsInterface object, an associative array of
"key", "secret", and an optional "token" key, `false` to use null
credentials, or a callable credentials provider used to create
credentials or return null. See Aws\\Credentials\\CredentialProvider for
a list of built-in credentials providers. If no credentials are
provided, the SDK will attempt to load them from the environment.
- token:
(Aws\\Token\\TokenInterface\|array\|bool\|callable) Specifies
the token used to authorize requests. Provide an
Aws\\Token\\TokenInterface object, an associative array of
"token" and an optional "expires" key, `false` to use no
token, or a callable token provider used to create a
token or return null. See Aws\\Token\\TokenProvider for
a list of built-in token providers. If no token is
provided, the SDK will attempt to load one from the environment.
- csm:
(Aws\\ClientSideMonitoring\\ConfigurationInterface\|array\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\ClientSideMonitoring\\ConfigurationInterface object, a callable
configuration provider used to create client-side monitoring configuration,
`false` to disable csm, or an associative array with the following keys:
enabled: (bool) Set to true to enable client-side monitoring, defaults
to false; host: (string) the host location to send monitoring events to,
defaults to 127.0.0.1; port: (int) The port used for the host connection,
defaults to 31000; client\_id: (string) An identifier for this project
- debug: (bool\|array) Set to true to display debug information when
sending requests. Alternatively, you can provide an associative array
with the following keys: logfn: (callable) Function that is invoked
with log messages; stream\_size: (int) When the size of a stream is
greater than this number, the stream data will not be logged (set to
"0" to not log any stream data); scrub\_auth: (bool) Set to false to
disable the scrubbing of auth data from the logged messages; http:
(bool) Set to false to disable the "debug" feature of lower level HTTP
adapters (e.g., verbose curl output).
- stats: (bool\|array) Set to true to gather transfer statistics on
requests sent. Alternatively, you can provide an associative array with
the following keys: retries: (bool) Set to false to disable reporting
on retries attempted; http: (bool) Set to true to enable collecting
statistics from lower level HTTP adapters (e.g., values returned in
GuzzleHttp\\TransferStats). HTTP handlers must support an
`http_stats_receiver` option for this to have an effect; timer: (bool)
Set to true to enable a command timer that reports the total wall clock
time spent on an operation in seconds.
- disable\_host\_prefix\_injection: (bool) Set to true to disable host prefix
injection logic for services that use it. This disables the entire
prefix injection, including the portions supplied by user-defined
parameters. Setting this flag will have no effect on services that do
not use host prefix injection.
- endpoint: (string) The full URI of the webservice. This is only
required when connecting to a custom endpoint (e.g., a local version
of S3).
- endpoint\_discovery: (Aws\\EndpointDiscovery\\ConfigurationInterface,
Aws\\CacheInterface, array, callable) Settings for endpoint discovery.
Provide an instance of Aws\\EndpointDiscovery\\ConfigurationInterface,
an instance Aws\\CacheInterface, a callable that provides a promise for
a Configuration object, or an associative array with the following
keys: enabled: (bool) Set to true to enable endpoint discovery, false
to explicitly disable it, defaults to false; cache\_limit: (int) The
maximum number of keys in the endpoints cache, defaults to 1000.
- endpoint\_provider: (callable) An optional PHP callable that
accepts a hash of options including a "service" and "region" key and
returns NULL or a hash of endpoint data, of which the "endpoint" key
is required. See Aws\\Endpoint\\EndpointProvider for a list of built-in
providers.
- handler: (callable) A handler that accepts a command object,
request object and returns a promise that is fulfilled with an
Aws\\ResultInterface object or rejected with an
Aws\\Exception\\AwsException. A handler does not accept a next handler
as it is terminal and expected to fulfill a command. If no handler is
provided, a default Guzzle handler will be utilized.
- http: (array, default=array(0)) Set to an array of SDK request
options to apply to each request (e.g., proxy, verify, etc.).
- http\_handler: (callable) An HTTP handler is a function that
accepts a PSR-7 request object and returns a promise that is fulfilled
with a PSR-7 response object or rejected with an array of exception
data. NOTE: This option supersedes any provided "handler" option.
- idempotency\_auto\_fill: (bool\|callable) Set to false to disable SDK to
populate parameters that enabled 'idempotencyToken' trait with a random
UUID v4 value on your behalf. Using default value 'true' still allows
parameter value to be overwritten when provided. Note: auto-fill only
works when cryptographically secure random bytes generator functions
(random\_bytes, openssl\_random\_pseudo\_bytes or mcrypt\_create\_iv) can be
found. You may also provide a callable source of random bytes.
- profile: (string) Allows you to specify which profile to use when
credentials are created from the AWS credentials file in your HOME
directory. This setting overrides the AWS\_PROFILE environment
variable. Note: Specifying "profile" will cause the "credentials" key
to be ignored.
- region: (string, required) Region to connect to. See
http://docs.aws.amazon.com/general/latest/gr/rande.html for a list of
available regions.
- retries: (int, Aws\\Retry\\ConfigurationInterface, Aws\\CacheInterface,
array, callable) Configures the retry mode and maximum number of
allowed retries for a client (pass 0 to disable retries). Provide an
integer for 'legacy' mode with the specified number of retries.
Otherwise provide an instance of Aws\\Retry\\ConfigurationInterface, an
instance of Aws\\CacheInterface, a callable function, or an array with
the following keys: mode: (string) Set to 'legacy', 'standard' (uses
retry quota management), or 'adapative' (an experimental mode that adds
client-side rate limiting to standard mode); max\_attempts (int) The
maximum number of attempts for a given request.
- scheme: (string, default=string(5) "https") URI scheme to use when
connecting connect. The SDK will utilize "https" endpoints (i.e.,
utilize SSL/TLS connections) by default. You can attempt to connect to
a service over an unencrypted "http" endpoint by setting `scheme` to
"http".
- signature\_provider: (callable) A callable that accepts a signature
version name (e.g., "v4"), a service name, and region, and
returns a SignatureInterface object or null. This provider is used to
create signers utilized by the client. See
Aws\\Signature\\SignatureProvider for a list of built-in providers
- signature\_version: (string) A string representing a custom
signature version to use with a service (e.g., v4). Note that
per/operation signature version MAY override this requested signature
version.
- use\_aws\_shared\_config\_files: (bool, default=bool(true)) Set to false to
disable checking for shared config file in '~/.aws/config' and
'~/.aws/credentials'. This will override the AWS\_CONFIG\_FILE
environment variable.
- validate: (bool, default=bool(true)) Set to false to disable
client-side parameter validation.
- version: (string, required) The version of the webservice to
utilize (e.g., 2006-03-01).
- account\_id\_endpoint\_mode: (string, default(preferred)) this option
decides whether credentials should resolve an accountId value,
which is going to be used as part of the endpoint resolution.
The valid values for this option are:
  - preferred: when this value is set then, a warning is logged when
    accountId is empty in the resolved identity.
  - required: when this value is set then, an exception is thrown when
    accountId is empty in the resolved identity.
  - disabled: when this value is set then, the validation for if accountId
    was resolved or not, is ignored.
- ua\_append: (string, array) To pass custom user agent parameters.
- app\_id: (string) an optional application specific identifier that can be set.
When set it will be appended to the User-Agent header of every request
in the form of App/{AppId}. This variable is sourced from environment
variable AWS\_SDK\_UA\_APP\_ID or the shared config profile attribute sdk\_ua\_app\_id.
See https://docs.aws.amazon.com/sdkref/latest/guide/settings-reference.html for
more information on environment variables and shared config settings.

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

##### Tags  [header link](class-aws-awsclient.md\#method___construct\#tags)

throwsInvalidArgumentException

if any required options are missing or
the service is not supported.

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### factory()  [header link](class-aws-awsclient.md\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient.md\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-awsclient.md\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient.md\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient.md\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclient.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-awsclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient.md\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient.md\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-awsclient.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient.md\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](class-aws-awsclient.md\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.IoTSiteWise.IoTSiteWiseClient.html#toc-methods)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-awsclient.md#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-awsclient.md#method_getArguments)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
  - [getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
  - [getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
  - [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
  - [getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-awsclient.md#method_getRegion)
  - [getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.IoTSiteWise.IoTSiteWiseClient.html#top)
