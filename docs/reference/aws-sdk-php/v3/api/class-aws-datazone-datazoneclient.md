Menu

- [Aws](namespace-aws.md)
- [DataZone](namespace-aws-datazone.md)

## DataZoneClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **Amazon DataZone** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2018-05-10**](api-datazone-2018-05-10.md)

  - [AcceptPredictions](api-datazone-2018-05-10-acceptpredictions.md)
  - [AcceptSubscriptionRequest](api-datazone-2018-05-10-acceptsubscriptionrequest.md)
  - [AddEntityOwner](api-datazone-2018-05-10-addentityowner.md)
  - [AddPolicyGrant](api-datazone-2018-05-10-addpolicygrant.md)
  - [AssociateEnvironmentRole](api-datazone-2018-05-10-associateenvironmentrole.md)
  - [AssociateGovernedTerms](api-datazone-2018-05-10-associategovernedterms.md)
  - [BatchGetAttributesMetadata](api-datazone-2018-05-10-batchgetattributesmetadata.md)
  - [BatchPutAttributesMetadata](api-datazone-2018-05-10-batchputattributesmetadata.md)
  - [CancelMetadataGenerationRun](api-datazone-2018-05-10-cancelmetadatagenerationrun.md)
  - [CancelSubscription](api-datazone-2018-05-10-cancelsubscription.md)
  - [CreateAccountPool](api-datazone-2018-05-10-createaccountpool.md)
  - [CreateAsset](api-datazone-2018-05-10-createasset.md)
  - [CreateAssetFilter](api-datazone-2018-05-10-createassetfilter.md)
  - [CreateAssetRevision](api-datazone-2018-05-10-createassetrevision.md)
  - [CreateAssetType](api-datazone-2018-05-10-createassettype.md)
  - [CreateConnection](api-datazone-2018-05-10-createconnection.md)
  - [CreateDataProduct](api-datazone-2018-05-10-createdataproduct.md)
  - [CreateDataProductRevision](api-datazone-2018-05-10-createdataproductrevision.md)
  - [CreateDataSource](api-datazone-2018-05-10-createdatasource.md)
  - [CreateDomain](api-datazone-2018-05-10-createdomain.md)
  - [CreateDomainUnit](api-datazone-2018-05-10-createdomainunit.md)
  - [CreateEnvironment](api-datazone-2018-05-10-createenvironment.md)
  - [CreateEnvironmentAction](api-datazone-2018-05-10-createenvironmentaction.md)
  - [CreateEnvironmentBlueprint](api-datazone-2018-05-10-createenvironmentblueprint.md)
  - [CreateEnvironmentProfile](api-datazone-2018-05-10-createenvironmentprofile.md)
  - [CreateFormType](api-datazone-2018-05-10-createformtype.md)
  - [CreateGlossary](api-datazone-2018-05-10-createglossary.md)
  - [CreateGlossaryTerm](api-datazone-2018-05-10-createglossaryterm.md)
  - [CreateGroupProfile](api-datazone-2018-05-10-creategroupprofile.md)
  - [CreateListingChangeSet](api-datazone-2018-05-10-createlistingchangeset.md)
  - [CreateProject](api-datazone-2018-05-10-createproject.md)
  - [CreateProjectMembership](api-datazone-2018-05-10-createprojectmembership.md)
  - [CreateProjectProfile](api-datazone-2018-05-10-createprojectprofile.md)
  - [CreateRule](api-datazone-2018-05-10-createrule.md)
  - [CreateSubscriptionGrant](api-datazone-2018-05-10-createsubscriptiongrant.md)
  - [CreateSubscriptionRequest](api-datazone-2018-05-10-createsubscriptionrequest.md)
  - [CreateSubscriptionTarget](api-datazone-2018-05-10-createsubscriptiontarget.md)
  - [CreateUserProfile](api-datazone-2018-05-10-createuserprofile.md)
  - [DeleteAccountPool](api-datazone-2018-05-10-deleteaccountpool.md)
  - [DeleteAsset](api-datazone-2018-05-10-deleteasset.md)
  - [DeleteAssetFilter](api-datazone-2018-05-10-deleteassetfilter.md)
  - [DeleteAssetType](api-datazone-2018-05-10-deleteassettype.md)
  - [DeleteConnection](api-datazone-2018-05-10-deleteconnection.md)
  - [DeleteDataExportConfiguration](api-datazone-2018-05-10-deletedataexportconfiguration.md)
  - [DeleteDataProduct](api-datazone-2018-05-10-deletedataproduct.md)
  - [DeleteDataSource](api-datazone-2018-05-10-deletedatasource.md)
  - [DeleteDomain](api-datazone-2018-05-10-deletedomain.md)
  - [DeleteDomainUnit](api-datazone-2018-05-10-deletedomainunit.md)
  - [DeleteEnvironment](api-datazone-2018-05-10-deleteenvironment.md)
  - [DeleteEnvironmentAction](api-datazone-2018-05-10-deleteenvironmentaction.md)
  - [DeleteEnvironmentBlueprint](api-datazone-2018-05-10-deleteenvironmentblueprint.md)
  - [DeleteEnvironmentBlueprintConfiguration](api-datazone-2018-05-10-deleteenvironmentblueprintconfiguration.md)
  - [DeleteEnvironmentProfile](api-datazone-2018-05-10-deleteenvironmentprofile.md)
  - [DeleteFormType](api-datazone-2018-05-10-deleteformtype.md)
  - [DeleteGlossary](api-datazone-2018-05-10-deleteglossary.md)
  - [DeleteGlossaryTerm](api-datazone-2018-05-10-deleteglossaryterm.md)
  - [DeleteListing](api-datazone-2018-05-10-deletelisting.md)
  - [DeleteProject](api-datazone-2018-05-10-deleteproject.md)
  - [DeleteProjectMembership](api-datazone-2018-05-10-deleteprojectmembership.md)
  - [DeleteProjectProfile](api-datazone-2018-05-10-deleteprojectprofile.md)
  - [DeleteRule](api-datazone-2018-05-10-deleterule.md)
  - [DeleteSubscriptionGrant](api-datazone-2018-05-10-deletesubscriptiongrant.md)
  - [DeleteSubscriptionRequest](api-datazone-2018-05-10-deletesubscriptionrequest.md)
  - [DeleteSubscriptionTarget](api-datazone-2018-05-10-deletesubscriptiontarget.md)
  - [DeleteTimeSeriesDataPoints](api-datazone-2018-05-10-deletetimeseriesdatapoints.md)
  - [DisassociateEnvironmentRole](api-datazone-2018-05-10-disassociateenvironmentrole.md)
  - [DisassociateGovernedTerms](api-datazone-2018-05-10-disassociategovernedterms.md)
  - [GetAccountPool](api-datazone-2018-05-10-getaccountpool.md)
  - [GetAsset](api-datazone-2018-05-10-getasset.md)
  - [GetAssetFilter](api-datazone-2018-05-10-getassetfilter.md)
  - [GetAssetType](api-datazone-2018-05-10-getassettype.md)
  - [GetConnection](api-datazone-2018-05-10-getconnection.md)
  - [GetDataExportConfiguration](api-datazone-2018-05-10-getdataexportconfiguration.md)
  - [GetDataProduct](api-datazone-2018-05-10-getdataproduct.md)
  - [GetDataSource](api-datazone-2018-05-10-getdatasource.md)
  - [GetDataSourceRun](api-datazone-2018-05-10-getdatasourcerun.md)
  - [GetDomain](api-datazone-2018-05-10-getdomain.md)
  - [GetDomainUnit](api-datazone-2018-05-10-getdomainunit.md)
  - [GetEnvironment](api-datazone-2018-05-10-getenvironment.md)
  - [GetEnvironmentAction](api-datazone-2018-05-10-getenvironmentaction.md)
  - [GetEnvironmentBlueprint](api-datazone-2018-05-10-getenvironmentblueprint.md)
  - [GetEnvironmentBlueprintConfiguration](api-datazone-2018-05-10-getenvironmentblueprintconfiguration.md)
  - [GetEnvironmentCredentials](api-datazone-2018-05-10-getenvironmentcredentials.md)
  - [GetEnvironmentProfile](api-datazone-2018-05-10-getenvironmentprofile.md)
  - [GetFormType](api-datazone-2018-05-10-getformtype.md)
  - [GetGlossary](api-datazone-2018-05-10-getglossary.md)
  - [GetGlossaryTerm](api-datazone-2018-05-10-getglossaryterm.md)
  - [GetGroupProfile](api-datazone-2018-05-10-getgroupprofile.md)
  - [GetIamPortalLoginUrl](api-datazone-2018-05-10-getiamportalloginurl.md)
  - [GetJobRun](api-datazone-2018-05-10-getjobrun.md)
  - [GetLineageEvent](api-datazone-2018-05-10-getlineageevent.md)
  - [GetLineageNode](api-datazone-2018-05-10-getlineagenode.md)
  - [GetListing](api-datazone-2018-05-10-getlisting.md)
  - [GetMetadataGenerationRun](api-datazone-2018-05-10-getmetadatagenerationrun.md)
  - [GetProject](api-datazone-2018-05-10-getproject.md)
  - [GetProjectProfile](api-datazone-2018-05-10-getprojectprofile.md)
  - [GetRule](api-datazone-2018-05-10-getrule.md)
  - [GetSubscription](api-datazone-2018-05-10-getsubscription.md)
  - [GetSubscriptionGrant](api-datazone-2018-05-10-getsubscriptiongrant.md)
  - [GetSubscriptionRequestDetails](api-datazone-2018-05-10-getsubscriptionrequestdetails.md)
  - [GetSubscriptionTarget](api-datazone-2018-05-10-getsubscriptiontarget.md)
  - [GetTimeSeriesDataPoint](api-datazone-2018-05-10-gettimeseriesdatapoint.md)
  - [GetUserProfile](api-datazone-2018-05-10-getuserprofile.md)
  - [ListAccountPools](api-datazone-2018-05-10-listaccountpools.md)
  - [ListAccountsInAccountPool](api-datazone-2018-05-10-listaccountsinaccountpool.md)
  - [ListAssetFilters](api-datazone-2018-05-10-listassetfilters.md)
  - [ListAssetRevisions](api-datazone-2018-05-10-listassetrevisions.md)
  - [ListConnections](api-datazone-2018-05-10-listconnections.md)
  - [ListDataProductRevisions](api-datazone-2018-05-10-listdataproductrevisions.md)
  - [ListDataSourceRunActivities](api-datazone-2018-05-10-listdatasourcerunactivities.md)
  - [ListDataSourceRuns](api-datazone-2018-05-10-listdatasourceruns.md)
  - [ListDataSources](api-datazone-2018-05-10-listdatasources.md)
  - [ListDomainUnitsForParent](api-datazone-2018-05-10-listdomainunitsforparent.md)
  - [ListDomains](api-datazone-2018-05-10-listdomains.md)
  - [ListEntityOwners](api-datazone-2018-05-10-listentityowners.md)
  - [ListEnvironmentActions](api-datazone-2018-05-10-listenvironmentactions.md)
  - [ListEnvironmentBlueprintConfigurations](api-datazone-2018-05-10-listenvironmentblueprintconfigurations.md)
  - [ListEnvironmentBlueprints](api-datazone-2018-05-10-listenvironmentblueprints.md)
  - [ListEnvironmentProfiles](api-datazone-2018-05-10-listenvironmentprofiles.md)
  - [ListEnvironments](api-datazone-2018-05-10-listenvironments.md)
  - [ListJobRuns](api-datazone-2018-05-10-listjobruns.md)
  - [ListLineageEvents](api-datazone-2018-05-10-listlineageevents.md)
  - [ListLineageNodeHistory](api-datazone-2018-05-10-listlineagenodehistory.md)
  - [ListMetadataGenerationRuns](api-datazone-2018-05-10-listmetadatagenerationruns.md)
  - [ListNotifications](api-datazone-2018-05-10-listnotifications.md)
  - [ListPolicyGrants](api-datazone-2018-05-10-listpolicygrants.md)
  - [ListProjectMemberships](api-datazone-2018-05-10-listprojectmemberships.md)
  - [ListProjectProfiles](api-datazone-2018-05-10-listprojectprofiles.md)
  - [ListProjects](api-datazone-2018-05-10-listprojects.md)
  - [ListRules](api-datazone-2018-05-10-listrules.md)
  - [ListSubscriptionGrants](api-datazone-2018-05-10-listsubscriptiongrants.md)
  - [ListSubscriptionRequests](api-datazone-2018-05-10-listsubscriptionrequests.md)
  - [ListSubscriptionTargets](api-datazone-2018-05-10-listsubscriptiontargets.md)
  - [ListSubscriptions](api-datazone-2018-05-10-listsubscriptions.md)
  - [ListTagsForResource](api-datazone-2018-05-10-listtagsforresource.md)
  - [ListTimeSeriesDataPoints](api-datazone-2018-05-10-listtimeseriesdatapoints.md)
  - [PostLineageEvent](api-datazone-2018-05-10-postlineageevent.md)
  - [PostTimeSeriesDataPoints](api-datazone-2018-05-10-posttimeseriesdatapoints.md)
  - [PutDataExportConfiguration](api-datazone-2018-05-10-putdataexportconfiguration.md)
  - [PutEnvironmentBlueprintConfiguration](api-datazone-2018-05-10-putenvironmentblueprintconfiguration.md)
  - [QueryGraph](api-datazone-2018-05-10-querygraph.md)
  - [RejectPredictions](api-datazone-2018-05-10-rejectpredictions.md)
  - [RejectSubscriptionRequest](api-datazone-2018-05-10-rejectsubscriptionrequest.md)
  - [RemoveEntityOwner](api-datazone-2018-05-10-removeentityowner.md)
  - [RemovePolicyGrant](api-datazone-2018-05-10-removepolicygrant.md)
  - [RevokeSubscription](api-datazone-2018-05-10-revokesubscription.md)
  - [Search](api-datazone-2018-05-10-search.md)
  - [SearchGroupProfiles](api-datazone-2018-05-10-searchgroupprofiles.md)
  - [SearchListings](api-datazone-2018-05-10-searchlistings.md)
  - [SearchTypes](api-datazone-2018-05-10-searchtypes.md)
  - [SearchUserProfiles](api-datazone-2018-05-10-searchuserprofiles.md)
  - [StartDataSourceRun](api-datazone-2018-05-10-startdatasourcerun.md)
  - [StartMetadataGenerationRun](api-datazone-2018-05-10-startmetadatagenerationrun.md)
  - [TagResource](api-datazone-2018-05-10-tagresource.md)
  - [UntagResource](api-datazone-2018-05-10-untagresource.md)
  - [UpdateAccountPool](api-datazone-2018-05-10-updateaccountpool.md)
  - [UpdateAssetFilter](api-datazone-2018-05-10-updateassetfilter.md)
  - [UpdateConnection](api-datazone-2018-05-10-updateconnection.md)
  - [UpdateDataSource](api-datazone-2018-05-10-updatedatasource.md)
  - [UpdateDomain](api-datazone-2018-05-10-updatedomain.md)
  - [UpdateDomainUnit](api-datazone-2018-05-10-updatedomainunit.md)
  - [UpdateEnvironment](api-datazone-2018-05-10-updateenvironment.md)
  - [UpdateEnvironmentAction](api-datazone-2018-05-10-updateenvironmentaction.md)
  - [UpdateEnvironmentBlueprint](api-datazone-2018-05-10-updateenvironmentblueprint.md)
  - [UpdateEnvironmentProfile](api-datazone-2018-05-10-updateenvironmentprofile.md)
  - [UpdateGlossary](api-datazone-2018-05-10-updateglossary.md)
  - [UpdateGlossaryTerm](api-datazone-2018-05-10-updateglossaryterm.md)
  - [UpdateGroupProfile](api-datazone-2018-05-10-updategroupprofile.md)
  - [UpdateProject](api-datazone-2018-05-10-updateproject.md)
  - [UpdateProjectProfile](api-datazone-2018-05-10-updateprojectprofile.md)
  - [UpdateRootDomainUnitOwner](api-datazone-2018-05-10-updaterootdomainunitowner.md)
  - [UpdateRule](api-datazone-2018-05-10-updaterule.md)
  - [UpdateSubscriptionGrantStatus](api-datazone-2018-05-10-updatesubscriptiongrantstatus.md)
  - [UpdateSubscriptionRequest](api-datazone-2018-05-10-updatesubscriptionrequest.md)
  - [UpdateSubscriptionTarget](api-datazone-2018-05-10-updatesubscriptiontarget.md)
  - [UpdateUserProfile](api-datazone-2018-05-10-updateuserprofile.md)

### Table of Contents  [header link](class-aws-datazone-datazoneclient-toc.md)

#### Methods  [header link](class-aws-datazone-datazoneclient-toc-methods.md)

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

### Methods  [header link](class-aws-datazone-datazoneclient-methods.md)

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
  - [Methods](class-aws-datazone-datazoneclient-toc-methods.md)
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

[Back To Top](class-aws-datazone-datazoneclient-top.md)
