import type { FileReport } from "./file.types";
import type { HostReport } from "./host.types";

export const NODE_WIDTH = 120;
export const NODE_HEIGHT = 30;

export type SearchResponse = FileReport | HostReport;
export type SearchTypeValue = (typeof SEARCH_TYPE)[keyof typeof SEARCH_TYPE];

export const SEARCH_TYPE = {
  INVALID_SEARCH: -1,
  HOST_SEARCH: 1,
  FILE_SEARCH: 2,
} as const;

export const SEARCH_URL = {
  HOST: "/host/report?host=",
  FILE: "/file/report?hash=",
} as const;
