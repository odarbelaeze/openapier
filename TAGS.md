# OpenAPI Tags Documentation

This document lists all the supported `@` tags and validation tags for the `openapier` tool.

## Spec-level Tags

These tags are used to define general information about the API (Info, Servers, etc.).

| Tag | Usage |
| :--- | :--- |
| `@externalDocs.description` | `@externalDocs.description <description>` |
| `@externalDocs.url` | `@externalDocs.url <url>` |
| `@info.contact.email` | `@info.contact.email <email>` |
| `@info.contact.name` | `@info.contact.name <name>` |
| `@info.contact.url` | `@info.contact.url <url>` |
| `@info.description` | `@info.description <description>` |
| `@info.license.identifier` | `@info.license.identifier <identifier>` |
| `@info.license.name` | `@info.license.name <name>` |
| `@info.license.url` | `@info.license.url <url>` |
| `@info.summary` | `@info.summary <summary>` |
| `@info.termsOfService` | `@info.termsOfService <url>` |
| `@info.title` | `@info.title <title>` |
| `@info.version` | `@info.version <version>` |
| `@security` | `@security <name> [scope1] [scope2] ...` |
| `@securityScheme` | `@securityScheme <name> <type> [<description>...]` |
| `@securityScheme.bearerFormat` | `@securityScheme.bearerFormat <securitySchemeName> <format>` |
| `@securityScheme.in` | `@securityScheme.in <securitySchemeName> <in>` |
| `@securityScheme.name` | `@securityScheme.name <securitySchemeName> <name>` |
| `@securityScheme.scheme` | `@securityScheme.scheme <securitySchemeName> <scheme>` |
| `@server.description` | `@server.description <description>` |
| `@server.url` | `@server.url <url>` |
| `@server.variable.default` | `@server.variable.default <variable> <default>` |
| `@server.variable.description` | `@server.variable.description <variable> <description>` |
| `@server.variable.enum` | `@server.variable.enum <variable> [value1] [value2] ...` |
| `@tag.description` | `@tag.description <description>` |
| `@tag.externalDocs.description` | `@tag.externalDocs.description <description>` |
| `@tag.externalDocs.url` | `@tag.externalDocs.url <url>` |
| `@tag.name` | `@tag.name <name>` |

## Operation-level Tags

These tags are used to define individual API operations (Paths, Parameters, Responses, etc.).

| Tag | Usage |
| :--- | :--- |
| `@deprecated` | `@deprecated` |
| `@description` | `@description <description>` |
| `@externaldocs.description` | `@externalDocs.description <description>` |
| `@externaldocs.url` | `@externalDocs.url <url>` |
| `@id` | `@id <operationId>` |
| `@param` | `@param <name> <type> <in> [description...]` |
| `@param.allowEmptyValue` | `@param.allowEmptyValue <param> [param]...` |
| `@param.allowReserved` | `@param.allowReserved <param> [param]...` |
| `@param.deprecated` | `@param.deprecated <param> [param]...` |
| `@param.required` | `@param.required <param> [param]...` |
| `@requestBody` | `@requestBody <content_type> <type> [description]` |
| `@response` | `@response <status_code> <content_type> <type> [description]` |
| `@router` | `@router <path> [<method>]` |
| `@security` | `@security <name> [scope1] [scope2] ...` |
| `@security.none` | `@security.none` |
| `@server.description` | `@server.description <description>` |
| `@server.url` | `@server.url <url>` |
| `@server.variable.default` | `@server.variable.default <variable> <default>` |
| `@server.variable.description` | `@server.variable.description <variable> <description>` |
| `@server.variable.enum` | `@server.variable.enum <variable> [value1] [value2] ...` |
| `@summary` | `@summary <summary>` |
| `@tags` | `@tags <tag1> [tag2]...` |

## Validation Tags

These tags are used in Go struct fields within the `validate` tag to define schema constraints.

| Tag | Usage |
| :--- | :--- |
| `alpha` | `alpha` |
| `alphanum` | `alphanum` |
| `alphanumspace` | `alphanumspace` |
| `alphaspace` | `alphaspace` |
| `ascii` | `ascii` |
| `base64` | `base64` |
| `base64rawurl` | `base64rawurl` |
| `base64url` | `base64url` |
| `cidr` | `cidr` |
| `cidrv4` | `cidrv4` |
| `cidrv6` | `cidrv6` |
| `cmyk` | `cmyk` |
| `contains` | `contains=value` |
| `cron` | `cron` |
| `cve` | `cve` |
| `datauri` | `datauri` |
| `datetime` | `datetime` |
| `e164` | `e164` |
| `email` | `email` |
| `endswith` | `endswith=x` |
| `eq` | `eq=x` |
| `fqdn` | `fqdn` |
| `gt` | `gt=x` |
| `gte` | `gte=x` |
| `hexadecimal` | `hexadecimal` |
| `hexcolor` | `hexcolor` |
| `hostname` | `hostname` |
| `hostname_port` | `hostname_port` |
| `hostname_rfc1123` | `hostname_rfc1123` |
| `hsl` | `hsl` |
| `hsla` | `hsla` |
| `http_url` | `http_url` |
| `https_url` | `https_url` |
| `ip` | `ip` |
| `ip4_addr` | `ip4_addr` |
| `ip6_addr` | `ip6_addr` |
| `ip_addr` | `ip_addr` |
| `ipv4` | `ipv4` |
| `ipv6` | `ipv6` |
| `isbn` | `isbn` |
| `iso3166_1_alpha2` | `iso3166_1_alpha2` |
| `iso3166_1_alpha3` | `iso3166_1_alpha3` |
| `iso4217` | `iso4217` |
| `json` | `json` |
| `jwt` | `jwt` |
| `latitude` | `latitude` |
| `len` | `len=x` |
| `longitude` | `longitude` |
| `lowercase` | `lowercase` |
| `lt` | `lt=x` |
| `lte` | `lte=x` |
| `mac` | `mac` |
| `max` | `max=x` |
| `min` | `min=x` |
| `ne` | `ne=value` |
| `noneof` | `noneof=value1 value2` |
| `number` | `number` |
| `numeric` | `numeric` |
| `oneof` | `oneof=val1 val2 val3` |
| `port` | `port` |
| `printascii` | `printascii` |
| `required` | `required` |
| `rgb` | `rgb` |
| `rgba` | `rgba` |
| `semver` | `semver` |
| `ssn` | `ssn` |
| `startswith` | `startswith=x` |
| `ulid` | `ulid` |
| `unique` | `unique` |
| `uppercase` | `uppercase` |
| `uri` | `uri` |
| `url` | `url` |
| `uuid` | `uuid` |
