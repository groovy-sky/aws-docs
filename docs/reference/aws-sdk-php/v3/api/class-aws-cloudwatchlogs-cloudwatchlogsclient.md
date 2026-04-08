Menu

- [Aws](namespace-aws.md)
- [CloudWatchLogs](namespace-aws-cloudwatchlogs.md)

## CloudWatchLogsClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon CloudWatch Logs** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2014-03-28**](api-logs-2014-03-28.md)

  - [AssociateKmsKey](api-logs-2014-03-28-associatekmskey.md)
  - [AssociateSourceToS3TableIntegration](api-logs-2014-03-28-associatesourcetos3tableintegration.md)
  - [CancelExportTask](api-logs-2014-03-28-cancelexporttask.md)
  - [CancelImportTask](api-logs-2014-03-28-cancelimporttask.md)
  - [CreateDelivery](api-logs-2014-03-28-createdelivery.md)
  - [CreateExportTask](api-logs-2014-03-28-createexporttask.md)
  - [CreateImportTask](api-logs-2014-03-28-createimporttask.md)
  - [CreateLogAnomalyDetector](api-logs-2014-03-28-createloganomalydetector.md)
  - [CreateLogGroup](api-logs-2014-03-28-createloggroup.md)
  - [CreateLogStream](api-logs-2014-03-28-createlogstream.md)
  - [CreateLookupTable](api-logs-2014-03-28-createlookuptable.md)
  - [CreateScheduledQuery](api-logs-2014-03-28-createscheduledquery.md)
  - [DeleteAccountPolicy](api-logs-2014-03-28-deleteaccountpolicy.md)
  - [DeleteDataProtectionPolicy](api-logs-2014-03-28-deletedataprotectionpolicy.md)
  - [DeleteDelivery](api-logs-2014-03-28-deletedelivery.md)
  - [DeleteDeliveryDestination](api-logs-2014-03-28-deletedeliverydestination.md)
  - [DeleteDeliveryDestinationPolicy](api-logs-2014-03-28-deletedeliverydestinationpolicy.md)
  - [DeleteDeliverySource](api-logs-2014-03-28-deletedeliverysource.md)
  - [DeleteDestination](api-logs-2014-03-28-deletedestination.md)
  - [DeleteIndexPolicy](api-logs-2014-03-28-deleteindexpolicy.md)
  - [DeleteIntegration](api-logs-2014-03-28-deleteintegration.md)
  - [DeleteLogAnomalyDetector](api-logs-2014-03-28-deleteloganomalydetector.md)
  - [DeleteLogGroup](api-logs-2014-03-28-deleteloggroup.md)
  - [DeleteLogStream](api-logs-2014-03-28-deletelogstream.md)
  - [DeleteLookupTable](api-logs-2014-03-28-deletelookuptable.md)
  - [DeleteMetricFilter](api-logs-2014-03-28-deletemetricfilter.md)
  - [DeleteQueryDefinition](api-logs-2014-03-28-deletequerydefinition.md)
  - [DeleteResourcePolicy](api-logs-2014-03-28-deleteresourcepolicy.md)
  - [DeleteRetentionPolicy](api-logs-2014-03-28-deleteretentionpolicy.md)
  - [DeleteScheduledQuery](api-logs-2014-03-28-deletescheduledquery.md)
  - [DeleteSubscriptionFilter](api-logs-2014-03-28-deletesubscriptionfilter.md)
  - [DeleteTransformer](api-logs-2014-03-28-deletetransformer.md)
  - [DescribeAccountPolicies](api-logs-2014-03-28-describeaccountpolicies.md)
  - [DescribeConfigurationTemplates](api-logs-2014-03-28-describeconfigurationtemplates.md)
  - [DescribeDeliveries](api-logs-2014-03-28-describedeliveries.md)
  - [DescribeDeliveryDestinations](api-logs-2014-03-28-describedeliverydestinations.md)
  - [DescribeDeliverySources](api-logs-2014-03-28-describedeliverysources.md)
  - [DescribeDestinations](api-logs-2014-03-28-describedestinations.md)
  - [DescribeExportTasks](api-logs-2014-03-28-describeexporttasks.md)
  - [DescribeFieldIndexes](api-logs-2014-03-28-describefieldindexes.md)
  - [DescribeImportTaskBatches](api-logs-2014-03-28-describeimporttaskbatches.md)
  - [DescribeImportTasks](api-logs-2014-03-28-describeimporttasks.md)
  - [DescribeIndexPolicies](api-logs-2014-03-28-describeindexpolicies.md)
  - [DescribeLogGroups](api-logs-2014-03-28-describeloggroups.md)
  - [DescribeLogStreams](api-logs-2014-03-28-describelogstreams.md)
  - [DescribeLookupTables](api-logs-2014-03-28-describelookuptables.md)
  - [DescribeMetricFilters](api-logs-2014-03-28-describemetricfilters.md)
  - [DescribeQueries](api-logs-2014-03-28-describequeries.md)
  - [DescribeQueryDefinitions](api-logs-2014-03-28-describequerydefinitions.md)
  - [DescribeResourcePolicies](api-logs-2014-03-28-describeresourcepolicies.md)
  - [DescribeSubscriptionFilters](api-logs-2014-03-28-describesubscriptionfilters.md)
  - [DisassociateKmsKey](api-logs-2014-03-28-disassociatekmskey.md)
  - [DisassociateSourceFromS3TableIntegration](api-logs-2014-03-28-disassociatesourcefroms3tableintegration.md)
  - [FilterLogEvents](api-logs-2014-03-28-filterlogevents.md)
  - [GetDataProtectionPolicy](api-logs-2014-03-28-getdataprotectionpolicy.md)
  - [GetDelivery](api-logs-2014-03-28-getdelivery.md)
  - [GetDeliveryDestination](api-logs-2014-03-28-getdeliverydestination.md)
  - [GetDeliveryDestinationPolicy](api-logs-2014-03-28-getdeliverydestinationpolicy.md)
  - [GetDeliverySource](api-logs-2014-03-28-getdeliverysource.md)
  - [GetIntegration](api-logs-2014-03-28-getintegration.md)
  - [GetLogAnomalyDetector](api-logs-2014-03-28-getloganomalydetector.md)
  - [GetLogEvents](api-logs-2014-03-28-getlogevents.md)
  - [GetLogFields](api-logs-2014-03-28-getlogfields.md)
  - [GetLogGroupFields](api-logs-2014-03-28-getloggroupfields.md)
  - [GetLogObject](api-logs-2014-03-28-getlogobject.md)
  - [GetLogRecord](api-logs-2014-03-28-getlogrecord.md)
  - [GetLookupTable](api-logs-2014-03-28-getlookuptable.md)
  - [GetQueryResults](api-logs-2014-03-28-getqueryresults.md)
  - [GetScheduledQuery](api-logs-2014-03-28-getscheduledquery.md)
  - [GetScheduledQueryHistory](api-logs-2014-03-28-getscheduledqueryhistory.md)
  - [GetTransformer](api-logs-2014-03-28-gettransformer.md)
  - [ListAggregateLogGroupSummaries](api-logs-2014-03-28-listaggregateloggroupsummaries.md)
  - [ListAnomalies](api-logs-2014-03-28-listanomalies.md)
  - [ListIntegrations](api-logs-2014-03-28-listintegrations.md)
  - [ListLogAnomalyDetectors](api-logs-2014-03-28-listloganomalydetectors.md)
  - [ListLogGroups](api-logs-2014-03-28-listloggroups.md)
  - [ListLogGroupsForQuery](api-logs-2014-03-28-listloggroupsforquery.md)
  - [ListScheduledQueries](api-logs-2014-03-28-listscheduledqueries.md)
  - [ListSourcesForS3TableIntegration](api-logs-2014-03-28-listsourcesfors3tableintegration.md)
  - [ListTagsForResource](api-logs-2014-03-28-listtagsforresource.md)
  - [ListTagsLogGroup](api-logs-2014-03-28-listtagsloggroup.md)
  - [PutAccountPolicy](api-logs-2014-03-28-putaccountpolicy.md)
  - [PutBearerTokenAuthentication](api-logs-2014-03-28-putbearertokenauthentication.md)
  - [PutDataProtectionPolicy](api-logs-2014-03-28-putdataprotectionpolicy.md)
  - [PutDeliveryDestination](api-logs-2014-03-28-putdeliverydestination.md)
  - [PutDeliveryDestinationPolicy](api-logs-2014-03-28-putdeliverydestinationpolicy.md)
  - [PutDeliverySource](api-logs-2014-03-28-putdeliverysource.md)
  - [PutDestination](api-logs-2014-03-28-putdestination.md)
  - [PutDestinationPolicy](api-logs-2014-03-28-putdestinationpolicy.md)
  - [PutIndexPolicy](api-logs-2014-03-28-putindexpolicy.md)
  - [PutIntegration](api-logs-2014-03-28-putintegration.md)
  - [PutLogEvents](api-logs-2014-03-28-putlogevents.md)
  - [PutLogGroupDeletionProtection](api-logs-2014-03-28-putloggroupdeletionprotection.md)
  - [PutMetricFilter](api-logs-2014-03-28-putmetricfilter.md)
  - [PutQueryDefinition](api-logs-2014-03-28-putquerydefinition.md)
  - [PutResourcePolicy](api-logs-2014-03-28-putresourcepolicy.md)
  - [PutRetentionPolicy](api-logs-2014-03-28-putretentionpolicy.md)
  - [PutSubscriptionFilter](api-logs-2014-03-28-putsubscriptionfilter.md)
  - [PutTransformer](api-logs-2014-03-28-puttransformer.md)
  - [StartLiveTail](api-logs-2014-03-28-startlivetail.md)
  - [StartQuery](api-logs-2014-03-28-startquery.md)
  - [StopQuery](api-logs-2014-03-28-stopquery.md)
  - [TagLogGroup](api-logs-2014-03-28-tagloggroup.md)
  - [TagResource](api-logs-2014-03-28-tagresource.md)
  - [TestMetricFilter](api-logs-2014-03-28-testmetricfilter.md)
  - [TestTransformer](api-logs-2014-03-28-testtransformer.md)
  - [UntagLogGroup](api-logs-2014-03-28-untagloggroup.md)
  - [UntagResource](api-logs-2014-03-28-untagresource.md)
  - [UpdateAnomaly](api-logs-2014-03-28-updateanomaly.md)
  - [UpdateDeliveryConfiguration](api-logs-2014-03-28-updatedeliveryconfiguration.md)
  - [UpdateLogAnomalyDetector](api-logs-2014-03-28-updateloganomalydetector.md)
  - [UpdateLookupTable](api-logs-2014-03-28-updatelookuptable.md)
  - [UpdateScheduledQuery](api-logs-2014-03-28-updatescheduledquery.md)

### Table of Contents  [header link](class-aws-cloudwatchlogs-cloudwatchlogsclient-toc.md)

#### Methods  [header link](class-aws-cloudwatchlogs-cloudwatchlogsclient-toc-methods.md)

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
: mixed [startLiveTailCheckingForResults()](class-aws-cloudwatchlogs-cloudwatchlogsclient-method-startlivetailcheckingforresults.md)
: GeneratorHelper method for 'startLiveTail' operation that checks for results.[waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](class-aws-cloudwatchlogs-cloudwatchlogsclient-methods.md)

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

#### startLiveTailCheckingForResults()  [header link](class-aws-cloudwatchlogs-cloudwatchlogsclient-method-startlivetailcheckingforresults.md)

Helper method for 'startLiveTail' operation that checks for results.

`
    public
                    startLiveTailCheckingForResults(array<string|int, mixed> $args) : Generator`

Initiates 'startLiveTail' operation with given arguments, and continuously
checks response stream for session updates or results, yielding each
stream chunk when results are not empty. This method abstracts from users
the need of checking if there are logs entry available to be watched, which means
that users will always get a next item to be iterated when more log entries are
available.

##### Parameters

$args
: array<string\|int, mixed>

Command arguments.

##### Return values

Generator
—

Yields session update or result stream chunks.

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
  - [Methods](class-aws-cloudwatchlogs-cloudwatchlogsclient-toc-methods.md)
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
  - [startLiveTailCheckingForResults()](class-aws-cloudwatchlogs-cloudwatchlogsclient-method-startlivetailcheckingforresults.md)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](class-aws-cloudwatchlogs-cloudwatchlogsclient-top.md)
