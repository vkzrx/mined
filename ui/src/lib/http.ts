import { config } from "~/lib/config";

type RequestParams = {
  url: string;
  options?: Omit<RequestInit, "method">;
};

export async function get<T>(params: RequestParams): Promise<T> {
  const url = config.serverBaseUrl + params.url;
  const response = await fetch(url, {
    ...params.options,
    method: "GET",
  });
  if (!response.ok) {
    throw new Error(`GET: Failed request to ${params.url}`);
  }
  const res = await response.json();
  return res;
}

export async function post<T>(params: RequestParams): Promise<T> {
  const url = config.serverBaseUrl + params.url;
  const response = await fetch(url, {
    ...params.options,
    method: "POST",
  });
  if (!response.ok) {
    throw new Error(`POST: Failed request to ${params.url}`);
  }
  const res = await response.json();
  return res;
}
