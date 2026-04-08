# Amazon Route 53 API actions by function

This topic lists all Route 53 and Route 53 Resolver API actions in groups by the function they perform.

###### Topics

- [Types of function](#API-actions-by-function-types)

- [Actions by function](#API-actions-by-function-actions)

## Types of function

**[DNS](#API-actions-by-function-dns)**

- [Public and private hosted zones](#actions-by-function-public-private-hosted-zones)

- [Public hosted zones](#actions-by-function-public-hosted-zones)

- [Private hosted zones](#actions-by-function-private-hosted-zones)

- [Resource record sets](#actions-by-function-resource-record-sets)

- [Public DNS query logs](#actions-by-function-public-dns-query-logs)

- [Reusable delegation sets](#actions-by-function-reusable-delegation-sets)

**[DNS—Traffic flow](#API-actions-by-function-dns-traffic-flow)**

- [Traffic policies](#actions-by-function-traffic-policies)

- [Traffic policy instances](#actions-by-function-traffic-policy-instances)

**[Domain registration](#API-actions-by-function-domain-registration)**

- [Register, renew, and transfer domains](#actions-by-function-register-renew-transfer-domains)

- [Transfer domains between AWS accounts](#actions-by-function-transfer-domains-between-accounts)

- [Get domain information](#actions-by-function-get-domain-info)

- [Change domain settings](#actions-by-function-change-domain-settings)

**[DNS—DNSSEC](#API-actions-by-function-dnssec)**

- [DNSSEC signing](#actions-by-function-dnssec-signing)

- [DNSSEC validation](#actions-by-function-dnssec-validation)

**[DNS—IP-based routing](#API-actions-by-function-ip-routing)****[Health checking](#API-actions-by-function-health-checking)**

- [Health checks](#actions-by-function-health-checks)

- [Health checker IP ranges](#actions-by-function-health-checks-ip-ranges)

**[Limits (quotas) for accounts, hosted zones, and reusable delegation sets](#API-actions-by-function-limits)****[Route 53 profiles](#API-actions-by-function-profiles)**

- [Route 53 profiles](#actions-by-function-profiles-managing)

- [Profile VPC associations](#actions-by-function-profiles-vpc-associations)

- [Profile resource associations](#actions-by-function-profiles-resource-associations)

**[Route 53 Resolver](#API-actions-by-function-resolver)**

- [Route 53 Resolver endpoints](#actions-by-function-resolver-endpoints)

- [Private DNS query logs](#actions-by-function-resolver-query-logs)

- [Route 53 Resolver rules](#actions-by-function-resolver-rules)

- [Route 53 Resolver DNS Firewall](#actions-by-function-resolver-dns-firewall)

- [Amazon Route 53 Resolver on Outposts](#actions-by-function-outpost-resolver)

- [Resolver configuration](#actions-by-function-resolver-configuration)

**[Tags](#API-actions-by-function-tags)**

- [Tags for hosted zones and health checks](#actions-by-function-tags-for-hosted-zones)

- [Tags for domains](#actions-by-function-tags-for-domains)

- [Tags for Route 53 Resolver](#actions-by-function-resolver-tags)

## Actions by function

[DNS](#API-actions-by-function-dns) \|
[DNS—Traffic flow](#API-actions-by-function-dns-traffic-flow) \|
[Domain registration](#API-actions-by-function-domain-registration) \|
[DNS—DNSSEC](#API-actions-by-function-dnssec) \|
[DNS—IP-based routing](#API-actions-by-function-ip-routing) \|
[Health checking](#API-actions-by-function-health-checking) \|
[Limits (quotas) for accounts, hosted zones, and reusable delegation sets](#API-actions-by-function-limits) \|
[Route 53 profiles](#API-actions-by-function-profiles) \|
[Route 53 Resolver](#API-actions-by-function-resolver) \|
[Tags](#API-actions-by-function-tags)

**DNS**

Public and private hosted zones

- [CreateHostedZone](api-createhostedzone.md)

- [DeleteHostedZone](api-deletehostedzone.md)

- [GetHostedZone](api-gethostedzone.md)

- [GetHostedZoneCount](api-gethostedzonecount.md)

- [ListHostedZones](api-listhostedzones.md)

- [ListHostedZonesByName](api-listhostedzonesbyname.md)

- [UpdateHostedZoneComment](api-updatehostedzonecomment.md)

Public hosted zones

- [TestDNSAnswer](api-testdnsanswer.md)

Private hosted zones

- [AssociateVPCWithHostedZone](api-associatevpcwithhostedzone.md)

- [DisassociateVPCFromHostedZone](api-disassociatevpcfromhostedzone.md)

- [CreateVPCAssociationAuthorization](api-createvpcassociationauthorization.md)

- [DeleteVPCAssociationAuthorization](api-deletevpcassociationauthorization.md)

- [ListHostedZonesByVPC](api-listhostedzonesbyvpc.md)

- [ListVPCAssociationAuthorizations](api-listvpcassociationauthorizations.md)

Resource record sets

- [ChangeResourceRecordSets](api-changeresourcerecordsets.md)

- [GetChange](api-getchange.md)

- [ListResourceRecordSets](api-listresourcerecordsets.md)

- [GetGeoLocation](api-getgeolocation.md)

- [ListGeoLocations](api-listgeolocations.md)

Public DNS query logs

- [CreateQueryLoggingConfig](api-createqueryloggingconfig.md)

- [DeleteQueryLoggingConfig](api-deletequeryloggingconfig.md)

- [GetQueryLoggingConfig](api-getqueryloggingconfig.md)

- [ListQueryLoggingConfigs](api-listqueryloggingconfigs.md)

Reusable delegation sets

- [CreateReusableDelegationSet](api-createreusabledelegationset.md)

- [DeleteReusableDelegationSet](api-deletereusabledelegationset.md)

- [GetReusableDelegationSet](api-getreusabledelegationset.md)

- [ListReusableDelegationSets](api-listreusabledelegationsets.md)

**DNS—Traffic flow**

Traffic policies

- [CreateTrafficPolicy](api-createtrafficpolicy.md)

- [CreateTrafficPolicyVersion](api-createtrafficpolicyversion.md)

- [DeleteTrafficPolicy](api-deletetrafficpolicy.md)

- [GetTrafficPolicy](api-gettrafficpolicy.md)

- [ListTrafficPolicies](api-listtrafficpolicies.md)

- [ListTrafficPolicyVersions](api-listtrafficpolicyversions.md)

- [UpdateTrafficPolicyComment](api-updatetrafficpolicycomment.md)

Traffic policy instances

- [CreateTrafficPolicyInstance](api-createtrafficpolicyinstance.md)

- [DeleteTrafficPolicyInstance](api-deletetrafficpolicyinstance.md)

- [GetTrafficPolicyInstance](api-gettrafficpolicyinstance.md)

- [GetTrafficPolicyInstanceCount](api-gettrafficpolicyinstancecount.md)

- [ListTrafficPolicyInstances](api-listtrafficpolicyinstances.md)

- [ListTrafficPolicyInstancesByHostedZone](api-listtrafficpolicyinstancesbyhostedzone.md)

- [ListTrafficPolicyInstancesByPolicy](api-listtrafficpolicyinstancesbypolicy.md)

- [UpdateTrafficPolicyInstance](api-updatetrafficpolicyinstance.md)

**Domain registration**

Register, renew, and transfer domains

- [RegisterDomain](api-domains-registerdomain.md)

- [RenewDomain](api-domains-renewdomain.md)

- [ResendContactReachabilityEmail](api-domains-resendcontactreachabilityemail.md)

- [RetrieveDomainAuthCode](api-domains-retrievedomainauthcode.md)

- [TransferDomain](api-domains-transferdomain.md)

- [DeleteDomain](api-domains-deletedomain.md)

- [PushDomain](api-domains-pushdomain.md)

- [ResendOperationAuthorization](api-domains-resendoperationauthorization.md)

Transfer domains between AWS accounts

- [AcceptDomainTransferFromAnotherAwsAccount](api-domains-acceptdomaintransferfromanotherawsaccount.md)

- [CancelDomainTransferToAnotherAwsAccount](api-domains-canceldomaintransfertoanotherawsaccount.md)

- [RejectDomainTransferFromAnotherAwsAccount](api-domains-rejectdomaintransferfromanotherawsaccount.md)

- [TransferDomainToAnotherAwsAccount](api-domains-transferdomaintoanotherawsaccount.md)

Get domain information

- [CheckDomainAvailability](api-domains-checkdomainavailability.md)

- [CheckDomainTransferability](api-domains-checkdomaintransferability.md)

- [GetContactReachabilityStatus](api-domains-getcontactreachabilitystatus.md)

- [GetDomainDetail](api-domains-getdomaindetail.md)

- [GetDomainSuggestions](api-domains-getdomainsuggestions.md)

- [GetOperationDetail](api-domains-getoperationdetail.md)

- [ListDomains](api-domains-listdomains.md)

- [ListOperations](api-domains-listoperations.md)

- [ListPrices](api-domains-listprices.md)

- [ViewBilling](api-domains-viewbilling.md)

Change domain settings

- [DisableDomainAutoRenew](api-domains-disabledomainautorenew.md)

- [DisableDomainTransferLock](api-domains-disabledomaintransferlock.md)

- [EnableDomainAutoRenew](api-domains-enabledomainautorenew.md)

- [EnableDomainTransferLock](api-domains-enabledomaintransferlock.md)

- [UpdateDomainContact](api-domains-updatedomaincontact.md)

- [UpdateDomainContactPrivacy](api-domains-updatedomaincontactprivacy.md)

- [UpdateDomainNameservers](api-domains-updatedomainnameservers.md)

**DNS—DNSSEC**

DNSSEC signing

- [ActivateKeySigningKey](api-activatekeysigningkey.md)

- [CreateKeySigningKey](api-createkeysigningkey.md)

- [DeactivateKeySigningKey](api-deactivatekeysigningkey.md)

- [GetDNSSEC](api-getdnssec.md)

- [AssociateDelegationSignerToDomain](api-domains-associatedelegationsignertodomain.md)

- [DisassociateDelegationSignerFromDomain](api-domains-disassociatedelegationsignerfromdomain.md)

DNSSEC validation

- [GetResolverDnssecConfig](api-route53resolver-getresolverdnssecconfig.md)

- [ListResolverDnssecConfigs](api-route53resolver-listresolverdnssecconfigs.md)

- [UpdateResolverDnssecConfig](api-route53resolver-updateresolverdnssecconfig.md)

**DNS—IP-based routing**

- [ChangeCidrCollection](api-changecidrcollection.md)

- [CreateCidrCollection](api-createcidrcollection.md)

- [DeleteCidrCollection](api-deletecidrcollection.md)

- [ListCidrBlocks](api-listcidrblocks.md)

- [ListCidrCollections](api-listcidrcollections.md)

- [ListCidrLocations](api-listcidrlocations.md)

**Health checking**

Health checks

- [CreateHealthCheck](api-createhealthcheck.md)

- [DeleteHealthCheck](api-deletehealthcheck.md)

- [GetHealthCheck](api-gethealthcheck.md)

- [GetHealthCheckCount](api-gethealthcheckcount.md)

- [GetHealthCheckLastFailureReason](api-gethealthchecklastfailurereason.md)

- [GetHealthCheckStatus](api-gethealthcheckstatus.md)

- [ListHealthChecks](api-listhealthchecks.md)

- [UpdateHealthCheck](api-updatehealthcheck.md)

Health checker IP ranges

- [GetCheckerIpRanges](api-getcheckeripranges.md)

**Limits (quotas) for**
**accounts, hosted zones, and reusable delegation sets**

- [GetAccountLimit](api-getaccountlimit.md)

- [GetHostedZoneLimit](api-gethostedzonelimit.md)

- [GetReusableDelegationSetLimit](api-getreusabledelegationsetlimit.md)

**Route 53 profiles**

Route 53 profiles

- [CreateProfile](api-route53profiles-createprofile.md)

- [DeleteProfile](api-route53profiles-deleteprofile.md)

- [GetProfile](api-route53profiles-getprofile.md)

- [ListProfiles](api-route53profiles-listprofiles.md)

Profile VPC associations

- [AssociateProfile](api-route53profiles-associateprofile.md)

- [DisassociateProfile](api-route53profiles-disassociateprofile.md)

- [GetProfileAssociation](api-route53profiles-getprofileassociation.md)

- [ListProfileAssociations](api-route53profiles-listprofileassociations.md)

Profile resource associations

- [AssociateResourceToProfile](api-route53profiles-associateresourcetoprofile.md)

- [DisassociateResourceFromProfile](api-route53profiles-disassociateresourcefromprofile.md)

- [GetProfileResourceAssociation](api-route53profiles-getprofileresourceassociation.md)

- [ListProfileResourceAssociations](api-route53profiles-listprofileresourceassociations.md)

- [UpdateProfileResourceAssociation](api-route53profiles-updateprofileresourceassociation.md)

**Route 53 Resolver**

Route 53 Resolver endpoints

- [AssociateResolverEndpointIpAddress](api-route53resolver-associateresolverendpointipaddress.md)

- [CreateResolverEndpoint](api-route53resolver-createresolverendpoint.md)

- [DeleteResolverEndpoint](api-route53resolver-deleteresolverendpoint.md)

- [DisassociateResolverEndpointIpAddress](api-route53resolver-disassociateresolverendpointipaddress.md)

- [GetResolverEndpoint](api-route53resolver-getresolverendpoint.md)

- [ListResolverEndpointIpAddresses](api-route53resolver-listresolverendpointipaddresses.md)

- [ListResolverEndpoints](api-route53resolver-listresolverendpoints.md)

- [UpdateResolverEndpoint](api-route53resolver-updateresolverendpoint.md)

DNS query logs

- [AssociateResolverQueryLogConfig](api-route53resolver-associateresolverquerylogconfig.md)

- [CreateResolverQueryLogConfig](api-route53resolver-createresolverquerylogconfig.md)

- [DeleteResolverQueryLogConfig](api-route53resolver-deleteresolverquerylogconfig.md)

- [DisassociateResolverQueryLogConfig](api-route53resolver-disassociateresolverquerylogconfig.md)

- [GetResolverQueryLogConfig](api-route53resolver-getresolverquerylogconfig.md)

- [GetResolverQueryLogConfigAssociation](api-route53resolver-getresolverquerylogconfigassociation.md)

- [GetResolverQueryLogConfigPolicy](api-route53resolver-getresolverquerylogconfigpolicy.md)

- [ListResolverQueryLogConfigAssociations](api-route53resolver-listresolverquerylogconfigassociations.md)

- [ListResolverQueryLogConfigs](api-route53resolver-listresolverquerylogconfigs.md)

- [PutResolverQueryLogConfigPolicy](api-route53resolver-putresolverquerylogconfigpolicy.md)

Route 53 Resolver rules

- [AssociateResolverRule](api-route53resolver-associateresolverrule.md)

- [CreateResolverRule](api-route53resolver-createresolverrule.md)

- [DeleteResolverRule](api-route53resolver-deleteresolverrule.md)

- [DisassociateResolverRule](api-route53resolver-disassociateresolverrule.md)

- [GetResolverRule](api-route53resolver-getresolverrule.md)

- [GetResolverRuleAssociation](api-route53resolver-getresolverruleassociation.md)

- [GetResolverRulePolicy](api-route53resolver-getresolverrulepolicy.md)

- [ListResolverRuleAssociations](api-route53resolver-listresolverruleassociations.md)

- [ListResolverRules](api-route53resolver-listresolverrules.md)

- [PutResolverRulePolicy](api-route53resolver-putresolverrulepolicy.md)

- [UpdateResolverRule](api-route53resolver-updateresolverrule.md)

Route 53 Resolver DNS Firewall

- [AssociateFirewallRuleGroup](api-route53resolver-associatefirewallrulegroup.md)

- [CreateFirewallDomainList](api-route53resolver-createfirewalldomainlist.md)

- [CreateFirewallRule](api-route53resolver-createfirewallrule.md)

- [CreateFirewallRuleGroup](api-route53resolver-createfirewallrulegroup.md)

- [DeleteFirewallDomainList](api-route53resolver-deletefirewalldomainlist.md)

- [DeleteFirewallRule](api-route53resolver-deletefirewallrule.md)

- [DeleteFirewallRuleGroup](api-route53resolver-deletefirewallrulegroup.md)

- [DisassociateFirewallRuleGroup](api-route53resolver-disassociatefirewallrulegroup.md)

- [GetFirewallConfig](api-route53resolver-getfirewallconfig.md)

- [GetFirewallDomainList](api-route53resolver-getfirewalldomainlist.md)

- [GetFirewallRuleGroup](api-route53resolver-getfirewallrulegroup.md)

- [GetFirewallRuleGroupAssociation](api-route53resolver-getfirewallrulegroupassociation.md)

- [GetFirewallRuleGroupPolicy](api-route53resolver-getfirewallrulegrouppolicy.md)

- [ImportFirewallDomains](api-route53resolver-importfirewalldomains.md)

- [ListFirewallConfigs](api-route53resolver-listfirewallconfigs.md)

- [ListFirewallDomainLists](api-route53resolver-listfirewalldomainlists.md)

- [ListFirewallDomains](api-route53resolver-listfirewalldomains.md)

- [ListFirewallRuleGroupAssociations](api-route53resolver-listfirewallrulegroupassociations.md)

- [ListFirewallRuleGroups](api-route53resolver-listfirewallrulegroups.md)

- [ListFirewallRules](api-route53resolver-listfirewallrules.md)

- [PutFirewallRuleGroupPolicy](api-route53resolver-putfirewallrulegrouppolicy.md)

- [UpdateFirewallConfig](api-route53resolver-updatefirewallconfig.md)

- [UpdateFirewallDomains](api-route53resolver-updatefirewalldomains.md)

- [UpdateFirewallRule](api-route53resolver-updatefirewallrule.md)

- [UpdateFirewallRuleGroupAssociation](api-route53resolver-updatefirewallrulegroupassociation.md)

Route 53 Resolver on Outposts

- [CreateOutpostResolver](api-route53resolver-createoutpostresolver.md)

- [DeleteOutpostResolver](api-route53resolver-deleteoutpostresolver.md)

- [GetOutpostResolver](api-route53resolver-getoutpostresolver.md)

- [ListOutpostResolvers](api-route53resolver-listoutpostresolvers.md)

- [UpdateOutpostResolver](api-route53resolver-updateoutpostresolver.md)

Route 53 Resolver configuration

- [GetResolverConfig](api-route53resolver-getresolverconfig.md)

- [ListResolverConfigs](api-route53resolver-listresolverconfigs.md)

- [UpdateResolverConfig](api-route53resolver-updateresolverconfig.md)

**Tags**

Tags for hosted zones and health checks

- [ChangeTagsForResource](api-changetagsforresource.md)

- [ListTagsForResource](api-listtagsforresource.md)

- [ListTagsForResources](api-listtagsforresources.md)

Tags for domains

- [DeleteTagsForDomain](api-domains-deletetagsfordomain.md)

- [ListTagsForDomain](api-domains-listtagsfordomain.md)

- [UpdateTagsForDomain](api-domains-updatetagsfordomain.md)

Tags for Route 53 Resolver

- [ListTagsForResource](api-route53resolver-listtagsforresource.md)

- [TagResource](api-route53resolver-tagresource.md)

- [UntagResource](api-route53resolver-untagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Welcome

Actions
