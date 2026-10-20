export const TENANT_BASE_DOMAIN =
  process.env.NEXT_PUBLIC_TENANT_BASE_DOMAIN || "edoutput.dev";

export const RESERVED_TENANT_SUBDOMAINS = (
  process.env.NEXT_PUBLIC_RESERVED_SUBDOMAINS ||
  "www,tourism-agency,gps-medical-assistance,business-manager"
)
  .split(",")
  .map((value) => value.trim().toLowerCase())
  .filter(Boolean);
