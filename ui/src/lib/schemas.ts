import { z } from "zod";

export const instanceSchema = z.object({
  name: z.string(),
  status: z.enum([
    "PROVISIONING",
    "REPAIRING",
    "RUNNING",
    "STAGING",
    "STOPPING",
    "SUSPENDING",
    "SUSPENDED",
    "TERMINATED",
  ]),
  zone: z
    .string()
    .url()
    .transform((val) => {
      const strs = val.split("/");
      return strs[strs.length - 1];
    }),
  networkIP: z.string().optional(),
  machineType: z.string().transform((val) => {
    const strs = val.split("/");
    return strs[strs.length - 1];
  }),
  cpuPlatform: z.string(),
  creationTime: z.string(),
});

export const instancesSchema = z.array(instanceSchema);

export type Instance = z.infer<typeof instanceSchema>;
