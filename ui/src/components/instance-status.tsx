import { Text, Responsive } from "@radix-ui/themes";
import {
  CheckCircledIcon,
  ExclamationTriangleIcon,
  LightningBoltIcon,
  MinusCircledIcon,
  PauseIcon,
  RocketIcon,
  StopwatchIcon,
} from "@radix-ui/react-icons";
import type { Instance } from "~/lib/schemas";

type InstanceStatusProps = {
  status: Instance["status"];
  size: Responsive<"1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9">;
};

export function InstanceStatus(props: InstanceStatusProps): JSX.Element {
  return props.status === "RUNNING" ? (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="green"
    >
      {props.status.toLowerCase()}
      <CheckCircledIcon className="ml-1" />
    </Text>
  ) : props.status === "STAGING" ? (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="teal"
    >
      {props.status.toLowerCase()}
      <RocketIcon className="ml-1" />
    </Text>
  ) : props.status === "PROVISIONING" ? (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="blue"
    >
      {props.status.toLowerCase()}
      <LightningBoltIcon className="ml-1" />
    </Text>
  ) : props.status === "REPAIRING" ? (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="tomato"
    >
      {props.status.toLowerCase()}
      <ExclamationTriangleIcon className="ml-1" />
    </Text>
  ) : props.status === "SUSPENDING" ? (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="amber"
    >
      {props.status.toLowerCase()}
      <StopwatchIcon className="ml-1" />
    </Text>
  ) : props.status === "SUSPENDED" ? (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="orange"
    >
      {props.status.toLowerCase()}
      <PauseIcon className="ml-1" />
    </Text>
  ) : props.status === "STOPPING" ? (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="amber"
    >
      {props.status.toLowerCase()}
      <StopwatchIcon className="ml-1" />
    </Text>
  ) : (
    <Text
      size={props.size}
      className="flex items-center capitalize"
      color="gray"
    >
      {props.status.toLowerCase()}
      <MinusCircledIcon className="ml-1" />
    </Text>
  );
}
