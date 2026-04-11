# GetGeoLocation

Gets information about whether a specified geographic location is supported for Amazon
Route 53 geolocation resource record sets.

Route 53 does not perform authorization for this API because it retrieves information
that is already available to the public.

Use the following syntax to determine whether a continent is supported for
geolocation:

`GET /2013-04-01/geolocation?continentcode=two-letter abbreviation for
					a continent
      `

Use the following syntax to determine whether a country is supported for
geolocation:

`GET /2013-04-01/geolocation?countrycode=two-character country
					code
      `

Use the following syntax to determine whether a subdivision of a country is supported
for geolocation:

`GET /2013-04-01/geolocation?countrycode=two-character country
					code&subdivisioncode=subdivision
			code
      `

## Request Syntax

```nohighlight

GET /2013-04-01/geolocation?continentcode=ContinentCode&countrycode=CountryCode&subdivisioncode=SubdivisionCode HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[continentcode](#API_GetGeoLocation_RequestSyntax)**

For geolocation resource record sets, a two-letter abbreviation that identifies a
continent. Amazon Route 53 supports the following continent codes:

- **AF**: Africa

- **AN**: Antarctica

- **AS**: Asia

- **EU**: Europe

- **OC**: Oceania

- **NA**: North America

- **SA**: South America

Length Constraints: Fixed length of 2.

**[countrycode](#API_GetGeoLocation_RequestSyntax)**

Amazon Route 53 uses the two-letter country codes that are specified in [ISO standard 3166-1\
alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).

Route 53 also supports the country code **UA** for
Ukraine.

Length Constraints: Minimum length of 1. Maximum length of 2.

**[subdivisioncode](#API_GetGeoLocation_RequestSyntax)**

The code for the subdivision, such as a particular state within the United States. For
a list of US state abbreviations, see [Appendix B: Two–Letter State and\
Possession Abbreviations](https://pe.usps.com/text/pub28/28apb.htm) on the United States Postal Service website. For a
list of all supported subdivision codes, use the [ListGeoLocations](api-listgeolocations.md)
API.

Length Constraints: Minimum length of 1. Maximum length of 3.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetGeoLocationResponse>
   <GeoLocationDetails>
      <ContinentCode>string</ContinentCode>
      <ContinentName>string</ContinentName>
      <CountryCode>string</CountryCode>
      <CountryName>string</CountryName>
      <SubdivisionCode>string</SubdivisionCode>
      <SubdivisionName>string</SubdivisionName>
   </GeoLocationDetails>
</GetGeoLocationResponse>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetGeoLocationResponse](#API_GetGeoLocation_ResponseSyntax)**

Root level tag for the GetGeoLocationResponse parameters.

Required: Yes

**[GeoLocationDetails](#API_GetGeoLocation_ResponseSyntax)**

A complex type that contains the codes and full continent, country, and subdivision
names for the specified geolocation code.

Type: [GeoLocationDetails](api-geolocationdetails.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The input is not valid.

**message**

HTTP Status Code: 400

**NoSuchGeoLocation**

Amazon Route 53 doesn't support the specified geographic location. For a list of
supported geolocation codes, see the [GeoLocation](api-geolocation.md) data
type.

**message**

HTTP Status Code: 404

## Examples

### Example Request

To determine whether France (FR) is supported for Route 53 geolocation, submit
the following request.

```

GET /2013-04-01/geolocation?countrycode=FR
```

### Example Response

The following response shows that France is supported for geolocation. If
France were not supported, Route 53 would return
`NoSuchGeoLocation`.

```

HTTP/1.1 200 OK
<?xml version="1.0" encoding="UTF-8"?>
<GetGeoLocationResponse xmlns="https://route53.amazonaws.com/doc/2013-04-01/">
   <GetGeoLocationDetails>
      <CountryCode>FR</CountryCode>
      <CountryName>France</CountryName>
   </GeoLocationDetails>
</GetGeoLocationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53-2013-04-01/getgeolocation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/getgeolocation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetDNSSEC

GetHealthCheck
