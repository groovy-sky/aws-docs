Menu

- [Aws](namespace-aws.md)
- [Iot](namespace-aws-iot.md)

## IotClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **AWS IoT** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2015-05-28**](api-iot-2015-05-28.md)

  - [AcceptCertificateTransfer](api-iot-2015-05-28-acceptcertificatetransfer.md)
  - [AddThingToBillingGroup](api-iot-2015-05-28-addthingtobillinggroup.md)
  - [AddThingToThingGroup](api-iot-2015-05-28-addthingtothinggroup.md)
  - [AssociateSbomWithPackageVersion](api-iot-2015-05-28-associatesbomwithpackageversion.md)
  - [AssociateTargetsWithJob](api-iot-2015-05-28-associatetargetswithjob.md)
  - [AttachPolicy](api-iot-2015-05-28-attachpolicy.md)
  - [AttachPrincipalPolicy](api-iot-2015-05-28-attachprincipalpolicy.md)
  - [AttachSecurityProfile](api-iot-2015-05-28-attachsecurityprofile.md)
  - [AttachThingPrincipal](api-iot-2015-05-28-attachthingprincipal.md)
  - [CancelAuditMitigationActionsTask](api-iot-2015-05-28-cancelauditmitigationactionstask.md)
  - [CancelAuditTask](api-iot-2015-05-28-cancelaudittask.md)
  - [CancelCertificateTransfer](api-iot-2015-05-28-cancelcertificatetransfer.md)
  - [CancelDetectMitigationActionsTask](api-iot-2015-05-28-canceldetectmitigationactionstask.md)
  - [CancelJob](api-iot-2015-05-28-canceljob.md)
  - [CancelJobExecution](api-iot-2015-05-28-canceljobexecution.md)
  - [ClearDefaultAuthorizer](api-iot-2015-05-28-cleardefaultauthorizer.md)
  - [ConfirmTopicRuleDestination](api-iot-2015-05-28-confirmtopicruledestination.md)
  - [CreateAuditSuppression](api-iot-2015-05-28-createauditsuppression.md)
  - [CreateAuthorizer](api-iot-2015-05-28-createauthorizer.md)
  - [CreateBillingGroup](api-iot-2015-05-28-createbillinggroup.md)
  - [CreateCertificateFromCsr](api-iot-2015-05-28-createcertificatefromcsr.md)
  - [CreateCertificateProvider](api-iot-2015-05-28-createcertificateprovider.md)
  - [CreateCommand](api-iot-2015-05-28-createcommand.md)
  - [CreateCustomMetric](api-iot-2015-05-28-createcustommetric.md)
  - [CreateDimension](api-iot-2015-05-28-createdimension.md)
  - [CreateDomainConfiguration](api-iot-2015-05-28-createdomainconfiguration.md)
  - [CreateDynamicThingGroup](api-iot-2015-05-28-createdynamicthinggroup.md)
  - [CreateFleetMetric](api-iot-2015-05-28-createfleetmetric.md)
  - [CreateJob](api-iot-2015-05-28-createjob.md)
  - [CreateJobTemplate](api-iot-2015-05-28-createjobtemplate.md)
  - [CreateKeysAndCertificate](api-iot-2015-05-28-createkeysandcertificate.md)
  - [CreateMitigationAction](api-iot-2015-05-28-createmitigationaction.md)
  - [CreateOTAUpdate](api-iot-2015-05-28-createotaupdate.md)
  - [CreatePackage](api-iot-2015-05-28-createpackage.md)
  - [CreatePackageVersion](api-iot-2015-05-28-createpackageversion.md)
  - [CreatePolicy](api-iot-2015-05-28-createpolicy.md)
  - [CreatePolicyVersion](api-iot-2015-05-28-createpolicyversion.md)
  - [CreateProvisioningClaim](api-iot-2015-05-28-createprovisioningclaim.md)
  - [CreateProvisioningTemplate](api-iot-2015-05-28-createprovisioningtemplate.md)
  - [CreateProvisioningTemplateVersion](api-iot-2015-05-28-createprovisioningtemplateversion.md)
  - [CreateRoleAlias](api-iot-2015-05-28-createrolealias.md)
  - [CreateScheduledAudit](api-iot-2015-05-28-createscheduledaudit.md)
  - [CreateSecurityProfile](api-iot-2015-05-28-createsecurityprofile.md)
  - [CreateStream](api-iot-2015-05-28-createstream.md)
  - [CreateThing](api-iot-2015-05-28-creatething.md)
  - [CreateThingGroup](api-iot-2015-05-28-createthinggroup.md)
  - [CreateThingType](api-iot-2015-05-28-createthingtype.md)
  - [CreateTopicRule](api-iot-2015-05-28-createtopicrule.md)
  - [CreateTopicRuleDestination](api-iot-2015-05-28-createtopicruledestination.md)
  - [DeleteAccountAuditConfiguration](api-iot-2015-05-28-deleteaccountauditconfiguration.md)
  - [DeleteAuditSuppression](api-iot-2015-05-28-deleteauditsuppression.md)
  - [DeleteAuthorizer](api-iot-2015-05-28-deleteauthorizer.md)
  - [DeleteBillingGroup](api-iot-2015-05-28-deletebillinggroup.md)
  - [DeleteCACertificate](api-iot-2015-05-28-deletecacertificate.md)
  - [DeleteCertificate](api-iot-2015-05-28-deletecertificate.md)
  - [DeleteCertificateProvider](api-iot-2015-05-28-deletecertificateprovider.md)
  - [DeleteCommand](api-iot-2015-05-28-deletecommand.md)
  - [DeleteCommandExecution](api-iot-2015-05-28-deletecommandexecution.md)
  - [DeleteCustomMetric](api-iot-2015-05-28-deletecustommetric.md)
  - [DeleteDimension](api-iot-2015-05-28-deletedimension.md)
  - [DeleteDomainConfiguration](api-iot-2015-05-28-deletedomainconfiguration.md)
  - [DeleteDynamicThingGroup](api-iot-2015-05-28-deletedynamicthinggroup.md)
  - [DeleteFleetMetric](api-iot-2015-05-28-deletefleetmetric.md)
  - [DeleteJob](api-iot-2015-05-28-deletejob.md)
  - [DeleteJobExecution](api-iot-2015-05-28-deletejobexecution.md)
  - [DeleteJobTemplate](api-iot-2015-05-28-deletejobtemplate.md)
  - [DeleteMitigationAction](api-iot-2015-05-28-deletemitigationaction.md)
  - [DeleteOTAUpdate](api-iot-2015-05-28-deleteotaupdate.md)
  - [DeletePackage](api-iot-2015-05-28-deletepackage.md)
  - [DeletePackageVersion](api-iot-2015-05-28-deletepackageversion.md)
  - [DeletePolicy](api-iot-2015-05-28-deletepolicy.md)
  - [DeletePolicyVersion](api-iot-2015-05-28-deletepolicyversion.md)
  - [DeleteProvisioningTemplate](api-iot-2015-05-28-deleteprovisioningtemplate.md)
  - [DeleteProvisioningTemplateVersion](api-iot-2015-05-28-deleteprovisioningtemplateversion.md)
  - [DeleteRegistrationCode](api-iot-2015-05-28-deleteregistrationcode.md)
  - [DeleteRoleAlias](api-iot-2015-05-28-deleterolealias.md)
  - [DeleteScheduledAudit](api-iot-2015-05-28-deletescheduledaudit.md)
  - [DeleteSecurityProfile](api-iot-2015-05-28-deletesecurityprofile.md)
  - [DeleteStream](api-iot-2015-05-28-deletestream.md)
  - [DeleteThing](api-iot-2015-05-28-deletething.md)
  - [DeleteThingGroup](api-iot-2015-05-28-deletethinggroup.md)
  - [DeleteThingType](api-iot-2015-05-28-deletethingtype.md)
  - [DeleteTopicRule](api-iot-2015-05-28-deletetopicrule.md)
  - [DeleteTopicRuleDestination](api-iot-2015-05-28-deletetopicruledestination.md)
  - [DeleteV2LoggingLevel](api-iot-2015-05-28-deletev2logginglevel.md)
  - [DeprecateThingType](api-iot-2015-05-28-deprecatethingtype.md)
  - [DescribeAccountAuditConfiguration](api-iot-2015-05-28-describeaccountauditconfiguration.md)
  - [DescribeAuditFinding](api-iot-2015-05-28-describeauditfinding.md)
  - [DescribeAuditMitigationActionsTask](api-iot-2015-05-28-describeauditmitigationactionstask.md)
  - [DescribeAuditSuppression](api-iot-2015-05-28-describeauditsuppression.md)
  - [DescribeAuditTask](api-iot-2015-05-28-describeaudittask.md)
  - [DescribeAuthorizer](api-iot-2015-05-28-describeauthorizer.md)
  - [DescribeBillingGroup](api-iot-2015-05-28-describebillinggroup.md)
  - [DescribeCACertificate](api-iot-2015-05-28-describecacertificate.md)
  - [DescribeCertificate](api-iot-2015-05-28-describecertificate.md)
  - [DescribeCertificateProvider](api-iot-2015-05-28-describecertificateprovider.md)
  - [DescribeCustomMetric](api-iot-2015-05-28-describecustommetric.md)
  - [DescribeDefaultAuthorizer](api-iot-2015-05-28-describedefaultauthorizer.md)
  - [DescribeDetectMitigationActionsTask](api-iot-2015-05-28-describedetectmitigationactionstask.md)
  - [DescribeDimension](api-iot-2015-05-28-describedimension.md)
  - [DescribeDomainConfiguration](api-iot-2015-05-28-describedomainconfiguration.md)
  - [DescribeEncryptionConfiguration](api-iot-2015-05-28-describeencryptionconfiguration.md)
  - [DescribeEndpoint](api-iot-2015-05-28-describeendpoint.md)
  - [DescribeEventConfigurations](api-iot-2015-05-28-describeeventconfigurations.md)
  - [DescribeFleetMetric](api-iot-2015-05-28-describefleetmetric.md)
  - [DescribeIndex](api-iot-2015-05-28-describeindex.md)
  - [DescribeJob](api-iot-2015-05-28-describejob.md)
  - [DescribeJobExecution](api-iot-2015-05-28-describejobexecution.md)
  - [DescribeJobTemplate](api-iot-2015-05-28-describejobtemplate.md)
  - [DescribeManagedJobTemplate](api-iot-2015-05-28-describemanagedjobtemplate.md)
  - [DescribeMitigationAction](api-iot-2015-05-28-describemitigationaction.md)
  - [DescribeProvisioningTemplate](api-iot-2015-05-28-describeprovisioningtemplate.md)
  - [DescribeProvisioningTemplateVersion](api-iot-2015-05-28-describeprovisioningtemplateversion.md)
  - [DescribeRoleAlias](api-iot-2015-05-28-describerolealias.md)
  - [DescribeScheduledAudit](api-iot-2015-05-28-describescheduledaudit.md)
  - [DescribeSecurityProfile](api-iot-2015-05-28-describesecurityprofile.md)
  - [DescribeStream](api-iot-2015-05-28-describestream.md)
  - [DescribeThing](api-iot-2015-05-28-describething.md)
  - [DescribeThingGroup](api-iot-2015-05-28-describethinggroup.md)
  - [DescribeThingRegistrationTask](api-iot-2015-05-28-describethingregistrationtask.md)
  - [DescribeThingType](api-iot-2015-05-28-describethingtype.md)
  - [DetachPolicy](api-iot-2015-05-28-detachpolicy.md)
  - [DetachPrincipalPolicy](api-iot-2015-05-28-detachprincipalpolicy.md)
  - [DetachSecurityProfile](api-iot-2015-05-28-detachsecurityprofile.md)
  - [DetachThingPrincipal](api-iot-2015-05-28-detachthingprincipal.md)
  - [DisableTopicRule](api-iot-2015-05-28-disabletopicrule.md)
  - [DisassociateSbomFromPackageVersion](api-iot-2015-05-28-disassociatesbomfrompackageversion.md)
  - [EnableTopicRule](api-iot-2015-05-28-enabletopicrule.md)
  - [GetBehaviorModelTrainingSummaries](api-iot-2015-05-28-getbehaviormodeltrainingsummaries.md)
  - [GetBucketsAggregation](api-iot-2015-05-28-getbucketsaggregation.md)
  - [GetCardinality](api-iot-2015-05-28-getcardinality.md)
  - [GetCommandExecution](api-iot-2015-05-28-getcommandexecution.md)
  - [GetDeviceCommand](api-iot-2015-05-28-getdevicecommand.md)
  - [GetEffectivePolicies](api-iot-2015-05-28-geteffectivepolicies.md)
  - [GetIndexingConfiguration](api-iot-2015-05-28-getindexingconfiguration.md)
  - [GetJobDocument](api-iot-2015-05-28-getjobdocument.md)
  - [GetLoggingOptions](api-iot-2015-05-28-getloggingoptions.md)
  - [GetOTAUpdate](api-iot-2015-05-28-getotaupdate.md)
  - [GetPackage](api-iot-2015-05-28-getpackage.md)
  - [GetPackageConfiguration](api-iot-2015-05-28-getpackageconfiguration.md)
  - [GetPackageVersion](api-iot-2015-05-28-getpackageversion.md)
  - [GetPercentiles](api-iot-2015-05-28-getpercentiles.md)
  - [GetPolicy](api-iot-2015-05-28-getpolicy.md)
  - [GetPolicyVersion](api-iot-2015-05-28-getpolicyversion.md)
  - [GetRegistrationCode](api-iot-2015-05-28-getregistrationcode.md)
  - [GetStatistics](api-iot-2015-05-28-getstatistics.md)
  - [GetThingConnectivityData](api-iot-2015-05-28-getthingconnectivitydata.md)
  - [GetTopicRule](api-iot-2015-05-28-gettopicrule.md)
  - [GetTopicRuleDestination](api-iot-2015-05-28-gettopicruledestination.md)
  - [GetV2LoggingOptions](api-iot-2015-05-28-getv2loggingoptions.md)
  - [ListActiveViolations](api-iot-2015-05-28-listactiveviolations.md)
  - [ListAttachedPolicies](api-iot-2015-05-28-listattachedpolicies.md)
  - [ListAuditFindings](api-iot-2015-05-28-listauditfindings.md)
  - [ListAuditMitigationActionsExecutions](api-iot-2015-05-28-listauditmitigationactionsexecutions.md)
  - [ListAuditMitigationActionsTasks](api-iot-2015-05-28-listauditmitigationactionstasks.md)
  - [ListAuditSuppressions](api-iot-2015-05-28-listauditsuppressions.md)
  - [ListAuditTasks](api-iot-2015-05-28-listaudittasks.md)
  - [ListAuthorizers](api-iot-2015-05-28-listauthorizers.md)
  - [ListBillingGroups](api-iot-2015-05-28-listbillinggroups.md)
  - [ListCACertificates](api-iot-2015-05-28-listcacertificates.md)
  - [ListCertificateProviders](api-iot-2015-05-28-listcertificateproviders.md)
  - [ListCertificates](api-iot-2015-05-28-listcertificates.md)
  - [ListCertificatesByCA](api-iot-2015-05-28-listcertificatesbyca.md)
  - [ListCommandExecutions](api-iot-2015-05-28-listcommandexecutions.md)
  - [ListCommands](api-iot-2015-05-28-listcommands.md)
  - [ListCustomMetrics](api-iot-2015-05-28-listcustommetrics.md)
  - [ListDetectMitigationActionsExecutions](api-iot-2015-05-28-listdetectmitigationactionsexecutions.md)
  - [ListDetectMitigationActionsTasks](api-iot-2015-05-28-listdetectmitigationactionstasks.md)
  - [ListDimensions](api-iot-2015-05-28-listdimensions.md)
  - [ListDomainConfigurations](api-iot-2015-05-28-listdomainconfigurations.md)
  - [ListFleetMetrics](api-iot-2015-05-28-listfleetmetrics.md)
  - [ListIndices](api-iot-2015-05-28-listindices.md)
  - [ListJobExecutionsForJob](api-iot-2015-05-28-listjobexecutionsforjob.md)
  - [ListJobExecutionsForThing](api-iot-2015-05-28-listjobexecutionsforthing.md)
  - [ListJobTemplates](api-iot-2015-05-28-listjobtemplates.md)
  - [ListJobs](api-iot-2015-05-28-listjobs.md)
  - [ListManagedJobTemplates](api-iot-2015-05-28-listmanagedjobtemplates.md)
  - [ListMetricValues](api-iot-2015-05-28-listmetricvalues.md)
  - [ListMitigationActions](api-iot-2015-05-28-listmitigationactions.md)
  - [ListOTAUpdates](api-iot-2015-05-28-listotaupdates.md)
  - [ListOutgoingCertificates](api-iot-2015-05-28-listoutgoingcertificates.md)
  - [ListPackageVersions](api-iot-2015-05-28-listpackageversions.md)
  - [ListPackages](api-iot-2015-05-28-listpackages.md)
  - [ListPolicies](api-iot-2015-05-28-listpolicies.md)
  - [ListPolicyPrincipals](api-iot-2015-05-28-listpolicyprincipals.md)
  - [ListPolicyVersions](api-iot-2015-05-28-listpolicyversions.md)
  - [ListPrincipalPolicies](api-iot-2015-05-28-listprincipalpolicies.md)
  - [ListPrincipalThings](api-iot-2015-05-28-listprincipalthings.md)
  - [ListPrincipalThingsV2](api-iot-2015-05-28-listprincipalthingsv2.md)
  - [ListProvisioningTemplateVersions](api-iot-2015-05-28-listprovisioningtemplateversions.md)
  - [ListProvisioningTemplates](api-iot-2015-05-28-listprovisioningtemplates.md)
  - [ListRelatedResourcesForAuditFinding](api-iot-2015-05-28-listrelatedresourcesforauditfinding.md)
  - [ListRoleAliases](api-iot-2015-05-28-listrolealiases.md)
  - [ListSbomValidationResults](api-iot-2015-05-28-listsbomvalidationresults.md)
  - [ListScheduledAudits](api-iot-2015-05-28-listscheduledaudits.md)
  - [ListSecurityProfiles](api-iot-2015-05-28-listsecurityprofiles.md)
  - [ListSecurityProfilesForTarget](api-iot-2015-05-28-listsecurityprofilesfortarget.md)
  - [ListStreams](api-iot-2015-05-28-liststreams.md)
  - [ListTagsForResource](api-iot-2015-05-28-listtagsforresource.md)
  - [ListTargetsForPolicy](api-iot-2015-05-28-listtargetsforpolicy.md)
  - [ListTargetsForSecurityProfile](api-iot-2015-05-28-listtargetsforsecurityprofile.md)
  - [ListThingGroups](api-iot-2015-05-28-listthinggroups.md)
  - [ListThingGroupsForThing](api-iot-2015-05-28-listthinggroupsforthing.md)
  - [ListThingPrincipals](api-iot-2015-05-28-listthingprincipals.md)
  - [ListThingPrincipalsV2](api-iot-2015-05-28-listthingprincipalsv2.md)
  - [ListThingRegistrationTaskReports](api-iot-2015-05-28-listthingregistrationtaskreports.md)
  - [ListThingRegistrationTasks](api-iot-2015-05-28-listthingregistrationtasks.md)
  - [ListThingTypes](api-iot-2015-05-28-listthingtypes.md)
  - [ListThings](api-iot-2015-05-28-listthings.md)
  - [ListThingsInBillingGroup](api-iot-2015-05-28-listthingsinbillinggroup.md)
  - [ListThingsInThingGroup](api-iot-2015-05-28-listthingsinthinggroup.md)
  - [ListTopicRuleDestinations](api-iot-2015-05-28-listtopicruledestinations.md)
  - [ListTopicRules](api-iot-2015-05-28-listtopicrules.md)
  - [ListV2LoggingLevels](api-iot-2015-05-28-listv2logginglevels.md)
  - [ListViolationEvents](api-iot-2015-05-28-listviolationevents.md)
  - [PutVerificationStateOnViolation](api-iot-2015-05-28-putverificationstateonviolation.md)
  - [RegisterCACertificate](api-iot-2015-05-28-registercacertificate.md)
  - [RegisterCertificate](api-iot-2015-05-28-registercertificate.md)
  - [RegisterCertificateWithoutCA](api-iot-2015-05-28-registercertificatewithoutca.md)
  - [RegisterThing](api-iot-2015-05-28-registerthing.md)
  - [RejectCertificateTransfer](api-iot-2015-05-28-rejectcertificatetransfer.md)
  - [RemoveThingFromBillingGroup](api-iot-2015-05-28-removethingfrombillinggroup.md)
  - [RemoveThingFromThingGroup](api-iot-2015-05-28-removethingfromthinggroup.md)
  - [ReplaceTopicRule](api-iot-2015-05-28-replacetopicrule.md)
  - [SearchIndex](api-iot-2015-05-28-searchindex.md)
  - [SetDefaultAuthorizer](api-iot-2015-05-28-setdefaultauthorizer.md)
  - [SetDefaultPolicyVersion](api-iot-2015-05-28-setdefaultpolicyversion.md)
  - [SetLoggingOptions](api-iot-2015-05-28-setloggingoptions.md)
  - [SetV2LoggingLevel](api-iot-2015-05-28-setv2logginglevel.md)
  - [SetV2LoggingOptions](api-iot-2015-05-28-setv2loggingoptions.md)
  - [StartAuditMitigationActionsTask](api-iot-2015-05-28-startauditmitigationactionstask.md)
  - [StartDetectMitigationActionsTask](api-iot-2015-05-28-startdetectmitigationactionstask.md)
  - [StartOnDemandAuditTask](api-iot-2015-05-28-startondemandaudittask.md)
  - [StartThingRegistrationTask](api-iot-2015-05-28-startthingregistrationtask.md)
  - [StopThingRegistrationTask](api-iot-2015-05-28-stopthingregistrationtask.md)
  - [TagResource](api-iot-2015-05-28-tagresource.md)
  - [TestAuthorization](api-iot-2015-05-28-testauthorization.md)
  - [TestInvokeAuthorizer](api-iot-2015-05-28-testinvokeauthorizer.md)
  - [TransferCertificate](api-iot-2015-05-28-transfercertificate.md)
  - [UntagResource](api-iot-2015-05-28-untagresource.md)
  - [UpdateAccountAuditConfiguration](api-iot-2015-05-28-updateaccountauditconfiguration.md)
  - [UpdateAuditSuppression](api-iot-2015-05-28-updateauditsuppression.md)
  - [UpdateAuthorizer](api-iot-2015-05-28-updateauthorizer.md)
  - [UpdateBillingGroup](api-iot-2015-05-28-updatebillinggroup.md)
  - [UpdateCACertificate](api-iot-2015-05-28-updatecacertificate.md)
  - [UpdateCertificate](api-iot-2015-05-28-updatecertificate.md)
  - [UpdateCertificateProvider](api-iot-2015-05-28-updatecertificateprovider.md)
  - [UpdateCommand](api-iot-2015-05-28-updatecommand.md)
  - [UpdateCustomMetric](api-iot-2015-05-28-updatecustommetric.md)
  - [UpdateDimension](api-iot-2015-05-28-updatedimension.md)
  - [UpdateDomainConfiguration](api-iot-2015-05-28-updatedomainconfiguration.md)
  - [UpdateDynamicThingGroup](api-iot-2015-05-28-updatedynamicthinggroup.md)
  - [UpdateEncryptionConfiguration](api-iot-2015-05-28-updateencryptionconfiguration.md)
  - [UpdateEventConfigurations](api-iot-2015-05-28-updateeventconfigurations.md)
  - [UpdateFleetMetric](api-iot-2015-05-28-updatefleetmetric.md)
  - [UpdateIndexingConfiguration](api-iot-2015-05-28-updateindexingconfiguration.md)
  - [UpdateJob](api-iot-2015-05-28-updatejob.md)
  - [UpdateMitigationAction](api-iot-2015-05-28-updatemitigationaction.md)
  - [UpdatePackage](api-iot-2015-05-28-updatepackage.md)
  - [UpdatePackageConfiguration](api-iot-2015-05-28-updatepackageconfiguration.md)
  - [UpdatePackageVersion](api-iot-2015-05-28-updatepackageversion.md)
  - [UpdateProvisioningTemplate](api-iot-2015-05-28-updateprovisioningtemplate.md)
  - [UpdateRoleAlias](api-iot-2015-05-28-updaterolealias.md)
  - [UpdateScheduledAudit](api-iot-2015-05-28-updatescheduledaudit.md)
  - [UpdateSecurityProfile](api-iot-2015-05-28-updatesecurityprofile.md)
  - [UpdateStream](api-iot-2015-05-28-updatestream.md)
  - [UpdateThing](api-iot-2015-05-28-updatething.md)
  - [UpdateThingGroup](api-iot-2015-05-28-updatethinggroup.md)
  - [UpdateThingGroupsForThing](api-iot-2015-05-28-updatethinggroupsforthing.md)
  - [UpdateThingType](api-iot-2015-05-28-updatethingtype.md)
  - [UpdateTopicRuleDestination](api-iot-2015-05-28-updatetopicruledestination.md)
  - [ValidateSecurityProfileBehaviors](api-iot-2015-05-28-validatesecurityprofilebehaviors.md)

### Table of Contents  [header link](class-aws-iot-iotclient-toc.md)

#### Methods  [header link](class-aws-iot-iotclient-toc-methods.md)

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

### Methods  [header link](class-aws-iot-iotclient-methods.md)

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
  - [Methods](class-aws-iot-iotclient-toc-methods.md)
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

[Back To Top](class-aws-iot-iotclient-top.md)
