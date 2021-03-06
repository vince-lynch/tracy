package configure

// DefaultConfig is the default configuration that is used when a configuration
// file is not found.
const DefaultConfig = `{
  "tracers": {
    "zzPLAINzz": "[[ID]]",
    "zzXSSzz": "\\\"'<[[ID]]>",
    "GEN-XSS": "\\\"'<[[ID]]>",
    "GEN-PLAIN": "[[ID]]"
   },
  "server-whitelist": ["127.0.0.1:3000"],
  "tracer-server": "127.0.0.1:7777",
  "external-hostname": "",
  "auto-launch": "false",
  "public-key-loc": "%s",
  "private-key-loc": "%s",
  "version": "%.2f"
}`
