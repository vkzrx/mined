import { Card, Text } from "@radix-ui/themes";
import { instanceSchema, type Instance } from "~/lib/schemas";
import { Toolbar } from "~/components/toolbar";
import * as http from "~/lib/http";
import { config } from "~/lib/config";

type GetInstanceArgs = {
  name: string;
  zone: string;
  project: string;
};

async function getInstance({
  name,
  zone,
  project,
}: GetInstanceArgs): Promise<Instance> {
  const response = await http.get({
    url: `/vms/${project}/${name}?zone=${zone}`,
    options: { cache: "no-store" },
  });
  return instanceSchema.parse(response);
}

type InstancePageProps = {
  params: { name: string };
  searchParams?: { project: string; zone: string };
};

export default async function InstancePage(props: InstancePageProps) {
  const instance = await getInstance({
    name: props.params.name,
    project: props.searchParams?.project || "",
    zone: props.searchParams?.zone || "",
  });

  return (
    <div className="grow">
      <Toolbar
        instance={{
          name: props.params.name,
          status: instance.status,
          networkIP: instance.networkIP,
          zone: config.instance.zone,
          project: config.instance.project,
        }}
      />
      <main className="p-8">
        <Card size="2" className="max-w-md">
          <Text size="5" color="iris" weight="medium">
            Details
          </Text>
          <div className="flex flex-col space-y-1 mt-2">
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                Status
              </Text>
              <Text as="span" className="capitalize">
                {instance.status.toLowerCase()}
              </Text>
            </div>
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                IP
              </Text>
              <Text as="span">{instance.networkIP || "None"}</Text>
            </div>
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                CPU
              </Text>
              <Text>{instance.cpuPlatform}</Text>
            </div>
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                Machine
              </Text>
              <Text>{instance.machineType}</Text>
            </div>
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                Project
              </Text>
              <Text>{props.searchParams?.project}</Text>
            </div>
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                Zone
              </Text>
              <Text>{instance.zone}</Text>
            </div>
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                Creation time
              </Text>
              <time dateTime={new Date(instance.creationTime).toISOString()}>
                {new Date(instance.creationTime).toLocaleTimeString("en-US", {
                  hour: "numeric",
                  minute: "numeric",
                  second: "numeric",
                })}
              </time>
            </div>
            <div className="flex">
              <Text as="span" className="w-2/5 font-semibold">
                Creation date
              </Text>
              <time dateTime={new Date(instance.creationTime).toISOString()}>
                {new Date(instance.creationTime).toLocaleDateString("en-US", {
                  month: "long",
                  day: "numeric",
                  year: "numeric",
                })}
              </time>
            </div>
          </div>
        </Card>
      </main>
    </div>
  );
}
