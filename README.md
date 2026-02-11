# uuid

Generate UUIDs from the command line.

```shell
Usage: uuid -v<version> [options]

Versions:
  -v1             Time-based (RFC 4122)
  -v2             DCE Security (uses current UID)
  -v3 <name>      Name-based MD5 (default namespace: dns)
  -v4             Random (RFC 4122)
  -v5 <name>      Name-based SHA1 (default namespace: dns)
  -v6             Reordered time-based (RFC 9562)
  -v7             Unix epoch time-based (RFC 9562)

Options for v3/v5:
  -ns <namespace>  Namespace: dns, url, oid, x500 (default: dns)
```
