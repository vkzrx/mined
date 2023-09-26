import { z } from "zod";

const schema = z.object({
  serverBaseUrl: z.string().url(),
  instance: z.object({
    project: z.string().min(1),
    zone: z.string().min(1),
  }),
});

const config = schema.parse({
  serverBaseUrl: process.env.SERVER_BASE_URL,
  instance: {
    project: process.env.INSTANCE_PROJECT_ID,
    zone: process.env.INSTANCE_ZONE,
  },
});

export { config };
