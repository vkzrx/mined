"use client";

import * as React from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { AlertDialog, Button } from "@radix-ui/themes";
import {
  ChevronLeftIcon,
  PauseIcon,
  PlayIcon,
  StopIcon,
} from "@radix-ui/react-icons";
import { InstanceStatus } from "./instance-status";
import * as actions from "~/app/actions";
import type { Instance } from "~/lib/schemas";

type ToolbarProps = {
  instance: {
    name: string;
    status: Instance["status"];
    networkIP?: string;
    zone: string;
    project: string;
  };
};

export function Toolbar(props: ToolbarProps) {
  const router = useRouter();

  async function startVM() {
    await actions.startVM({
      name: props.instance.name,
      project: props.instance.project,
      zone: props.instance.zone,
    });
    router.refresh();
  }

  async function stopVM() {
    await actions.stopVM({
      name: props.instance.name,
      project: props.instance.project,
      zone: props.instance.zone,
    });
    router.refresh();
  }

  return (
    <div className="flex h-16 items-center space-x-8 px-8 border-b border-b-gray-800">
      <div className="flex items-center">
        <Link href="/" className="group">
          <ChevronLeftIcon className="w-6 h-6 p-1 mr-1 rounded-full duration-100 group-hover:bg-gray-800 group-hover:text-indigo-500" />
        </Link>
        <div className="font-bold">{props.instance.name}</div>
      </div>
      <InstanceStatus status={props.instance.status} size="2" />
      <form action={startVM}>
        <Button
          type="submit"
          variant="surface"
          aria-disabled={props.instance.status !== "TERMINATED"}
          disabled={props.instance.status !== "TERMINATED"}
        >
          <PlayIcon />
          Start
        </Button>
      </form>

      <AlertDialog.Root>
        <AlertDialog.Trigger>
          <Button
            type="button"
            variant="surface"
            color="amber"
            aria-disabled={props.instance.status !== "RUNNING"}
            disabled={props.instance.status !== "RUNNING"}
          >
            <PauseIcon />
            Suspend
          </Button>
        </AlertDialog.Trigger>
        <AlertDialog.Content className="max-w-md">
          <AlertDialog.Title>Suspend virtual machine</AlertDialog.Title>
          <AlertDialog.Description size="2">
            Are you sure? The players currently connected to the server will be
            disconnected.
          </AlertDialog.Description>
          <div className="flex justify-end mt-4 gap-3">
            <AlertDialog.Cancel>
              <Button type="button" variant="soft" color="gray">
                Cancel
              </Button>
            </AlertDialog.Cancel>
            <AlertDialog.Action>
              <form action={stopVM}>
                <Button
                  type="submit"
                  variant="solid"
                  color="amber"
                  onClick={stopVM}
                >
                  Confirm
                </Button>
              </form>
            </AlertDialog.Action>
          </div>
        </AlertDialog.Content>
      </AlertDialog.Root>

      <AlertDialog.Root>
        <AlertDialog.Trigger>
          <Button
            type="button"
            variant="surface"
            color="red"
            aria-disabled={props.instance.status !== "RUNNING"}
            disabled={props.instance.status !== "RUNNING"}
          >
            <StopIcon />
            Stop
          </Button>
        </AlertDialog.Trigger>
        <AlertDialog.Content className="max-w-md">
          <AlertDialog.Title>Stop virtual machine</AlertDialog.Title>
          <AlertDialog.Description size="2">
            Are you sure? The players currently connected to the server will be
            disconnected.
          </AlertDialog.Description>
          <div className="flex justify-end mt-4 gap-3">
            <AlertDialog.Cancel>
              <Button type="button" variant="soft" color="gray">
                Cancel
              </Button>
            </AlertDialog.Cancel>
            <AlertDialog.Action>
              <form action={stopVM}>
                <Button type="submit" variant="solid" color="red">
                  Confirm
                </Button>
              </form>
            </AlertDialog.Action>
          </div>
        </AlertDialog.Content>
      </AlertDialog.Root>
    </div>
  );
}
