import React, { useState } from "react";
import "../styles/searchForm.css";
import { IoSearch } from "react-icons/io5";
import { RiLoader2Fill } from "react-icons/ri";
import { searchTypes } from "../utils/searchTypes";
import {
  SEARCH_TYPE,
  SEARCH_URL,
  type SearchResponse,
  type SearchTypeValue,
} from "../types/common.types";

interface SearchFormProps {
  onSearchReports: (data: SearchResponse) => void;
  onSearchType: (type: SearchTypeValue) => void;
  errMsg: string;
}

export function SearchForm({
  onSearchReports,
  onSearchType,
  errMsg,
}: SearchFormProps) {
  const [searchValue, setSearchValue] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [searchErrMsg, setSearchErrMsg] = useState<string>(errMsg);

  const handleSearchSubmit = async (
    event: React.SubmitEvent<HTMLFormElement>,
  ) => {
    setIsLoading(true);
    event.preventDefault();

    const formData = new FormData(event.currentTarget);
    const rawQuery = formData.get("searchQuery");

    try {
      if (typeof rawQuery !== "string") {
        setSearchErrMsg("please input SHA256 or IPAddress or Domain");
        return;
      }
      const query = rawQuery.trim();

      if (!query) {
        setSearchErrMsg("please input SHA256 or IPAddress or Domain");
        return;
      }

      switch (searchTypes(query)) {
        case SEARCH_TYPE.HOST_SEARCH: {
          const hostReportResponse = await fetch(
            SEARCH_URL.HOST + `${encodeURIComponent(query)}`,
          );

          if (!hostReportResponse.ok) {
            setSearchErrMsg(
              `server error: status: ${hostReportResponse.status} ${hostReportResponse.statusText}`,
            );
            return;
          }
          const ipReport = await hostReportResponse.json();
          setSearchErrMsg("");
          onSearchReports(ipReport);
          onSearchType(SEARCH_TYPE.HOST_SEARCH);
          break;
        }

        case SEARCH_TYPE.FILE_SEARCH: {
          const fileReportResponse = await fetch(
            SEARCH_URL.FILE + `${encodeURIComponent(query)}`,
          );

          if (!fileReportResponse.ok) {
            setSearchErrMsg(
              `server error: status: ${fileReportResponse.status} ${fileReportResponse.statusText}`,
            );
            return;
          }

          const data = await fileReportResponse.json();
          console.log("file report data", data);
          setSearchErrMsg("");
          onSearchReports(data);
          onSearchType(SEARCH_TYPE.FILE_SEARCH);
          break;
        }
        default:
          setSearchErrMsg("不適切な検索文字列");
          break;
      }
    } catch (e) {
      setSearchErrMsg("error: " + e);
      return;
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="searchFormContainer">
      <h1 className="tool-title">OSINTツール</h1>
      <div>
        <p>ドメイン、IPアドレス、ファイル調査</p>
      </div>
      <form onSubmit={handleSearchSubmit}>
        <input
          className="search-form"
          type="search"
          name="searchQuery"
          placeholder="SHA256, IPAddress, Domain..."
          value={searchValue}
          onChange={(e) => setSearchValue(e.target.value)}
        />
        <button
          className={`search-btn ${isLoading ? "loading" : ""}`}
          type="submit"
          disabled={isLoading}
        >
          {isLoading ? (
            <>
              <span className="spinner">
                <RiLoader2Fill />
              </span>
              <span>検索中</span>
            </>
          ) : (
            <>
              <IoSearch />
              <span>検索</span>
            </>
          )}
        </button>
      </form>
      {searchErrMsg !== "" ? (
        <span className="error">{searchErrMsg}</span>
      ) : (
        <></>
      )}
    </div>
  );
}
