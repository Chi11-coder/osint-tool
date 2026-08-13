import z from "zod";
import { SEARCH_TYPE } from "../types/common.types";

const domainRegex = new RegExp(
  /^([a-zA-Z0-9]([a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$/,
);
const fileRegex = new RegExp(/^[a-fA-F0-9]{64}$/);

type SearchTypeValue = (typeof SEARCH_TYPE)[keyof typeof SEARCH_TYPE];
/**
 * searchTypes
 * ユーザーが入力した文字列がIPアドレス、ドメイン、SHA256であるか判定する
 * @param inputValue string
 */
export function searchTypes(inputValue: string): SearchTypeValue {
  const ipv4 = z.ipv4();
  if (ipv4.safeParse(inputValue).success) {
    return SEARCH_TYPE.HOST_SEARCH;
  }

  const isDomain = inputValue.match(domainRegex);
  if (isDomain) {
    return SEARCH_TYPE.HOST_SEARCH;
  }

  const isFile = inputValue.match(fileRegex);
  if (isFile) {
    return SEARCH_TYPE.FILE_SEARCH;
  }

  return SEARCH_TYPE.INVALID_SEARCH;
}
