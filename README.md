# external-dns-infoblox-webhook

Infoblox provider based on in-tree provider for ExternalDNS. Supported records:

| Record Type | Status     |
|-------------|------------|
| A           | supported  |
| CNAME       | supported  |
| TXT         | supported  |
| PTR         | not tested |


## Quick start

Required environment variables:

| Environment Variable        | Default value | Required |
|-----------------------------|---------------|----------|
| INFOBLOX_HOST               | localhost     | true     |
| INFOBLOX_PORT               | 443           | true     |   
| INFOBLOX_WAPI_USER          |               | true     |
| INFOBLOX_WAPI_PASSWORD      |               | true     |
| INFOBLOX_VERSION            |               | true     |
| INFOBLOX_SSL_VERIFY         | true          | false    |
| INFOBLOX_DRY_RUN            | false         | false    |
| INFOBLOX_VIEW               | default       | false    |
| INFOBLOX_MAX_RESULTS        | 1500          | false    |
| INFOBLOX_CREATE_PTR         | false         | false    |
| INFOBLOX_DEFAULT_TTL        | 300           | false    |
