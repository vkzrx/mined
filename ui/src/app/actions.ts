"use server";

import * as http from "~/lib/http";

type StartVMParams = {
  name: string;
  project: string;
  zone: string;
};

export async function startVM(params: StartVMParams) {
  await http.post<{ message: string }>({
    url: `/vms/start/${params.project}/${params.name}?zone=${params.zone}`,
    options: { cache: "no-cache" },
  });
}

type StopVMParams = {
  name: string;
  project: string;
  zone: string;
};

export async function stopVM(params: StopVMParams) {
  await http.post<{ message: string }>({
    url: `/vms/stop/${params.project}/${params.name}?zone=${params.zone}`,
    options: { cache: "no-cache" },
  });
}
