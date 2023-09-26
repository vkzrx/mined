import Link from "next/link";
import { Card, Text } from "@radix-ui/themes";
import { InstanceStatus } from "~/components/instance-status";
import { instancesSchema, type Instance } from "~/lib/schemas";
import * as http from "~/lib/http";
import { config } from "~/lib/config";

type GetInstancesParams = {
  project: string;
  zone: string;
};

async function getInstances({
  project,
  zone,
}: GetInstancesParams): Promise<Instance[]> {
  const response = await http.get({
    url: `/vms/${project}?zone=${zone}`,
    options: { cache: "no-store" },
  });
  return instancesSchema.parse(response);
}

export default async function Home() {
  const instances = await getInstances({
    project: config.instance.project,
    zone: config.instance.zone,
  });
  return (
    <div className="grow min-h-screen flex flex-col">
      <div className="flex h-16 items-center space-x-8 px-8 border-b border-b-gray-800 font-bold">
        Instances
      </div>
      <main className="flex flex-col p-8">
        {instances.map((instance) => (
          <Card key={instance.name} asChild size="2" className="max-w-xs">
            <Link
              href={`/instance/${instance.name}?project=${config.instance.project}&zone=${instance.zone}`}
            >
              <div className="flex flex-col">
                <div className="flex mb-4">
                  <Text size="3" weight="bold" mr="4">
                    {instance.name}
                  </Text>
                  <InstanceStatus status={instance.status} size="2" />
                </div>
                <div className="flex justify-between">
                  <Text size="2">{instance.zone}</Text>
                  <Text size="2">{instance.networkIP}</Text>
                </div>
              </div>
            </Link>
          </Card>
        ))}
      </main>
    </div>
  );
}
